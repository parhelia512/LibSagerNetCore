package libcore

import (
	ss_common "github.com/v2fly/v2ray-core/v5/proxy/shadowsocks/common"
	"github.com/v2fly/v2ray-core/v5/proxy/sip003"
	"github.com/v2fly/v2ray-core/v5/proxy/sip003/self"
	"github.com/v2fly/v2ray-core/v5/transport/internet"

	"libcore/clash/transport/simple-obfs"
)

var _ sip003.StreamPlugin = (*obfsLocalPlugin)(nil)

func init() {
	sip003.RegisterPlugin("obfs-local", func() sip003.Plugin {
		return new(obfsLocalPlugin)
	})
}

type obfsLocalPlugin struct {
	tls  bool
	host string
	port string
}

func (p *obfsLocalPlugin) Init(_ string, _ string, _ string, remotePort string, pluginOpts string, _ []string, _ *ss_common.MemoryAccount) error {
	options, err := self.ParsePluginOptions(pluginOpts)
	if err != nil {
		return newError("obfs-local: failed to parse plugin options").Base(err)
	}

	mode := "http"

	if s, ok := options.Get("obfs"); ok {
		mode = s
	}

	if s, ok := options.Get("obfs-host"); ok {
		p.host = s
	}

	switch mode {
	case "http":
	case "tls":
		p.tls = true
	default:
		return newError("unknown obfs mode ", mode)
	}

	p.port = remotePort

	return nil
}

func (p *obfsLocalPlugin) StreamConn(connection internet.Connection) internet.Connection {
	if !p.tls {
		return obfs.NewHTTPObfs(connection, p.host, p.port)
	} else {
		return obfs.NewTLSObfs(connection, p.host)
	}
}

func (p *obfsLocalPlugin) Close() error {
	return nil
}
