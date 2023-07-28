# 什么是 Services Weaver ？
Services Weaver是一个用于开发、部署和管理分布式应用程序的程序设计框架。您可以在本地机器上运行、测试和调试Service Weaver应用程序，然后使用一个命令将应用程序部署到云中。
```console
$ go run .                      # Run locally.
$ weaver gke deploy weaver.toml # Run in the cloud.
```
一个Service Weaver应用程序由许多个**组件**组成。组件表示为常规的Go [interface][go_interfaces]，组件通过调用这些接口定义的方法来相互交互。这使得开发Service Weaver应用程序变得容易。您不需要编写任何网络或序列化代码;你只要写Go。Service Weaver还提供了用于记录 logging, metrics, tracing, routing, testing等的库。

您可以像运行单个命令一样轻松地部署Service Weaver应用程序。在幕后，Service Weaver将沿着组件边界剖析您的二进制文件，允许不同的组件在不同的机器上运行。Service Weaver将为您复制、自动伸缩和共同定位这些分布式组件。它还将代表您管理所有网络细节，确保不同的组件可以相互通信，并且客户端可以与您的应用程序通信。

请参阅 [Installation](#安装) 部分，在您的机器上安装Service Weaver，或者阅读 [教程](#教程) 部分，了解如何开发Service Weaver应用程序。

# 安装
确保你已经[安装 Go][go_install]，版本1.20或更高。然后，执行如下命令安装weaver命令:
```console
go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
```
`go install` 将`weaver`命令安装到`$GOBIN`中，默认为`$GOPATH/bin`。确保这个目录包含在您的`PATH`中。你可以做到这一点，例如，通过在你的`.bashrc`中添加以下代码并运行`source ~/.bashrc`:

```console
$ export PATH="$PATH:$GOPATH/bin"
```
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
**注意**:对于GKE部署，您还应该安装`weaver gke`命令(有关详细信息，请参阅 [GKE](#gke) 部分):

```console
$ go install github.com/ServiceWeaver/weaver-gke/cmd/weaver-gke@latest
```
**注意**:如果你在macOS上安装`weaver`和`weaver gke`命令时遇到问题，你可能需要在install命令前加上`export CGO_ENABLED=1; export CC=gcc`。例如:

```console
$ export CGO_ENABLED=1; export CC=gcc; go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
```
# 教程
在本节中，我们将向您展示如何开发`Service Weaver`应用程序。要安装`Service Weaver`并学习本教程，请参阅 [Installation](#安装) 部分。本教程中提供的完整源代码可以在 [here][hello_app] 找到。
## Components
Service Weaver的核心抽象是**组件**。组件就像一个参与者，Service Weaver应用程序是由一组组件实现的。具体地说，组件用一个常规的Go接口表示，组件通过调用这些接口定义的方法来相互交互。




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
[emojis]: https://emojis.serviceweaver.dev/
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
[weavertest.Fake]: https://pkg.go.dev/github.com/ServiceWeaver/weaver/weavertest#Fake
[workshop]: https://github.com/serviceweaver/workshops
[xdg]: https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
