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
		var err error
		p, err = pool.New(addr, pool.DefaultOptions)
		if err != nil {
			log.Fatalf("failed to new pool: %v", err)
			return nil, err
		}
		if poolClient == nil {
			poolClient = make(map[string]pool.Pool)
		}
		poolClient[addr] = p
	}
	conn, err := p.Get()
	if err != nil {
		log.Fatalf("failed to get conn: %v", err)
	}
	return conn, err
}

func GetMetadataCtx(traceId, source, token string, kv ...string) context.Context {
	wkv := []string{"trace_id", traceId, "project_source", source, "token", token}
	if len(kv) > 0 {
		wkv = append(wkv, kv...)
	}
	md := metadata.Pairs(wkv...)
	mdCtx := metadata.NewOutgoingContext(context.Background(), md)
	return mdCtx
}
