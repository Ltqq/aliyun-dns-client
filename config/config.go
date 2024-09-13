package config

type Config struct {
	Openapi *OpenapiConfig
	Record  *RecordConfig
}

type OpenapiConfig struct {
	RegionID string
	AK       string
	SK       string
}

type RecordConfig struct {
	RR               string // 主机记录 ddns
	Type             string // 解析记录类型 AAAA(ipv6)
	TTL              int64  // default 600s
	Priority         int64  // mx记录优先级 default 1
	DefaultRecordID  string // 默认记录id 控制台查看
	DefaultInterface string // 默认网卡
}
