package model

type UserInfo struct {
	// 错误信息，一般正常情况就是个“ok”
	Error string `json:"error"`

	// 登录时间
	AddTime int `json:"add_time"`

	// 产品策略
	BillingName string `json:"billing_name"`

	// 实际上是当前时间（GMT时间，不是东八区时间）
	KeepaliveTime int `json:"keepalive_time"`

	// 出入流量
	BytesIn  int `json:"bytes_in"`
	BytesOut int `json:"bytes_out"`

	// 总的在线设备和当前ip
	OnlineDeviceTotal string `json:"online_device_total"`
	OnlineIp          string `json:"online_ip"`

	// 产品名
	ProductsName string `json:"products_name"`

	// 总流量和时长
	SumBytes   int64 `json:"sum_bytes"`
	SumSeconds int   `json:"sum_seconds"`

	// 用户组
	GroupId string `json:"group_id"`

	// 用户MAC地址和用户名
	UserMac  string `json:"user_mac"`
	UserName string `json:"user_name"`

	// 下面的字段是未登录的时候才会有的字段
	// 当前设备ip，ClientIp和OnlineIp两个其实是一样的，只取一个就好了
	ClientIp string `json:"client_ip"`
	SrunVer  string `json:"srun_ver"`

	// 时间戳
	St int `json:"st"`
}
