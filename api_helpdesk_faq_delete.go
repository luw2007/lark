// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteFAQ 该接口用于删除知识库。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/delete
func (r *HelpdeskService) DeleteFAQ(ctx context.Context, request *DeleteFAQReq, options ...MethodOptionFunc) (*DeleteFAQResp, *Response, error) {
	if r.cli.mock.mockHelpdeskDeleteFAQ != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#DeleteFAQ mock enable")
		return r.cli.mock.mockHelpdeskDeleteFAQ(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Helpdesk#DeleteFAQ call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#DeleteFAQ request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:       "DELETE",
		URL:          "https://open.feishu.cn/open-apis/helpdesk/v1/faqs/:id",
		Body:         request,
		MethodOption: newMethodOption(options),

		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(deleteFAQResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Helpdesk#DeleteFAQ DELETE https://open.feishu.cn/open-apis/helpdesk/v1/faqs/:id failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Helpdesk#DeleteFAQ DELETE https://open.feishu.cn/open-apis/helpdesk/v1/faqs/:id failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Helpdesk", "DeleteFAQ", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#DeleteFAQ success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockHelpdeskDeleteFAQ(f func(ctx context.Context, request *DeleteFAQReq, options ...MethodOptionFunc) (*DeleteFAQResp, *Response, error)) {
	r.mockHelpdeskDeleteFAQ = f
}

func (r *Mock) UnMockHelpdeskDeleteFAQ() {
	r.mockHelpdeskDeleteFAQ = nil
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
