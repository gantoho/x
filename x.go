package x

import (
	"fmt"
	"net/http"
)

type Context struct {
}

type X struct {
	net *http.Server
}

func (x *X) GET(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	handle(http.MethodGet, pattern, handler)
}

func (x *X) POST(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	handle(http.MethodPost, pattern, handler)
}

func handle(httpMethod string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, handler)
}

func (x *X) Run() error {
	err := x.net.ListenAndServe()
	if err != nil {
		fmt.Printf("🈲 RUN ERROR!!!")
		return err
	}
	fmt.Printf("🚩 RUN SUCCESS! PORT: 3128.")
	return err
}

func New(Addr string) *X {
	if Addr == "" {
		Addr = ":3128"
	}
	x := &X{
		net: &http.Server{
			Addr:    Addr,
			Handler: nil,
		},
	}
	return x
}
