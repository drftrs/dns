package drftrs

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/mholt/caddy"
)

func init() {
	caddy.RegisterPlugin("drftrs", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})

}

func setup(c *caddy.Controller) error {
	c.Next() // Skip directive name
	if !c.NextArg() {
		return c.ArgErr()
	}

	blacklistUrl := c.Val()
	blacklist, err := CompileBlacklist(blacklistUrl)
	if err != nil {
		return c.Errf("Error downloading blacklist: %s", err)
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return &DNS{
			Next:      next,
			Blacklist: blacklist,
		}
	})

	return nil
}
