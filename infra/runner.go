// Code generated .* DO NOT EDIT.

package infra

import(
  "context"
  "io"
  jaeger "github.com/uber/jaeger-client-go"
  opentracing "github.com/opentracing/opentracing-go"
  "github.com/paulhindemith/category/pkg/adapter"
  "github.com/paulhindemith/board_generator/arrows/board_generator"
  "github.com/paulhindemith/board_generator/arrows/publisher"
)
type SpanObject interface {
  SetSpan(opentracing.Span)
}

type BoardGeneratorMessage struct {
  Object *board_generator.Board
  SpanContext opentracing.SpanContext
}
type PublisherMessage struct {
  Object *publisher.Empty
  SpanContext opentracing.SpanContext
}
type PublisherEndingRunner struct {
  Dom chan *PublisherMessage
  ErrChan chan error
}
func (r *PublisherEndingRunner) Init(ctx context.Context) *PublisherEndingRunner {
  return r
}
func(r *PublisherEndingRunner) Start() {
  for range r.Dom {
  }
  r.ErrChan <- io.EOF
}
type BoardGeneratorStartingRunner struct {
  Arrow board_generator.BoardGenerator
  Cod chan *BoardGeneratorMessage
  closed bool
  closedSeen bool
}
func (s *BoardGeneratorStartingRunner) Init(ctx context.Context) *BoardGeneratorStartingRunner {
  return s
}
func(r *BoardGeneratorStartingRunner) Start() {
  var coterminal = &adapter.Coterminal{}
  defer close(r.Cod)
  for {
    span := opentracing.GlobalTracer().StartSpan("board_generator")
    coterminal.SetSpan(span)
    object := r.Arrow(coterminal)
    span.Finish()
    if r.closed && !r.closedSeen {
      r.closedSeen = true
      coterminal.Stop()
    }
    if object == nil {
      continue
    }
    if object.Err() == io.EOF {
      return
    }
    r.Cod <- &BoardGeneratorMessage {
      Object: object,
      SpanContext: span.(*jaeger.Span).SpanContext(),
    }
  }

}
func(r *BoardGeneratorStartingRunner) Close() {
  r.closed = true
}
type PublisherRunner struct {
  Arrow func(*board_generator.Board)*publisher.Empty
  Dom chan *BoardGeneratorMessage
  Cod chan *PublisherMessage
}
func (s *PublisherRunner) Init(ctx context.Context) *PublisherRunner {
  return s
}
func(r *PublisherRunner) Start() {
  for msg := range r.Dom {
    span := opentracing.GlobalTracer().StartSpan("publisher", opentracing.FollowsFrom(msg.SpanContext))
    if so, ok := interface{}(msg.Object).(SpanObject); ok {
      so.SetSpan(span)
    }
    object := r.Arrow(msg.Object)
    span.Finish()
    if object == nil {
      continue
    }
    r.Cod <- &PublisherMessage {
      Object: object,
      SpanContext: span.(*jaeger.Span).SpanContext(),
    }
  }
  // 本当は複数closeされるであろうcodについて全てのcloseを待ってから次へと進むべき。
  defer func(){
    if err := recover(); err != nil {
    }
  }()
  close(r.Cod)
}
func(r *PublisherRunner) Close() {
  close(r.Dom)
}
