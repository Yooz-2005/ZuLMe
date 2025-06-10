package server

import (
	"admin_srv/internal/logic"
	admin "admin_srv/proto_admin"
	"context"
)

type ServerAdmin struct {
	admin.UnimplementedAdminServer
}

// MerchantApprove 商户审核
func (s ServerAdmin) MerchantApprove(ctx context.Context, in *admin.MerchantApproveRequest) (*admin.MerchantApproveResponse, error) {
	return logic.MerchantApprove(ctx, in)
}

// MerchantUpdate 商户更新
func (s ServerAdmin) MerchantUpdate(ctx context.Context, in *admin.MerchantUpdateRequest) (*admin.MerchantUpdateResponse, error) {
	return logic.MerchantUpdate(ctx, in)
}

// MerchantDelete 商户删除
func (s ServerAdmin) MerchantDelete(ctx context.Context, in *admin.MerchantDeleteRequest) (*admin.MerchantDeleteResponse, error) {
	return logic.MerchantDelete(ctx, in)
}

// MerchantList 商户列表
func (s ServerAdmin) MerchantList(ctx context.Context, in *admin.MerchantListRequest) (*admin.MerchantListResponse, error) {
	return logic.MerchantList(ctx, in)
}

// MerchantDetail 商户详情
func (s ServerAdmin) MerchantDetail(ctx context.Context, in *admin.MerchantDetailRequest) (*admin.MerchantDetailResponse, error) {
	return logic.MerchantDetail(ctx, in)
}
