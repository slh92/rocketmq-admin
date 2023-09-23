/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package admin

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/apache/rocketmq-client-go/v2/internal"
	"github.com/apache/rocketmq-client-go/v2/internal/remote"
	"github.com/apache/rocketmq-client-go/v2/internal/utils"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
)

type Admin interface {
	CreateTopic(ctx context.Context, opts ...OptionCreate) error
	DeleteTopic(ctx context.Context, opts ...OptionDelete) error

	GetAllSubscriptionGroup(ctx context.Context, brokerAddr string, timeoutMillis time.Duration) (*SubscriptionGroupWrapper, error)
	FetchAllTopicList(ctx context.Context) (*TopicList, error)
	QueryTopicRouteInfo(topic string) (*internal.TopicRouteData, error)
	QueryConsumersByTopic(topic string) (*internal.TopicRouteData, error)
	QueryConsumerConnectsByGroup(ctx context.Context, group string) (*ConsumerConnection, error)
	//GetBrokerClusterInfo(ctx context.Context) (*remote.RemotingCommand, error)
	FetchPublishMessageQueues(ctx context.Context, topic string) ([]*primitive.MessageQueue, error)
	Close() error
}

// TODO: move outdated context to ctx
type adminOptions struct {
	internal.ClientOptions
}

type AdminOption func(options *adminOptions)

func defaultAdminOptions() *adminOptions {
	opts := &adminOptions{
		ClientOptions: internal.DefaultClientOptions(),
	}
	opts.GroupName = "TOOLS_ADMIN"
	opts.InstanceName = time.Now().String()
	return opts
}

// WithResolver nameserver resolver to fetch nameserver addr
func WithResolver(resolver primitive.NsResolver) AdminOption {
	return func(options *adminOptions) {
		options.Resolver = resolver
	}
}

func WithCredentials(c primitive.Credentials) AdminOption {
	return func(options *adminOptions) {
		options.ClientOptions.Credentials = c
	}
}

// WithNamespace set the namespace of admin
func WithNamespace(namespace string) AdminOption {
	return func(options *adminOptions) {
		options.ClientOptions.Namespace = namespace
	}
}

type MqAdmin struct {
	Cli internal.RMQClient

	opts *adminOptions

	closeOnce sync.Once
}

// NewAdmin initialize admin
func NewAdmin(opts ...AdminOption) (*MqAdmin, error) {
	defaultOpts := defaultAdminOptions()
	for _, opt := range opts {
		opt(defaultOpts)
	}
	namesrv, err := internal.NewNamesrv(defaultOpts.Resolver, defaultOpts.RemotingClientConfig)
	defaultOpts.Namesrv = namesrv
	if err != nil {
		return nil, err
	}

	cli := internal.GetOrNewRocketMQClient(defaultOpts.ClientOptions, nil)
	if cli == nil {
		return nil, fmt.Errorf("GetOrNewRocketMQClient faild")
	}
	defaultOpts.Namesrv = cli.GetNameSrv()
	//log.Printf("Client: %#v", namesrv.srvs)
	return &MqAdmin{
		Cli:  cli,
		opts: defaultOpts,
	}, nil
}

func (a *MqAdmin) GetAllSubscriptionGroup(ctx context.Context, brokerAddr string, timeoutMillis time.Duration) (*SubscriptionGroupWrapper, error) {
	cmd := remote.NewRemotingCommand(internal.ReqGetAllSubscriptionGroupConfig, nil, nil)
	a.Cli.RegisterACL()
	response, err := a.Cli.InvokeSync(ctx, brokerAddr, cmd, timeoutMillis)
	if err != nil {
		rlog.Error("Get all group list error", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
		return nil, err
	} else {
		rlog.Info("Get all group list success", map[string]interface{}{})
	}
	var subscriptionGroupWrapper SubscriptionGroupWrapper
	_, err = subscriptionGroupWrapper.Decode(response.Body, &subscriptionGroupWrapper)
	if err != nil {
		rlog.Error("Get all group list decode error", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
		return nil, err
	}
	return &subscriptionGroupWrapper, nil
}

func (a *MqAdmin) QueryGroupList() ([]*GroupConsumeInfo, error) {
	consumerGroupSet := []string{}
	clusterInfo, _ := a.getBrokerClusterInfo(context.Background())
	for _, brokerData := range clusterInfo.BrokerAddrTable {
		wrapper, err := a.GetAllSubscriptionGroup(context.Background(), brokerData.SelectBrokerAddr(), 3*time.Second)
		if err != nil {
			return nil, err
		}
		for key, _ := range wrapper.SubscriptionGroupTable {
			consumerGroupSet = append(consumerGroupSet, key)
		}
	}
	groupList := []*GroupConsumeInfo{}
	for _, group := range consumerGroupSet {
		res, err := a.QueryGroup(group)
		if err != nil {
			return nil, err
		}
		groupList = append(groupList, res)
	}
	return groupList, nil
}

func (a *MqAdmin) QueryGroup(group string) (*GroupConsumeInfo, error) {
	groupConsumeInfo := &GroupConsumeInfo{}
	consumeStats := a.examineConsumeStats(group)
	conn, err := a.QueryConsumerConnectsByGroup(context.Background(), group)
	if err != nil {
		return nil, err
	}
	groupConsumeInfo.Group = group
	if consumeStats != nil {
		groupConsumeInfo.ConsumeTps = int(consumeStats.ConsumeTps)
		groupConsumeInfo.DiffTotal = consumeStats.ComputeTotalDiff()
	}
	if conn != nil {
		groupConsumeInfo.Count = len(conn.ConnectionSet)
		groupConsumeInfo.MessageModel = conn.MessageModel
		groupConsumeInfo.ConsumeType = conn.ConsumeType
		groupConsumeInfo.Version = ""
	}
	return groupConsumeInfo, nil
}

func (a *MqAdmin) examineConsumeStats(group string) *ConsumeStats {
	return a.examineConsumeStatsWithTopic(group, "")
}

func (a *MqAdmin) examineConsumeStatsWithTopic(group string, topic string) *ConsumeStats {
	routeInfo, _ := a.QueryTopicRouteInfo(utils.MixAllUtil.GetRetryTopic(group))
	result := &ConsumeStats{}
	for _, bd := range routeInfo.BrokerDataList {
		addr := bd.SelectBrokerAddr()
		if addr != "" {
			header := &internal.GetConsumeStatsRequestHeader{
				ConsumerGroup: group,
				Topic:         topic,
			}
			cmd := remote.NewRemotingCommand(internal.ReqGetConsumerStatsFromServer, header, nil)
			response, err := a.Cli.InvokeSync(context.Background(), addr, cmd, 3*time.Second)
			if err != nil {
				rlog.Error("Fetch all consumerConnectiion list error", map[string]interface{}{
					rlog.LogKeyUnderlayError: err,
				})
				return nil
			} else {
				rlog.Info("Fetch all consumerConnectiion list success", map[string]interface{}{})
			}
			var stats ConsumeStats
			_, err = stats.Decode(response.Body, &stats)
			if err != nil {
				rlog.Error("Fetch all consumerConnectiion list decode error", map[string]interface{}{
					rlog.LogKeyUnderlayError: err,
				})
				return nil
			}
			for key, val := range stats.OffsetTable {
				result.OffsetTable[key] = val
			}
			result.ConsumeTps += stats.ConsumeTps
		}
	}
	return result
}

func (a *MqAdmin) FetchAllTopicList(ctx context.Context) (*TopicList, error) {
	cmd := remote.NewRemotingCommand(internal.ReqGetAllTopicListFromNameServer, nil, nil)
	response, err := a.Cli.InvokeSync(ctx, a.Cli.GetNameSrv().AddrList()[0], cmd, 3*time.Second)
	if err != nil {
		rlog.Error("Fetch all topic list error", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
		return nil, err
	} else {
		rlog.Info("Fetch all topic list success", map[string]interface{}{})
	}
	var topicList TopicList
	_, err = topicList.Decode(response.Body, &topicList)
	if err != nil {
		rlog.Error("Fetch all topic list decode error", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
		return nil, err
	}
	return &topicList, nil
}

func (a *MqAdmin) QueryConsumerConnectsByGroup(ctx context.Context, group string) (*ConsumerConnection, error) {
	header := &internal.GetConsumerConnectionListRequestHeader{
		ConsumerGroup: group,
	}
	cmd := remote.NewRemotingCommand(internal.ReqGetConsumerConnectionList, header, nil)
	response, err := a.Cli.InvokeSync(ctx, a.Cli.GetNameSrv().AddrList()[0], cmd, 3*time.Second)
	if err != nil {
		rlog.Error("Fetch all consumerConnectiion list error", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
		return nil, err
	} else {
		rlog.Info("Fetch all consumerConnectiion list success", map[string]interface{}{})
	}
	var connect ConsumerConnection
	_, err = connect.Decode(response.Body, &connect)
	if err != nil {
		rlog.Error("Fetch all consumerConnectiion list decode error", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
		return nil, err
	}
	return &connect, nil
}

// CreateTopic create topic.
// TODO: another implementation like sarama, without brokerAddr as input
func (a *MqAdmin) CreateTopic(ctx context.Context, opts ...OptionCreate) error {
	cfg := defaultTopicConfigCreate()
	for _, apply := range opts {
		apply(&cfg)
	}

	request := &internal.CreateTopicRequestHeader{
		Topic:           cfg.Topic,
		DefaultTopic:    cfg.DefaultTopic,
		ReadQueueNums:   cfg.ReadQueueNums,
		WriteQueueNums:  cfg.WriteQueueNums,
		Perm:            cfg.Perm,
		TopicFilterType: cfg.TopicFilterType,
		TopicSysFlag:    cfg.TopicSysFlag,
		Order:           cfg.Order,
	}

	cmd := remote.NewRemotingCommand(internal.ReqCreateTopic, request, nil)
	_, err := a.Cli.InvokeSync(ctx, cfg.BrokerAddr, cmd, 5*time.Second)
	if err != nil {
		rlog.Error("create topic error", map[string]interface{}{
			rlog.LogKeyTopic:         cfg.Topic,
			rlog.LogKeyBroker:        cfg.BrokerAddr,
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("create topic success", map[string]interface{}{
			rlog.LogKeyTopic:  cfg.Topic,
			rlog.LogKeyBroker: cfg.BrokerAddr,
		})
	}
	return err
}

// DeleteTopicInBroker delete topic in broker.
func (a *MqAdmin) deleteTopicInBroker(ctx context.Context, topic string, brokerAddr string) (*remote.RemotingCommand, error) {
	request := &internal.DeleteTopicRequestHeader{
		Topic: topic,
	}

	cmd := remote.NewRemotingCommand(internal.ReqDeleteTopicInBroker, request, nil)
	return a.Cli.InvokeSync(ctx, brokerAddr, cmd, 5*time.Second)
}

// DeleteTopicInNameServer delete topic in nameserver.
func (a *MqAdmin) deleteTopicInNameServer(ctx context.Context, topic string, nameSrvAddr string) (*remote.RemotingCommand, error) {
	request := &internal.DeleteTopicRequestHeader{
		Topic: topic,
	}

	cmd := remote.NewRemotingCommand(internal.ReqDeleteTopicInNameSrv, request, nil)
	return a.Cli.InvokeSync(ctx, nameSrvAddr, cmd, 5*time.Second)
}

// DeleteTopic delete topic in both broker and nameserver.
func (a *MqAdmin) DeleteTopic(ctx context.Context, opts ...OptionDelete) error {
	cfg := defaultTopicConfigDelete()
	for _, apply := range opts {
		apply(&cfg)
	}
	//delete topic in broker
	if cfg.BrokerAddr == "" {
		a.Cli.GetNameSrv().UpdateTopicRouteInfo(cfg.Topic)
		cfg.BrokerAddr = a.Cli.GetNameSrv().FindBrokerAddrByTopic(cfg.Topic)
	}

	if _, err := a.deleteTopicInBroker(ctx, cfg.Topic, cfg.BrokerAddr); err != nil {
		rlog.Error("delete topic in broker error", map[string]interface{}{
			rlog.LogKeyTopic:         cfg.Topic,
			rlog.LogKeyBroker:        cfg.BrokerAddr,
			rlog.LogKeyUnderlayError: err,
		})
		return err
	}

	//delete topic in nameserver
	if len(cfg.NameSrvAddr) == 0 {
		a.Cli.GetNameSrv().UpdateTopicRouteInfo(cfg.Topic)
		cfg.NameSrvAddr = a.Cli.GetNameSrv().AddrList()
		_, _, err := a.Cli.GetNameSrv().UpdateTopicRouteInfo(cfg.Topic)
		if err != nil {
			rlog.Error("delete topic in nameserver error", map[string]interface{}{
				rlog.LogKeyTopic:         cfg.Topic,
				rlog.LogKeyUnderlayError: err,
			})
		}
		cfg.NameSrvAddr = a.Cli.GetNameSrv().AddrList()
	}

	for _, nameSrvAddr := range cfg.NameSrvAddr {
		if _, err := a.deleteTopicInNameServer(ctx, cfg.Topic, nameSrvAddr); err != nil {
			rlog.Error("delete topic in nameserver error", map[string]interface{}{
				"nameServer":             nameSrvAddr,
				rlog.LogKeyTopic:         cfg.Topic,
				rlog.LogKeyUnderlayError: err,
			})
			return err
		}
	}
	rlog.Info("delete topic success", map[string]interface{}{
		"nameServer":      cfg.NameSrvAddr,
		rlog.LogKeyTopic:  cfg.Topic,
		rlog.LogKeyBroker: cfg.BrokerAddr,
	})
	return nil
}

func (a *MqAdmin) FetchPublishMessageQueues(ctx context.Context, topic string) ([]*primitive.MessageQueue, error) {
	return a.Cli.GetNameSrv().FetchPublishMessageQueues(utils.WrapNamespace(a.opts.Namespace, topic))
}

func (a *MqAdmin) QueryTopicRouteInfo(topic string) (*internal.TopicRouteData, error) {
	return a.Cli.GetNameSrv().FindRouteInfoByTopic(topic)
}

func (a *MqAdmin) getBrokerClusterInfo(ctx context.Context) (*ClusterInfo, error) {
	cmd := remote.NewRemotingCommand(internal.ReqGetBrokerClusterInfo, nil, nil)
	res, err := a.Cli.InvokeSync(ctx, "", cmd, 5*time.Second)
	if err != nil {
		rlog.Error("Fetch all clusterinfo list error", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
	}
	var cluster ClusterInfo
	_, err = cluster.Decode(res.Body, &cluster)
	if err != nil {
		rlog.Error("Fetch all consumerConnectiion list decode error", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
		return nil, err
	}
	return &cluster, nil
}

func (a *MqAdmin) Close() error {
	a.closeOnce.Do(func() {
		a.Cli.Shutdown()
	})
	return nil
}
