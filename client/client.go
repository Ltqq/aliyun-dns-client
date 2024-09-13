package client

import (
	"aliyun-dns-client/config"
	dns "github.com/alibabacloud-go/alidns-20150109/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"sync"
)

var Client *dns.Client
var once = sync.Once{}

func InitClient(cfg *config.OpenapiConfig) error {
	var initErr error
	once.Do(func() {
		c := &openapi.Config{
			AccessKeyId:     &cfg.AK,       // AccessKey ID
			AccessKeySecret: &cfg.SK,       // AccessKey Secret
			RegionId:        &cfg.RegionID, // 可用区ID
		}
		_result, err := dns.NewClient(c)
		if err != nil {
			initErr = err // 保存错误并在函数外返回
			return
		}
		Client = _result
	})

	return initErr
}
