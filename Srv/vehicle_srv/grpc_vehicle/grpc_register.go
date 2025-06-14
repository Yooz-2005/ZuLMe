package grpc_vehicle

import (
	"ZuLMe/ZuLMe/Srv/vehicle_srv/internal"
	"google.golang.org/grpc"
)

func RegisterVehicleServices(s *grpc.Server) {
	internal.RegisterVehicleServer(s)
}
