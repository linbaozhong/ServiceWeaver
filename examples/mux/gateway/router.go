package gateway

import (
	"context"
	"examples/mux/gateway/user"
	"github.com/ServiceWeaver/weaver"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	weaver.Implements[weaver.Main]
	lis      weaver.Listener
	user     weaver.Ref[user.IRouter]
	userinfo weaver.Ref[user.IRouter]
}

func (p *Server) Run(ctx context.Context) error {
	r := mux.NewRouter()

	g := r.PathPrefix("/v1").Subrouter()

	gCtx := context.WithValue(ctx, "group", g)

	p.user.Get().RegisterRouter(gCtx)
	p.userinfo.Get().RegisterRouter(gCtx)

	srv := http.Server{
		Handler: r,
	}

	return srv.Serve(p.lis)
}
