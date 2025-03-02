package handlers

import (
	"context"
	"examples/httprouter/components/reverse"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type hello struct {
}

var revers weaver.Ref[reverse.T]

func init() {
	Instances = append(Instances, &hello{})
}

func (p *hello) RegisterRouter(g *httprouter.Router) {

	g.Handler(http.MethodGet, "/hello", weaver.InstrumentHandler("hello", http.HandlerFunc(p.hello)))
	g.Handler(http.MethodGet, "/hello/hi/:name", weaver.InstrumentHandler("hello/hi", http.HandlerFunc(p.hi)))
}

func (p *hello) hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	reversed, err := revers.Get().Reverse(context.Background(), name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!\n", reversed)
}

func (p *hello) hi(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	var name = "World"
	if len(params) > 0 {
		name = params[0].Value
	}
	reversed, err := revers.Get().Reverse(context.Background(), name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi, %s!\n", reversed)
}
