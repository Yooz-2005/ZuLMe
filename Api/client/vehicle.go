package client

import (
	"context"
	vehicle "vehicle_srv/proto_vehicle"

	"google.golang.org/grpc"
)

type HandlerVehicle func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error)

func VehicleClient(ctx context.Context, handlerVehicle HandlerVehicle) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8004", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := vehicle.NewVehicleClient(dial)
	res, err := handlerVehicle(ctx, client)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	return res, nil
}
