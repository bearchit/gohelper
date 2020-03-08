package kithelper

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

func EndpointLoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log("took", time.Since(begin))
			}(time.Now())

			return next(ctx, request)
		}
	}
}

func HttpRequestLoggingMiddleware(logger log.Logger) httptransport.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		var buf bytes.Buffer
		tee := io.TeeReader(r.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		r.Body = ioutil.NopCloser(&buf)

		bodyStr := func() string {
			if len(body) <= 0 {
				return "<empty>"
			}

			return string(body)
		}

		logger.Log("method", r.Method, "uri", r.RequestURI, "body", bodyStr())
		return ctx
	}
}

func HttpResponseLoggingMiddleware(logger log.Logger) httptransport.ServerFinalizerFunc {
	return func(ctx context.Context, code int, r *http.Request) {
		logger.Log("method", r.Method, "uri", r.RequestURI, "status", code)
	}
}
