package logic

import (
	user "ZuLMe/ZuLMe/Srv/user_srv/proto_user"
	"ZuLMe/ZuLMe/models/model_mysql"
	"errors"
	"time"
)

func CollectVehicle(in *user.CollectVehicleRequest) (*user.CollectVehicleResponse, error) {
	v := &model_mysql.Vehicle{}
	err := v.GetByID(uint(in.VehicleId))
	if err != nil {
		return nil, errors.New("车辆不存在")
	}
	f := &model_mysql.Favourite{
		UserId:        int32(in.UserId),
		VehicleId:     int32(in.VehicleId),
		VehicleName:   v.Brand,
		Image:         v.Images,
		FavouriteTime: time.Now(),
	}
	if !f.IsCollectVehicle(in.VehicleId, in.UserId) {
		// 未收藏，执行收藏
		err = f.CollectVehicle()
		if err != nil {
			return &user.CollectVehicleResponse{
				UserId:  in.UserId,
				Message: "收藏失败",
			}, err
		}
		return &user.CollectVehicleResponse{
			UserId:  in.UserId,
			Message: "收藏成功",
		}, nil
	} else {
		// 已收藏，取消收藏
		err = f.CancelCollectVehicle(in.VehicleId, in.UserId)
		if err != nil {
			return &user.CollectVehicleResponse{
				UserId:  in.UserId,
				Message: "取消收藏失败",
			}, err
		}
		return &user.CollectVehicleResponse{
			UserId:  in.UserId,
			Message: "取消收藏成功",
		}, nil
	}
}
