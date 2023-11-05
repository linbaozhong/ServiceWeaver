package gateways

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/kataras/iris/v12"
	"iris/components/reverse"
	"iris/gateways/company"
	"iris/gateways/user"
	"iris/routers"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	companyApp = routers.NewApp("company")
	userApp    = routers.NewApp("user")
)

type Server struct {
	weaver.Implements[weaver.Main]
	lisCompany weaver.Listener
	lisUser    weaver.Listener

	Reverser weaver.Ref[reverse.Reverser]
}

func Run(ctx context.Context, server *Server) error {
	return nil
}

func (p *Server) Init(ctx context.Context) error {
	_ctx, cancel := context.WithCancel(ctx)

	sigs := make(chan os.Signal, 1)
	done := make(chan struct{})

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	go func() {
		<-sigs
		cancel()

		close(done)
	}()

	go p.serveCompany(_ctx)
	go p.serveUser(_ctx)

	//
	<-done
	return nil
}

func (p *Server) serveCompany(ctx context.Context) error {
	go func(ctx2 context.Context) {
		for {
			select {
			case <-ctx2.Done():
				timeout := 5 * time.Second
				_ctx, cancel := context.WithTimeout(context.Background(), timeout)
				defer cancel()

				companyApp.Application().Shutdown(_ctx)
			default:

			}
		}
	}(ctx)

	v1 := companyApp.Application().Party("/v1")
	company.Server.Register(v1, p.Reverser.Get())

	e := companyApp.Application().Run(iris.Listener(p.lisCompany),
		iris.WithLogLevel("debug"))
	if e != nil {
		p.Logger(ctx).Error(e.Error())
	}

	return nil
}

func (p *Server) serveUser(ctx context.Context) error {
	go func(ctx2 context.Context) {
		for {
			select {
			case <-ctx2.Done():
				timeout := 5 * time.Second
				_ctx, cancel := context.WithTimeout(context.Background(), timeout)
				defer cancel()

				userApp.Application().Shutdown(_ctx)
			default:

			}
		}
	}(ctx)

	v1 := userApp.Application().Party("/v1")
	user.Server.Register(v1, p.Reverser.Get())

	e := userApp.Application().Run(iris.Listener(p.lisUser),
		iris.WithLogLevel("debug"))
	if e != nil {
		p.Logger(ctx).Error(e.Error())
	}

	return nil
}
