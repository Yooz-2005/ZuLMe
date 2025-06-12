package server

import order "order_srv/proto_order"

type ServerOrder struct {
	order.UnimplementedOrderServer
}
