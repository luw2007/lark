// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// BatchCreateRecord 该接口用于在数据表中新增多条记录
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_create
func (r *BitableService) BatchCreateRecord(ctx context.Context, request *BatchCreateRecordReq, options ...MethodOptionFunc) (*BatchCreateRecordResp, *Response, error) {
	if r.cli.mock.mockBitableBatchCreateRecord != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#BatchCreateRecord mock enable")
		return r.cli.mock.mockBitableBatchCreateRecord(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Bitable#BatchCreateRecord call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#BatchCreateRecord request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:       "POST",
		URL:          "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_create",
		Body:         request,
		MethodOption: newMethodOption(options),

		NeedUserAccessToken: true,
	}
	resp := new(batchCreateRecordResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Bitable#BatchCreateRecord POST https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_create failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Bitable#BatchCreateRecord POST https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_create failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Bitable", "BatchCreateRecord", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#BatchCreateRecord success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockBitableBatchCreateRecord(f func(ctx context.Context, request *BatchCreateRecordReq, options ...MethodOptionFunc) (*BatchCreateRecordResp, *Response, error)) {
	r.mockBitableBatchCreateRecord = f
}

func (r *Mock) UnMockBitableBatchCreateRecord() {
	r.mockBitableBatchCreateRecord = nil
}

type BatchCreateRecordReq struct {
	UserIDType *IDType                       `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	AppToken   string                        `path:"app_token" json:"-"`     // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
	TableID    string                        `path:"table_id" json:"-"`      // table id, 示例值："tblsRc9GRRXKqhvW"
	Records    []*BatchCreateRecordReqRecord `json:"records,omitempty"`      // 记录
}

type BatchCreateRecordReqRecord struct {
	RecordID *string                `json:"record_id,omitempty"` // 记录 id, 示例值："recqwIwhc6"
	Fields   map[string]interface{} `json:"fields,omitempty"`    // 记录字段
}

type batchCreateRecordResp struct {
	Code int                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *BatchCreateRecordResp `json:"data,omitempty"` //
}

type BatchCreateRecordResp struct {
	Records []*BatchCreateRecordRespRecord `json:"records,omitempty"` // 记录
}

type BatchCreateRecordRespRecord struct {
	RecordID string                 `json:"record_id,omitempty"` // 记录 id
	Fields   map[string]interface{} `json:"fields,omitempty"`    // 记录字段
}
