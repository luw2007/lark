// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
	"io"
)

// CropFaceVerifyImage
//
// ::: note
// 无源人脸比对接入需申请白名单，接入前请联系飞书开放平台工作人员，邮箱：openplatform@bytedance.com。
// :::
// 无源人脸比对流程，开发者后台通过调用此接口对基准图片做规范校验及处理。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/face/facial-image-cropping
func (r *HumanAuthAPI) CropFaceVerifyImage(ctx context.Context, request *CropFaceVerifyImageReq) (*CropFaceVerifyImageResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/face_verify/v1/crop_face_image",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                true,
	}
	resp := new(cropFaceVerifyImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("HumanAuth", "CropFaceVerifyImage", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type CropFaceVerifyImageReq struct {
	RawImage io.Reader `json:"raw_image,omitempty"` // 带有头像的人脸照片文件名称
}

type cropFaceVerifyImageResp struct {
	Code int                      `json:"code,omitempty"` // 返回码，非0为失败
	Msg  string                   `json:"msg,omitempty"`  // 返回信息，返回码的描述
	Data *CropFaceVerifyImageResp `json:"data,omitempty"` // 业务数据
}

type CropFaceVerifyImageResp struct {
	FaceImage string `json:"face_image,omitempty"` // BASE64(裁剪后的人脸基准图片)，code为0时返回
}
