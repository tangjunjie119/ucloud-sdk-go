//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet UpdateFirewall

package unet

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UpdateFirewallRequest is request schema for UpdateFirewall action
type UpdateFirewallRequest struct {
	request.CommonBase

	// 防火墙资源ID
	FWId *string `required:"true"`

	// 防火墙规则，例如：TCP|22|192.168.1.1/22|DROP|LOW，第一个参数代表协议：第二个参数代表端口号，第三个参数为ip，第四个参数为ACCEPT（接受）和DROP（拒绝），第五个参数优先级：HIGH（高），MEDIUM（中），LOW（低）
	Rule []string `required:"true"`
}

// UpdateFirewallResponse is response schema for UpdateFirewall action
type UpdateFirewallResponse struct {
	response.CommonBase

	// 防火墙id
	FWId string
}

// NewUpdateFirewallRequest will create request of UpdateFirewall action.
func (c *UNetClient) NewUpdateFirewallRequest() *UpdateFirewallRequest {
	req := &UpdateFirewallRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// UpdateFirewall - 更新防火墙规则
func (c *UNetClient) UpdateFirewall(req *UpdateFirewallRequest) (*UpdateFirewallResponse, error) {
	var err error
	var res UpdateFirewallResponse

	err = c.client.InvokeAction("UpdateFirewall", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
