package httphelper

import (
	"errors"
	"net/http"
)

type Cookier interface {
	Cookie(name string) (*http.Cookie, error)
}

func GetCookieAliases(cookier Cookier, name ...string) (*http.Cookie, error) {
	for _, x := range name {
		c, err := cookier.Cookie(x)
		if err == nil {
			return c, nil
		}
	}
	return nil, errors.New("there is no any cookie")
}
