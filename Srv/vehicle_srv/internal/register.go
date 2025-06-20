package internal

import (
	"google.golang.org/grpc"
	vehicle "vehicle_srv/proto_vehicle"
	"vehicle_srv/server"
)

func RegisterVehicleServer(ser *grpc.Server) {
	vehicle.RegisterVehicleServer(ser, &server.ServerVehicle{})
}
