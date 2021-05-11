// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetShiftByName
//
// 通过班次的名称获取班次信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_by_name
func (r *AttendanceAPI) GetShiftByName(ctx context.Context, request *GetShiftByNameReq) (*GetShiftByNameResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/shifts/query",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getShiftByNameResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Attendance", "GetShiftByName", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetShiftByNameReq struct {
	ShiftName string `query:"shift_name" json:"-"` // 班次名称，示例值："早班"
}

type getShiftByNameResp struct {
	Code int                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetShiftByNameResp `json:"data,omitempty"` // -
}

type GetShiftByNameResp struct {
	ShiftID           string                               `json:"shift_id,omitempty"`              // 班次 ID
	ShiftName         string                               `json:"shift_name,omitempty"`            // 班次名称
	PunchTimes        int                                  `json:"punch_times,omitempty"`           // 打卡次数
	IsFlexible        bool                                 `json:"is_flexible,omitempty"`           // 是否弹性打卡
	FlexibleMinutes   int                                  `json:"flexible_minutes,omitempty"`      // 弹性打卡时间
	NoNeedOff         bool                                 `json:"no_need_off,omitempty"`           // 是否下班免打卡
	PunchTimeRule     *GetShiftByNameRespPunchTimeRule     `json:"punch_time_rule,omitempty"`       // 打卡规则
	LateOffLateOnRule *GetShiftByNameRespLateOffLateOnRule `json:"late_off_late_on_rule,omitempty"` // 晚走晚到规则
	RestTimeRule      *GetShiftByNameRespRestTimeRule      `json:"rest_time_rule,omitempty"`        // 休息时间规则
}

type GetShiftByNameRespPunchTimeRule struct {
	OnTime              string `json:"on_time,omitempty"`                // 上班时间
	OffTime             string `json:"off_time,omitempty"`               // 下班时间
	LateMinutesAsLate   int    `json:"late_minutes_as_late,omitempty"`   // 晚到多长时间记为迟到
	LateMinutesAsLack   int    `json:"late_minutes_as_lack,omitempty"`   // 晚到多长时间记为缺卡
	OnAdvanceMinutes    int    `json:"on_advance_minutes,omitempty"`     // 最早可提前多长时间打上班卡
	EarlyMinutesAsEarly int    `json:"early_minutes_as_early,omitempty"` // 早走多长时间记为早退
	EarlyMinutesAsLack  int    `json:"early_minutes_as_lack,omitempty"`  // 早走多长时间记为缺卡
	OffDelayMinutes     int    `json:"off_delay_minutes,omitempty"`      // 最晚可延后多长时间打下班卡
}

type GetShiftByNameRespLateOffLateOnRule struct {
	LateOffMinutes int `json:"late_off_minutes,omitempty"` // 晚走多长时间
	LateOnMinutes  int `json:"late_on_minutes,omitempty"`  // 晚到多长时间
}

type GetShiftByNameRespRestTimeRule struct {
	RestBeginTime string `json:"rest_begin_time,omitempty"` // 休息开始时间
	RestEndTime   string `json:"rest_end_time,omitempty"`   // 休息结束时间
}
