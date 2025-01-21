package proxy

import (
	"context"
	"log"

	"github.com/shimingyah/pool"
	"google.golang.org/grpc/metadata"
)

var poolClient map[string]pool.Pool

func GetConnect(addr string) (pool.Conn, error) {
	p, ok := poolClient[addr]
	if !ok {
		p, err := pool.New(addr, pool.DefaultOptions)
		if err != nil {
			log.Fatalf("failed to new pool: %v", err)
			return nil, err
		}
		poolClient[addr] = p
	}
	conn, err := p.Get()
	if err != nil {
		log.Fatalf("failed to get conn: %v", err)
	}
	return conn, err
}

func GetMetadataCtx(trace_id, source string, kv ...string) context.Context {
	md := metadata.Pairs(kv...)
	mdCtx := metadata.NewOutgoingContext(context.Background(), md)
	return mdCtx
}
