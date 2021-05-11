// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// SetRoomConfig 设置一个范围内的会议室配置。
//
// 根据设置范围传入对应的参数
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_config/set
func (r *VCAPI) SetRoomConfig(ctx context.Context, request *SetRoomConfigReq) (*SetRoomConfigResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/room_configs/set",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(setRoomConfigResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "SetRoomConfig", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type SetRoomConfigReq struct {
	Scope      int                         `json:"scope,omitempty"`       // 设置节点范围, 示例值：5, 可选值有: `1`：租户, `2`：国家/地区, `3`：城市, `4`：建筑, `5`：楼层, `6`：会议室
	CountryID  *string                     `json:"country_id,omitempty"`  // 国家/地区ID scope为1，2时需要此参数, 示例值："1"
	DistrictID *string                     `json:"district_id,omitempty"` // 城市ID scope为2时需要此参数, 示例值："2"
	BuildingID *string                     `json:"building_id,omitempty"` // 建筑ID scope为3，4时需要此参数, 示例值："3"
	FloorName  *string                     `json:"floor_name,omitempty"`  // 楼层 scope为4时需要此参数, 示例值："4"
	RoomID     *string                     `json:"room_id,omitempty"`     // 会议室ID scope为5时需要此参数, 示例值："67687262867363"
	RoomConfig *SetRoomConfigReqRoomConfig `json:"room_config,omitempty"` // 会议室设置
}

type SetRoomConfigReqRoomConfig struct {
	RoomBackground    *string                                   `json:"room_background,omitempty"`    // 飞书会议室背景图, 示例值："https://lf1-ttcdn-tos.pstatp.com/obj/xxx"
	DisplayBackground *string                                   `json:"display_background,omitempty"` // 飞书签到板背景图, 示例值："https://lf1-ttcdn-tos.pstatp.com/obj/xxx"
	DigitalSignage    *SetRoomConfigReqRoomConfigDigitalSignage `json:"digital_signage,omitempty"`    // 飞书会议室数字标牌
}

type SetRoomConfigReqRoomConfigDigitalSignage struct {
	Enable       *bool                                               `json:"enable,omitempty"`        // 是否开启数字标牌功能, 示例值：true
	Mute         *bool                                               `json:"mute,omitempty"`          // 是否静音播放, 示例值：true
	StartDisplay *int                                                `json:"start_display,omitempty"` // 日程会议开始前n分钟结束播放, 示例值：3
	StopDisplay  *int                                                `json:"stop_display,omitempty"`  // 会议结束后n分钟开始播放, 示例值：3
	Materials    []*SetRoomConfigReqRoomConfigDigitalSignageMaterial `json:"materials,omitempty"`     // 素材列表
}

type SetRoomConfigReqRoomConfigDigitalSignageMaterial struct {
	ID           *string `json:"id,omitempty"`            // 素材ID, 示例值："7847784676276"
	Name         *string `json:"name,omitempty"`          // 素材名称, 示例值："name"
	MaterialType *int    `json:"material_type,omitempty"` // 素材类型, 示例值：0, 可选值有: `1`：图片, `2`：视频, `3`：GIF
	URL          *string `json:"url,omitempty"`           // 素材url, 示例值："url"
	Duration     *int    `json:"duration,omitempty"`      // 播放时长（单位sec）, 示例值：15
	Cover        *string `json:"cover,omitempty"`         // 素材封面url, 示例值："url"
	Md5          *string `json:"md5,omitempty"`           // 素材文件md5, 示例值："md5"
}

type setRoomConfigResp struct {
	Code int                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *SetRoomConfigResp `json:"data,omitempty"`
}

type SetRoomConfigResp struct{}
