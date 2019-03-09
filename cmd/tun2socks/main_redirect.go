// +build redirect

package main

import (
	"github.com/eycorsican/go-tun2socks/core"
	"github.com/eycorsican/go-tun2socks/proxy/redirect"
)

func init() {
	args.addFlag("proxyServer")
	args.addFlag("udpTimeout")

	registerHandlerCreater("redirect", func() {
		core.RegisterTCPConnectionHandler(redirect.NewTCPHandler(*args.ProxyServer))
		core.RegisterUDPConnectionHandler(redirect.NewUDPHandler(*args.ProxyServer, *args.UdpTimeout))
	})
}