//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UHost ResetUHostInstancePassword

package uhost

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// ResetUHostInstancePasswordRequest is request schema for ResetUHostInstancePassword action
type ResetUHostInstancePasswordRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// UHost实例ID
	UHostId *string `required:"true"`

	// UHost新密码（密码格式使用BASE64编码）
	Password *string `required:"true"`
}

// ResetUHostInstancePasswordResponse is response schema for ResetUHostInstancePassword action
type ResetUHostInstancePasswordResponse struct {
	response.CommonBase

	// UHost实例ID
	UhostId string
}

// NewResetUHostInstancePasswordRequest will create request of ResetUHostInstancePassword action.
func (c *UHostClient) NewResetUHostInstancePasswordRequest() *ResetUHostInstancePasswordRequest {
	req := &ResetUHostInstancePasswordRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ResetUHostInstancePassword - 重置UHost实例的管理员密码。
func (c *UHostClient) ResetUHostInstancePassword(req *ResetUHostInstancePasswordRequest) (*ResetUHostInstancePasswordResponse, error) {
	var err error
	var res ResetUHostInstancePasswordResponse

	err = c.client.InvokeAction("ResetUHostInstancePassword", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}