package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/peterlearn/kratos/pkg/conf/env"
	"github.com/peterlearn/kratos/pkg/naming"
	"github.com/peterlearn/kratos/pkg/naming/discovery"
	"github.com/peterlearn/kratos/pkg/net/rpc/warden/resolver"
	"net/http"
	"os"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle scene")
	_, _ = w.Write([]byte("Hello scene!\n"))
}

func main() {
	flag.Parse()
	http.HandleFunc("/hello", handleHello)
	go http.ListenAndServe(":8888", nil)
	register()
	select {}
}

const (
	ServiceName = "scene"
	Namespace   = "test"
)

func register() {
	if os.Getenv("DISCOVERY") != "" {
		_, err := RegisterDiscovery(fmt.Sprintf("%s.%s", Namespace, ServiceName))
		if err != nil {
			panic(err)
		}
	}
}

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
