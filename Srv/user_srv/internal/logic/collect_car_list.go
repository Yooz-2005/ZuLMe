package logic

import (
	"models/model_mysql"
	user "user_srv/proto_user"
)

func CollectVehicleList(in *user.CollectVehicleListRequest) (*user.CollectVehicleListResponse, error) {
	f := &model_mysql.Favourite{}
	collect, err := f.GetUserCollectVehicle(in.UserId)
	if err != nil {
		return nil, err
	}
	var collectVehicle []*user.Vehicle
	for _, v := range collect {
		collectVehicle = append(collectVehicle, &user.Vehicle{
			VehicleId:   int64(v.VehicleId),
			VehicleName: v.VehicleName,
			Image:       v.Image,
		})
	}
	return &user.CollectVehicleListResponse{
		UserId:      in.UserId,
		VehicleList: collectVehicle,
	}, nil
}
