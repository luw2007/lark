package lark

import (
	"context"
)

// CreateCalendar
//
// 该接口用于为当前身份（应用 / 用户）创建一个共享日历。
// 身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar&method=create)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/create
func (r *CalendarAPI) CreateCalendar(ctx context.Context, request *CreateCalendarReq) (*CreateCalendarResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(createCalendarResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "CreateCalendar", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type CreateCalendarReq struct {
	Summary      *string `json:"summary,omitempty"`       // 日历标题,**示例值**："测试日历",**数据校验规则**：,- 最大长度：`255` 字符
	Description  *string `json:"description,omitempty"`   // 日历描述,**示例值**："使用开放接口创建日历",**数据校验规则**：,- 最大长度：`255` 字符
	Permissions  *string `json:"permissions,omitempty"`   // 日历公开范围,**示例值**："private",**可选值有**：,- `private`：私密,- `show_only_free_busy`：仅展示忙闲信息,- `public`：他人可查看日程详情
	Color        *int    `json:"color,omitempty"`         // 日历颜色，颜色RGB值的int32表示。客户端展示时会映射到色板上最接近的一种颜色。仅对当前身份生效,**示例值**：-1
	SummaryAlias *string `json:"summary_alias,omitempty"` // 日历备注名，修改或添加后仅对当前身份生效,**示例值**："日历备注名",**数据校验规则**：,- 最大长度：`255` 字符
}

type createCalendarResp struct {
	Code int                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *CreateCalendarResp `json:"data,omitempty"` // \-
}

type CreateCalendarResp struct {
	Calendar *CreateCalendarRespCalendar `json:"calendar,omitempty"` // 新创建的日历实体
}

type CreateCalendarRespCalendar struct {
	CalendarID   string       `json:"calendar_id,omitempty"`    // 日历ID
	Summary      string       `json:"summary,omitempty"`        // 日历标题,**数据校验规则**：,- 最大长度：`255` 字符
	Description  string       `json:"description,omitempty"`    // 日历描述,**数据校验规则**：,- 最大长度：`255` 字符
	Permissions  string       `json:"permissions,omitempty"`    // 日历公开范围,**可选值有**：,- `private`：私密,- `show_only_free_busy`：仅展示忙闲信息,- `public`：他人可查看日程详情
	Color        int          `json:"color,omitempty"`          // 日历颜色，颜色RGB值的int32表示。客户端展示时会映射到色板上最接近的一种颜色。仅对当前身份生效
	Type         string       `json:"type,omitempty"`           // 日历类型,**可选值有**：,- `unknown`：未知类型,- `primary`：用户或应用的主日历,- `shared`：由用户或应用创建的共享日历,- `google`：用户绑定的谷歌日历,- `resource`：会议室日历,- `exchange`：用户绑定的Exchange日历
	SummaryAlias string       `json:"summary_alias,omitempty"`  // 日历备注名，修改或添加后仅对当前身份生效,**数据校验规则**：,- 最大长度：`255` 字符
	IsDeleted    bool         `json:"is_deleted,omitempty"`     // 对于当前身份，日历是否已经被标记为删除,**默认值**：`false`
	IsThirdParty bool         `json:"is_third_party,omitempty"` // 当前日历是否是第三方数据；三方日历及日程只支持读，不支持写入,**默认值**：`false`
	Role         CalendarRole `json:"role,omitempty"`           // 当前身份对于该日历的访问权限,**可选值有**：,- `unknown`：未知权限,- `free_busy_reader`：游客，只能看到忙碌/空闲信息,- `reader`：订阅者，查看所有日程详情,- `writer`：编辑者，创建及修改日程,- `owner`：管理员，管理日历及共享设置
}
