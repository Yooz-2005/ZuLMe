package internal

import (
	vehicle "ZuLMe/ZuLMe/Srv/vehicle_srv/proto_vehicle"
	"ZuLMe/ZuLMe/Srv/vehicle_srv/server"
	"google.golang.org/grpc"
)

func RegisterVehicleServer(ser *grpc.Server) {
	vehicle.RegisterVehicleServer(ser, &server.ServerVehicle{})
}
