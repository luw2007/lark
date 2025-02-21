// Code generated by lark_sdk_gen. DO NOT EDIT.

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Message_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Message

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendEphemeralMessage(ctx, &lark.SendEphemeralMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendUrgentAppMessage(ctx, &lark.SendUrgentAppMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendUrgentSmsMessage(ctx, &lark.SendUrgentSmsMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendUrgentPhoneMessage(ctx, &lark.SendUrgentPhoneMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendRawMessage(ctx, &lark.SendRawMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendRawMessageOld(ctx, &lark.SendRawMessageOldReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchSendOldRawMessage(ctx, &lark.BatchSendOldRawMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.ReplyRawMessage(ctx, &lark.ReplyRawMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteMessage(ctx, &lark.DeleteMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchDeleteMessage(ctx, &lark.BatchDeleteMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateMessage(ctx, &lark.UpdateMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateMessageDelay(ctx, &lark.UpdateMessageDelayReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMessageReadUserList(ctx, &lark.GetMessageReadUserListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBatchSentMessageReadUser(ctx, &lark.GetBatchSentMessageReadUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBatchSentMessageProgress(ctx, &lark.GetBatchSentMessageProgressReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMessageList(ctx, &lark.GetMessageListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMessageFile(ctx, &lark.GetMessageFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMessage(ctx, &lark.GetMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteEphemeralMessage(ctx, &lark.DeleteEphemeralMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Message

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageSendEphemeralMessage(func(ctx context.Context, request *lark.SendEphemeralMessageReq, options ...lark.MethodOptionFunc) (*lark.SendEphemeralMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageSendEphemeralMessage()

			_, _, err := moduleCli.SendEphemeralMessage(ctx, &lark.SendEphemeralMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageSendUrgentAppMessage(func(ctx context.Context, request *lark.SendUrgentAppMessageReq, options ...lark.MethodOptionFunc) (*lark.SendUrgentAppMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageSendUrgentAppMessage()

			_, _, err := moduleCli.SendUrgentAppMessage(ctx, &lark.SendUrgentAppMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageSendUrgentSmsMessage(func(ctx context.Context, request *lark.SendUrgentSmsMessageReq, options ...lark.MethodOptionFunc) (*lark.SendUrgentSmsMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageSendUrgentSmsMessage()

			_, _, err := moduleCli.SendUrgentSmsMessage(ctx, &lark.SendUrgentSmsMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageSendUrgentPhoneMessage(func(ctx context.Context, request *lark.SendUrgentPhoneMessageReq, options ...lark.MethodOptionFunc) (*lark.SendUrgentPhoneMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageSendUrgentPhoneMessage()

			_, _, err := moduleCli.SendUrgentPhoneMessage(ctx, &lark.SendUrgentPhoneMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageSendRawMessage(func(ctx context.Context, request *lark.SendRawMessageReq, options ...lark.MethodOptionFunc) (*lark.SendRawMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageSendRawMessage()

			_, _, err := moduleCli.SendRawMessage(ctx, &lark.SendRawMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageSendRawMessageOld(func(ctx context.Context, request *lark.SendRawMessageOldReq, options ...lark.MethodOptionFunc) (*lark.SendRawMessageOldResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageSendRawMessageOld()

			_, _, err := moduleCli.SendRawMessageOld(ctx, &lark.SendRawMessageOldReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageBatchSendOldRawMessage(func(ctx context.Context, request *lark.BatchSendOldRawMessageReq, options ...lark.MethodOptionFunc) (*lark.BatchSendOldRawMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageBatchSendOldRawMessage()

			_, _, err := moduleCli.BatchSendOldRawMessage(ctx, &lark.BatchSendOldRawMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageReplyRawMessage(func(ctx context.Context, request *lark.ReplyRawMessageReq, options ...lark.MethodOptionFunc) (*lark.ReplyRawMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageReplyRawMessage()

			_, _, err := moduleCli.ReplyRawMessage(ctx, &lark.ReplyRawMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageDeleteMessage(func(ctx context.Context, request *lark.DeleteMessageReq, options ...lark.MethodOptionFunc) (*lark.DeleteMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageDeleteMessage()

			_, _, err := moduleCli.DeleteMessage(ctx, &lark.DeleteMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageBatchDeleteMessage(func(ctx context.Context, request *lark.BatchDeleteMessageReq, options ...lark.MethodOptionFunc) (*lark.BatchDeleteMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageBatchDeleteMessage()

			_, _, err := moduleCli.BatchDeleteMessage(ctx, &lark.BatchDeleteMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageUpdateMessage(func(ctx context.Context, request *lark.UpdateMessageReq, options ...lark.MethodOptionFunc) (*lark.UpdateMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageUpdateMessage()

			_, _, err := moduleCli.UpdateMessage(ctx, &lark.UpdateMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageUpdateMessageDelay(func(ctx context.Context, request *lark.UpdateMessageDelayReq, options ...lark.MethodOptionFunc) (*lark.UpdateMessageDelayResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageUpdateMessageDelay()

			_, _, err := moduleCli.UpdateMessageDelay(ctx, &lark.UpdateMessageDelayReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageGetMessageReadUserList(func(ctx context.Context, request *lark.GetMessageReadUserListReq, options ...lark.MethodOptionFunc) (*lark.GetMessageReadUserListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageGetMessageReadUserList()

			_, _, err := moduleCli.GetMessageReadUserList(ctx, &lark.GetMessageReadUserListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageGetBatchSentMessageReadUser(func(ctx context.Context, request *lark.GetBatchSentMessageReadUserReq, options ...lark.MethodOptionFunc) (*lark.GetBatchSentMessageReadUserResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageGetBatchSentMessageReadUser()

			_, _, err := moduleCli.GetBatchSentMessageReadUser(ctx, &lark.GetBatchSentMessageReadUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageGetBatchSentMessageProgress(func(ctx context.Context, request *lark.GetBatchSentMessageProgressReq, options ...lark.MethodOptionFunc) (*lark.GetBatchSentMessageProgressResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageGetBatchSentMessageProgress()

			_, _, err := moduleCli.GetBatchSentMessageProgress(ctx, &lark.GetBatchSentMessageProgressReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageGetMessageList(func(ctx context.Context, request *lark.GetMessageListReq, options ...lark.MethodOptionFunc) (*lark.GetMessageListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageGetMessageList()

			_, _, err := moduleCli.GetMessageList(ctx, &lark.GetMessageListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageGetMessageFile(func(ctx context.Context, request *lark.GetMessageFileReq, options ...lark.MethodOptionFunc) (*lark.GetMessageFileResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageGetMessageFile()

			_, _, err := moduleCli.GetMessageFile(ctx, &lark.GetMessageFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageGetMessage(func(ctx context.Context, request *lark.GetMessageReq, options ...lark.MethodOptionFunc) (*lark.GetMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageGetMessage()

			_, _, err := moduleCli.GetMessage(ctx, &lark.GetMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMessageDeleteEphemeralMessage(func(ctx context.Context, request *lark.DeleteEphemeralMessageReq, options ...lark.MethodOptionFunc) (*lark.DeleteEphemeralMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMessageDeleteEphemeralMessage()

			_, _, err := moduleCli.DeleteEphemeralMessage(ctx, &lark.DeleteEphemeralMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Message

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendEphemeralMessage(ctx, &lark.SendEphemeralMessageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendUrgentAppMessage(ctx, &lark.SendUrgentAppMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendUrgentSmsMessage(ctx, &lark.SendUrgentSmsMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendUrgentPhoneMessage(ctx, &lark.SendUrgentPhoneMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendRawMessage(ctx, &lark.SendRawMessageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendRawMessageOld(ctx, &lark.SendRawMessageOldReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchSendOldRawMessage(ctx, &lark.BatchSendOldRawMessageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.ReplyRawMessage(ctx, &lark.ReplyRawMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteMessage(ctx, &lark.DeleteMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchDeleteMessage(ctx, &lark.BatchDeleteMessageReq{
				BatchMessageID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateMessage(ctx, &lark.UpdateMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateMessageDelay(ctx, &lark.UpdateMessageDelayReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMessageReadUserList(ctx, &lark.GetMessageReadUserListReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBatchSentMessageReadUser(ctx, &lark.GetBatchSentMessageReadUserReq{
				BatchMessageID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBatchSentMessageProgress(ctx, &lark.GetBatchSentMessageProgressReq{
				BatchMessageID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMessageList(ctx, &lark.GetMessageListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMessageFile(ctx, &lark.GetMessageFileReq{
				MessageID: "x",
				FileKey:   "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMessage(ctx, &lark.GetMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteEphemeralMessage(ctx, &lark.DeleteEphemeralMessageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Message
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendEphemeralMessage(ctx, &lark.SendEphemeralMessageReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendUrgentAppMessage(ctx, &lark.SendUrgentAppMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendUrgentSmsMessage(ctx, &lark.SendUrgentSmsMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendUrgentPhoneMessage(ctx, &lark.SendUrgentPhoneMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendRawMessage(ctx, &lark.SendRawMessageReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendRawMessageOld(ctx, &lark.SendRawMessageOldReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchSendOldRawMessage(ctx, &lark.BatchSendOldRawMessageReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.ReplyRawMessage(ctx, &lark.ReplyRawMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteMessage(ctx, &lark.DeleteMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchDeleteMessage(ctx, &lark.BatchDeleteMessageReq{
				BatchMessageID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateMessage(ctx, &lark.UpdateMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateMessageDelay(ctx, &lark.UpdateMessageDelayReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMessageReadUserList(ctx, &lark.GetMessageReadUserListReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBatchSentMessageReadUser(ctx, &lark.GetBatchSentMessageReadUserReq{
				BatchMessageID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBatchSentMessageProgress(ctx, &lark.GetBatchSentMessageProgressReq{
				BatchMessageID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMessageList(ctx, &lark.GetMessageListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMessageFile(ctx, &lark.GetMessageFileReq{
				MessageID: "x",
				FileKey:   "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMessage(ctx, &lark.GetMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteEphemeralMessage(ctx, &lark.DeleteEphemeralMessageReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})
	})
}
