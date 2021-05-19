// Code generated by lark_sdk_gen. DO NOT EDIT.

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Drive_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Drive

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateDriveMemberPermission(ctx, &lark.CreateDriveMemberPermissionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.TransferDriveMemberPermission(ctx, &lark.TransferDriveMemberPermissionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateDrivePublicPermission(ctx, &lark.UpdateDrivePublicPermissionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetDriveMemberPermissionList(ctx, &lark.GetDriveMemberPermissionListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteDriveMemberPermission(ctx, &lark.DeleteDriveMemberPermissionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateDriveMemberPermission(ctx, &lark.UpdateDriveMemberPermissionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CheckDriveMemberPermission(ctx, &lark.CheckDriveMemberPermissionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateDrivePublicPermissionV2(ctx, &lark.UpdateDrivePublicPermissionV2Req{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetDrivePublicPermissionV2(ctx, &lark.GetDrivePublicPermissionV2Req{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetDriveCommentList(ctx, &lark.GetDriveCommentListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetDriveComment(ctx, &lark.GetDriveCommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateDriveComment(ctx, &lark.CreateDriveCommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteDriveComment(ctx, &lark.DeleteDriveCommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateDriveCommentPatch(ctx, &lark.UpdateDriveCommentPatchReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetSheetMeta(ctx, &lark.GetSheetMetaReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateSheetProperty(ctx, &lark.UpdateSheetPropertyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchUpdateSheet(ctx, &lark.BatchUpdateSheetReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.ImportSheet(ctx, &lark.ImportSheetReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.PrependSheetValue(ctx, &lark.PrependSheetValueReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.AppendSheetValue(ctx, &lark.AppendSheetValueReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.InsertSheetDimensionRange(ctx, &lark.InsertSheetDimensionRangeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.AddSheetDimensionRange(ctx, &lark.AddSheetDimensionRangeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateSheetDimensionRange(ctx, &lark.UpdateSheetDimensionRangeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteSheetDimensionRange(ctx, &lark.DeleteSheetDimensionRangeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateSheetConditionFormat(ctx, &lark.CreateSheetConditionFormatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetSheetConditionFormat(ctx, &lark.GetSheetConditionFormatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateSheetConditionFormat(ctx, &lark.UpdateSheetConditionFormatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteSheetConditionFormat(ctx, &lark.DeleteSheetConditionFormatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateSheetProtectedDimension(ctx, &lark.CreateSheetProtectedDimensionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetSheetProtectedDimension(ctx, &lark.GetSheetProtectedDimensionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateSheetProtectedDimension(ctx, &lark.UpdateSheetProtectedDimensionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteSheetProtectedDimension(ctx, &lark.DeleteSheetProtectedDimensionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		moduleCli := cli.Drive

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveCreateDriveMemberPermission(func(ctx context.Context, request *lark.CreateDriveMemberPermissionReq, options ...lark.MethodOptionFunc) (*lark.CreateDriveMemberPermissionResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveCreateDriveMemberPermission()

			_, _, err := moduleCli.CreateDriveMemberPermission(ctx, &lark.CreateDriveMemberPermissionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveTransferDriveMemberPermission(func(ctx context.Context, request *lark.TransferDriveMemberPermissionReq, options ...lark.MethodOptionFunc) (*lark.TransferDriveMemberPermissionResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveTransferDriveMemberPermission()

			_, _, err := moduleCli.TransferDriveMemberPermission(ctx, &lark.TransferDriveMemberPermissionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveUpdateDrivePublicPermission(func(ctx context.Context, request *lark.UpdateDrivePublicPermissionReq, options ...lark.MethodOptionFunc) (*lark.UpdateDrivePublicPermissionResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveUpdateDrivePublicPermission()

			_, _, err := moduleCli.UpdateDrivePublicPermission(ctx, &lark.UpdateDrivePublicPermissionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveGetDriveMemberPermissionList(func(ctx context.Context, request *lark.GetDriveMemberPermissionListReq, options ...lark.MethodOptionFunc) (*lark.GetDriveMemberPermissionListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveGetDriveMemberPermissionList()

			_, _, err := moduleCli.GetDriveMemberPermissionList(ctx, &lark.GetDriveMemberPermissionListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveDeleteDriveMemberPermission(func(ctx context.Context, request *lark.DeleteDriveMemberPermissionReq, options ...lark.MethodOptionFunc) (*lark.DeleteDriveMemberPermissionResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveDeleteDriveMemberPermission()

			_, _, err := moduleCli.DeleteDriveMemberPermission(ctx, &lark.DeleteDriveMemberPermissionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveUpdateDriveMemberPermission(func(ctx context.Context, request *lark.UpdateDriveMemberPermissionReq, options ...lark.MethodOptionFunc) (*lark.UpdateDriveMemberPermissionResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveUpdateDriveMemberPermission()

			_, _, err := moduleCli.UpdateDriveMemberPermission(ctx, &lark.UpdateDriveMemberPermissionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveCheckDriveMemberPermission(func(ctx context.Context, request *lark.CheckDriveMemberPermissionReq, options ...lark.MethodOptionFunc) (*lark.CheckDriveMemberPermissionResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveCheckDriveMemberPermission()

			_, _, err := moduleCli.CheckDriveMemberPermission(ctx, &lark.CheckDriveMemberPermissionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveUpdateDrivePublicPermissionV2(func(ctx context.Context, request *lark.UpdateDrivePublicPermissionV2Req, options ...lark.MethodOptionFunc) (*lark.UpdateDrivePublicPermissionV2Resp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveUpdateDrivePublicPermissionV2()

			_, _, err := moduleCli.UpdateDrivePublicPermissionV2(ctx, &lark.UpdateDrivePublicPermissionV2Req{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveGetDrivePublicPermissionV2(func(ctx context.Context, request *lark.GetDrivePublicPermissionV2Req, options ...lark.MethodOptionFunc) (*lark.GetDrivePublicPermissionV2Resp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveGetDrivePublicPermissionV2()

			_, _, err := moduleCli.GetDrivePublicPermissionV2(ctx, &lark.GetDrivePublicPermissionV2Req{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveGetDriveCommentList(func(ctx context.Context, request *lark.GetDriveCommentListReq, options ...lark.MethodOptionFunc) (*lark.GetDriveCommentListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveGetDriveCommentList()

			_, _, err := moduleCli.GetDriveCommentList(ctx, &lark.GetDriveCommentListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveGetDriveComment(func(ctx context.Context, request *lark.GetDriveCommentReq, options ...lark.MethodOptionFunc) (*lark.GetDriveCommentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveGetDriveComment()

			_, _, err := moduleCli.GetDriveComment(ctx, &lark.GetDriveCommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveCreateDriveComment(func(ctx context.Context, request *lark.CreateDriveCommentReq, options ...lark.MethodOptionFunc) (*lark.CreateDriveCommentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveCreateDriveComment()

			_, _, err := moduleCli.CreateDriveComment(ctx, &lark.CreateDriveCommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveDeleteDriveComment(func(ctx context.Context, request *lark.DeleteDriveCommentReq, options ...lark.MethodOptionFunc) (*lark.DeleteDriveCommentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveDeleteDriveComment()

			_, _, err := moduleCli.DeleteDriveComment(ctx, &lark.DeleteDriveCommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveUpdateDriveCommentPatch(func(ctx context.Context, request *lark.UpdateDriveCommentPatchReq, options ...lark.MethodOptionFunc) (*lark.UpdateDriveCommentPatchResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveUpdateDriveCommentPatch()

			_, _, err := moduleCli.UpdateDriveCommentPatch(ctx, &lark.UpdateDriveCommentPatchReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveGetSheetMeta(func(ctx context.Context, request *lark.GetSheetMetaReq, options ...lark.MethodOptionFunc) (*lark.GetSheetMetaResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveGetSheetMeta()

			_, _, err := moduleCli.GetSheetMeta(ctx, &lark.GetSheetMetaReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveUpdateSheetProperty(func(ctx context.Context, request *lark.UpdateSheetPropertyReq, options ...lark.MethodOptionFunc) (*lark.UpdateSheetPropertyResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveUpdateSheetProperty()

			_, _, err := moduleCli.UpdateSheetProperty(ctx, &lark.UpdateSheetPropertyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveBatchUpdateSheet(func(ctx context.Context, request *lark.BatchUpdateSheetReq, options ...lark.MethodOptionFunc) (*lark.BatchUpdateSheetResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveBatchUpdateSheet()

			_, _, err := moduleCli.BatchUpdateSheet(ctx, &lark.BatchUpdateSheetReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveImportSheet(func(ctx context.Context, request *lark.ImportSheetReq, options ...lark.MethodOptionFunc) (*lark.ImportSheetResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveImportSheet()

			_, _, err := moduleCli.ImportSheet(ctx, &lark.ImportSheetReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDrivePrependSheetValue(func(ctx context.Context, request *lark.PrependSheetValueReq, options ...lark.MethodOptionFunc) (*lark.PrependSheetValueResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDrivePrependSheetValue()

			_, _, err := moduleCli.PrependSheetValue(ctx, &lark.PrependSheetValueReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveAppendSheetValue(func(ctx context.Context, request *lark.AppendSheetValueReq, options ...lark.MethodOptionFunc) (*lark.AppendSheetValueResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveAppendSheetValue()

			_, _, err := moduleCli.AppendSheetValue(ctx, &lark.AppendSheetValueReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveInsertSheetDimensionRange(func(ctx context.Context, request *lark.InsertSheetDimensionRangeReq, options ...lark.MethodOptionFunc) (*lark.InsertSheetDimensionRangeResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveInsertSheetDimensionRange()

			_, _, err := moduleCli.InsertSheetDimensionRange(ctx, &lark.InsertSheetDimensionRangeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveAddSheetDimensionRange(func(ctx context.Context, request *lark.AddSheetDimensionRangeReq, options ...lark.MethodOptionFunc) (*lark.AddSheetDimensionRangeResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveAddSheetDimensionRange()

			_, _, err := moduleCli.AddSheetDimensionRange(ctx, &lark.AddSheetDimensionRangeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveUpdateSheetDimensionRange(func(ctx context.Context, request *lark.UpdateSheetDimensionRangeReq, options ...lark.MethodOptionFunc) (*lark.UpdateSheetDimensionRangeResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveUpdateSheetDimensionRange()

			_, _, err := moduleCli.UpdateSheetDimensionRange(ctx, &lark.UpdateSheetDimensionRangeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveDeleteSheetDimensionRange(func(ctx context.Context, request *lark.DeleteSheetDimensionRangeReq, options ...lark.MethodOptionFunc) (*lark.DeleteSheetDimensionRangeResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveDeleteSheetDimensionRange()

			_, _, err := moduleCli.DeleteSheetDimensionRange(ctx, &lark.DeleteSheetDimensionRangeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveCreateSheetConditionFormat(func(ctx context.Context, request *lark.CreateSheetConditionFormatReq, options ...lark.MethodOptionFunc) (*lark.CreateSheetConditionFormatResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveCreateSheetConditionFormat()

			_, _, err := moduleCli.CreateSheetConditionFormat(ctx, &lark.CreateSheetConditionFormatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveGetSheetConditionFormat(func(ctx context.Context, request *lark.GetSheetConditionFormatReq, options ...lark.MethodOptionFunc) (*lark.GetSheetConditionFormatResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveGetSheetConditionFormat()

			_, _, err := moduleCli.GetSheetConditionFormat(ctx, &lark.GetSheetConditionFormatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveUpdateSheetConditionFormat(func(ctx context.Context, request *lark.UpdateSheetConditionFormatReq, options ...lark.MethodOptionFunc) (*lark.UpdateSheetConditionFormatResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveUpdateSheetConditionFormat()

			_, _, err := moduleCli.UpdateSheetConditionFormat(ctx, &lark.UpdateSheetConditionFormatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveDeleteSheetConditionFormat(func(ctx context.Context, request *lark.DeleteSheetConditionFormatReq, options ...lark.MethodOptionFunc) (*lark.DeleteSheetConditionFormatResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveDeleteSheetConditionFormat()

			_, _, err := moduleCli.DeleteSheetConditionFormat(ctx, &lark.DeleteSheetConditionFormatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveCreateSheetProtectedDimension(func(ctx context.Context, request *lark.CreateSheetProtectedDimensionReq, options ...lark.MethodOptionFunc) (*lark.CreateSheetProtectedDimensionResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveCreateSheetProtectedDimension()

			_, _, err := moduleCli.CreateSheetProtectedDimension(ctx, &lark.CreateSheetProtectedDimensionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveGetSheetProtectedDimension(func(ctx context.Context, request *lark.GetSheetProtectedDimensionReq, options ...lark.MethodOptionFunc) (*lark.GetSheetProtectedDimensionResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveGetSheetProtectedDimension()

			_, _, err := moduleCli.GetSheetProtectedDimension(ctx, &lark.GetSheetProtectedDimensionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveUpdateSheetProtectedDimension(func(ctx context.Context, request *lark.UpdateSheetProtectedDimensionReq, options ...lark.MethodOptionFunc) (*lark.UpdateSheetProtectedDimensionResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveUpdateSheetProtectedDimension()

			_, _, err := moduleCli.UpdateSheetProtectedDimension(ctx, &lark.UpdateSheetProtectedDimensionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockDriveDeleteSheetProtectedDimension(func(ctx context.Context, request *lark.DeleteSheetProtectedDimensionReq, options ...lark.MethodOptionFunc) (*lark.DeleteSheetProtectedDimensionResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockDriveDeleteSheetProtectedDimension()

			_, _, err := moduleCli.DeleteSheetProtectedDimension(ctx, &lark.DeleteSheetProtectedDimensionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Drive

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateDriveMemberPermission(ctx, &lark.CreateDriveMemberPermissionReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.TransferDriveMemberPermission(ctx, &lark.TransferDriveMemberPermissionReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateDrivePublicPermission(ctx, &lark.UpdateDrivePublicPermissionReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetDriveMemberPermissionList(ctx, &lark.GetDriveMemberPermissionListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteDriveMemberPermission(ctx, &lark.DeleteDriveMemberPermissionReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateDriveMemberPermission(ctx, &lark.UpdateDriveMemberPermissionReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CheckDriveMemberPermission(ctx, &lark.CheckDriveMemberPermissionReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateDrivePublicPermissionV2(ctx, &lark.UpdateDrivePublicPermissionV2Req{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetDrivePublicPermissionV2(ctx, &lark.GetDrivePublicPermissionV2Req{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetDriveCommentList(ctx, &lark.GetDriveCommentListReq{
				FileToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetDriveComment(ctx, &lark.GetDriveCommentReq{
				FileToken: "x",
				CommentID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateDriveComment(ctx, &lark.CreateDriveCommentReq{
				FileToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteDriveComment(ctx, &lark.DeleteDriveCommentReq{
				FileToken: "x",
				CommentID: "x",
				ReplyID:   "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateDriveCommentPatch(ctx, &lark.UpdateDriveCommentPatchReq{
				FileToken: "x",
				CommentID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetSheetMeta(ctx, &lark.GetSheetMetaReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateSheetProperty(ctx, &lark.UpdateSheetPropertyReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchUpdateSheet(ctx, &lark.BatchUpdateSheetReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.ImportSheet(ctx, &lark.ImportSheetReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.PrependSheetValue(ctx, &lark.PrependSheetValueReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.AppendSheetValue(ctx, &lark.AppendSheetValueReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.InsertSheetDimensionRange(ctx, &lark.InsertSheetDimensionRangeReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.AddSheetDimensionRange(ctx, &lark.AddSheetDimensionRangeReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateSheetDimensionRange(ctx, &lark.UpdateSheetDimensionRangeReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteSheetDimensionRange(ctx, &lark.DeleteSheetDimensionRangeReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateSheetConditionFormat(ctx, &lark.CreateSheetConditionFormatReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetSheetConditionFormat(ctx, &lark.GetSheetConditionFormatReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateSheetConditionFormat(ctx, &lark.UpdateSheetConditionFormatReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteSheetConditionFormat(ctx, &lark.DeleteSheetConditionFormatReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateSheetProtectedDimension(ctx, &lark.CreateSheetProtectedDimensionReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetSheetProtectedDimension(ctx, &lark.GetSheetProtectedDimensionReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateSheetProtectedDimension(ctx, &lark.UpdateSheetProtectedDimensionReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteSheetProtectedDimension(ctx, &lark.DeleteSheetProtectedDimensionReq{
				SpreadsheetToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}
