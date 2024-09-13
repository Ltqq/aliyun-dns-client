// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"aliyun-dns-client/client"
	"aliyun-dns-client/config"
	"aliyun-dns-client/handler"
	dns "github.com/alibabacloud-go/alidns-20150109/v2/client"
)

func main() {
	var cfg = &config.Config{
		Openapi: &config.OpenapiConfig{
			RegionID: "",
			AK:       "",
			SK:       "",
		},
		Record: &config.RecordConfig{
			RR:               "ddns",
			Type:             "AAAA",
			TTL:              600,
			Priority:         1,
			DefaultRecordID:  "917264544980150272",
			DefaultInterface: "eth0",
		},
	}
	err := client.InitClient(cfg.Openapi)
	if err != nil {
		panic(err)
	}

	iPv6 := handler.NewIPv6(cfg.Record.DefaultInterface)
	updateDomainHandler := handler.NewUpdateDomainHandler(client.Client)
	err = RecordUpdate(iPv6, updateDomainHandler, cfg.Record)
	if err != nil {
		panic(err)
	}

}

func RecordUpdate(handler handler.HostIPHandler, openapiHandler handler.OpenapiHandler, cfg *config.RecordConfig) error {
	hostIP, err := handler.GetHostIP()
	if err != nil {
		return err
	}
	err = openapiHandler.DoRequest(&dns.UpdateDomainRecordRequest{
		RecordId: &cfg.DefaultRecordID,
		RR:       &cfg.RR,
		Type:     &cfg.Type,
		Value:    &hostIP,
		TTL:      &cfg.TTL,
		Priority: &cfg.Priority,
	})
	if err != nil {
		return err
	}
	return nil
}
