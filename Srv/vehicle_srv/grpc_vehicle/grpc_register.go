package grpc_vehicle

import (
	"google.golang.org/grpc"
	"vehicle_srv/internal"
)

func RegisterVehicleServices(s *grpc.Server) {
	internal.RegisterVehicleServer(s)
}
