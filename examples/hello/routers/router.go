package routers

import (
	"context"
	"examples/hello/components"
	"examples/hello/handlers"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/kataras/iris/v12"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type T interface {
	InitRoute(ctx context.Context) error
}
type router struct {
	weaver.Implements[T]
	reverser weaver.Ref[components.Reverser]
}

func (r *router) InitRoute(ctx context.Context) error {
	// Get a network listener on address "localhost:12345".
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := r.Listener("hello", opts)
	if err != nil {
		r.Logger().Error(err.Error())
		return err
	}
	fmt.Printf("hello listener available on %v\n", lis)
	//
	a := iris.New()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)
	stopCh := make(chan struct{})

	c, cancel := context.WithCancel(ctx)

	go func(ctx context.Context, stopCh chan struct{}) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("IRIS")
				s := time.Since(time.Now())
				fmt.Println(s.Seconds())
				time.Sleep(5 * time.Second)
				fmt.Println(s.Seconds())
				stopCh <- struct{}{}
			}
		}
	}(c, stopCh)

	handlers.Hello(r.reverser.Get(), a)

	go func(application *iris.Application) {
		err = application.Run(iris.Listener(lis), iris.WithoutInterruptHandler)
		if err != nil {
			r.Logger().Error(err.Error())
		}
	}(a)
	<-sig
	cancel()
	<-stopCh

	return err
}

func stop(ctx context.Context) {
	s := time.Since(time.Now())
	fmt.Println(s.Seconds())
	time.Sleep(4 * time.Second)
	fmt.Println(s.Seconds())
	ctx.Deadline()
	fmt.Println(s.Seconds())

}
