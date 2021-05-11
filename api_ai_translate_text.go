// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// TranslateText 支持中日英（zh、ja、en）三语互译
//
// 单租户限流：20QPS，同租户下的应用没有限流，共享本租户的 20QPS 限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/translation-v1/text/translate
func (r *AIAPI) TranslateText(ctx context.Context, request *TranslateTextReq) (*TranslateTextResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/translation/v1/text/translate",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(translateTextResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("AI", "TranslateText", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type TranslateTextReq struct {
	SourceLanguage string                    `json:"source_language,omitempty"` // 源语言, 示例值："zh"
	Text           string                    `json:"text,omitempty"`            // 源文本, 示例值："尝试使用一下飞书吧"
	TargetLanguage string                    `json:"target_language,omitempty"` // 目标语言, 示例值："en"
	Glossary       *TranslateTextReqGlossary `json:"glossary,omitempty"`        // 请求级术语表，携带术语，仅在本次翻译中生效（最多能携带 128个术语词）
}

type TranslateTextReqGlossary struct {
	From string `json:"from,omitempty"` // 原文, 示例值："飞书"
	To   string `json:"to,omitempty"`   // 译文, 示例值："Lark"
}

type translateTextResp struct {
	Code int                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *TranslateTextResp `json:"data,omitempty"` //
}

type TranslateTextResp struct {
	Text string `json:"text,omitempty"` // 翻译后的文本
}
