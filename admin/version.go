package admin

type ClientVersion string

const (
	V3_0_0_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_0_ALPHA1         = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_0_BETA1          = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_0_BETA2          = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_0_BETA3          = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_0_BETA4          = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_0_BETA5          = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_0_BETA6_SNAPSHOT = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_0_BETA6          = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_0_BETA7_SNAPSHOT = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_0_BETA7          = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_0_BETA8_SNAPSHOT = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_0_BETA8          = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_0_BETA9_SNAPSHOT = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_0_BETA9          = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_0_FINAL          = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_1_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_1                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_2_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_2                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_3_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_3                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_4_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_4                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_5_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_5                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_6_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_6                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_7_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_7                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_8_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_8                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_9_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_9                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_10_SNAPSHOT      = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_10               = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_11_SNAPSHOT      = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_11               = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_12_SNAPSHOT      = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_12               = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_13_SNAPSHOT      = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_13               = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_14_SNAPSHOT      = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_14               = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_15_SNAPSHOT      = ClientVersion("V4_0_0_SNAPSHOT")
	V3_0_15               = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_0_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_0                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_1_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_1                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_2_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_2                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_3_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_3                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_4_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_4                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_5_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_5                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_6_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_6                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_7_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_7                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_8_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_8                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_9_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_1_9                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_0_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_0                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_1_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_1                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_2_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_2                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_3_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_3                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_4_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_4                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_5_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_5                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_6_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_6                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_7_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_7                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_8_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_8                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_9_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_2_9                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_3_1_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_3_1                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_3_2_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_3_2                = ClientVersion("V4_0_0_SNAPSHOT")
	V3_3_3_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V3_3_3                = ClientVersion("V3_3_3")
	V3_3_4_SNAPSHOT       = ClientVersion("V3_3_4_SNAPSHOT")
	V3_3_4                = ClientVersion("V3_3_4")
	V3_3_5_SNAPSHOT       = ClientVersion("V3_3_5_SNAPSHOT")
	V3_3_5                = ClientVersion("V3_3_5")
	V3_3_6_SNAPSHOT       = ClientVersion("V3_3_6_SNAPSHOT")
	V3_3_6                = ClientVersion("V3_3_6")
	V3_3_7_SNAPSHOT       = ClientVersion("V3_3_7_SNAPSHOT")
	V3_3_7                = ClientVersion("V3_3_7")
	V3_3_8_SNAPSHOT       = ClientVersion("V3_3_8_SNAPSHOT")
	V3_3_8                = ClientVersion("V3_3_8")
	V3_3_9_SNAPSHOT       = ClientVersion("V3_3_9_SNAPSHOT")
	V3_3_9                = ClientVersion("V3_3_9")
	V3_4_1_SNAPSHOT       = ClientVersion("V3_4_1_SNAPSHOT")
	V3_4_1                = ClientVersion("V3_4_1")
	V3_4_2_SNAPSHOT       = ClientVersion("V3_4_2_SNAPSHOT")
	V3_4_2                = ClientVersion("V3_4_2")
	V3_4_3_SNAPSHOT       = ClientVersion("V3_4_3_SNAPSHOT")
	V3_4_3                = ClientVersion("V3_4_3")
	V3_4_4_SNAPSHOT       = ClientVersion("V3_4_4_SNAPSHOT")
	V3_4_4                = ClientVersion("V3_4_4")
	V3_4_5_SNAPSHOT       = ClientVersion("V3_4_5_SNAPSHOT")
	V3_4_5                = ClientVersion("V3_4_5")
	V3_4_6_SNAPSHOT       = ClientVersion("V3_4_6_SNAPSHOT")
	V3_4_6                = ClientVersion("V3_4_6")
	V3_4_7_SNAPSHOT       = ClientVersion("V3_4_7_SNAPSHOT")
	V3_4_7                = ClientVersion("V3_4_7")
	V3_4_8_SNAPSHOT       = ClientVersion("V3_4_8_SNAPSHOT")
	V3_4_8                = ClientVersion("V3_4_8")
	V3_4_9_SNAPSHOT       = ClientVersion("V3_4_9_SNAPSHOT")
	V3_4_9                = ClientVersion("V3_4_9")
	V3_5_1_SNAPSHOT       = ClientVersion("V3_5_1_SNAPSHOT")
	V3_5_1                = ClientVersion("V3_5_1")
	V3_5_2_SNAPSHOT       = ClientVersion("V3_5_2_SNAPSHOT")
	V3_5_2                = ClientVersion("V3_5_2")
	V3_5_3_SNAPSHOT       = ClientVersion("V3_5_3_SNAPSHOT")
	V3_5_3                = ClientVersion("V3_5_3")
	V3_5_4_SNAPSHOT       = ClientVersion("V3_5_4_SNAPSHOT")
	V3_5_4                = ClientVersion("V3_5_4")
	V3_5_5_SNAPSHOT       = ClientVersion("V3_5_5_SNAPSHOT")
	V3_5_5                = ClientVersion("V3_5_5")
	V3_5_6_SNAPSHOT       = ClientVersion("V3_5_6_SNAPSHOT")
	V3_5_6                = ClientVersion("V3_5_6")
	V3_5_7_SNAPSHOT       = ClientVersion("V3_5_7_SNAPSHOT")
	V3_5_7                = ClientVersion("V3_5_7")
	V3_5_8_SNAPSHOT       = ClientVersion("V3_5_8_SNAPSHOT")
	V3_5_8                = ClientVersion("V3_5_8")
	V3_5_9_SNAPSHOT       = ClientVersion("V3_5_9_SNAPSHOT")
	V3_5_9                = ClientVersion("V3_5_9")
	V3_6_1_SNAPSHOT       = ClientVersion("V3_6_1_SNAPSHOT")
	V3_6_1                = ClientVersion("V3_6_1")
	V3_6_2_SNAPSHOT       = ClientVersion("V3_6_2_SNAPSHOT")
	V3_6_2                = ClientVersion("V3_6_2")
	V3_6_3_SNAPSHOT       = ClientVersion("V3_6_3_SNAPSHOT")
	V3_6_3                = ClientVersion("V3_6_3")
	V3_6_4_SNAPSHOT       = ClientVersion("V3_6_4_SNAPSHOT")
	V3_6_4                = ClientVersion("V3_6_4")
	V3_6_5_SNAPSHOT       = ClientVersion("V3_6_5_SNAPSHOT")
	V3_6_5                = ClientVersion("V3_6_5")
	V3_6_6_SNAPSHOT       = ClientVersion("V3_6_6_SNAPSHOT")
	V3_6_6                = ClientVersion("V3_6_6")
	V3_6_7_SNAPSHOT       = ClientVersion("V3_6_7_SNAPSHOT")
	V3_6_7                = ClientVersion("V3_6_7")
	V3_6_8_SNAPSHOT       = ClientVersion("V3_6_8_SNAPSHOT")
	V3_6_8                = ClientVersion("V3_6_8")
	V3_6_9_SNAPSHOT       = ClientVersion("V3_6_9_SNAPSHOT")
	V3_6_9                = ClientVersion("V3_6_9")
	V3_7_1_SNAPSHOT       = ClientVersion("V3_7_1_SNAPSHOT")
	V3_7_1                = ClientVersion("V3_7_1")
	V3_7_2_SNAPSHOT       = ClientVersion("V3_7_2_SNAPSHOT")
	V3_7_2                = ClientVersion("V3_7_2")
	V3_7_3_SNAPSHOT       = ClientVersion("V3_7_3_SNAPSHOT")
	V3_7_3                = ClientVersion("V3_7_3")
	V3_7_4_SNAPSHOT       = ClientVersion("V3_7_4_SNAPSHOT")
	V3_7_4                = ClientVersion("V3_7_4")
	V3_7_5_SNAPSHOT       = ClientVersion("V3_7_5_SNAPSHOT")
	V3_7_5                = ClientVersion("V3_7_5")
	V3_7_6_SNAPSHOT       = ClientVersion("V3_7_6_SNAPSHOT")
	V3_7_6                = ClientVersion("V3_7_6")
	V3_7_7_SNAPSHOT       = ClientVersion("V3_7_7_SNAPSHOT")
	V3_7_7                = ClientVersion("V3_7_7")
	V3_7_8_SNAPSHOT       = ClientVersion("V3_7_8_SNAPSHOT")
	V3_7_8                = ClientVersion("V3_7_8")
	V3_7_9_SNAPSHOT       = ClientVersion("V3_7_9_SNAPSHOT")
	V3_7_9                = ClientVersion("V3_7_9")
	V3_8_1_SNAPSHOT       = ClientVersion("V3_8_1_SNAPSHOT")
	V3_8_1                = ClientVersion("V3_8_1")
	V3_8_2_SNAPSHOT       = ClientVersion("V3_8_2_SNAPSHOT")
	V3_8_2                = ClientVersion("V3_8_2")
	V3_8_3_SNAPSHOT       = ClientVersion("V3_8_3_SNAPSHOT")
	V3_8_3                = ClientVersion("V3_8_3")
	V3_8_4_SNAPSHOT       = ClientVersion("V3_8_4_SNAPSHOT")
	V3_8_4                = ClientVersion("V3_8_4")
	V3_8_5_SNAPSHOT       = ClientVersion("V3_8_5_SNAPSHOT")
	V3_8_5                = ClientVersion("V3_8_5")
	V3_8_6_SNAPSHOT       = ClientVersion("V3_8_6_SNAPSHOT")
	V3_8_6                = ClientVersion("V3_8_6")
	V3_8_7_SNAPSHOT       = ClientVersion("V3_8_7_SNAPSHOT")
	V3_8_7                = ClientVersion("V3_8_7")
	V3_8_8_SNAPSHOT       = ClientVersion("V3_8_8_SNAPSHOT")
	V3_8_8                = ClientVersion("V3_8_8")
	V3_8_9_SNAPSHOT       = ClientVersion("V3_8_9_SNAPSHOT")
	V3_8_9                = ClientVersion("V3_8_9")
	V3_9_1_SNAPSHOT       = ClientVersion("V3_9_1_SNAPSHOT")
	V3_9_1                = ClientVersion("V3_9_1")
	V3_9_2_SNAPSHOT       = ClientVersion("V3_9_2_SNAPSHOT")
	V3_9_2                = ClientVersion("V3_9_2")
	V3_9_3_SNAPSHOT       = ClientVersion("V3_9_3_SNAPSHOT")
	V3_9_3                = ClientVersion("V3_9_3")
	V3_9_4_SNAPSHOT       = ClientVersion("V3_9_4_SNAPSHOT")
	V3_9_4                = ClientVersion("V3_9_4")
	V3_9_5_SNAPSHOT       = ClientVersion("V3_9_5_SNAPSHOT")
	V3_9_5                = ClientVersion("V3_9_5")
	V3_9_6_SNAPSHOT       = ClientVersion("V3_9_6_SNAPSHOT")
	V3_9_6                = ClientVersion("V3_9_6")
	V3_9_7_SNAPSHOT       = ClientVersion("V3_9_7_SNAPSHOT")
	V3_9_7                = ClientVersion("V3_9_7")
	V3_9_8_SNAPSHOT       = ClientVersion("V3_9_8_SNAPSHOT")
	V3_9_8                = ClientVersion("V3_9_8")
	V3_9_9_SNAPSHOT       = ClientVersion("V3_9_9_SNAPSHOT")
	V3_9_9                = ClientVersion("V3_9_9")
	V4_0_0_SNAPSHOT       = ClientVersion("V4_0_0_SNAPSHOT")
	V4_0_0                = ClientVersion("V4_0_0")
	V4_1_0_SNAPSHOT       = ClientVersion("V4_1_0_SNAPSHOT")
	V4_1_0                = ClientVersion("V4_1_0")
	V4_2_0_SNAPSHOT       = ClientVersion("V4_2_0_SNAPSHOT")
	V4_2_0                = ClientVersion("V4_2_0")
	V4_3_0_SNAPSHOT       = ClientVersion("V4_3_0_SNAPSHOT")
	V4_3_0                = ClientVersion("V4_3_0")
	V4_4_0_SNAPSHOT       = ClientVersion("V4_4_0_SNAPSHOT")
	V4_4_0                = ClientVersion("V4_4_0")
	V4_5_0_SNAPSHOT       = ClientVersion("V4_5_0_SNAPSHOT")
	V4_6_0                = ClientVersion("V4_6_0")
	V4_7_0                = ClientVersion("V4_7_0")
	V4_8_0                = ClientVersion("V4_8_0")
	V4_9_0                = ClientVersion("V4_9_0")
	V4_9_1                = ClientVersion("V4_9_1")
	V4_9_2                = ClientVersion("V4_9_2")
	V4_9_3                = ClientVersion("V4_9_3")
	V4_9_4                = ClientVersion("V4_9_4")
	V4_9_5                = ClientVersion("V4_9_5")
	V4_9_6                = ClientVersion("V4_9_6")
	V4_9_7                = ClientVersion("V4_9_7")
)