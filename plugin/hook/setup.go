package example

import (
	"github.com/caddyserver/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/pkg/log"
)

func init() {
	log.Info("init hook")
	plugin.Register("hook", setup) }

func setup(c *caddy.Controller) error {
	c.Next() // 'hook'
	if c.NextArg() {
		return plugin.Error("hook", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Hook{}
	})

	return nil
}