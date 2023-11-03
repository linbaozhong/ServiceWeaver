package gateways

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"iris/gateways/company"
	"iris/gateways/user"
	"os"
	"os/signal"
)

type Server struct {
	weaver.Implements[weaver.Main]
	user    weaver.Ref[user.Server]
	company weaver.Ref[company.Server]
}

func Run(ctx context.Context, server *Server) error {
	server.start(ctx)
	return nil
}

func (p *Server) start(ctx context.Context) {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Notify(sigs, os.Interrupt, os.Kill)

	go func() {
		go p.user.Get().Run(context.Background())
		go p.company.Get().Run(context.Background())

		<-sigs

		p.user.Get().Shutdown(context.Background())
		p.company.Get().Shutdown(context.Background())

		close(done)
	}()
	//优雅地关闭
	<-done
	fmt.Println("closed...")
}
