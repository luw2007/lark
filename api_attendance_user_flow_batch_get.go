// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// BatchGetUserFlow
//
// 批量查询授权内员工的实际打卡记录。例如，企业给一个员工设定的班次是上午 9 点和下午 6 点各打一次上下班卡，但是员工在这期间打了多次卡，该接口会把所有的打卡记录都返回。
// 如果只需获取打卡结果，而不需要打卡的详细数据，可使用“获取打卡结果”的接口。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//BatchQueryCheckinFlowHistory
func (r *AttendanceAPI) BatchGetUserFlow(ctx context.Context, request *BatchGetUserFlowReq) (*BatchGetUserFlowResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_flows/query",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(batchGetUserFlowResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Attendance", "BatchGetUserFlow", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type BatchGetUserFlowReq struct {
	EmployeeType  EmployeeType `query:"employee_type" json:"-"`   // 请求体中的 user_ids 的员工工号类型，可用值：【employee_id（员工的 employeeId），employee_no（员工工号）】，示例值：“employee_id”
	UserIDs       []string     `json:"user_ids,omitempty"`        // employee_no 或 employee_id 列表
	CheckTimeFrom string       `json:"check_time_from,omitempty"` // 查询的起始时间，时间戳
	CheckTimeTo   string       `json:"check_time_to,omitempty"`   // 查询的结束时间，时间戳
}

type batchGetUserFlowResp struct {
	Code int                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *BatchGetUserFlowResp `json:"data,omitempty"` // -
}

type BatchGetUserFlowResp struct {
	UserFlowResults []*BatchGetUserFlowRespUserFlowResult `json:"user_flow_results,omitempty"` // 打卡记录列表
}

type BatchGetUserFlowRespUserFlowResult struct {
	UserID       string   `json:"user_id,omitempty"`       // 员工工号
	CreatorID    string   `json:"creator_id,omitempty"`    // 打卡记录创建者的 employee_no
	LocationName string   `json:"location_name,omitempty"` // 打卡位置名称信息
	CheckTime    string   `json:"check_time,omitempty"`    // 打卡时间，精确到秒的时间戳
	Comment      string   `json:"comment,omitempty"`       // 打卡备注
	RecordID     string   `json:"record_id,omitempty"`     // 打卡记录 ID
	Longitude    float64  `json:"longitude,omitempty"`     // 打卡经度
	Latitude     float64  `json:"latitude,omitempty"`      // 打卡纬度
	Ssid         string   `json:"ssid,omitempty"`          // 打卡 Wi-Fi 的 SSID
	Bssid        string   `json:"bssid,omitempty"`         // 打卡 Wi-Fi 的 MAC 地址
	IsField      bool     `json:"is_field,omitempty"`      // 是否为外勤打卡
	IsWifi       bool     `json:"is_wifi,omitempty"`       // 是否为 Wi-Fi 打卡
	Type         int      `json:"type,omitempty"`          // 记录生成方式，可用值：【0（用户自己打卡），1（管理员修改），2（用户补卡），3（系统自动生成），4（下班免打卡），5（考勤机打卡），6（极速打卡），7（考勤开放平台导入）】
	PhotoUrls    []string `json:"photo_urls,omitempty"`    // 打卡照片列表
}
