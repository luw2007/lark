// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateIdentity 该接口用于录入实名认证的身份信息，在唤起有源活体认证前，需要使用该接口进行实名认证。
//
// 实名认证接口会有计费管理，接入前请联系飞书开放平台工作人员，邮箱：openplatform@bytedance.com。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/identity/create
func (r *HumanAuthService) CreateIdentity(ctx context.Context, request *CreateIdentityReq, options ...MethodOptionFunc) (*CreateIdentityResp, *Response, error) {
	if r.cli.mock.mockHumanAuthCreateIdentity != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] HumanAuth#CreateIdentity mock enable")
		return r.cli.mock.mockHumanAuthCreateIdentity(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] HumanAuth#CreateIdentity call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] HumanAuth#CreateIdentity request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/human_authentication/v1/identities",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createIdentityResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] HumanAuth#CreateIdentity POST https://open.feishu.cn/open-apis/human_authentication/v1/identities failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] HumanAuth#CreateIdentity POST https://open.feishu.cn/open-apis/human_authentication/v1/identities failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("HumanAuth", "CreateIdentity", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] HumanAuth#CreateIdentity success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockHumanAuthCreateIdentity(f func(ctx context.Context, request *CreateIdentityReq, options ...MethodOptionFunc) (*CreateIdentityResp, *Response, error)) {
	r.mockHumanAuthCreateIdentity = f
}

func (r *Mock) UnMockHumanAuthCreateIdentity() {
	r.mockHumanAuthCreateIdentity = nil
}

type CreateIdentityReq struct {
	UserID       string  `query:"user_id" json:"-"`       // 用户的唯一标识, 示例值: "ou_2eb5483cb377daa5054bc6f86e2089a5"
	UserIDType   *IDType `query:"user_id_type" json:"-"`  // 用户ID类型, 示例值: "open_id", 可选值有: `open_id`：用户的open id, `union_id`：用户的union id, `user_id`：用户的user id, 默认值: `open_id`, 当值为 `user_id`，字段权限要求: 获取用户 userid
	IdentityName string  `json:"identity_name,omitempty"` // user identity name, 示例值: "张三"
	IdentityCode string  `json:"identity_code,omitempty"` // user identity code, 示例值: "4xxxxxxxx"
	Mobile       *string `json:"mobile,omitempty"`        // user mobile, 示例值: "13xxxxxxx"
}

type createIdentityResp struct {
	Code int                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *CreateIdentityResp `json:"data,omitempty"` //
}

type CreateIdentityResp struct {
	VerifyUid string `json:"verify_uid,omitempty"` // 用户绑定实名身份的uid
}
