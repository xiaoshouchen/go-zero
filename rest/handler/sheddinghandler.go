package handler

import (
	"net/http"
	"sync"

	"github.com/xiaoshouchen/go-zero/core/load"
	"github.com/xiaoshouchen/go-zero/core/logx"
	"github.com/xiaoshouchen/go-zero/core/stat"
	"github.com/xiaoshouchen/go-zero/rest/httpx"
	"github.com/xiaoshouchen/go-zero/rest/internal/response"
)

const serviceType = "api"

var (
	sheddingStat *load.SheddingStat
	lock         sync.Mutex
)

// SheddingHandler returns a middleware that does load shedding.
func SheddingHandler(shedder load.Shedder, metrics *stat.Metrics) func(http.Handler) http.Handler {
	if shedder == nil {
		return func(next http.Handler) http.Handler {
			return next
		}
	}

	ensureSheddingStat()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sheddingStat.IncrementTotal()
			promise, err := shedder.Allow()
			if err != nil {
				metrics.AddDrop()
				sheddingStat.IncrementDrop()
				logx.Errorf("[http] dropped, %s - %s - %s",
					r.RequestURI, httpx.GetRemoteAddr(r), r.UserAgent())
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}

			cw := &response.WithCodeResponseWriter{Writer: w}
			defer func() {
				if cw.Code == http.StatusServiceUnavailable {
					promise.Fail()
				} else {
					sheddingStat.IncrementPass()
					promise.Pass()
				}
			}()
			next.ServeHTTP(cw, r)
		})
	}
}

func ensureSheddingStat() {
	lock.Lock()
	if sheddingStat == nil {
		sheddingStat = load.NewSheddingStat(serviceType)
	}
	lock.Unlock()
}
