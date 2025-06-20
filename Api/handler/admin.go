package handler

import (
	"Api/client"
	"Api/request"
	admin "admin_srv/proto_admin"
	"context"
)

func MerchantApprove(ctx context.Context, req *request.MerchantApproveRequest) (*admin.MerchantApproveResponse, error) {
	adminClient, err := client.AdminClient(ctx, func(ctx context.Context, in admin.AdminClient) (interface{}, error) {
		approveReq := &admin.MerchantApproveRequest{
			Id:     req.ID,
			Status: req.Status,
		}
		return in.MerchantApprove(ctx, approveReq)
	})
	if err != nil {
		return nil, err
	}
	return adminClient.(*admin.MerchantApproveResponse), nil
}

func MerchantUpdate(ctx context.Context, req *request.MerchantUpdateRequest) (*admin.MerchantUpdateResponse, error) {
	adminClient, err := client.AdminClient(ctx, func(ctx context.Context, in admin.AdminClient) (interface{}, error) {
		updateReq := &admin.MerchantUpdateRequest{
			Id:           req.ID,
			Name:         req.Name,
			Phone:        req.Phone,
			Email:        req.Email,
			Location:     req.Location,
			BusinessTime: req.BusinessTime,
			Longitude:    req.Longitude,
			Latitude:     req.Latitude,
		}
		return in.MerchantUpdate(ctx, updateReq)
	})
	if err != nil {
		return nil, err
	}
	return adminClient.(*admin.MerchantUpdateResponse), nil
}

func MerchantDelete(ctx context.Context, req *request.MerchantDeleteRequest) (*admin.MerchantDeleteResponse, error) {
	adminClient, err := client.AdminClient(ctx, func(ctx context.Context, in admin.AdminClient) (interface{}, error) {
		deleteReq := &admin.MerchantDeleteRequest{
			Id: req.ID,
		}
		return in.MerchantDelete(ctx, deleteReq)
	})
	if err != nil {
		return nil, err
	}
	return adminClient.(*admin.MerchantDeleteResponse), nil
}

func MerchantList(ctx context.Context, req *request.MerchantListRequest) (*admin.MerchantListResponse, error) {
	adminClient, err := client.AdminClient(ctx, func(ctx context.Context, in admin.AdminClient) (interface{}, error) {
		listReq := &admin.MerchantListRequest{
			Page:         req.Page,
			PageSize:     req.PageSize,
			Keyword:      req.Keyword,
			StatusFilter: req.StatusFilter,
		}
		return in.MerchantList(ctx, listReq)
	})
	if err != nil {
		return nil, err
	}
	return adminClient.(*admin.MerchantListResponse), nil
}

func MerchantDetail(ctx context.Context, req *request.MerchantDetailRequest) (*admin.MerchantDetailResponse, error) {
	adminClient, err := client.AdminClient(ctx, func(ctx context.Context, in admin.AdminClient) (interface{}, error) {
		detailReq := &admin.MerchantDetailRequest{
			Id: req.ID,
		}
		return in.MerchantDetail(ctx, detailReq)
	})
	if err != nil {
		return nil, err
	}
	return adminClient.(*admin.MerchantDetailResponse), nil
}
