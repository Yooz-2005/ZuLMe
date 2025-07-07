package handler

import (
	"Api/client"
	proto_coupon "Srv/coupon_srv/proto_coupon"
	"context"
)

// GrantCoupon 发放优惠券
func GrantCoupon(ctx context.Context, req *proto_coupon.GrantCouponRequest) (*proto_coupon.GrantCouponResponse, error) {
	couponClient, err := client.CouponClient(ctx, func(ctx context.Context, in proto_coupon.CouponServiceClient) (interface{}, error) {
		result, err := in.GrantCoupon(ctx, req)
		if err != nil {
			return nil, err
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}
	return couponClient.(*proto_coupon.GrantCouponResponse), nil
}

// GetUserCoupons 获取用户优惠券列表
func GetUserCoupons(ctx context.Context, req *proto_coupon.GetUserCouponsRequest) (*proto_coupon.GetUserCouponsResponse, error) {
	couponClient, err := client.CouponClient(ctx, func(ctx context.Context, in proto_coupon.CouponServiceClient) (interface{}, error) {
		coupons, err := in.GetUserCoupons(ctx, req)
		if err != nil {
			return nil, err
		}
		return coupons, nil
	})
	if err != nil {
		return nil, err
	}
	return couponClient.(*proto_coupon.GetUserCouponsResponse), nil
}

// GetAvailableCoupons 获取用户可用优惠券
func GetAvailableCoupons(ctx context.Context, req *proto_coupon.GetAvailableCouponsRequest) (*proto_coupon.GetAvailableCouponsResponse, error) {
	couponClient, err := client.CouponClient(ctx, func(ctx context.Context, in proto_coupon.CouponServiceClient) (interface{}, error) {
		coupons, err := in.GetAvailableCoupons(ctx, req)
		if err != nil {
			return nil, err
		}
		return coupons, nil
	})
	if err != nil {
		return nil, err
	}
	return couponClient.(*proto_coupon.GetAvailableCouponsResponse), nil
}

// ValidateCoupon 验证优惠券
func ValidateCoupon(ctx context.Context, req *proto_coupon.ValidateCouponRequest) (*proto_coupon.ValidateCouponResponse, error) {
	couponClient, err := client.CouponClient(ctx, func(ctx context.Context, in proto_coupon.CouponServiceClient) (interface{}, error) {
		result, err := in.ValidateCoupon(ctx, req)
		if err != nil {
			return nil, err
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}
	return couponClient.(*proto_coupon.ValidateCouponResponse), nil
}

// UseCoupon 使用优惠券
func UseCoupon(ctx context.Context, req *proto_coupon.UseCouponRequest) (*proto_coupon.UseCouponResponse, error) {
	couponClient, err := client.CouponClient(ctx, func(ctx context.Context, in proto_coupon.CouponServiceClient) (interface{}, error) {
		result, err := in.UseCoupon(ctx, req)
		if err != nil {
			return nil, err
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}
	return couponClient.(*proto_coupon.UseCouponResponse), nil
}
