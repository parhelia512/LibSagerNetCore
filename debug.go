package libcore

import (
	"net/http"
	_ "net/http/pprof"

	"libcore/comm"
)

type DebugInstance struct {
	server *http.Server
}

func NewDebugInstance(addr string) *DebugInstance {
	s := &http.Server{
		Addr: addr,
	}
	go func() {
		_ = s.ListenAndServe()
	}()
	return &DebugInstance{s}
}

func (d *DebugInstance) Close() {
	comm.CloseIgnore(d.server)
}
