package main

import (
	"context"
	"flag"
	"github.com/peterlearn/kratos/pkg/conf/env"
	"github.com/peterlearn/kratos/pkg/naming"
	"github.com/peterlearn/kratos/pkg/naming/discovery"
	"github.com/peterlearn/kratos/pkg/net/rpc/warden/resolver"
	"os"
	"strconv"
)

var (
	_version  int64
	_serverid int64
)

func init() {
	addFlag(flag.CommandLine)
}

func addFlag(fs *flag.FlagSet) {
	v := os.Getenv("DISCOVERY_GRPC_ADDR")
	v = os.Getenv("VERSION")
	version, _ := strconv.ParseInt(v, 10, 64)
	fs.Int64Var(&_version, "version", version, "service version , default: 0")
	v = os.Getenv("SERVERID")
	serverid, _ := strconv.ParseInt(v, 10, 64)
	fs.Int64Var(&_serverid, "serverid", serverid, "serverid , default: 0")
}

// 注册discovery
func RegisterDiscovery(appId string) (context.CancelFunc, error) {
	resolver.Register(discovery.Builder())

	hn, _ := os.Hostname()
	dis := discovery.New(nil)
	ins := &naming.Instance{
		Zone:     env.Zone,
		Env:      env.DeployEnv,
		AppID:    appId,
		Hostname: hn,
		Addrs: []string{
			"127.0.0.1:8888",
		},
	}
	var (
		cancel context.CancelFunc
		err    error
	)
	if cancel, err = dis.Register(context.Background(), ins); err != nil {
		return nil, err
	}

	return cancel, nil
}
