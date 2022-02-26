module github.com/paulhindemith/board_generator

go 1.15

replace (
	github.com/paulhindemith/category => ../category
	github.com/paulhindemith/ethical_man => ../ethical_man
	github.com/paulhindemith/pasculli => ../pasculli
	github.com/paulhindemith/pkg-strategist => ../pkg-strategist
	github.com/paulhindemith/scraper => ../scraper
)

require (
	github.com/carterjones/signalr v0.3.5
	github.com/devigned/signalr-go v0.0.0-20190520231709-aa89e914a439
	github.com/goccy/go-graphviz v0.0.9
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/websocket v1.4.2
	github.com/opentracing/opentracing-go v1.2.0
	github.com/paulhindemith/category v0.0.0-00010101000000-000000000000
	github.com/paulhindemith/pasculli v0.0.0-00010101000000-000000000000
	github.com/paulhindemith/pkg-strategist v0.0.0-00010101000000-000000000000
	github.com/paulhindemith/scraper v0.0.0-00010101000000-000000000000
	github.com/philippseith/signalr v0.4.1
	github.com/prometheus/client_golang v1.11.0
	github.com/spf13/viper v1.7.1
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	go.uber.org/multierr v1.6.0
	go.uber.org/zap v1.17.0
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)
