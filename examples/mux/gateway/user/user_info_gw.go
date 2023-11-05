package user

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/gorilla/mux"
	"net/http"
)

type userinfo struct {
	weaver.Implements[IRouter]
}

func (p userinfo) RegisterRouter(ctx context.Context) error {
	r := ctx.Value("group").(*mux.Router)
	g := r.PathPrefix("/user_info").Subrouter()
	g.HandleFunc("/get", p.get)
	return nil
}

func (p userinfo) get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("my name is userinfo"))
}
