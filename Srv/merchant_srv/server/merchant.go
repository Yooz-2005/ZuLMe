package server

import merchant "merchant_srv/proto_merchant"

type ServerMerchant struct {
	merchant.UnimplementedMerchantServer
}
