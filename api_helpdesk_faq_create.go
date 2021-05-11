// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateFAQ 该接口用于创建知识库。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/create
func (r *HelpdeskAPI) CreateFAQ(ctx context.Context, request *CreateFAQReq) (*CreateFAQResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/faqs",
		Body:                  request,
		NeedTenantAccessToken: false,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      true,
		IsFile:                false,
	}
	resp := new(createFAQResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Helpdesk", "CreateFAQ", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type CreateFAQReq struct {
	Faq *CreateFAQReqFaq `json:"faq,omitempty"` // 知识库详情
}

type CreateFAQReqFaq struct {
	CategoryID     *string  `json:"category_id,omitempty"`     // 知识库分类ID, 示例值："6836004780707807251"
	Question       string   `json:"question,omitempty"`        // 问题, 示例值："问题"
	Answer         *string  `json:"answer,omitempty"`          // 答案, 示例值："答案"
	AnswerRichtext *string  `json:"answer_richtext,omitempty"` // 富文本答案和答案必须有一个必填。Json Array格式，富文本结构请见[了解更多: 富文本](/ssl:ttdoc/ukTMukTMukTM/uITM0YjLyEDN24iMxQjN), 示例值："[{,                        "content": "这只是一个测试，医保问题",,                        "type": "text",                    }]"
	Tags           []string `json:"tags,omitempty"`            // 关联词
}

type createFAQResp struct {
	Code int            `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string         `json:"msg,omitempty"`  // 错误描述
	Data *CreateFAQResp `json:"data,omitempty"` //
}

type CreateFAQResp struct {
	Faq *CreateFAQRespFaq `json:"faq,omitempty"` // 知识库详情
}

type CreateFAQRespFaq struct {
	FaqID          string                      `json:"faq_id,omitempty"`          // 知识库ID
	ID             string                      `json:"id,omitempty"`              // 知识库旧版ID，请使用faq_id
	HelpdeskID     string                      `json:"helpdesk_id,omitempty"`     // 服务台ID
	Question       string                      `json:"question,omitempty"`        // 问题
	Answer         string                      `json:"answer,omitempty"`          // 答案
	AnswerRichtext string                      `json:"answer_richtext,omitempty"` // 富文本答案
	CreateTime     int                         `json:"create_time,omitempty"`     // 创建时间
	UpdateTime     int                         `json:"update_time,omitempty"`     // 修改时间
	Categories     []*HelpdeskCategory         `json:"categories,omitempty"`      // 分类
	Tags           []string                    `json:"tags,omitempty"`            // 关联词列表
	ExpireTime     int                         `json:"expire_time,omitempty"`     // 失效时间
	UpdateUser     *CreateFAQRespFaqUpdateUser `json:"update_user,omitempty"`     // 更新用户
	CreateUser     *CreateFAQRespFaqCreateUser `json:"create_user,omitempty"`     // 创建用户
}

type CreateFAQRespFaqUpdateUser struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarUrl string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
}

type CreateFAQRespFaqCreateUser struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarUrl string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
}
