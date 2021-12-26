package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mongodb MongoConf       // manual code
	Cache   cache.CacheConf // manual code
}

type MongoConf struct {
	Uri         string
	Db          string
	MaxPoolSize uint64
}
