// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// SetPermissionMeetingRecording 将一个会议的录制文件授权给组织、用户或公开到公网
//
// 会议结束后并且收到了"录制完成"的事件方可进行授权；会议owner（通过开放平台预约的会议即为预约人）才有权限操作
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/set_permission
func (r *VCAPI) SetPermissionMeetingRecording(ctx context.Context, request *SetPermissionMeetingRecordingReq) (*SetPermissionMeetingRecordingResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "PATCH",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording/set_permission",
		Body:                  request,
		NeedTenantAccessToken: false,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(setPermissionMeetingRecordingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "SetPermissionMeetingRecording", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type SetPermissionMeetingRecordingReq struct {
	UserIDType        *IDType                                             `query:"user_id_type" json:"-"`       // 用户ID类型, 示例值: "open_id", 可选值有: `open_id`：用户的open id, `union_id`：用户的union id, `user_id`：用户的user id, 默认值: `open_id`, 当值为 `user_id`，字段权限要求: 获取用户 userid
	MeetingID         string                                              `path:"meeting_id" json:"-"`          // 会议ID, 示例值: "6911188411932033028"
	Token             *string                                             `json:"token,omitempty"`              // 录指文件的token, 示例值: "obcn37dxcftoc3656rgyejm7"
	PermissionObjects []*SetPermissionMeetingRecordingReqPermissionObject `json:"permission_objects,omitempty"` // 授权对象列表
}

type SetPermissionMeetingRecordingReqPermissionObject struct {
	ID             *string `json:"id,omitempty"`              // 授权对象ID, 示例值: "ou_3ec3f6a28a0d08c45d895276e8e5e19b"
	PermissionType *int    `json:"permission_type,omitempty"` // 授权类型, 示例值: 0, 可选值有: `0`：用户授权, `1`：群组授权, `2`：租户内授权（id字段不填）, `3`：公网授权（id字段不填）
	Permission     *int    `json:"permission,omitempty"`      // 权限, 示例值: 1, 可选值有: `1`：查看
}

type setPermissionMeetingRecordingResp struct {
	Code int                                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                             `json:"msg,omitempty"`  // 错误描述
	Data *SetPermissionMeetingRecordingResp `json:"data,omitempty"`
}

type SetPermissionMeetingRecordingResp struct{}
