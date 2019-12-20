package handle

import (
	"context"
	"github.com/micro/go-micro/util/log"
	proto "seckill/proto"
)



type Product struct {}

func (p *Product) Detail(ctx context.Context, req *proto.DetailRequest, rsp *proto.Detailesponse) error {
	log.Info(req)
	rsp.Total  = 1100
	return nil
}