package user

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/gorilla/mux"
	"net/http"
)

type IRouter interface {
	RegisterRouter(ctx context.Context) error
}

type user struct {
	weaver.Implements[IRouter]
}

func (p user) RegisterRouter(ctx context.Context) error {
	r := ctx.Value("group").(*mux.Router)
	g := r.PathPrefix("/user").Subrouter()
	g.HandleFunc("/get", p.get)
	g.HandleFunc("/set", p.set)
	return nil
}

func (p user) get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("my name is lbz"))
}

func (p user) set(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("my name is linbaozhong"))
}
