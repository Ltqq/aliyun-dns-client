package handler

import (
	dns "github.com/alibabacloud-go/alidns-20150109/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type OpenapiHandler interface {
	DoRequest(request interface{}) (_err error)
}

type UpdateDomainHandler struct {
	client *dns.Client
}

func NewUpdateDomainHandler(client *dns.Client) *UpdateDomainHandler {
	return &UpdateDomainHandler{client: client}
}

func (u *UpdateDomainHandler) DoRequest(request interface{}) (_err error) {
	var req = request.(*dns.UpdateDomainRecordRequest)
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := u.client.UpdateDomainRecord(req)
		if _err != nil {
			return _err
		}

		console.Log(tea.String("-------------------修改解析记录--------------------"))
		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}
