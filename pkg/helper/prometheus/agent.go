package prometheus

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	once    sync.Once
	enabled atomic.Value
)

// Enabled returns if prometheus is enabled.
func Enabled() bool {
	return enabled.Load().(bool) == true
}

// StartAgent starts a prometheus agent.
func StartAgent(c Config) {
	if len(c.Host) == 0 {
		return
	}

	once.Do(func() {
		enabled.Store(true)
		go func() {
			http.Handle(c.Path, promhttp.Handler())
			addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
			fmt.Printf("Starting prometheus agent at %s\n", addr)
			if err := http.ListenAndServe(addr, nil); err != nil {
				fmt.Printf("prometheus start failure:%s", err)
			}
		}()
	})
}
