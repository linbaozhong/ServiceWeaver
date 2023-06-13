<div hidden class="todo">
TODO: Link to code snippets to make sure they are compilable and runnable.
</div>

# What is Service Weaver? 

Service Weaver is a programming framework for writing, deploying, and managing
distributed applications. You can run, test, and debug a Service Weaver application
locally on your machine, and then deploy the application to the cloud with a
single command.

Service Weaver是一个用于编写、部署和管理分布式应用程序的编程框架。您可以在本地机器上运行、测试和调试Service Weaver应用程序，然后使用一个命令将应用程序部署到云中。

```console
$ go run .                      # Run locally.
$ weaver gke deploy weaver.toml # Run in the cloud.
```

A Service Weaver application is composed of a number of **components**. A
component is represented as a regular Go [interface][go_interfaces], and
components interact with each other by calling the methods defined by these
interfaces. This makes writing Service Weaver applications easy. You don't have
to write any networking or serialization code; you just write Go. Service Weaver
also provides libraries for logging, metrics, tracing, routing, testing, and
more.

Service Weaver应用程序由许多组件组成。组件表现为常规的Go接口，组件通过调用这些接口定义的方法来相互交互。这使得编写Service Weaver应用程序变得容易。您不需要编写任何网络或序列化代码;你只要写Go。Service Weaver还提供了用于日志记录、计量、跟踪、路由、测试等的库。

You can deploy a Service Weaver application as easily as running a single command. Under
the covers, Service Weaver will dissect your binary along component boundaries, allowing
different components to run on different machines. Service Weaver will replicate,
autoscale, and co-locate these distributed components for you. It will also
manage all the networking details on your behalf, ensuring that different
components can communicate with each other and that clients can communicate with
your application.

您可以像运行单个命令一样轻松地部署Service Weaver应用程序。在幕后，Service Weaver将沿着组件边界剖析您的二进制文件，允许不同的组件在不同的机器上运行。Service Weaver将为您复制、自动伸缩和共同定位这些分布式组件。它还将代表您管理所有网络细节，确保不同的组件可以相互通信，并且客户端可以与您的应用程序通信。

Refer to the [Installation](#installation) section to install Service Weaver on
your machine, or read the [Step by Step Tutorial](#step-by-step-tutorial)
section for a tutorial on how to write Service Weaver applications.

请参阅 [Installation](#installation) 部分，在您的机器上安装Service Weaver，或者阅读 [Step by Step Tutorial](#step-by-step-tutorial)部分，了解如何编写Service Weaver应用程序。

# Installation

Ensure you have [Go installed][go_install], version 1.20 or higher. Then, run
the following to install the `weaver` command:

确保你已经 [Go installed][go_install]，版本1.20或更高。然后，执行如下命令安装weaver命令:

```console
$ go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
```

`go install` installs the `weaver` command to `$GOBIN`, which defaults to
`$GOPATH/bin`. Make sure this directory is included in your `PATH`. You can
accomplish this, for example, by adding the following to your `.bashrc` and
running `source ~/.bashrc`:

`go install` 将`weaver`命令安装到`$GOBIN`中，默认为`$GOPATH/bin`。确保这个目录包含在您的`PATH`中。你可以做到这一点，例如，通过在你的`.bashrc`中添加以下代码并运行`source ~/.bashrc`:

```console
$ export PATH="$PATH:$GOPATH/bin"
```

If the installation was successful, you should be able to run `weaver --help`:

如果安装成功，您应该能够运行`weaver --help`:

```console
$ weaver --help
USAGE

  weaver generate                 // weaver code generator
  weaver version                  // show weaver version
  weaver single    <command> ...  // for single process deployments
  weaver multi     <command> ...  // for multiprocess deployments
  ...
```

**Note**: For GKE deployments you should also install the `weaver gke` command
(see the [GKE](#gke) section for details):

**注意**:对于GKE部署，您还应该安装`weaver gke`命令(有关详细信息，请参阅 [GKE](#gke) 部分):

```console
$ go install github.com/ServiceWeaver/weaver-gke/cmd/weaver-gke@latest
```

**Note**: If you run into issues installing `weaver` and `weaver gke` commands on
macOS, you may want to prefix the install command with `export CGO_ENABLED=1; export CC=gcc`.
For example:

**注意**:如果你在macOS上安装`weaver`和`weaver gke`命令时遇到问题，你可能需要在install命令前加上`export CGO_ENABLED=1; export CC=gcc`。例如:

```console
$ export CGO_ENABLED=1; export CC=gcc; go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
```

# Step by Step Tutorial

In this section, we show you how to write Service Weaver applications. To
install Service Weaver and follow along, refer to the
[Installation](#installation) section. The full source code presented in this
tutorial can be found [here][hello_app].

在本节中，我们将向您展示如何编写`Service Weaver`应用程序。要安装`Service Weaver`并遵循本教程，请参阅 [Installation](#installation) 部分。本教程中提供的完整源代码可以在 [here][hello_app]找到。

## Components

Service Weaver's core abstraction is the **component**. A component is like an
[actor][actors], and a Service Weaver application is implemented of a set of
components. Concretely, a component is represented with a regular Go
[interface][go_interfaces], and components interact with each other by calling
the methods defined by these interfaces.

Service Weaver的核心抽象是组件。组件就像一个参与者，Service Weaver应用程序是由一组组件实现的。具体地说，组件用一个常规的Go接口表示，组件通过调用这些接口定义的方法来相互交互。

In this section, we'll define a simple `hello` component that just prints
a string and returns. First, run `go mod init hello` to create a go module.

在本节中，我们将定义一个简单的`hello`组件，它只打印字符串并返回。首先，运行go mod init hello创建一个go模块。

```console
$ mkdir hello/
$ cd hello/
$ go mod init hello
```

Then, create a file called `hello.go` with the following contents:

然后，创建一个名为 `hello.go` 的文件。内容如下:

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ServiceWeaver/weaver"
)

func main() {
    if err := weaver.Run(context.Background()); err != nil {
        log.Fatal(err)
    }
}

// app is the main component of the application. weaver.Run creates
// it and calls its Main method.
type app struct{
    weaver.Implements[weaver.Main]
}

// Main is called by weaver.Run and contains the body of the application.
func (*app) Main(context.Context) error {
    fmt.Println("Hello")
    return nil
}
```

`weaver.Run` initializes and runs the Service Weaver application.  In
particular, `weaver.Run` finds the main component, creates it, and calls its
Main method. In this example,`app` is the main component since it
contains a `weaver.Implements[weaver.Main]` field.

`weaver.Run` 初始化并运行Service Weaver应用程序。尤其是，`weaver.Run` 找到主组件，创建它，并调用它的 Main 方法。在这个例子中，`app`是主要组件，因为它包含了一个 `weaver.Implements[weaver.Main]` 域。

Before we build and run the app, we need to run Service Weaver's code generator,
called `weaver generate`. `weaver generate` writes a `weaver_gen.go` file that
contains code needed by the Service Weaver runtime. We'll elaborate on what
exactly `weaver generate` does and why we need to run it later. Finally, run the
app!

在构建和运行应用程序之前，我们需要运行Service Weaver的代码生成器，称为 `weaver generate`。 `weaver generate`写入一个包含Service Weaver运行时所需的代码的文件 `weaver_gen.go` 。我们将详细说明`weaver generate`究竟做了什么，以及为什么稍后需要运行它。最后，运行应用程序!

```console
$ go mod tidy
$ weaver generate .
$ go run .
Hello
```

Components are the core abstraction of Service Weaver. All code in a Service
Weaver application runs as part of some component. The main advantage of
components is that they decouple how you *write* your code from how you *run*
your code. They let you write your application as a monolith, but when you go to
run your code, you can run components in a separate process or on a different
machine entirely. Here's a diagram illustrating this concept:

组件是Service Weaver的核心抽象。Service Weaver应用程序中的所有代码都作为某个组件的一部分运行。组件的主要优点是它们将编写代码的方式与运行代码的方式分离开来。它们允许您将应用程序编写为一个整体，但是当您要运行代码时，您可以在单独的进程中或完全在不同的机器上运行组件。下面的图表说明了这个概念:

![A diagram showing off various types of Service Weaver deployments](assets/images/components.svg)

When we `go run` a Service Weaver application, all components run together in a
single process, and method calls between components are executed as regular Go
method calls. In a moment, we'll describe how to run each component in a
separate process with method calls between components executed as RPCs.

当我们 `go run` 一个Service Weaver应用程序时，所有组件都在单个进程中一起运行，组件之间的方法调用作为常规的go方法调用执行。稍后，我们将描述如何在单独的进程中运行每个组件，并将组件之间的方法调用作为rpc执行。

## Multiple Components

In a Service Weaver application, any component can call any other component. To
demonstrate this, we introduce a second `Reverser` component. Create a file
`reverser.go` with the following contents:

在Service Weaver应用程序中，任何组件都可以调用任何其他组件。为了演示这一点，我们引入第二个 `Reverser` 组件。创建文件 `Reverser.go` 。内容如下:

```go
package main

import (
    "context"

    "github.com/ServiceWeaver/weaver"
)

// Reverser component.
type Reverser interface {
    Reverse(context.Context, string) (string, error)
}

// Implementation of the Reverser component.
type reverser struct{
    weaver.Implements[Reverser]
}

func (r *reverser) Reverse(_ context.Context, s string) (string, error) {
    runes := []rune(s)
    n := len(runes)
    for i := 0; i < n/2; i++ {
        runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
    }
    return string(runes), nil
}
```

The `Reverser` component is represented by a `Reverser` interface with,
unsurprisingly, a `Reverse` method that reverses strings. The `reverser` struct
is our implementation of the `Reverser` component (as indicated by the
`weaver.Implements[Reverser]` field it contains).

`Reverser`组件由一个`Reverser`接口表示，毫无疑问，它带有一个`Reverse`方法来反转字符串。`reverser` struct 是我们对 `Reverser` 组件的实现(如`weaver.Implements[Reverser]`所示)。

Next, edit the app component in `main.go` to use the `Reverser` component:

接下来，在 `main.go` 中编辑app组件去使用 `Reverser` 组件:

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ServiceWeaver/weaver"
)

func main() {
    if err := weaver.Run(context.Background()); err != nil {
        log.Fatal(err)
    }
}

type app struct{
    weaver.Implements[weaver.Main]
    reverser weaver.Ref[Reverser]
}

func (app *app) Main(ctx context.Context) error {
    // Call the Reverse method.
    var r Reverser = app.reverser.Get()
    reversed, err := r.Reverse(ctx, "!dlroW ,olleH")
    if err != nil {
        return err
    }
    fmt.Println(reversed)
    return nil
}
```

The `app` struct has a new field of type `weaver.Ref[Reverser]` that provides
access to the `Reverser` component.

`app` struct 有一个新的字段，类型为 `weaver.Ref[Reverser]` ，它提供了对 `Reverser` 组件的访问。

In general, if component X uses component Y, the implementation struct for X
should contain a field of type `weaver.Ref[Y]`. When an X component instance is
created, Service Weaver will automatically create the Y component as well and
will fill the `weaver.Ref[Y]` field with a handle to the Y component.  The
implementation of X can call `Get()` on the `weaver.Ref[Y]` field to get the Y
component, as demonstrated by the following lines in the preceding examples:

通常，如果组件X使用组件Y，则X的struct实现应该包含类型为 `weaver.Ref[Y]` 的字段。当创建X组件实例时，Service Weaver也将自动创建Y组件并填充 `weaver.Ref[Y]` 字段和Y组件的句柄。X的实现可以在 `weaver.Ref[Y]` 上调用Get()方法获取Y组件，如前面示例中的以下行所示:

```go
    var r Reverser = app.reverser.Get()
    reversed, err := r.Reverse(ctx, "!dlroW ,olleH")
```

## Listeners

Service Weaver is designed for writing serving systems. In this section, we'll
augment our app to serve HTTP traffic using a network listener. Rewrite
`main.go` with the following contents:

Service Weaver是为编写服务系统而设计的。在本节中，我们将扩展我们的应用程序，使用网络侦听器来提供HTTP访问。用以下内容重写`main.go`:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "github.com/ServiceWeaver/weaver"
)

func main() {
    if err := weaver.Run(context.Background()); err != nil {
        log.Fatal(err)
    }
}

type app struct {
    weaver.Implements[weaver.Main]
    reverser weaver.Ref[Reverser]
    hello    weaver.Listener
}

func (app *app) Main(ctx context.Context) error {
    // The hello listener will listen on a random port chosen by the operating
    // system. This behavior can be changed in the config file.
    fmt.Printf("hello listener available on %v\n", app.hello)

    // Serve the /hello endpoint.
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        if name == "" {
            name = "World"
        }
        reversed, err := app.reverser.Get().Reverse(ctx, name)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "Hello, %s!\n", reversed)
    })
    return http.Serve(app.hello, nil)
}
```

Here's an explanation of the code:

- The `hello` field in the `app` struct declares a network listener, similar to
  [`net.Listen`][net_listen].
- `http.HandleFunc(...)` registers an HTTP handler for the `/hello?name=<name>`
  endpoint that returns a reversed greeting by calling the `Reverser.Reverse`
  method.
- `http.HandleFunc(...)` 为`/hello?name=<name>`注册一个HTTP处理程序，通过调用 `Reverser.Reverse` 返回反转的问候语。
- `http.Serve(lis, nil)` runs the HTTP server on the provided listener.
- `http.Serve(lis, nil)` 在提供的侦听器上运行HTTP服务器。

By default, all application listeners listen on a random port chosen by the
operating system. Here, we want to change this default behavior and assign a
fixed local listener port for the `hello` listener. To do so,
we create a [TOML](https://toml.io) config file named `weaver.toml` with
the following contents:

```toml
[listeners]
hello = {local_address = "localhost:12345"}
```

Run `go mod tidy` and then `SERVICEWEAVER_CONFIG=weaver.toml go run .`.
The program should print out the name of the application and a unique
deployment id. It should then block serving HTTP requests on `localhost:12345`.

运行 `go mod tidy` ，然后 `go run .`。程序应该打印出应用程序的名称和唯一的部署id。然后，它应该阻塞在`localhost:12345`上提供HTTP请求。

```console
$ go mod tidy
$ go run .
╭───────────────────────────────────────────────────╮
│ app        : hello                                │
│ deployment : 5c9753e4-c476-4f93-97a0-0ea599184178 │
╰───────────────────────────────────────────────────╯
hello listener available on 127.0.0.1:12345
...
```

In a separate terminal, curl the server to receive a reversed greeting:

在另外一个终端上，curl服务器接收一个反向的问候语:

```console
$ curl "localhost:12345/hello?name=Weaver"
Hello, revaeW!
```

Run `weaver single status` to view the status of the Service Weaver application.
The status shows every deployment, component, and listener.

运行`weaver single status`查看Service weaver应用程序的状态。状态显示每个部署、组件和侦听器。

```console
$ weaver single status
╭────────────────────────────────────────────────────╮
│ DEPLOYMENTS                                        │
├───────┬──────────────────────────────────────┬─────┤
│ APP   │ DEPLOYMENT                           │ AGE │
├───────┼──────────────────────────────────────┼─────┤
│ hello │ 5c9753e4-c476-4f93-97a0-0ea599184178 │ 1s  │
╰───────┴──────────────────────────────────────┴─────╯
╭────────────────────────────────────────────────────╮
│ COMPONENTS                                         │
├───────┬────────────┬────────────────┬──────────────┤
│ APP   │ DEPLOYMENT │ COMPONENT      │ REPLICA PIDS │
├───────┼────────────┼────────────────┼──────────────┤
│ hello │ 5c9753e4   │ main           │ 691625       │
│ hello │ 5c9753e4   │ hello.Reverser │ 691625       │
╰───────┴────────────┴────────────────┴──────────────╯
╭─────────────────────────────────────────────────╮
│ LISTENERS                                       │
├───────┬────────────┬──────────┬─────────────────┤
│ APP   │ DEPLOYMENT │ LISTENER │ ADDRESS         │
├───────┼────────────┼──────────┼─────────────────┤
│ hello │ 5c9753e4   │ hello    │ 127.0.0.1:12345 │
╰───────┴────────────┴──────────┴─────────────────╯
```

You can also run `weaver single dashboard` to open a dashboard in a web browser.

你也可以运行 `weaver single dashboard` 在web浏览器中打开一个仪表板。

## Multiprocess Execution

We've seen how to run a Service Weaver application in a single process with `go
run`. Now, we'll run our application in multiple processes, with method calls
between components executed as RPCs. First, create a [TOML](https://toml.io)
config file named `weaver.toml` with the following contents:

我们已经看到了如何使用 `go
run`在单个进程中运行Service Weaver应用程序。现在，我们将在多个进程中运行应用程序，组件之间的方法调用作为rpc执行。首先，创建一个名为 `weaver.toml` 的 [TOML](https://toml.io) 配置文件。内容如下:

```toml
[serviceweaver]
binary = "./hello"

[listeners]
hello = {local_address = "localhost:12345"}
```

This config file specifies the binary of the Service Weaver application, as
well as a fixed address for the hello listener. Next, build and run the app
using `weaver multi deploy`:

这个配置文件指定Service Weaver应用程序的二进制文件。接下来，使用“weaver multi deploy”构建并运行应用:

```console
$ go build                        # build the ./hello binary
$ weaver multi deploy weaver.toml # deploy the application
╭───────────────────────────────────────────────────╮
│ app        : hello                                │
│ deployment : 6b285407-423a-46cc-9a18-727b5891fc57 │
╰───────────────────────────────────────────────────╯
S1205 10:21:15.450917 stdout  26b601c4] hello listener available on 127.0.0.1:12345
S1205 10:21:15.454387 stdout  88639bf8] hello listener available on 127.0.0.1:12345
```

**Note**: `weaver multi` replicates every component twice, which is why you see
two log entries. We elaborate on replication more in the
[Components](#components) section later.

**注意**:`weaver multi` 将每个组件复制两次，这就是为什么您看到两个日志条目。我们在
[Components](#components) 小节详细说明。

In a separate terminal, curl the server:

在另外一个终端上，curl服务器：

```console
$ curl "localhost:12345/hello?name=Weaver"
Hello, revaeW!
```

When the main component receives your `/hello` HTTP request, it calls the
`reverser.Reverse` method. This method call is executed as an RPC to the
`Reverser` component running in a different process. Remember earlier when we
ran `weaver generate`, the Service Weaver code generator? One thing that `weaver
generate` does is generate RPC clients and servers for every component to make
this communication possible.

当主组件接收到你的`/hello` HTTP请求时，它调用`reverser.Reverse`方法。此方法调用作为RPC执行到在不同进程中运行的 `Reverser` 组件。还记得前面我们运行 `weaver generate` (Service weaver代码生成器)的时候吗? `weaver generate`所做的一件事是为每个组件生成RPC客户端和服务器，以使这种通信成为可能。

Run `weaver multi status` to view the status of the Service Weaver application.
Note that the `main` and `Reverser` components are replicated twice, and every
replica is run in its own OS process.

运行 `weaver multi status` 查看Service weaver应用程序的状态。请注意， `main` 组件和 `Reverser` 组件被复制两次，并且每个副本都在其自己的操作系统进程中运行。

```console
$ weaver multi status
╭────────────────────────────────────────────────────╮
│ DEPLOYMENTS                                        │
├───────┬──────────────────────────────────────┬─────┤
│ APP   │ DEPLOYMENT                           │ AGE │
├───────┼──────────────────────────────────────┼─────┤
│ hello │ 6b285407-423a-46cc-9a18-727b5891fc57 │ 3s  │
╰───────┴──────────────────────────────────────┴─────╯
╭──────────────────────────────────────────────────────╮
│ COMPONENTS                                           │
├───────┬────────────┬────────────────┬────────────────┤
│ APP   │ DEPLOYMENT │ COMPONENT      │ REPLICA PIDS   │
├───────┼────────────┼────────────────┼────────────────┤
│ hello │ 6b285407   │ main           │ 695110, 695115 │
│ hello │ 6b285407   │ hello.Reverser │ 695136, 695137 │
╰───────┴────────────┴────────────────┴────────────────╯
╭─────────────────────────────────────────────────╮
│ LISTENERS                                       │
├───────┬────────────┬──────────┬─────────────────┤
│ APP   │ DEPLOYMENT │ LISTENER │ ADDRESS         │
├───────┼────────────┼──────────┼─────────────────┤
│ hello │ 6b285407   │ hello    │ 127.0.0.1:12345 │
╰───────┴────────────┴──────────┴─────────────────╯
```

You can also run `weaver multi dashboard` to open a dashboard in a web browser.

您还可以运行 `weaver multi dashboard` 在web浏览器中打开仪表板。

## Deploying to the Cloud

The ability to run Service Weaver applications locally&mdash;either in a single
process with `go run` or across multiple processes with `weaver multi
deploy`&mdash;makes it easy to quickly develop, debug, and test your
applications. When your application is ready for production, however, you'll
often want to deploy it to the cloud. Service Weaver makes this easy too.

For example, we can deploy our "Hello, World" application to [Google Kubernetes
Engine][gke], Google Cloud's hosted Kubernetes offering, as easily as running a
single command (see the [GKE](#gke) section for details):

```console
$ weaver gke deploy weaver.toml
```

When you run this command, Service Weaver will

- wrap your application binary into a container;
- upload the container to the cloud project of your choosing;
- create and provision the appropriate Kubernetes clusters;
- set up all load balancers and networking infrastructure; and
- deploy your application on Kubernetes, with components distributed across
  machines in multiple regions.

Service Weaver also integrates your application with existing cloud tooling.
Logs are uploaded to [Google Cloud Logging][cloud_logging], metrics are uploaded
to [Google Cloud Monitoring][cloud_metrics], traces are uploaded to [Google
Cloud Tracing][cloud_trace], etc.

## Next Steps

- Continue reading to get a better understanding of [components](#components)
  and learn about other fundamental features of Service Weaver like
  [logging](#logging), [metrics](#metrics), [routing](#routing), and so on.
- Dive deeper into the various ways you can deploy a Service Weaver application,
  including [single process](#single-process), [multiprocess](#multiprocess),
  and [GKE](#gke) deployers.
- Read through [example Service Weaver applications][weaver_examples] that
  demonstrate what Service Weaver has to offer.
- Check out [Service Weaver's source code on GitHub][weaver_github].
- Read [our blog](/blog).
- Chat with us on [Discord](https://discord.gg/FzbQ3SM8R5) or send us an
  [email](serviceweaver@google.com).

# Components

**Components** are Service Weaver's core abstraction. A component is a
long-lived, possibly replicated entity that exposes a set of methods.
Concretely, a component is represented as a Go interface and corresponding
implementation of that interface. Consider the following `Adder` component for
example:

**组件**是Service Weaver的核心抽象。组件是一个持久的、可能被复制的实体，它公开了一组方法。具体地说，组件被表示为Go接口和该接口的相应实现。以下Adder组件为例:

```go
type Adder interface {
    Add(context.Context, int, int) (int, error)
}

type adder struct {
    weaver.Implements[Adder]
}

func (*adder) Add(_ context.Context, x, y int) (int, error) {
    return x + y, nil
}
```

`Adder` defines the component's interface, and `adder` defines the component's
implementation. The two are linked with the embedded `weaver.Implements[Adder]`
field. You can call `weaver.Ref[Adder].Get()` to get a client to the `Adder`
component. The returned client implements the component's interface, so you can
invoke the component's methods as you would any regular Go method. When you
invoke a component's method, the method call is performed by one of the possibly
many component replicas.

`Adder` 定义组件的接口， `adder` 定义组件的实现。两者以嵌入 `weaver.Implements[Adder]`字段相连接。你可以在客户端调用 `weaver.Ref[Adder].Get()` 获取到Adder组件。返回的客户机实现组件的接口，因此您可以像调用任何常规Go方法一样调用组件的方法。当您调用组件的方法时，该方法调用将由可能存在的多个组件副本中的一个执行。

Components are generally long-lived, but the Service Weaver runtime may scale up
or scale down the number of replicas of a component over time based on load.
Similarly, component replicas may fail and get restarted. Service Weaver may
also move component replicas around, co-locating two chatty components in the
same OS process, for example, so that communication between the components is
done locally rather than over the network.

组件通常是持久存活的，但是Service Weaver运行时可能会根据负载随时间增加或减少组件的副本数量。类似地，组件副本可能会失败并重新启动。Service Weaver还可以移动组件副本，例如，将两个聊天组件放在同一个操作系统进程中，这样组件之间的通信就可以在本地完成，而不是通过网络。

When invoking a component's method, be prepared that it may be executed via
a remote procedure call. As a result, your call may fail with a network error
instead of an application error. If you don't want to deal with network errors,
you can explicitly place the two components in the same
[colocation group](#config-files), ensuring that they always run in the
same OS process.

当调用组件的方法时，它可能会通过远程过程调用执行。因此，您的调用可能会因为网络错误而不是应用程序错误而失败。如果不想处理网络错误，可以显式地将这两个组件放在同一个主机托管[colocation group](#config-files) 中，确保它们始终在同一个操作系统进程中运行。

## Interfaces

Every method in a component interface must receive a `context.Context` as its
first argument and return an `error` as its final result. All other arguments
must be [serializable](#serializable-types). These are all valid component
methods:

组件接口中的每个方法都必须接收`context.Context`作为它的第一个参数，并返回一个 `error` 作为它的最终结果。所有其他参数必须是 [serializable](#serializable-types)的。这些都是有效的组件方法:

```go
a(context.Context) error
b(context.Context, int) error
c(context.Context) (int, error)
d(context.Context, int) (int, error)
```

These are all *invalid* component methods:

这些都是无效的组件方法:

```go
a() error                          // no context.Context argument
b(context.Context)                 // no error result
c(int, context.Context) error      // first argument isn't context.Context
d(context.Context) (error, int)    // final result isn't error
e(context.Context, chan int) error // chan int isn't serializable
```

## Implementation

A component implementation must be a struct that looks like:

组件实现必须是一个结构体，如下所示:

```go
type foo struct{
    weaver.Implements[Foo]
    // ...
}
```

-   It must be a struct.
- 它必须是一个结构体。
-   It must embed a `weaver.Implements[T]` field where `T` is the component
    interface it implements.
- 它必须嵌入一个 `weaver.Implements[T]` 字段，其中T是它实现的组件接口。

`weaver.Implements[T]` implements the `weaver.Instance` interface and therefore
every component implementation (including `foo`) also implements
`weaver.Instance`.

`weaver.Implements[T]` 实现 `weaver.Instance` 实例接口，因此每个组件实现(包括`foo`)也实现了`weaver.Instance`。

If a component implementation implements an `Init(context.Context) error`
method, it will be called when an instance of the component is created.

如果组件实现实现了 `Init(context.Context) error`方法，它将在创建组件实例时被调用。

```go
func (f *foo) Init(context.Context) error {
    // ...
}
```

## Semantics

When implementing a component, there are three semantic details to keep in mind:

在实现组件时，需要记住三个语义细节:

1.  A component's state is not persisted.
2.  A component's methods may be invoked concurrently.
3.  There may be multiple replicas of a component.


1. 组件的状态不会被持久化。
2. 组件的方法可以并发调用。
3. 一个组件可能有多个副本。

Take the following `Cache` component for example, which maintains an in-memory
key-value cache.

以下面的 `Cache` 组件为例，它维护内存中的键值缓存。

```go
type Cache interface {
    Put(ctx context.Context, key, value string) error
    Get(ctx context.Context, key string) (string, error)
}

type cache struct {
    mu sync.Mutex
    data map[string]string
}

func (c *Cache) Put(_ context.Context, key, value string) error {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
    return nil
}

func (c *Cache) Get(_ context.Context, key string) (string, error) {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.data[key], nil
}
```

Noting the points above:

注意到以上几点:

1.  A `Cache`'s state is not persisted, so if a `Cache` replica fails, its data
    is lost. Any state that needs to be persisted should be persisted
    explicitly.
2.  A `Cache`'s methods may be invoked concurrently, so it's essential that we
    guard access to `data` with the mutex `mu`.
3.  There may be multiple replicas of a `Cache` component, so it is not
    guaranteed that one client's `Get` will be routed to the same replica as
    another client's `Put`. For this example, this means that the `Cache` has
    [weak consistency][weak_consistency].


1. `Cache`的状态不是持久化的，所以如果`Cache`副本失败了，它的数据将被丢弃。任何需要持久化的状态都应该显式持久化。
2. `Cache`的方法可能被并发调用，所以我们必须用互斥锁mu保护对`data`的访问
3. `Cache`组件可能有多个副本，因此不能保证一个客户机的 `Get` 将被路由到与另一个客户机的`Put`相同的副本。对于本例，这意味着缓存具有弱一致性。

If a remote method call fails to execute properly&mdash;because of a machine
crash or a network partition, for example&mdash;it returns an error with an
embedded `weaver.RemoteCallError`. Here's an illustrative example:

如果远程方法调用无法正常执行(例如，由于机器崩溃或网络分区)，它将返回一个带有嵌入式 `weaver.RemoteCallError`的错误。这里有一个说明性的例子:

```go
// Call the cache.Get method.
value, err := cache.Get(ctx, "key")
if errors.Is(err, weaver.RemoteCallError) {
    // cache.Get did not execute properly.
} else if err != nil {
    // cache.Get executed properly, but returned an error.
} else {
    // cache.Get executed properly and did not return an error.
}
```

Note that if a method call returns an error with an embedded
`weaver.RemoteCallError`, it does *not* mean that the method never executed. The
method may have executed partially or fully. Thus, you must be careful retrying
method calls that result in a `weaver.RemoteCallError`. Ensuring that all
methods are either read-only or idempotent is one way to ensure safe retries,
for example. Service Weaver does not automatically retry method calls that fail.

## Listeners

A component implementation may wish to use one or more network listeners, e.g.,
to serve HTTP network traffic. To do so, named `weaver.Listener` fields must
be added to the implementation struct. For example, the following component
implementation creates two network listeners:

```go
type impl struct{
    weaver.Implements[MyComponent]
    foo weaver.Listener
    Bar weaver.Listener
}
```

With Service Weaver, listeners are named. By default, listeners are named
after their corresponding struct fields (e.g., `"foo"` and `"bar"` in the
above example). Alternatively, a special ````weaver:"name"```` struct tag
can be added to the struct field to specify the listener name explicitly:

```go
type impl struct{
    weaver.Implements[MyComponent]
    foo weaver.Listener
    lis weaver.Listener `weaver:"bar"`
}
```

Listener names are always lowercased and must be unique inside a given
application binary, regardless of which components they are specified in. For
example, it is illegal to declare a Listener field `"foo"` in two different
component implementations structs, unless one is renamed using the
````weaver:"name"```` struct tag.

By default, all application listeners will listen on a random port chosen
by the operating system. This behavior, as well as other customization options,
can be controlled in the [configuration file](#config-file). For example, the
following config will assign addresses `"localhost:12345"` and
`"localhost:12346"` to `"foo"` and `"bar"`, respectively.

```toml
[listeners]
foo = {local_address = "localhost:12345"}
bar = {local_address = "localhost:12346"}
```

## Config

Service Weaver uses [config files](#config-files), written in [TOML](#toml), to
configure how applications are run. A minimal config file, for example, simply
lists the application binary:

Service Weaver使用用 [config files](#config-files) 编写 [TOML](#toml) 配置文件来配置应用程序的运行方式。例如，一个最小的配置文件只是列出了应用程序的二进制文件:

```toml
[serviceweaver]
binary = "./hello"
```

A config file may additionally contain listener-specific configuration sections,
which allow you to configure the [network listeners](#components-listeners) in
your application. For example, consider the listener `"foo"` declared in
a component implementation.

```go
type impl struct{
    weaver.Implements[MyComponent]
    foo weaver.Listener
}
```

By default, listener `"foo"` will listen on a random port chosen by the
operating system. To configure it to listen on a specific local port, we can
add the following section to the config file:

```toml
[listeners]
foo = {local_address = "localhost:12345}
```

A config file may also contain component-specific configuration
sections, which allow you to configure the components in your application. For
example, consider the following `Greeter` component.

配置文件可能还包含特定于组件的配置部分，这些部分允许您配置应用程序中的组件。例如，参考以下的Greeter组件。

```go
type Greeter interface {
    Greet(context.Context, string) (string, error)
}

type greeter struct {
    weaver.Implements[Greeter]
}

func (g *greeter) Greet(_ context.Context, name string) (string, error) {
    return fmt.Sprintf("Hello, %s!", name), nil
}
```

Rather than hard-coding the greeting `"Hello"`, we can provide a greeting in a
config file. First, we define a options struct.

我们可以在配置文件中提供问候语，而不是硬编码 `"Hello"`。首先，我们定义一个options结构体。

```go
type greeterOptions struct {
    Greeting string
}
```

Next, we associate the options struct with the `greeter` implementation by
embedding the `weaver.WithConfig[T]` struct.

接下来，我们通过嵌入 `weaver.WithConfig[T]` struct与带有`greeter`实现的options struct联系起来。

```go
type greeter struct {
    weaver.Implements[Greeter]
    weaver.WithConfig[greeterOptions]
}
```

Now, we can add a `Greeter` section to the config file. The section is keyed by
the full path-prefixed name of the component.

现在，我们可以在配置文件中添加一个 `Greeter` 部分。该部分由组件的完整路径前缀名称作为关键字。

```toml
["example.com/mypkg/Greeter"]
Greeting = "Bonjour"
```

When the `Greeter` component is created, Service Weaver will automatically parse
the `Greeter` section of the config file into a `greeterOptions` struct. You can
access the populated struct via the `Config` method of the embedded `WithConfig`
struct. For example:

创建 `Greeter` 组件时，Service Weaver会自动将配置文件中的 `Greeter` 部分解析为 `greeterOptions` struct。您可以通过嵌入的`WithConfig` struct的 `Config` 方法访问填充的结构体。例如:

```go
func (g *greeter) Greet(_ context.Context, name string) (string, error) {
    greeting := g.Config().Greeting
    if greeting == "" {
        greeting = "Hello"
    }
    return fmt.Sprintf("%s, %s!", greeting, name), nil
}
```

<div hidden class="todo">
    Move the next part to the Single Process section and forward link.
</div>

If you run an application directly (i.e. using `go run`), you can pass the
config file using the `SERVICEWEAVER_CONFIG` environment variable:

如果你直接运行应用程序(即使用 `go run`)，你可以使用 `SERVICEWEAVER_CONFIG` 环境变量传递配置文件:

```console
$ SERVICEWEAVER_CONFIG=weaver.toml go run .
```

Or, use `weaver single deploy`:

```console
$ weaver single deploy weaver.toml
```

# Logging

<div hidden class="todo">
TODO(mwhittaker): Pick a better name for node ids?
</div>

Service Weaver provides a logging API, `weaver.Logger`. By using Service
Weaver's logging API, you can cat, tail, search, and filter logs from every one
of your Service Weaver applications (past or present). Service Weaver also
integrates the logs into the environment where your application is deployed. If
you [deploy a Service Weaver application to Google Cloud](#gke), for example,
logs are automatically exported to [Google Cloud Logging][cloud_logging].

Service Weaver提供了一个日志API `weaver.Logger`。通过使用Service Weaver的日志API，您可以记录、跟踪、搜索和过滤来自每个Service Weaver应用程序(过去或现在)的日志。Service Weaver还将日志集成到部署应用程序的环境中。例如，如果您将Service Weaver应用程序部署到Google Cloud，则日志将自动导出到Google Cloud Logging。

Use the `Logger` method of a component implementation to get a logger scoped to
the component. For example:

使用组件实现的 `Logger` 方法来获取组件范围内的日志记录器。例如:

```go
type Adder interface {
    Add(context.Context, int, int) (int, error)
}

type adder struct {
    weaver.Implements[Adder]
}

func (a *adder) Add(_ context.Context, x, y int) (int, error) {
    // adder embeds weaver.Implements[Adder] which provides the Logger method.
    logger := a.Logger()
    logger.Debug("A debug log.")
    logger.Info("An info log.")
    logger.Error("An error log.", fmt.Errorf("an error"))
    return x + y, nil
}
```

Logs look like this:

日志像这样的:

```console
D1103 08:55:15.650138 main.Adder 73ddcd04 adder.go:12] A debug log.
I1103 08:55:15.650149 main.Adder 73ddcd04 adder.go:13] An info log.
E1103 08:55:15.650158 main.Adder 73ddcd04 adder.go:14] An error log. err="an error"
```

The first character of a log line indicates whether the log is a [D]ebug,
[I]nfo, or [E]rror log entry. Then comes the date in `MMDD` format, followed by
the time. Then comes the component name followed by a logical node id. If two
components are co-located in the same OS process, they are given the same node
id. Then comes the file and line where the log was produced, followed finally by
the contents of the log.

日志行的第一个字符表示该日志是[D]ebug、[I] info还是[E] error日志条目。然后是MMDD格式的日期，后面跟着时间。然后是组件名和逻辑节点id。如果两个组件位于同一操作系统进程中，则为它们分配相同的节点id。然后是生成日志的文件和行，最后是日志的内容。

Service Weaver also allows you to attach key-value attributes to log entries.
These attributes can be useful when searching and filtering logs.

Service Weaver还允许您将键值属性附加到日志条目。这些属性在搜索和过滤日志时非常有用。

```go
logger.Info("A log with attributes.", "foo", "bar")  // adds foo="bar"
```

If you find yourself adding the same set of key-value attributes repeatedly, you
can pre-create a logger that will add those attributes to all log entries:

如果你发现自己反复添加相同的键值属性集，你可以预先创建一个记录器，将这些属性添加到所有日志条目:

```go
fooLogger = logger.With("foo", "bar")
fooLogger.Info("A log with attributes.")  // adds foo="bar"
```

**Note**: You can also add normal print statements to your code. These prints
will be captured and logged by Service Weaver, but they won't be associated with
a particular component, they won't have `file:line` information, and they won't
have any attributes, so we recommend you use a `weaver.Logger` whenever
possible.

**注意**:您还可以在代码中添加普通的打印语句。这些打印将被Service Weaver捕获和记录，但它们不会与特定的组件相关联，它们不会有 `file:line` 信息，也不会有任何属性，因此我们建议您使用 `weaver.Logger` 。

```console
S1027 14:40:55.210541 stdout d772dcad] This was printed by fmt.Println
```

Refer to the deployer-specific documentation to learn how to search and filter
logs for [single process](#single-process-logging),
[multiprocess](#multiprocess-logging), and [GKE](#gke-logging) deployments.

# Metrics

Service Weaver provides an API for [metrics][metric_types]; specifically
[counters][prometheus_counter], [gauges][prometheus_gauge], and
[histograms][prometheus_histogram].

- A **counter** is a number that can only increase over time. It never
  decreases. You can use a counter to measure things like the number of HTTP
  requests your program has processed so far.
- A **gauge** is a number that can increase *or* decrease over time. You can use
  a gauge to measure things like the current amount of memory your program is
  using, in bytes.
- A **histogram** is a collection of numbers that are grouped into buckets. You
  can use a histogram to measure things like the latency of every HTTP request
  your program has received so far.

Service Weaver integrates these metrics into the environment where your application is
deployed. If you [deploy a Service Weaver application to Google Cloud](#gke), for
example, metrics are automatically exported to the [Google Cloud Metrics
Explorer][metrics_explorer] where they can be queried, aggregated, and graphed.

Here's an example of how to add metrics to a simple `Adder` component.

```go
var (
    addCount = metrics.NewCounter(
        "add_count",
        "The number of times Adder.Add has been called",
    )
    addConcurrent = metrics.NewGauge(
        "add_concurrent",
        "The number of concurrent Adder.Add calls",
    )
    addSum = metrics.NewHistogram(
        "add_sum",
        "The sums returned by Adder.Add",
        []float64{1, 10, 100, 1000, 10000},
    )
)

type Adder interface {
    Add(context.Context, int, int) (int, error)
}

type adder struct {
    weaver.Implements[Adder]
}

func (*adder) Add(_ context.Context, x, y int) (int, error) {
    addCount.Add(1.0)
    addConcurrent.Add(1.0)
    defer addConcurrent.Sub(1.0)
    addSum.Put(float64(x + y))
    return x + y, nil
}
```

Refer to the deployer-specific documentation to learn how to view metrics for
[single process](#single-process-metrics), [multiprocess](#multiprocess-metrics),
and [GKE](#gke-metrics) deployments.

## Labels

Metrics can also have a set of key-value labels. Service Weaver represents
labels using structs. Here's an example of how to declare and use a labeled
counter to count the parity of the argument to a `Halve` method.

```go
type halveLabels struct {
    Parity string // "odd" or "even"
}

var (
    halveCounts = metrics.NewCounterMap[halveLabels](
        "halve_count",
        "The number of values that have been halved",
    )
    oddCount = halveCounts.Get(halveLabels{"odd"})
    evenCount = halveCounts.Get(halveLabels{"even"})
)

type Halver interface {
    Halve(context.Context, int) (int, error)
}

type halver struct {
    weaver.Implements[Halver]
}

func (halver) Halve(_ context.Context, val int) (int, error) {
    if val % 2 == 0 {
        evenCount.Add(1)
    } else {
        oddCount.Add(1)
    }
    return val / 2, nil
}
```

To adhere to [popular metric naming conventions][prometheus_naming], Service
Weaver lowercases the first letter of every label by default. The `Parity` field
for example is exported as `parity`. You can override this behavior and provide
a custom label name using a `weaver` annotation.

```go
type labels struct {
    Foo string                           // exported as "foo"
    Bar string `weaver:"my_custom_name"` // exported as "my_custom_name"
}
```

## Auto-Generated Metrics

Service Weaver automatically creates and maintains the following set of metrics, which
measure the count, latency, and chattiness of every remote component method
invocation. Every metric is labeled by the calling component as well as the
invoked component and method.

-   `serviceweaver_remote_method_count`: Count of Service Weaver component
    method invocations.
-   `serviceweaver_remote_method_error_count`: Count of Service Weaver component
    method invocations that result in an error.
-   `serviceweaver_remote_method_latency_micros`: Duration, in microseconds, of
    Service Weaver component method execution.
-   `serviceweaver_remote_method_bytes_request`: Number of bytes in Service
    Weaver component method requests.
-   `serviceweaver_remote_method_bytes_reply`: Number of bytes in Service Weaver
    component method replies.

**Note**: These metrics only measure *remote* method calls. Local method calls,
like those between two co-located components, are not measured.

## HTTP Metrics

Service Weaver declares the following set of HTTP related metrics.

-   `serviceweaver_http_request_count`: Count of HTTP requests.
-   `serviceweaver_http_error_count`: Count of HTTP requests resulting in a 4XX or 5XX
    response. This metric is also labeled with the returned status code.
-   `serviceweaver_http_request_latency_micros`: Duration, in microseconds, of HTTP
    request execution.
-   `serviceweaver_http_request_bytes_received`: Estimated number of bytes *received* by
    an HTTP handler.
-   `serviceweaver_http_request_bytes_returned`: Estimated number of bytes *returned* by
    an HTTP handler.

If you pass an [`http.Handler`](https://pkg.go.dev/net/http#Handler) to the
`weaver.InstrumentHandler` function, it will return a new `http.Handler` that
updates these metrics automatically, labeled with the provided label. For
example:

```go
// Metrics are recorded for fooHandler with label "foo".
var mux http.ServeMux
var fooHandler http.Handler = ...
mux.Handle("/foo", weaver.InstrumentHandler("foo", fooHandler))
```

# Tracing

Service Weaver relies on [OpenTelemetry][otel] to trace your application.
Service Weaver exports these traces into the environment where your application
is deployed. If you [deploy a Service Weaver application to Google Cloud](#gke),
for example, traces are automatically exported to [Google Cloud
Trace][cloud_trace]. Here's an example of how to enable tracing for a simple
`Hello, World!` application.

```go
import (
    "context"
    "fmt"
    "log"
    "net/http"

    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
    "github.com/ServiceWeaver/weaver"
)

func main() {
    if err := weaver.Run(context.Background()); err != nil {
        log.Fatal(err)
    }
}

type app struct {
    weaver.Implements[weaver.Main]
    lis weaver.Listener
}

func (app *app) Main(ctx context.Context) error {
    fmt.Printf("hello listener available on %v\n", app.lis)

    // Serve the /hello endpoint.
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %s!\n", r.URL.Query().Get("name"))
    })

    // Create an otel handler to enable tracing.
    otelHandler := otelhttp.NewHandler(http.DefaultServeMux, "http")
    return http.Serve(lis, otelHandler)
}
```

This code does the following:

- `http.HandleFunc("/hello", ...)` registers a handler with the default HTTP
  mux, called `http.DefaultServeMux`.
- `otelhttp.NewHandler(http.DefaultServeMux, "http")` returns a new HTTP handler
  that wraps the default HTTP mux.
- `http.Serve(lis, otelHandler)` serves HTTP traffic on `lis` using the
  OpenTelemetry handler.

Using the OpenTelemetry HTTP handler enables tracing. Once tracing is enabled,
all HTTP requests and resulting component method calls will be automatically
traced. Service Weaver will collect and export the traces for you. Refer to the
deployer-specific documentation for [single process](#single-process-tracing),
[multiprocess](#multiprocess-tracing), and [GKE](#gke-tracing) to learn about
deployer specific exporters.

The step above is all you need to get started with tracing. If you want to add
more application-specific details to your traces, you can add attributes,
events, and errors using the context passed to registered HTTP handlers and
component methods. For example, in our `hello` example, you can add an event as
follows:

```go
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!\n", r.URL.Query().Get("name"))
    trace.SpanFromContext(r.Context()).AddEvent("writing response",
        trace.WithAttributes(
            label.String("content", "hello "),
            label.String("answer", r.URL.Query().Get("name")),
        ))
})
```

Refer to [OpenTelemetry Go: All you need to know][otel_all_you_need] to learn
more about how to add more application-specific details to your traces.

# Profiling

Service Weaver allows you to profile an entire Service Weaver application, even one that is
deployed in multiple processes across multiple machines. Service Weaver profiles every
individual binary and aggregates them into a single profile that captures the
performance of the application as a whole. Refer to the deployer-specific
documentation for details on how to collect profiles for [single
process](#single-process-profiling), [multiprocess](#multiprocess-profiling),
and [GKE](#gke-profiling) deployments.

# Routing

By default, when a client invokes a remote component's method, this method call
will be performed by one of possibly many component replicas, selected
arbitrarily. It is sometimes beneficial for method invocations to be routed to
*a particular* replica based on the arguments provided to the method. For
example, consider a `Cache` component that maintains an in-memory cache in front
of an underlying disk-backed key-value store:

默认情况下，当客户端调用远程组件的方法时，该方法调用将由任意选择的多个组件副本之一执行。根据提供给方法的参数将方法调用路由到*特定的*副本有时是有益的。例如，考虑一个 `Cache` 组件，它在底层磁盘支持的键值存储前面维护内存中的缓存:

```go
type Cache interface {
    Get(ctx context.Context, key string) (string, error)
    Put(ctx context.Context, key, value string) error
}

type cache struct {
    weaver.Implements[Cache]
    // ...
}
```

To increase the cache hit ratio, we may want to route every request for a given
key to the same replica. Service Weaver supports this affinity based routing by allowing
the application to specify a router type associated with the component
implementation. For example:

为了提高缓存命中率，我们可能希望将对给定键的每个请求路由到相同的副本。Service Weaver通过允许应用程序指定与组件实现相关联的路由器类型来支持这种基于关联的路由。例如:

```go
type cacheRouter struct{}
func (cacheRouter) Get(_ context.Context, key string) string { return key }
func (cacheRouter) Put(_ context.Context, key, value string) string { return key }
```

For every component method that needs to be routed (e.g., `Get` and `Put`), the
router type should implement an equivalent method (i.e., same name and
argument types) whose return type is the routing key. When a component's routed
method is invoked, its corresponding router method is invoked to produce a
routing key. Method invocations that produce the same key are routed to the same
replica.

对于每个需要路由的组件方法(例如`Get`和`Put`)，路由器类型应该实现一个等价的方法(例如，相同的名称和参数类型)，其返回类型是路由键。当一个组件的路由方法被调用时，它对应的路由器方法被调用来产生一个路由密钥。产生相同键的方法调用被路由到相同的副本。

A routing key can be

路由密钥可以是

-   any integer (e.g., `int`, `int32`), float (i.e. `float32`, `float64`), or
    string; or
-   a struct where every field is an integer, float, or string (e.g., `struct{x
    int; y string}`).

Every router method must return the same routing key type. The following, for
example, is invalid:

每个路由器方法必须返回相同的路由密钥类型。例如:

```go
// ERROR: Get returns a string, but Put returns an int.
func (cacheRouter) Get(_ context.Context, key string) string { return key }
func (cacheRouter) Put(_ context.Context, key, value string) int { return 42 }
```

To associate a router with its component, embed a `weaver.WithRouter[T]` field in
the component implementation where `T` is the type of the router.

要将路由器与其组件关联，需要嵌入一个 `weaver.WithRouter[T]` 字段，其中T是路由器的类型。

```go
type cache struct {
    weaver.Implements[Cache]
    weaver.WithRouter[cacheRouter]
    // ...
}
```

**NOTE**: Routing is done on a best-effort basis. Service Weaver will try to route
method invocations with the same key to the same replica, but this is *not*
guaranteed. As a corollary, you should *never* depend on routing for
correctness. Only use routing to increase performance in the common case.

**注意**:路由是在尽力而为的基础上完成的。Service Weaver将尝试将具有相同键的方法调用路由到相同的副本，但这*不能*保证。因此，*永远*不应该依赖路由的正确性。在一般情况下，只使用路由来提高性能。

Also note that if a component invokes a method on a co-located component, the
method call will always be executed by the co-located component and won't be
routed.

还要注意的是，如果组件调用位于同一位置的组件上的方法，则该方法调用将始终由位于同一位置的组件执行，而不会被路由。

# Storage

We expect most Service Weaver applications to persist their data in some way. For
example, an e-commerce application may store its products catalog and user
information in a database and access them while serving user requests.

我们期望大多数Service Weaver应用程序以某种方式持久化它们的数据。例如，电子商务应用程序可以将其产品目录和用户信息存储在数据库中，并在服务用户请求时访问它们。

By default, Service Weaver leaves the storage and retrieval of application data
up to the developer. If you're using a database, for example, you have to create
the database, pre-populate it with data, and write the code to access the
database from your Service Weaver application.

默认情况下，Service Weaver将应用程序数据的存储和检索留给开发人员。例如，如果您正在使用数据库，则必须创建数据库，用数据预先填充它，并编写代码以从Service Weaver应用程序访问数据库。

Below is an example of how database information can be passed to a simple
`Adder` component using a [config file](#components-config). First, the config
file:

下面是如何使用配置文件将数据库信息传递给简单的`Adder`组件的示例。首先，配置文件:

```toml
["example.com/mypkg/Adder"]
Driver = "mysql"
Source = "root:@tcp(localhost:3306)/"
```

And the application that uses it:

```go
type Adder interface {
    Add(context.Context, int, int) (int, error)
}

type adder struct {
    weaver.Implements[Adder]
    weaver.WithConfig[config]

    db *sql.DB
}

type config struct {
    Driver string // Name of the DB driver.
    Source string // DB data source.
}

func (a *adder) Init(_ context.Context) error {
    db, err := sql.Open(a.Config().Driver, a.Config().Source)
    a.db = db
    return err
}

func (a *Adder) Add(ctx context.Context, x, y int) (int, error) {
    // Check in the database first.
    var sum int
    const q = "SELECT sum FROM table WHERE x=? AND y=?;"
    if err := a.db.QueryRowContext(ctx, q, x, y).Scan(&sum); err == nil {
        return sum, nil
    }

    // Make a best-effort attempt to store in the database.
    q = "INSERT INTO table(x, y, sum) VALUES (?, ?, ?);"
    a.db.ExecContext(ctx, q, x, y, x + y)
    return x + y, nil
}
```

A similar process can be followed to pass database information using Go flags or
environment variables.

可以遵循类似的方式来使用Go flags 或环境变量传递数据库信息。

# Testing

Service Weaver includes a `weavertest` package that you can use to test your
Service Weaver applications. The package provides a `Runner` type with `Test`
and `Bench` methods. Tests use `Runner.Test` instead of `weaver.Run`. To test an
`Adder` component with an `Add` method, for example, create an `adder_test.go`
file with the following contents.

```go
package main

import (
    "context"
    "testing"

    "github.com/ServiceWeaver/weaver"
    "github.com/ServiceWeaver/weaver/weavertest"
)

func TestAdd(t *testing.T) {
     runner := weavertest.Local  // A runner that runs components in a single process
     runner.Test(t, func(t *testing.T, adder Adder) {
         ctx := context.Background()
         got, err := adder.Add(ctx, 1, 2)
         if err != nil {
             t.Fatal(err)
         }
         if want := 3; got != want {
             t.Fatalf("got %q, want %q", got, want)
         }
     })
}
```

Run `go test` to run the test. `runner.Test` will create a sub-test and within
it will create an `Adder` component and pass it to the supplied function. Tests
that want to exercise multiple components can pass a function with a separate
argument per component. Each of those components will be created and passed to
the function.

```go
func TestArithmetic(t *testing.T) {
    weavertest.Local.Test(t, func(t *testing.T, adder Adder, multiplier Multiplier) {
        // ...
    })
}
```

`weavertest` provides a set of builtin Runners that differ in how they partition
components across processes and how the components communicate with each other:

1. **weavertest.Local**: Every component will be placed in the test process, and
   all component method calls will use local procedure calls, happens when you
   `go run` a Service Weaver application.
2. **weavertest.Multi**: Every component will be placed in a
   different process. This is similar to what happens when you run `weaver multi
   deploy`.
3. **weavertest.RPC**: Every component will be placed in the test process, but
   all component method calls will use remote even though the callee is
   local. This mode is most useful when collecting profiles or coverage data.

Tests run using `weavertest.Local` are easier to debug and troubleshoot, but do
not test distributed execution. You should test with different runners to get
the best of both worlds (each Runner.Test call will create a new sub-test):

```go
func TestAdd(t *testing.T) {
    for _, runner := range weavertest.AllRunners() {
        runner.Test(t, func(t *testing.T, adder Adder) {
            // ...
        })
    }
}
```

You can also provide the contents of a [config file](#config-files) to a runner
by setting the `Runner.Config` field:

```go
func TestArithmetic(t *testing.T) {
    runner := weavertest.Local()
    runner.Name = "Custom"
    runner.Config = `[serviceweaver] ...`
    runner.Test(t, func(t *testing.T, adder Adder, multiplier Multiplier) {
        // ...
    })
}
```

# Versioning

Serving systems evolve over time. Whether you're fixing bugs or adding new
features, it is inevitable that you will have to roll out a new version of your
system to replace the currently running version. To maintain the availability of
their systems, people typically perform **rolling updates**, where the nodes in
a deployment are updated from the old version to the new version one by one.

During a rolling update, nodes running the old version of the code will have to
communicate with other nodes running the new version of the code. Ensuring that
a system is correct despite the possibility of these cross-version interactions
is very challenging. In
[*Understanding and Detecting Software Upgrade Failures in Distributed Systems*][update_failures_paper],
Zhang et al. perform a case study of 123 failed updates in 8 widely used
systems. They found that the majority of failures were caused by the
interactions between multiple versions of a system:

>    _About two thirds of update failures are caused by interaction between two
>    software versions that hold incompatible data syntax or semantics
>    assumption._


Service Weaver takes a different approach to rollouts and sidesteps these
complex cross-version interactions. Service Weaver ensures that client requests
are executed entirely within a single version of a system. A component in one
version will *never* communicate with a component in a different version. This
eliminates the leading cause of update failures, allowing you to roll out new
versions of your Service Weaver application safely and with less headache.

Avoiding cross-version communication is trivial for applications deployed using
[`go run`](#single-process) or [`weaver multi deploy`](#multiprocess) because
every deployment runs independently from one another. Refer to the
[GKE Deployments](#gke-multi-region) and
[GKE Versioning](#gke-versioning) sections to learn how Service Weaver uses a combination
of [blue/green deployments][blue_green] and autoscaling to slowly shift traffic
from an old version of a Service Weaver application running on GKE to a new version,
avoiding cross-version communication in a resource-efficient manner.

# Single Process

## Getting Started

The simplest and easiest way to deploy a Service Weaver application is to run it
directly via `go run`. When you `go run` a Service Weaver application, every
component is co-located in a single process, and method calls between components
are executed as regular Go method calls. Refer to the [Step by Step
Tutorial](#step-by-step-tutorial) section for a full example.

```console
$ go run .
```

If you run an application using `go run`, you can provide a config file using
the `SERVICEWEAVER_CONFIG` environment variable:

```console
$ SERVICEWEAVER_CONFIG=weaver.toml go run .
```

Or, you can use the `weaver single deploy` command. `weaver single deploy` is
practically identical to `go run .`, but it makes it easier to provide a config
file.

```console
$ weaver single deploy weaver.toml
```

You can run `weaver single status` to view the status of all active Service
Weaver applications deployed using `go run`.

```console
$ weaver single status
╭────────────────────────────────────────────────────╮
│ DEPLOYMENTS                                        │
├───────┬──────────────────────────────────────┬─────┤
│ APP   │ DEPLOYMENT                           │ AGE │
├───────┼──────────────────────────────────────┼─────┤
│ hello │ a4bba25b-6312-4af1-beec-447c33b8e805 │ 26s │
│ hello │ a4d4c71b-a99f-4ade-9586-640bd289158f │ 19s │
│ hello │ bc663a25-c70e-440d-b022-04a83708c616 │ 12s │
╰───────┴──────────────────────────────────────┴─────╯
╭─────────────────────────────────────────────────────╮
│ COMPONENTS                                          │
├───────┬────────────┬─────────────────┬──────────────┤
│ APP   │ DEPLOYMENT │ COMPONENT       │ REPLICA PIDS │
├───────┼────────────┼─────────────────┼──────────────┤
│ hello │ a4bba25b   │ main            │ 123450       │
│ hello │ a4bba25b   │ hello.Reverser  │ 123450       │
│ hello │ a4d4c71b   │ main            │ 903510       │
│ hello │ a4d4c71b   │ hello.Reverser  │ 903510       │
│ hello │ bc663a25   │ main            │ 489102       │
│ hello │ bc663a25   │ hello.Reverser  │ 489102       │
╰───────┴────────────┴─────────────────┴──────────────╯
╭────────────────────────────────────────────╮
│ LISTENERS                                  │
├───────┬────────────┬──────────┬────────────┤
│ APP   │ DEPLOYMENT │ LISTENER │ ADDRESS    │
├───────┼────────────┼──────────┼────────────┤
│ hello │ a4bba25b   │ hello    │ [::]:33541 │
│ hello │ a4d4c71b   │ hello    │ [::]:41619 │
│ hello │ bc663a25   │ hello    │ [::]:33319 │
╰───────┴────────────┴──────────┴────────────╯
```

You can also run `weaver single dashboard` to open a dashboard in a web browser.

## Listeners

You can add `weaver.Listener` fields to the component implementation to trigger
creation of network listeners (see the
[Step by Step Tutorial](#step-by-step-tutorial) section for context).

```go
type app struct {
    weaver.Implements[weaver.Main]
    hello    weaver.Listener
}
```

When you deploy an application using `go run`, the network listeners will be
automatically created by the Service Weaver runtime. Each listener will listen
on a random port chosen by the operating system, unless concrete local addresses
have been specified in the [config file](#components-config).

## Logging

When you deploy a Service Weaver application with `go run`, [logs](#logging) are
printed to standard out. These logs are not persisted. You can optionally save
the logs for later analysis using basic shell constructs:

```console
$ go run . | tee mylogs.txt
```

## Metrics

Run `weaver single dashboard` to open a dashboard in a web browser. The
dashboard has a page for every Service Weaver application deployed via `go run
.`.  Every deployment's page has a link to the deployment's [metrics](#metrics).
The metrics are exported in [Prometheus format][prometheus] and looks something
like this:

```txt
# Metrics in Prometheus text format [1].
#
# To visualize and query the metrics, make sure Prometheus is installed on
# your local machine and then add the following stanza to your Prometheus yaml
# config file:
#
# scrape_configs:
# - job_name: 'prometheus-serviceweaver-scraper'
#   scrape_interval: 5s
#   metrics_path: /debug/serviceweaver/prometheus
#   static_configs:
#     - targets: ['127.0.0.1:43087']
#
# [1]: https://prometheus.io

# HELP example_count An example counter.
# TYPE example_count counter
example_count{serviceweaver_node="bbc9beb5"} 42
example_count{serviceweaver_node="00555c38"} 9001

# ┌─────────────────────────────────────┐
# │ SERVICEWEAVER AUTOGENERATED METRICS │
# └─────────────────────────────────────┘
# HELP serviceweaver_method_count Count of Service Weaver component method invocations
# TYPE serviceweaver_method_count counter
serviceweaver_method_count{caller="main",component="main.Example",serviceweaver_node="9fa07495",method="Foo"} 0
serviceweaver_method_count{caller="main",component="main.Example",serviceweaver_node="ee76816d",method="Foo"} 1
...
```

As the header explains, you can visualize and query the metrics by installing
Prometheus and configuring it, using the provided stanza, to periodically scrape
the `/debug/serviceweaver/prometheus` endpoint of the provided target
(`127.0.0.1:43087` in the example above). You can also inspect the metrics
manually. The metrics page shows the latest value of every metric in your
application followed by [the metrics that Service Weaver automatically creates
for you](#metrics-auto-generated-metrics).

## Profiling

Use the `weaver single profile` command to collect a profile of your Service Weaver
application. Invoke the command with the id of your deployment. For example,
imagine you `go run` your Service Weaver application and it gets a deployment id
`28807368-1101-41a3-bdcb-9625e0f02ca0`.

```console
$ go run .
╭───────────────────────────────────────────────────╮
│ app        : hello                                │
│ deployment : 28807368-1101-41a3-bdcb-9625e0f02ca0 │
╰───────────────────────────────────────────────────╯
```

In a separate terminal, you can run the `weaver single profile` command.

```console
$ weaver single profile 28807368               # Collect a CPU profile.
$ weaver single profile --duration=1m 28807368 # Adjust the duration of the profile.
$ weaver single profile --type=heap 28807368   # Collect a heap profile.
```

`weaver single profile` prints out the filename of the collected profile. You can
use the `go tool pprof` command to visualize and analyze the profile. For
example:

```console
$ profile=$(weaver single profile <deployment>) # Collect the profile.
$ go tool pprof -http=localhost:9000 $profile   # Visualize the profile.
```

Refer to `weaver single profile --help` for more details. Refer to `go tool pprof
--help` for more information on how to use pprof to analyze your profiles. Refer
to [*Profiling Go Programs*][pprof_blog] for a tutorial.

## Tracing

Run `weaver single dashboard` to open a dashboard in a web browser. The
dashboard has a page for every Service Weaver application deployed via `go run
.`.  Every deployment's page has a link to the deployment's [traces](#tracing)
accessible via [Perfetto][perfetto]. Here's an example of what the tracing page
looks like:

![An example trace page](assets/images/trace_single.png)

Refer to [Perfetto UI Docs](https://perfetto.dev/docs/visualization/perfetto-ui)
to learn more about how to use the tracing UI.

# Multiprocess

## Getting Started

You can use `weaver multi` to deploy a Service Weaver application across
multiple processes on your local machine, with each component replica running in
a separate OS process. Create [a config file](#config-files), say `weaver.toml`,
that points to your compiled Service Weaver application.

```toml
[serviceweaver]
binary = "./your_compiled_serviceweaver_binary"
```

Deploy the application using `weaver multi deploy`:

```console
$ weaver multi deploy weaver.toml
```

Refer to the [Step by Step Tutorial](#step-by-step-tutorial) section for a full
example.

When `weaver multi deploy` terminates (e.g., when you press `ctrl+c`), the
application is destroyed and all processes are terminated.

You can run `weaver multi status` to view the status of all active Service Weaver
applications deployed using `weaver multi`.

```console
$ weaver multi status
╭────────────────────────────────────────────────────╮
│ DEPLOYMENTS                                        │
├───────┬──────────────────────────────────────┬─────┤
│ APP   │ DEPLOYMENT                           │ AGE │
├───────┼──────────────────────────────────────┼─────┤
│ hello │ a4bba25b-6312-4af1-beec-447c33b8e805 │ 26s │
│ hello │ a4d4c71b-a99f-4ade-9586-640bd289158f │ 19s │
│ hello │ bc663a25-c70e-440d-b022-04a83708c616 │ 12s │
╰───────┴──────────────────────────────────────┴─────╯
╭───────────────────────────────────────────────────────╮
│ COMPONENTS                                            │
├───────┬────────────┬─────────────────┬────────────────┤
│ APP   │ DEPLOYMENT │ COMPONENT       │ REPLICA PIDS   │
├───────┼────────────┼─────────────────┼────────────────┤
│ hello │ a4bba25b   │ main            │ 695110, 695115 │
│ hello │ a4bba25b   │ hello.Reverser  │ 193720, 398751 │
│ hello │ a4d4c71b   │ main            │ 847020, 292745 │
│ hello │ a4d4c71b   │ hello.Reverser  │ 849035, 897452 │
│ hello │ bc663a25   │ main            │ 245702, 157455 │
│ hello │ bc663a25   │ hello.Reverser  │ 997520, 225023 │
╰───────┴────────────┴─────────────────┴────────────────╯
╭────────────────────────────────────────────╮
│ LISTENERS                                  │
├───────┬────────────┬──────────┬────────────┤
│ APP   │ DEPLOYMENT │ LISTENER │ ADDRESS    │
├───────┼────────────┼──────────┼────────────┤
│ hello │ a4bba25b   │ hello    │ [::]:33541 │
│ hello │ a4d4c71b   │ hello    │ [::]:41619 │
│ hello │ bc663a25   │ hello    │ [::]:33319 │
╰───────┴────────────┴──────────┴────────────╯
```

You can also run `weaver multi dashboard` to open a dashboard in a web browser.

## Listeners

You can add `weaver.Listener` fields to the component implementation to trigger
creation of network listeners (see the
[Step by Step Tutorial](#step-by-step-tutorial) section for context).

```go
type app struct {
    weaver.Implements[weaver.Main]
    hello    weaver.Listener
}
```

When you deploy an application using `weaver multi deploy`, the network
listeners will be automatically created by the Service Weaver runtime.
In particular, for each listener specified in the application binary,
the runtime:

1. Creates a localhost network listener listening on a random port chosen
   by the operating system (i.e. listening on `localhost:0`).
2. Ensures that an HTTP proxy is created. This proxy forwards traffic to the
   listener. In fact, the proxy balances traffic across every replica of the
   listener. (Recall that components may be replicated, and so every component
   replica will have a different instance of the listener.)
   The proxy address is by default `localhost:0`, or the local address assigned
   to the listener in the config file, if any.

## Logging

`weaver multi deploy` logs to stdout. It additionally persists all log entries in
a set of files in `/tmp/serviceweaver/logs/weaver-multi`. Every file contains a stream of
log entries encoded as protocol buffers. You can cat, follow, and filter these
logs using `weaver multi logs`. For example:

```shell
# Display all of the application logs
weaver multi logs

# Follow all of the logs (similar to tail -f).
weaver multi logs --follow

# Display all of the logs for the "todo" app.
weaver multi logs 'app == "todo"'

# Display all of the debug logs for the "todo" app.
weaver multi logs 'app=="todo" && level=="debug"'

# Display all of the logs for the "todo" app in files called foo.go.
weaver multi logs 'app=="todo" && source.contains("foo.go")'

# Display all of the logs that contain the string "error".
weaver multi logs 'msg.contains("error")'

# Display all of the logs that match a regex.
weaver multi logs 'msg.matches("error: file .* already closed")'

# Display all of the logs that have an attribute "foo" with value "bar".
weaver multi logs 'attrs["foo"] == "bar"'

# Display all of the logs in JSON format. This is useful if you want to
# perform some sort of post-processing on the logs.
weaver multi logs --format=json

# Display all of the logs, including internal system logs that are hidden by
# default.
weaver multi logs --system
```

Refer to `weaver multi logs --help` for a full explanation of the query language,
along with many more examples.

## Metrics

Run `weaver multi dashboard` to open a dashboard in a web browser. The dashboard
has a page for every Service Weaver application deployed via `weaver muli
deploy`.  Every deployment's page has a link to the deployment's
[metrics](#metrics). The metrics are exported in [Prometheus
format][prometheus] and looks something like this:

```txt
# Metrics in Prometheus text format [1].
#
# To visualize and query the metrics, make sure Prometheus is installed on
# your local machine and then add the following stanza to your Prometheus yaml
# config file:
#
# scrape_configs:
# - job_name: 'prometheus-serviceweaver-scraper'
#   scrape_interval: 5s
#   metrics_path: /debug/serviceweaver/prometheus
#   static_configs:
#     - targets: ['127.0.0.1:43087']
#
#
# [1]: https://prometheus.io

# HELP example_count An example counter.
# TYPE example_count counter
example_count{serviceweaver_node="bbc9beb5"} 42
example_count{serviceweaver_node="00555c38"} 9001

# ┌─────────────────────────────────────┐
# │ SERVICEWEAVER AUTOGENERATED METRICS │
# └─────────────────────────────────────┘
# HELP serviceweaver_method_count Count of Service Weaver component method invocations
# TYPE serviceweaver_method_count counter
serviceweaver_method_count{caller="main",component="main.Example",serviceweaver_node="9fa07495",method="Foo"} 0
serviceweaver_method_count{caller="main",component="main.Example",serviceweaver_node="ee76816d",method="Foo"} 1
...
```

As the header explains, you can visualize and query the metrics by installing
Prometheus and configuring it, using the provided stanza, to periodically scrape
the `/debug/serviceweaver/prometheus` endpoint of the provided target (e.g.,
`127.0.0.1:43087`). You can also inspect the metrics manually. The metrics page
shows the latest value of every metric in your application followed by [the
metrics that Service Weaver automatically creates for
you](#metrics-auto-generated-metrics).

## Profiling

Use the `weaver multi profile` command to collect a profile of your Service Weaver
application. Invoke the command with the id of your deployment. For example,
imagine you `weaver multi deploy` your Service Weaver application and it gets a deployment
id `28807368-1101-41a3-bdcb-9625e0f02ca0`.

```console
$ weaver multi deploy weaver.toml
╭───────────────────────────────────────────────────╮
│ app        : hello                                │
│ deployment : 28807368-1101-41a3-bdcb-9625e0f02ca0 │
╰───────────────────────────────────────────────────╯
```

In a separate terminal, you can run the `weaver multi profile` command.

```console
$ weaver multi profile 28807368               # Collect a CPU profile.
$ weaver multi profile --duration=1m 28807368 # Adjust the duration of the profile.
$ weaver multi profile --type=heap 28807368   # Collect a heap profile.
```

`weaver multi profile` prints out the filename of the collected profile. You can
use the `go tool pprof` command to visualize and analyze the profile. For
example:

```console
$ profile=$(weaver multi profile <deployment>) # Collect the profile.
$ go tool pprof -http=localhost:9000 $profile # Visualize the profile.
```

Refer to `weaver multi profile --help` for more details. Refer to `go tool pprof
--help` for more information on how to use pprof to analyze your profiles. Refer
to [*Profiling Go Programs*][pprof_blog] for a tutorial.

## Tracing

Run `weaver multi dashboard` to open a dashboard in a web browser. The
dashboard has a page for every Service Weaver application deployed via
`weaver multi deploy`. Every deployment's page has a link to the deployment's
[traces](#tracing) accessible via [Perfetto][perfetto]. Here's an example of
what the tracing page looks like:

![An example trace page](assets/images/trace_multi.png)

Trace events are grouped by colocation group and their corresponding replicas.
Each event has a label associated with it, based on whether the event was due to
a local or remote call. Note that the user can filter the set of events for a
particular trace by clicking on an event's `traceID` and choosing `Find slices
with the same arg value`.

Refer to [Perfetto UI Docs](https://perfetto.dev/docs/visualization/perfetto-ui)
to learn more about how to use the tracing UI.

# GKE

[Google Kubernetes Engine (GKE)][gke] is a Google Cloud managed service that
implements the full [Kubernetes][kubernetes] API. It supports autoscaling and
multi-cluster development, and allows you to run containerized applications in
the cloud.

You can use `weaver gke` to deploy a Service Weaver application to GKE, with components
running on different machines across multiple cloud regions. The `weaver gke`
command does a lot of the heavy lifting to set up GKE on your behalf. It
containerizes your application; it creates the appropriate GKE clusters; it
plumbs together all the networking infrastructure; and so on. This makes
deploying your Service Weaver application to the cloud as easy as running `weaver gke
deploy`. In this section, we show you how to deploy your application using
`weaver gke`. Refer to the [Local GKE](#local-gke) section to see how to simulate
a GKE deployment locally on your machine.

## Installation

First, [ensure you have Service Weaver installed](#installation). Next, install
the `weaver-gke` command:

```console
$ go install github.com/ServiceWeaver/weaver-gke/cmd/weaver-gke@latest
```

Install the `gcloud` command to your local machine. To do so, follow [these
instructions][gcloud_install], or run the following command and follow its
prompts:

```console
$ curl https://sdk.cloud.google.com | bash
```

After installing `gcloud`, install the required GKE authentication plugin:

```console
$ gcloud components install gke-gcloud-auth-plugin
```

, and then run the following command to initialize your local environment:

```console
$ gcloud init
```

The above command will prompt you to select the Google account and cloud project
you wish to use. If you don't have a cloud project, the command will prompt you
to create one. Make sure to select a unique project name or the command will
fail. If that happens, follow [these instructions][gke_create_project] to create
a new project, or simply run:

```console
$ gcloud projects create my-unique-project-name
```

Before you can use your cloud project, however, you must add a billing account
to it. Go to [this page][gcloud_billing] to create a new billing account, and
[this page][gcloud_billing_projects] to associate a billing account with your
cloud project.

## Getting Started

Consider again the "Hello, World!" Service Weaver application from the [Step by
Step Tutorial](#step-by-step-tutorial) section. The application runs an HTTP
server on a listener named `hello` with a `/hello?name=<name>` endpoint that
returns a `Hello, <name>!` greeting. To deploy this application to GKE, first
create a [Service Weaver config file](#config-files), say `weaver.toml`, with
the following contents:

```toml
[serviceweaver]
binary = "./hello"

[gke]
regions = ["us-west1"]
public_listener = [
  {name = "hello", hostname = "hello.com"},
]
```

The `[serviceweaver]` section of the config file specifies the compiled Service
Weaver binary. The `[gke]` section configures the regions where the application
is deployed (`us-west1` in this example). It also declares which listeners
should be **public**, i.e., which listeners should be accessible from the public
internet. By default, all listeners are **private**, i.e., accessible only from
the cloud project's internal network. In our example, we declare that the
`hello` listener is public.

All listeners deployed to GKE are configured to be health-checked by GKE
load-balancers on the `/debug/weaver/healthz` URL path. ServiceWeaver
automatically registers a health-check handler under this URL path in the
default ServerMux, so the `hello` application requires no changes.

Deploy the application using `weaver gke deploy`:

```console
$ GOOS=linux GOARCH=amd64 go build
$ weaver gke deploy weaver.toml
...
Deploying the application... Done
Version "8e1c640a-d87b-4020-b3dd-4efc1850756c" of app "hello" started successfully.
Note that stopping this binary will not affect the app in any way.
Tailing the logs...
...
```

The first time you deploy a Service Weaver application to a cloud project, the process
may be slow, since Service Weaver needs to configure your cloud project, create the
appropriate GKE clusters, etc. Subsequent deployments should be significantly
faster.

When `weaver gke` deploys your application, it creates a global, externally
accessibly load balancer that forwards traffic to the public listeners in your
application. `weaver gke deploy` prints out the IP address of this load balancer
as well as instructions on how to interact with it:

```text
NOTE: The applications' public listeners will be accessible via an
L7 load-balancer managed by Service Weaver running at the public IP address:

    http://34.149.225.62

This load-balancer uses hostname-based routing to route request to the
appropriate listeners. As a result, all HTTP(s) requests reaching this
load-balancer must have the correct "Host" header field populated. This can be
achieved in one of two ways:
...
```

For an application running in production, you will likely want to configure DNS
to map your domain name (e.g. `hello.com`), to the address of the load balancer
(e.g., `http://34.149.225.62`). When testing and debugging an application,
however, we can also simply curl the load balancer with the appropriate hostname
header. Since we configured our application to associate host name `hello.com`
with the `hello` listener, we use the following command:

```console
$ curl --header 'Host: hello.com' "http://34.149.225.63/hello?name=Weaver"
Hello, Weaver!
```

We can inspect the Service Weaver applications running on GKE using the `weaver gke
status` command.

```console
$ weaver gke status
╭───────────────────────────────────────────────────────────────╮
│ Deployments                                                   │
├───────┬──────────────────────────────────────┬───────┬────────┤
│ APP   │ DEPLOYMENT                           │ AGE   │ STATUS │
├───────┼──────────────────────────────────────┼───────┼────────┤
│ hello │ 20c1d756-80b5-42a7-9e73-b0d3e717516e │ 1m10s │ ACTIVE │
╰───────┴──────────────────────────────────────┴───────┴────────╯
╭──────────────────────────────────────────────────────────╮
│ COMPONENTS                                               │
├───────┬────────────┬──────────┬────────────────┬─────────┤
│ APP   │ DEPLOYMENT │ LOCATION │ COMPONENT      │ HEALTHY │
├───────┼────────────┼──────────┼────────────────┼─────────┤
│ hello │ 20c1d756   │ us-west1 │ hello.Reverser │ 2/2     │
│ hello │ 20c1d756   │ us-west1 │ main           │ 2/2     │
╰───────┴────────────┴──────────┴────────────────┴─────────╯
╭─────────────────────────────────────────────────────────────────────────────────────╮
│ TRAFFIC                                                                             │
├───────────┬────────────┬───────┬────────────┬──────────┬─────────┬──────────────────┤
│ HOST      │ VISIBILITY │ APP   │ DEPLOYMENT │ LOCATION │ ADDRESS │ TRAFFIC FRACTION │
├───────────┼────────────┼───────┼────────────┼──────────┼─────────┼──────────────────┤
│ hello.com │ public     │ hello │ 20c1d756   │ us-west1 │         │ 0.5              │
├───────────┼────────────┼───────┼────────────┼──────────┼─────────┼──────────────────┤
│ hello.com │ public     │ hello │ 20c1d756   │ us-west1 │         │ 0.5              │
╰───────────┴────────────┴───────┴────────────┴──────────┴─────────┴──────────────────╯
╭────────────────────────────╮
│ ROLLOUT OF hello           │
├─────────────────┬──────────┤
│                 │ us-west1 │
├─────────────────┼──────────┤
│ TIME            │ 20c1d756 │
│ Feb 27 21:23:07 │ 1.00     │
╰─────────────────┴──────────╯
```

`weaver gke status` reports information about every app, deployment, component,
and listener in your cloud project. In this example, we have a single deployment
(with id `20c1d756`) of the `hello` app. Our app has two components (`main` and
`hello.Reverser`) each with two healthy replicas running in the `us-west1`
region. The two replicas of the `main` component each export a `hello` listener.
The global load balancer that we curled earlier balances traffic evenly across
these two listeners. The final section of the output details the rollout
schedule of the application. We'll discuss rollouts later in the
[Rollouts](#gke-multi-region) section. You can also run `weaver gke dashboard`
to open a dashboard in a web browser.

<div hidden class="todo">
TODO(mwhittaker): Remove rollout section?
</div>

**Note**: `weaver gke` configures GKE to autoscale your application. As the load
on your application increases, the number of replicas of the overloaded
components will increase. Conversely, as the load on your application decreases,
the number of replicas decreases. Service Weaver can independently scale the different
components of your application, meaning that heavily loaded components can be
scaled up while lesser loaded components can simultaneously be scaled down.

You can use the `weaver gke kill` command to kill your deployed application.

```console
$ weaver gke kill hello
WARNING: You are about to kill every active deployment of the "hello" app.
The deployments will be killed immediately and irrevocably. Are you sure you
want to proceed?

Enter (y)es to continue: y
```

## Logging

`weaver gke deploy` logs to stdout. It additionally exports all log entries to
[Cloud Logging][cloud_logging].  You can cat, follow, and filter these logs from
the command line using `weaver gke logs`. For example:

```shell
# Display all of the application logs
weaver gke logs

# Follow all of the logs (similar to tail -f).
weaver gke logs --follow

# Display all of the logs for the "todo" app.
weaver gke logs 'app == "todo"'

# Display all of the debug logs for the "todo" app.
weaver gke logs 'app=="todo" && level=="debug"'

# Display all of the logs for the "todo" app in files called foo.go.
weaver gke logs 'app=="todo" && source.contains("foo.go")'

# Display all of the logs that contain the string "error".
weaver gke logs 'msg.contains("error")'

# Display all of the logs that match a regex.
weaver gke logs 'msg.matches("error: file .* already closed")'

# Display all of the logs that have an attribute "foo" with value "bar".
weaver gke logs 'attrs["foo"] == "bar"'

# Display all of the logs in JSON format. This is useful if you want to
# perform some sort of post-processing on the logs.
weaver gke logs --format=json

# Display all of the logs, including internal system logs that are hidden by
# default.
weaver gke logs --system
```

Refer to `weaver gke logs --help` for a full explanation of the query language,
along with many more examples.

You can also run `weaver gke dashboard` to open a dashboard in a web browser.
The dashboard has a page for every Service Weaver application deployed via
`weaver gke deploy`. Every deployment's page has a link to the deployment's logs
on [Google Cloud's Logs Explorer][logs_explorer] as shown below.

![A screenshot of Service Weaver logs in the Logs Explorer](assets/images/logs_explorer.png)

## Metrics

`weaver gke` exports metrics to the
[Google Cloud Monitoring console][cloud_metrics]. You can view and graph these
metrics using the [Cloud Metrics Explorer][metrics_explorer]. When you open the
Metrics Explorer, click `SELECT A METRIC`.

![A screenshot of the Metrics Explorer](assets/images/cloud_metrics_1.png)

All Service Weaver metrics are exported under the `custom.googleapis.com` domain. Query
for `serviceweaver` to view these metrics and select the metric you're interested in.

![A screenshot of selecting a metric in Metrics Explorer](assets/images/cloud_metrics_2.png)

You can use the Metrics Explorer to graph the metric you selected.

![A screenshot of a metric graph in Metrics Explorer](assets/images/cloud_metrics_3.png)

Refer to the [Cloud Metrics][cloud_metrics] documentation for more information.

## Profiling

Use the `weaver gke profile` command to collect a profile of your Service Weaver
application. Invoke the command with the name (and optionally version) of the
app you wish to profile. For example:

```console
# Collect a CPU profile of the latest version of the hello app.
$ weaver gke profile hello

# Collect a CPU profile of a specific version of the hello app.
$ weaver gke profile --version=8e1c640a-d87b-4020-b3dd-4efc1850756c hello

# Adjust the duration of a CPU profile.
$ weaver gke profile --duration=1m hello

# Collect a heap profile.
$ weaver gke profile --type=heap hello
```

`weaver gke profile` prints out the filename of the collected profile. You can
use the `go tool pprof` command to visualize and analyze the profile. For
example:

```console
$ profile=$(weaver gke profile <app>)         # Collect the profile.
$ go tool pprof -http=localhost:9000 $profile # Visualize the profile.
```

Refer to `weaver gke profile --help` for more details.

## Tracing

Run `weaver gke dashboard` to open a dashboard in a web browser. The
dashboard has a page for every Service Weaver application deployed via
`weaver gke deploy`. Every deployment's page has a link to the deployment's
[traces](#tracing) accessible via [Google Cloud Trace][trace_service] as shown
below.

![A screenshot of a Google Cloud Trace page](assets/images/trace_gke.png)

## Multi-Region

`weaver gke` allows you to deploy a Service Weaver application to multiple
[cloud regions](https://cloud.google.com/compute/docs/regions-zones). Simply
include the regions where you want to deploy in your config file. For example:

```toml
[gke]
regions = ["us-west1", "us-east1", "asia-east2", "europe-north1"]
```

When `weaver gke` deploys an application to multiple regions, it intentionally
does not deploy the application to every region right away. Instead, it performs
a **slow rollout** of the application. `weaver gke` first deploys the application
to a small subset of the regions, which act as [canaries][canary]. The
application runs in the canary clusters for some time before being rolled out to
a larger subset of regions. `weaver gke` continues this incremental
rollout---iteratively increasing the number of regions where the application is
deployed---until the application has been rolled out to every region specified
in the config file. Within each region, `weaver gke` also slowly shifts traffic
from old application versions to new versions. We discuss this in
[the next section](#versioning).

By slowly rolling out an application across regions, `weaver gke` allows you to
catch buggy releases early and mitigate the amount of damage they can cause. The
`rollout` field in a [config file](#config-files) determines the length of a
slow rollout. For example:

```toml
[serviceweaver]
rollout = "1h" # Perform a one hour slow rollout.
...
```

<div hidden class="todo">
TODO(mwhittaker): Remove this part?
</div>

You can monitor the rollout of an application using `weaver gke status`. For
example, here is the rollout schedule produced by `weaver gke status` for a one
hour deployment of the `hello` app across the us-central1, us-west1, us-south1,
and us-east1 regions.

```console
[ROLLOUT OF hello]
                 us-west1  us-central1  us-south1  us-east1
TIME             a838cf1d  a838cf1d     a838cf1d   a838cf1d
Nov  8 22:47:30  1.00      0.00         0.00       0.00
        +15m00s  0.50      0.50         0.00       0.00
        +30m00s  0.33      0.33         0.33       0.00
        +45m00s  0.25      0.25         0.25       0.25
```

Every row in the schedule shows the fraction of traffic each region receives
from the global load balancer. The top row is the current traffic assignment,
and each subsequent row shows the projected traffic assignment at some point in
the future. Noting that only regions with a deployed application receive
traffic, we can see the application is initially deployed in us-west1, then
slowly rolls out to us-central1, us-south1, and us-east1 in 15 minute
increments.

Also note that while the global load balancer balances traffic across regions,
once a request is received within a region, it is processed entirely within that
region. As with slow rollouts and canarying, avoiding cross-region communication
is a form of [isolation][isolation] that helps minimize the blast radius of a
misbehaving application.

## Versioning

To roll out a new version of your application as a replacement of an existing
version, simply rebuild your application and run `weaver gke deploy` again.
`weaver gke` will slowly roll out the new version of the application to the
regions provided in the config file, as described in the previous section. In
addition to slowly rolling out *across* regions, `weaver gke` also slowly rolls
out *within* regions. Within each region, `weaver gke` updates the global load
balancer to slowly shift traffic from the old version of the application to the
new version.

<div hidden class="todo">
TODO(mwhittaker): Remove this part?
</div>

We can again use `weaver gke status` to monitor the rollout of a new application
version. For example, here is the rollout schedule produced by `weaver gke
status` for a one hour update of the `hello` app across the us-west1 and
us-east1 regions. The new version of the app `45a521a3` is replacing the old
version `def1f485`.

```console
[ROLLOUT OF hello]
                 us-west1  us-west1  us-east1  us-east1
TIME             def1f485  45a521a3  def1f485  45a521a3
Nov  9 00:54:59  0.45      0.05      0.50      0.00
         +4m46s  0.38      0.12      0.50      0.00
         +9m34s  0.25      0.25      0.50      0.00
        +14m22s  0.12      0.38      0.50      0.00
        +19m10s  0.00      0.50      0.50      0.00
        +29m58s  0.00      0.50      0.45      0.05
        +34m46s  0.00      0.50      0.38      0.12
        +39m34s  0.00      0.50      0.25      0.25
        +44m22s  0.00      0.50      0.12      0.38
        +49m10s  0.00      0.50      0.00      0.50
```

Every row in the schedule shows the fraction of traffic that every deployment
receives in every region. The schedule shows that the new application is rolled
out in us-west1 before us-east1. Initially, the new version receives
increasingly more traffic in the us-west1 region, transitioning from 5% of the
global traffic (10% of the us-west1 traffic) to 50% of the global traffic (100%
of the us-west1 traffic) over the course of roughly 20 minutes. Ten minutes
later, this process repeats in us-east1 over the course of another 20 minutes
until the new version is receiving 100% of the global traffic. After the full
one hour rollout is complete, the old version is considered obsolete and is
deleted automatically.

**Note**: While the load balancer balances traffic across application versions,
once a request is received, it is processed entirely by the version that
received it. There is no cross-version communication.

Superficially, `weaver gke`'s rollout scheme seems to require a lot of resources
because it runs two copies of the application side-by-side. In reality,
`weaver gke`'s use of autoscaling makes this type of
[blue/green rollout][blue_green] resource efficient. As traffic is shifted away
from the old version, its load decreases, and the autoscaler reduces its
resource allocation. Simultaneously, as the new version receives more traffic,
its load increases, and the autoscaler begins to increase its resource
allocation. These two transitions cancel out causing the rollout to use a
roughly constant number of resources.

<div hidden class="todo">
TODO(mwhittaker): What if the new version doesn't have the same regions as
the old version? Explain what happens in this case.
</div>

## Config

You can configure `weaver gke` using the `[gke]` section of a
[config file](#config-files).

```toml
[gke]
project = "my-google-cloud-project"
account = "my_account@gmail.com"
regions = ["us-west1", "us-east1"]
public_listener = [
    {name = "cat", hostname = "cat.com"},
    {name = "hat", hostname = "hat.gg"},
]
```

| Field | Required? | Description |
| --- | --- | --- |
| project | optional | Name of the Google Cloud Project in which to deploy the Service Weaver application. If absent, the currently active project is used (i.e. `gcloud config get-value project`) |
| account | required | Google Cloud account used to deploy the Service Weaver application. If absent, the currently active account is used (i.e. `gcloud config get-value account`). |
| regions | optional | Regions in which the Service Weaver application should be deployed. Defaults to `["us-west1"]`. |
| public_listener | optional | The application's public listeners along with their corresponding hostnames. |

# Local GKE

[`weaver gke`](#gke) lets you deploy Service Weaver applications to GKE. `weaver gke-local`
is a drop-in replacement for `weaver gke` that allows you to simulate GKE
deployments locally on your machine. Every `weaver gke` command can be replaced
with an equivalent `weaver gke-local` command. `weaver gke deploy` becomes
`weaver gke-local deploy`; `weaver gke status` becomes `weaver gke-local status`;
and so on. `weaver gke-local` runs your components in simulated GKE clusters and
launches a local proxy to emulate GKE's global load balancer. `weaver gke-local`
also uses [the same config as a `weaver gke`](#gke-config), meaning that after you
test your application locally using `weaver gke-local`, you can deploy the same
application to GKE without any code *or* config changes.

## Installation

First, [ensure you have Service Weaver installed](#installation). Next, install
the `weaver-gke-local` command:

```console
$ go install github.com/ServiceWeaver/weaver-gke/cmd/weaver-gke-local@latest
```

## Getting Started

In the [`weaver gke`](#gke-getting-started) section, we deployed a "Hello,
World!" application to GKE using `weaver gke deploy`. We can deploy the same app
locally using `weaver gke-local deploy`:

```console
$ cat weaver.toml
[serviceweaver]
binary = "./hello"

[gke]
regions = ["us-west1"]
public_listener = [
  {name = "hello", hostname = "hello.com"},
]

$ weaver gke-local deploy weaver.toml
Deploying the application... Done
Version "a2bc7a7a-fcf6-45df-91fe-6e6af171885d" of app "hello" started successfully.
Note that stopping this binary will not affect the app in any way.
Tailing the logs...
...
```

You can run `weaver gke-local status` to check the status of all the applications
deployed using `weaver gke-local`.

```console
$ weaver gke-local status
╭─────────────────────────────────────────────────────────────╮
│ Deployments                                                 │
├───────┬──────────────────────────────────────┬─────┬────────┤
│ APP   │ DEPLOYMENT                           │ AGE │ STATUS │
├───────┼──────────────────────────────────────┼─────┼────────┤
│ hello │ af09030c-b3a6-4d15-ba47-cd9e9e9ec2e7 │ 13s │ ACTIVE │
╰───────┴──────────────────────────────────────┴─────┴────────╯
╭──────────────────────────────────────────────────────────╮
│ COMPONENTS                                               │
├───────┬────────────┬──────────┬────────────────┬─────────┤
│ APP   │ DEPLOYMENT │ LOCATION │ COMPONENT      │ HEALTHY │
├───────┼────────────┼──────────┼────────────────┼─────────┤
│ hello │ af09030c   │ us-west1 │ hello.Reverser │ 2/2     │
│ hello │ af09030c   │ us-west1 │ main           │ 2/2     │
╰───────┴────────────┴──────────┴────────────────┴─────────╯
╭─────────────────────────────────────────────────────────────────────────────────────────────╮
│ TRAFFIC                                                                                     │
├───────────┬────────────┬───────┬────────────┬──────────┬─────────────────┬──────────────────┤
│ HOST      │ VISIBILITY │ APP   │ DEPLOYMENT │ LOCATION │ ADDRESS         │ TRAFFIC FRACTION │
├───────────┼────────────┼───────┼────────────┼──────────┼─────────────────┼──────────────────┤
│ hello.com │ public     │ hello │ af09030c   │ us-west1 │ 127.0.0.1:46539 │ 0.5              │
│ hello.com │ public     │ hello │ af09030c   │ us-west1 │ 127.0.0.1:43439 │ 0.5              │
╰───────────┴────────────┴───────┴────────────┴──────────┴─────────────────┴──────────────────╯
╭────────────────────────────╮
│ ROLLOUT OF hello           │
├─────────────────┬──────────┤
│                 │ us-west1 │
├─────────────────┼──────────┤
│ TIME            │ af09030c │
│ Feb 27 20:33:10 │ 1.00     │
╰─────────────────┴──────────╯
```

The output is, unsurprisingly, identical to that of `weaver gke status`. There is
information about every app, component, and listener. Note that for this
example, `weaver gke-local` is running the "Hello, World!" application in a fake
us-west1 "region", as specified in the `weaver.toml` config file.

`weaver gke-local` runs a proxy on port 8000 that simulates the global load
balancer used by `weaver gke`. We can curl the proxy in the same way we curled
the global load balancer. Since we configured our application to associate host
name `hello.com` with the `hello` listener, we use the following command:

```console
$ curl --header 'Host: hello.com' "localhost:8000/hello?name=Weaver"
Hello, Weaver!
```

You can use the `weaver gke-local kill` command to kill your deployed
application.

```console
$ weaver gke-local kill hello
WARNING: You are about to kill every active deployment of the "hello" app.
The deployments will be killed immediately and irrevocably. Are you sure you
want to proceed?

Enter (y)es to continue: y
```

<div hidden class="todo">
TODO(mwhittaker): Have `weaver gke-local` print instructions on how to curl the
proxy.
</div>

## Logging

`weaver gke-local deploy` logs to stdout. It additionally persists all log
entries in a set of files in `/tmp/serviceweaver/logs/weaver-gke-local`. Every file
contains a stream of log entries encoded as protocol buffers. You can cat,
follow, and filter these logs using `weaver gke-local logs`. For example:

```shell
# Display all of the application logs
weaver gke-local logs

# Follow all of the logs (similar to tail -f).
weaver gke-local logs --follow

# Display all of the logs for the "todo" app.
weaver gke-local logs 'app == "todo"'

# Display all of the debug logs for the "todo" app.
weaver gke-local logs 'app=="todo" && level=="debug"'

# Display all of the logs for the "todo" app in files called foo.go.
weaver gke-local logs 'app=="todo" && source.contains("foo.go")'

# Display all of the logs that contain the string "error".
weaver gke-local logs 'msg.contains("error")'

# Display all of the logs that match a regex.
weaver gke-local logs 'msg.matches("error: file .* already closed")'

# Display all of the logs that have an attribute "foo" with value "bar".
weaver gke-local logs 'attrs["foo"] == "bar"'

# Display all of the logs in JSON format. This is useful if you want to
# perform some sort of post-processing on the logs.
weaver gke-local logs --format=json

# Display all of the logs, including internal system logs that are hidden by
# default.
weaver gke-local logs --system
```

Refer to `weaver gke-local logs --help` for a full explanation of the query
language, along with many more examples.

## Metrics

In addition to running the proxy on port 8000 (see the [Getting
Started](#local-gke-getting-started)), `weaver gke-local` also runs a status
server on port 8001. This server's `/metrics` endpoint exports the metrics of
all running Service Weaver applications in [Prometheus format][prometheus],
which looks like this:

```console
# HELP example_count An example counter.
# TYPE example_count counter
example_count{serviceweaver_node="bbc9beb5"} 42
example_count{serviceweaver_node="00555c38"} 9001
```

To visualize and query the metrics, make sure Prometheus is installed on your
local machine and then add the following stanza to your Prometheus yaml config
file:

```yaml
scrape_configs:
- job_name: 'prometheus-serviceweaver-scraper'
  scrape_interval: 5s
  metrics_path: /metrics
  static_configs:
    - targets: ['localhost:8001']
```

## Profiling

Use the `weaver gke-local profile` command to collect a profile of your Service Weaver
application. Invoke the command with the name (and optionally version) of the
app you wish to profile. For example:

```shell
# Collect a CPU profile of the latest version of the hello app.
$ weaver gke-local profile hello

# Collect a CPU profile of a specific version of the hello app.
$ weaver gke-local profile --version=8e1c640a-d87b-4020-b3dd-4efc1850756c hello

# Adjust the duration of a CPU profile.
$ weaver gke-local profile --duration=1m hello

# Collect a heap profile.
$ weaver gke-local profile --type=heap hello
```

`weaver gke-local profile` prints out the filename of the collected profile. You
can use the `go tool pprof` command to visualize and analyze the profile. For
example:

```console
$ profile=$(weaver gke-local profile <app>)    # Collect the profile.
$ go tool pprof -http=localhost:9000 $profile # Visualize the profile.
```

Refer to `weaver gke-local profile --help` for more details.

## Tracing

Run `weaver gke-local dashboard` to open a dashboard in a web browser. The
dashboard has a page for every Service Weaver application deployed via
`weaver gke-local deploy`. Every deployment's page has a link to the
deployment's [traces](#tracing) accessible via [Perfetto][perfetto]. Here's an
example of what the tracing page looks like:

![An example trace page](assets/images/trace_gke_local.png)

Refer to [Perfetto UI Docs](https://perfetto.dev/docs/visualization/perfetto-ui)
to learn more about how to use the tracing UI.

## Versioning

Recall that `weaver gke` performs slow rollouts
[across regions](#gke-multi-region) and
[across application versions](#versioning). `weaver gke-local` simulates this
behavior locally. When you `weaver gke-local deploy` an application, the
application is first rolled out to a number of canary regions before being
slowly rolled out to all regions. And within a region, the locally running proxy
slowly shifts traffic from old versions of the application to the new version of
the application. You can use `weaver gke-local status`, exactly like how you use
`weaver gke status`, to monitor the rollouts of your applications.

# Serializable Types

When you invoke a component's method, the arguments to the method (and the
results returned by the method) may be serialized and sent over the network.
Thus, a component's methods may only receive and return types that Service
Weaver knows how to serialize, types we call **serializable**. If a component
method receives or returns a type that isn't serializable, `weaver generate`
will raise an error during code generation time. The following types are
serializable:

-   All primitive types (e.g., `int`, `bool`, `string`) are serializable.
-   Pointer type `*t` is serializable if `t` is serializable.
-   Array type `[N]t` is serializable if `t` is serializable.
-   Slice type `[]t` is serializable if `t` is serializable.
-   Map type `map[k]v` is serializable if `k` and `v` are serializable.
-   Named type `t` in `type t u` is serializable if it is not recursive and one
    or more of the following are true:
    -   `t` is a protocol buffer (i.e. `*t` implements `proto.Message`);
    -   `t` implements [`encoding.BinaryMarshaler`][binary_marshaler] and
        [`encoding.BinaryUnmarshaler`][binary_unmarshaler];
    -   `u` is serializable; or
    -   `u` is a struct type that embeds `weaver.AutoMarshal` (see below).

The following types are not serializable:

-   Chan type `chan t` is *not* serializable.
-   Struct literal type `struct{...}` is *not* serializable.
-   Function type `func(...)` is *not* serializable.
-   Interface type `interface{...}` is *not* serializable.

**Note**: Named struct types that don't implement `proto.Message` or
`BinaryMarshaler` and `BinaryUnmarshaler` are *not* serializable by default.
However, they can trivially be made serializable by embedding
`weaver.AutoMarshal`.

```go
type Pair struct {
    weaver.AutoMarshal
    x, y int
}
```

The `weaver.AutoMarshal` embedding instructs `weaver generate` to generate
serialization methods for the struct. Note, however, that `weaver.AutoMarshal`
cannot magically make *any type* serializable. For example, `weaver generate`
will raise an error for the following code because the `NotSerializable` struct
is fundamentally not serializable.

```go
// ERROR: NotSerializable cannot be made serializable.
type NotSerializable struct {
    weaver.AutoMarshal
    f func()   // functions are not serializable
    c chan int // chans are not serializable
}
```

Also note that `weaver.AutoMarshal` can *not* be embedded in generic structs.

```go
// ERROR: Cannot embed weaver.AutoMarshal in a generic struct.
type Pair[A any] struct {
    weaver.AutoMarshal
    x A
    y A
}
```

To serialize generic structs, implement `BinaryMarshaler` and
`BinaryUnmarshaler`.

Finally note that while [Service Weaver requires every component method to
return an `error`](#components-interfaces), `error` is not a
serializable type. Service Weaver serializes `error`s in a way that does not
preserve any custom `Is` or `As` methods.

# weaver generate

`weaver generate` is Service Weaver's code generator. Before you compile and run a Service Weaver
application, you should run `weaver generate` to generate the code Service Weaver needs
to run an application. For example, `weaver generate` generates code to marshal
and unmarshal any types that may be sent over the network.

From the command line, `weaver generate` accepts a list of package paths. For
example, `weaver generate . ./foo` will generate code for the Service Weaver applications
in the current directory and in the `./foo` directory. For every package, the
generated code is placed in a `weaver_gen.go` file in the package's directory.
Running `weaver generate .  ./foo`, for example, will create `./weaver_gen.go`
and `./foo/weaver_gen.go`. You specify packages for `weaver generate` in the same
way you specify packages for `go build`, `go test`, `go vet`, etc. Run `go help
packages` for more information.

While you can invoke `weaver generate` directly, we recommend that you instead
place a line of the following form in one of the `.go` files in the root of
your module:

```go
//go:generate weaver generate ./...
```

Then, you can use the [`go generate`][go_generate] command to generate all of
the `weaver_gen.go` files in your module.

# Config Files

Service Weaver config files are written in [TOML](https://toml.io/en/) and look
something like this:

```toml
[serviceweaver]
name = "hello"
binary = "./hello"
args = ["these", "are", "command", "line", "arguments"]
env = ["PUT=your", "ENV=vars", "HERE="]
colocate = [
    ["main/Rock", "main/Paper", "main/Scissors"],
    ["github.com/example/sandy/PeanutButter", "github.com/example/sandy/Jelly"],
]
rollout = "1m"
```

A config file includes a `[serviceweaver]` section followed by a subset of the
following fields:

| Field | Required? | Description |
| --- | --- | --- |
| name | optional | Name of the Service Weaver application. If absent, the name of the app is derived from the name of the binary. |
| binary | required | Compiled Service Weaver application. The binary path, if not absolute, should be relative to the directory that contains the config file. |
| args | optional | Command line arguments passed to the binary. |
| env | optional | Environment variables that are set before the binary executes. |
| colocate | optional | List of colocation groups. When two components in the same colocation group are deployed, they are deployed in the same OS process, where all method calls between them are performed as regular Go method calls. To avoid ambiguity, components must be prefixed by their full package path (e.g., `github.com/example/sandy/`). Note that the full package path of the main package in an executable is `main`. |
| rollout | optional | How long it will take to roll out a new version of the application. See the [GKE Deployments](#gke-multi-region) section for more information on rollouts. |

A config file may additionally contain listener-specific and component-specific
configuration sections. See the [Component Config](#components-config) section
for details.

<div hidden class="todo">
Architecture
TODO: Explain the internals of Service Weaver.
</div>

# FAQ

### Do I need to worry about network errors when using Service Weaver?

Yes. While Service Weaver allows you to *write* your application as a single
binary, a distributed deployer (e.g., [multiprocess](#multiprocess),
[gke](#gke)), may place your components on separate processes/machines.
This means that method calls between those components will be executed as remote
procedure calls, resulting in possible network errors surfacing in your
application.

To be safe, we recommend that you assume that all cross-component method calls
involve a network, regardless of the actual component placement. If this is
overly burdensome, you can explicitly place relevant components in the same
[colocation group](#config-files), ensuring that they always run in the same OS
process.

**Note**: Service Weaver guarantees that all system errors are surfaced to the
application code as `weaver.RemoteCallError`, which can be handled as described
in an [earlier section](#components-semantics).

### What types of distributed applications does Service Weaver target?

Service Weaver primarily targets distributed serving systems. These are online
systems that need to handle user requests as they arrive. A web application or
an API server are serving systems, for example. Service Weaver tailors its
feature set and runtime assumptions towards serving systems in the following
ways:

* *Network servers are integrated into the framework*. The application can
easily obtain a network listener and create an HTTP server on top of it.
* *Rollouts are built into the framework*. The user specifies the rollout
duration and the framework gradually shifts network traffic from the old
version to the new.
* *All components are replicated*. A request for a component can go to any one
of its replicas. Replicas may automatically be scaled up and down depending on
the load.

### What about data-processing applications? Can I use Service Weaver for those?

In theory, you may be able to use Service Weaver for data-processing
applications, though you will find that it provides little support for some of
the common data-processing features such as checkpointing, failure recovery,
restarts etc.

Additionally, Service Weaver's replication model means that component replicas
may automatically be scaled up and down depending on the load. This is likely
something that you wouldn't want in your data-processing application. This
scale-up/scale-down behavior translates even to the application's `main()`
function and may cause your data-processing program to run multiple times.

### Why doesn't Service Weaver provide its own data storage?

Different applications have different storage needs (e.g., global replication,
performance, SQL/NoSQL). There are also a [myriad][db_engines] of storage
systems out there that make different tradeoffs along various dimensions
(e.g., price, performance, API).

We didn't feel like we could provide enough value by inserting ourselves
into the application's data model. We also didn't want to restrict how
applications interact with their data (e.g., offline DB updates). For those
reasons, we left the choice of data storage up to the application.

### Doesn't the lack of data storage integration limit the portability of Service Weaver applications?

Yes, to a degree. If you use a globally reachable data storage system, then you
can truly run your application anywhere, removing any portability concerns.

If, however, you run your storage system inside your deployment environment
(e.g., a MySQL instance running in the Cloud VPN), then if you start your
application in a different environment (e.g., your desktop), it may not have
access to the storage system. In such cases, we generally recommend that you
create different storage systems for different application environments, and
use Service Weaver [config files](#config-files) to point your application to
the right storage system for the given execution environment.

If you're using SQL, Go's [sql package][sql_package] helps isolate your
code from some differences in the underlying storage systems. See the
Service Weaver's [chat application example][chat_example] for how to setup
your application to use the environment-local storage systems.

### Does the Service Weaver versioning approach mean I will end up running multiple instances of my app during a rollout? Isn't that expensive?

As we described in the GKE [versioning](#versioning) section, we utilize the
combination of auto-scaling and blue/green deployment to minimize the cost
of running multiple versions of the same application during rollouts.

In general, it is up to the deployer implementation to ensure that the
rollout cost is minimized. We envision that most cloud deployers will use a
similar technique to GKE to minimize their rollout costs. Other deployers
may choose to simply run full per-version serving trees, like the
[multiprocess](#multiprocess) deployer.

### Service Weaver's microservice development model is quite unique. Is it making a stand against traditional microservices development?

No. We acknowledge that there are still valid reasons why developers may
choose to run separate binaries for different microservices (e.g.,
different teams controlling their own binaries). We believe, however,
that Service Weaver's *modular monolith* model is applicable to a lot of
common use-cases and can be used in conjunction with the traditional
microservices model.

For example, a team may decide to unify all of the services in their control
into a single Service Weaver application. Cross-team interactions will still
be handled in the traditional model, with all of the versioning and development
implications that come with that model.

### Isn't writing "monoliths" a step in the wrong direction for distributed application development?

Service Weaver is trying to encourage a *modular monolith* model, where
the application is written as a single modularized binary that runs as separate
microservices. This is different from the monolith model, where the binary runs
as a single (replicated) service.

We believe that the Service Weaver's *modular monolith* model has the best of
both worlds: the ease of development of monolithic applications, with the
runtime benefits of microservices.

[actors]: https://en.wikipedia.org/wiki/Actor_model
[binary_marshaler]: https://pkg.go.dev/encoding#BinaryMarshaler
[binary_unmarshaler]: https://pkg.go.dev/encoding#BinaryUnmarshaler
[blue_green]: https://docs.aws.amazon.com/whitepapers/latest/overview-deployment-options/bluegreen-deployments.html
[canary]: https://sre.google/workbook/canarying-releases/
[chat_example]: https://github.com/ServiceWeaver/weaver/tree/main/examples/chat/
[chrome_tracing]: https://docs.google.com/document/d/1CvAClvFfyA5R-PhYUmn5OOQtYMH4h6I0nSsKchNAySU/preview
[cloud_logging]: https://cloud.google.com/logging
[cloud_metrics]: https://cloud.google.com/monitoring/api/metrics_gcp
[cloud_trace]: https://cloud.google.com/trace
[db_engines]: https://db-engines.com/en/ranking
[gcloud_billing]: https://console.cloud.google.com/billing
[gcloud_billing_projects]: https://console.cloud.google.com/billing/projects
[gcloud_install]: https://cloud.google.com/sdk/docs/install
[gke]: https://cloud.google.com/kubernetes-engine
[gke_create_project]: https://cloud.google.com/resource-manager/docs/creating-managing-projects#gcloud
[go_generate]: https://pkg.go.dev/cmd/go/internal/generate
[go_install]: https://go.dev/doc/install
[go_interfaces]: https://go.dev/tour/methods/9
[hello_app]: https://github.com/ServiceWeaver/weaver/tree/main/examples/hello
[http_pprof]: https://pkg.go.dev/net/http/pprof
[isolation]: https://sre.google/workbook/canarying-releases/#dependencies-and-isolation
[kubernetes]: https://kubernetes.io/
[logs_explorer]: https://cloud.google.com/logging/docs/view/logs-explorer-interface
[metric_types]: https://prometheus.io/docs/concepts/metric_types/
[metrics_explorer]: https://cloud.google.com/monitoring/charts/metrics-explorer
[n_queens]: https://en.wikipedia.org/wiki/Eight_queens_puzzle
[net_listen]: https://pkg.go.dev/net#Listen
[otel]: https://opentelemetry.io/docs/instrumentation/go/getting-started/
[otel_all_you_need]: https://lightstep.com/blog/opentelemetry-go-all-you-need-to-know#adding-detail
[perfetto]: https://ui.perfetto.dev/
[pprof]: https://github.com/google/pprof
[pprof_blog]: https://go.dev/blog/pprof
[prometheus]: https://prometheus.io
[prometheus_counter]: https://prometheus.io/docs/concepts/metric_types/#counter
[prometheus_gauge]: https://prometheus.io/docs/concepts/metric_types/#gauge
[prometheus_histogram]: https://prometheus.io/docs/concepts/metric_types/#histogram
[prometheus_naming]: https://prometheus.io/docs/practices/naming/
[sql_package]: https://pkg.go.dev/database/sql
[trace_service]: https://cloud.google.com/trace
[update_failures_paper]: https://scholar.google.com/scholar?cluster=4116586908204898847
[weak_consistency]: https://mwhittaker.github.io/consistency_in_distributed_systems/1_baseball.html
[weaver_examples]: https://github.com/ServiceWeaver/weaver/tree/main/examples
[weaver_github]: https://github.com/ServiceWeaver/weaver
[xdg]: https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
