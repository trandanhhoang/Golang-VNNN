package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	CheckInventoryService zrpc.RpcClientConf
	DataSource            string          // manual code
	Cache                 cache.CacheConf // manual code
}
