package lark

import (
	"context"
)

// SendRawMessage 给指定用户或者会话发送消息，支持文本、富文本、卡片、群名片、个人名片、图片、视频、音频、文件、表情包。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 给用户发送消息，需要机器人对用户有可见性
// - 给群组发送消息，需要机器人在群中
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/create
func (r *MessageAPI) SendRawMessage(ctx context.Context, request *SendRawMessageReq) (*SendRawMessageResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/messages",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(sendRawMessageResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Message", "SendRawMessage", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type SendRawMessageReq struct {
	ReceiveIDType *IDType `query:"receive_id_type" json:"-"` // 消息接收者id类型 open_id/user_id/union_id/email/chat_id,**示例值**："open_id",**可选值有**：,- `open_id`：以open_id来识别用户,- `user_id`：以user_id来识别用户,- `union_id`：以union_id来识别用户,- `email`：以email来识别用户,- `chat_id`：以chat_id来识别群聊
	ReceiveID     *string `json:"receive_id,omitempty"`      // 依据receive_id_type的值，填写对应的消息接收者id,**示例值**："ou_7d8a6e6df7621556ce0d21922b676706ccs"
	Content       string  `json:"content,omitempty"`         // 消息内容 json格式，格式说明参考: [发送消息content说明](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json),**示例值**："{\"text\":\"<at user_id=\\\"ou_155184d1e73cbfb8973e5a9e698e74f2\\\">Tom</at> test content\"}"
	MsgType       MsgType `json:"msg_type,omitempty"`        // 消息类型，包括：text、post、image、file、audio、media、sticker、interactive、share_chat、share_user,**示例值**："text"
}

type sendRawMessageResp struct {
	Code int                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *SendRawMessageResp `json:"data,omitempty"` // \-
}

type SendRawMessageResp struct {
	MessageID      string       `json:"message_id,omitempty"`       // 消息id open_message_id
	RootID         string       `json:"root_id,omitempty"`          // 根消息id open_message_id
	ParentID       string       `json:"parent_id,omitempty"`        // 父消息的id open_message_id
	MsgType        MsgType      `json:"msg_type,omitempty"`         // 消息类型 text post card image等等
	CreateTime     string       `json:"create_time,omitempty"`      // 消息生成的时间戳(毫秒)
	UpdateTime     string       `json:"update_time,omitempty"`      // 消息更新的时间戳
	Deleted        bool         `json:"deleted,omitempty"`          // 消息是否被撤回
	Updated        bool         `json:"updated,omitempty"`          // 消息是否被更新
	ChatID         string       `json:"chat_id,omitempty"`          // 所属的群
	Sender         *Sender      `json:"sender,omitempty"`           // 发送者，可以是用户或应用
	Body           *MessageBody `json:"body,omitempty"`             // 消息内容，json结构，格式说明参考： [消息content说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/events/message_content)
	Mentions       []*Mention   `json:"mentions,omitempty"`         // 被艾特的人或应用的id
	UpperMessageID string       `json:"upper_message_id,omitempty"` // 上一层级的消息id open_message_id
}
