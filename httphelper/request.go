package httphelper

import "net/http"

const (
	HeaderXForwardedProto    = "X-Forwarded-Proto"
	HeaderXForwardedProtocol = "X-Forwarded-Protocol"
	HeaderXForwardedSsl      = "X-Forwarded-Ssl"
	HeaderXUrlScheme         = "X-Url-Scheme"
)

func IsTLS(r *http.Request) bool {
	return r.TLS != nil
}

func Scheme(r *http.Request) string {
	if IsTLS(r) {
		return "https"
	}

	if scheme := r.Header.Get(HeaderXForwardedProto); scheme != "" {
		return scheme
	}
	if scheme := r.Header.Get(HeaderXForwardedProtocol); scheme != "" {
		return scheme
	}
	if ssl := r.Header.Get(HeaderXForwardedSsl); ssl == "on" {
		return "https"
	}
	if scheme := r.Header.Get(HeaderXUrlScheme); scheme != "" {
		return scheme
	}

	return "httphelper"
}
