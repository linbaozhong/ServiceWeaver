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
	"github.com/ServiceWeaver/weaver"
	"hello/routers"
	"log"
)

func main() {
	// 启动应用程序，例如 HTTP 服务等
	e := weaver.Run[*app](context.Background(), func(ctx context.Context, a *app) error {
		return a.Main(ctx)
	})
	if e != nil {
		log.Fatal(e)
	}

}

//go:generate ./weaver generate ./...

type app struct {
	weaver.Implements[weaver.Main]
	router weaver.Ref[routers.T]
}

func (app *app) Main(ctx context.Context) error {
	e := app.router.Get().InitRouter(ctx)
	if e != nil {
		app.Logger().Error(e.Error())
	}
	return e
}
