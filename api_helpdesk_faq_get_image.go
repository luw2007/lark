// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
	"io"
)

// GetFAQImage 该接口用于获取知识库图像。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/faq_image
func (r *HelpdeskAPI) GetFAQImage(ctx context.Context, request *GetFAQImageReq) (*GetFAQImageResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/faqs/:id/image/:image_key",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      true,
		IsFile:                false,
	}
	resp := new(getFAQImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Helpdesk", "GetFAQImage", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetFAQImageReq struct {
	ID       string `path:"id" json:"-"`        // 知识库ID, 示例值："12345"
	ImageKey string `path:"image_key" json:"-"` // 图像key, 示例值："img_b07ffac0-19c1-48a3-afca-599f8ea825fj"
}

type getFAQImageResp struct {
	IsFile bool             `json:"is_file,omitempty"`
	Code   int              `json:"code,omitempty"`
	Msg    string           `json:"msg,omitempty"`
	Data   *GetFAQImageResp `json:"data,omitempty"`
}

func (r *getFAQImageResp) IsFileType() bool {
	return r.IsFile
}

func (r *getFAQImageResp) SetFile(file io.Reader) {
	r.Data = &GetFAQImageResp{
		File: file,
	}
}

type GetFAQImageResp struct {
	File io.Reader `json:"file,omitempty"`
}
