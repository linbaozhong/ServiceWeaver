// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"context"
	"examples/hello/routers"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)
	stopCh := make(chan struct{})

	c, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context, stopCh chan struct{}) {
		for {
			select {
			case <-sig:
				fmt.Println("MAIN")
				s := time.Since(time.Now())
				fmt.Println(s.Seconds())
				time.Sleep(5 * time.Second)
				fmt.Println(s.Seconds())
				cancel()

				stopCh <- struct{}{}
			}
		}
	}(c, stopCh)

	if err := weaver.Run[*app](c, func(ctx context.Context, a *app) error {
		return a.Main(ctx)
	}); err != nil {
		log.Fatal(err)
	}

	<-stopCh
}

//go:generate ./weaver generate ./...

type app struct {
	weaver.Implements[weaver.Main]
	router weaver.Ref[routers.T]
}

func (app *app) Main(ctx context.Context) error {

	return app.router.Get().InitRoute(ctx)

}
