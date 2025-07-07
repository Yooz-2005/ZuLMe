package server

import (
	"context"
	"coupon_srv/internal/logic"
	coupon "coupon_srv/proto_coupon"
)

type ServerCoupon struct {
	coupon.UnimplementedCouponServiceServer
}

func (s ServerCoupon) GrantCoupon(ctx context.Context, req *coupon.GrantCouponRequest) (*coupon.GrantCouponResponse, error) {
	grantCoupon, err := logic.GrantCoupon(req)
	if err != nil {
		return nil, err
	}
	return grantCoupon, nil
}

func (s ServerCoupon) GetUserCoupons(ctx context.Context, req *coupon.GetUserCouponsRequest) (*coupon.GetUserCouponsResponse, error) {
	userCoupons, err := logic.GetUserCoupons(req)
	if err != nil {
		return nil, err
	}
	return userCoupons, nil
}

func (s ServerCoupon) GetAvailableCoupons(ctx context.Context, req *coupon.GetAvailableCouponsRequest) (*coupon.GetAvailableCouponsResponse, error) {
	availableCoupons, err := logic.GetAvailableCoupons(req)
	if err != nil {
		return nil, err
	}
	return availableCoupons, nil
}

func (s ServerCoupon) ValidateCoupon(ctx context.Context, req *coupon.ValidateCouponRequest) (*coupon.ValidateCouponResponse, error) {
	validateCoupon, err := logic.ValidateCoupon(req)
	if err != nil {
		return nil, err
	}
	return validateCoupon, nil
}

func (s ServerCoupon) UseCoupon(ctx context.Context, req *coupon.UseCouponRequest) (*coupon.UseCouponResponse, error) {
	useCoupon, err := logic.UseCoupon(req)
	if err != nil {
		return nil, err
	}
	return useCoupon, nil
}

// ===== 活动管理接口实现 =====

func (s ServerCoupon) CreateActivity(ctx context.Context, req *coupon.CreateActivityRequest) (*coupon.CreateActivityResponse, error) {
	createActivity, err := logic.CreateActivity(req)
	if err != nil {
		return nil, err
	}
	return createActivity, nil
}

func (s ServerCoupon) UpdateActivity(ctx context.Context, req *coupon.UpdateActivityRequest) (*coupon.UpdateActivityResponse, error) {
	updateActivity, err := logic.UpdateActivity(req)
	if err != nil {
		return nil, err
	}
	return updateActivity, nil
}

func (s ServerCoupon) DeleteActivity(ctx context.Context, req *coupon.DeleteActivityRequest) (*coupon.DeleteActivityResponse, error) {
	deleteActivity, err := logic.DeleteActivity(req)
	if err != nil {
		return nil, err
	}
	return deleteActivity, nil
}

func (s ServerCoupon) GetActivity(ctx context.Context, req *coupon.GetActivityRequest) (*coupon.GetActivityResponse, error) {
	getActivity, err := logic.GetActivity(req)
	if err != nil {
		return nil, err
	}
	return getActivity, nil
}

func (s ServerCoupon) GetActivityList(ctx context.Context, req *coupon.GetActivityListRequest) (*coupon.GetActivityListResponse, error) {
	getActivityList, err := logic.GetActivityList(req)
	if err != nil {
		return nil, err
	}
	return getActivityList, nil
}

func (s ServerCoupon) DistributeActivityCoupons(ctx context.Context, req *coupon.DistributeActivityCouponsRequest) (*coupon.DistributeActivityCouponsResponse, error) {
	distributeActivityCoupons, err := logic.DistributeActivityCoupons(req)
	if err != nil {
		return nil, err
	}
	return distributeActivityCoupons, nil
}
