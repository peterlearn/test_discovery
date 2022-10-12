package main

import (
	"flag"
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
