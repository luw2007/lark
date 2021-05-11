// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteFAQ 该接口用于删除知识库。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/delete
func (r *HelpdeskAPI) DeleteFAQ(ctx context.Context, request *DeleteFAQReq) (*DeleteFAQResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/faqs/:id",
		Body:                  request,
		NeedTenantAccessToken: false,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      true,
		IsFile:                false,
	}
	resp := new(deleteFAQResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Helpdesk", "DeleteFAQ", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteFAQReq struct {
	ID string `path:"id" json:"-"` // id, 示例值："12345"
}

type deleteFAQResp struct {
	Code int            `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string         `json:"msg,omitempty"`  // 错误描述
	Data *DeleteFAQResp `json:"data,omitempty"`
}

type DeleteFAQResp struct{}
