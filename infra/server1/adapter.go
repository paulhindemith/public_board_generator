// Code generated .* DO NOT EDIT.

package server1

import(
  "sync"
  "go.uber.org/multierr"
  "golang.org/x/sync/errgroup"
  "context"
  "go.uber.org/zap"
  "reflect"
  "io"
  "github.com/paulhindemith/pkg-strategist/grpc_server"
  "github.com/paulhindemith/pkg-strategist/grpc_server_ready"
  "github.com/paulhindemith/pkg-strategist/grpc_client_conn"
  adapter "github.com/paulhindemith/category/pkg/adapter"
  s_adapter "github.com/paulhindemith/pkg-strategist/adapter"
  "github.com/paulhindemith/pkg-strategist/banner"
  "github.com/paulhindemith/pkg-strategist/logging"
  ctx_rdy "github.com/paulhindemith/category/pkg/ready"
  ctx_init "github.com/paulhindemith/category/pkg/initialize"
  ctx_eg "github.com/paulhindemith/category/pkg/errgroup"
  ctx_errchan "github.com/paulhindemith/category/pkg/errchan"
  "github.com/paulhindemith/board_generator/infra"
  "github.com/paulhindemith/board_generator/arrows/board_generator"
  "github.com/paulhindemith/board_generator/arrows/publisher"
)

type Adapter struct {
  logger *zap.Logger
  routines []adapter.Routine
  firsts []adapter.FirstRoutine
  finishInitialize context.CancelFunc
  eg  *errgroup.Group
  rdy_eg *errgroup.Group
  cancel context.CancelFunc
  grpc_rdy chan struct{}
}

var ErrChan = make(chan error, 1)
var BoardGeneratorChan = make(chan *infra.BoardGeneratorMessage, 10)
var PublisherChan = make(chan *infra.PublisherMessage, 10)
func NewAdapter(pctx context.Context) s_adapter.Adapter {
  ctx, cancel := context.WithCancel(context.Background())
  iniCtx, finishInitialize := context.WithCancel(context.Background())
  a := &Adapter{
    logger: logging.FromContext(pctx),
    grpc_rdy: grpc_server_ready.FromContext(pctx),
    cancel: cancel,
    finishInitialize: finishInitialize,
    eg: &errgroup.Group{},
    rdy_eg: &errgroup.Group{},
    firsts: make([]adapter.FirstRoutine, 0),
  }
  ctx = grpc_server.WithContext(ctx, grpc_server.FromContext(pctx))
  ctx = grpc_client_conn.WithContext(ctx, grpc_client_conn.FromContext(pctx))
  ctx = logging.WithContext(ctx, a.logger)
  ctx = ctx_init.WithContext(ctx, iniCtx)
  ctx = ctx_rdy.WithContext(ctx, a.rdy_eg)
  ctx = ctx_eg.WithContext(ctx, a.eg)
  ctx = ctx_errchan.WithContext(ctx, ErrChan)
  a.routines = make([]adapter.Routine, 0)
  // runners
  var oneRoutine adapter.FirstRoutine
  a.routines = append(a.routines, (&infra.PublisherEndingRunner{Dom: PublisherChan, ErrChan: ErrChan}).Init(ctx))
  oneRoutine = (&infra.BoardGeneratorStartingRunner{Arrow: board_generator.NewArrow(ctx).(board_generator.BoardGenerator), Cod: BoardGeneratorChan }).Init(ctx)
  a.firsts = append(a.firsts, oneRoutine)
  a.routines = append(a.routines, oneRoutine)
  a.routines = append(a.routines, (&infra.PublisherRunner{Dom: BoardGeneratorChan, Arrow: publisher.NewArrow(ctx).(func(*board_generator.Board)*publisher.Empty), Cod: PublisherChan}).Init(ctx))
  // converters
    // 通常のConverter
    a.routines = append(a.routines, (&infra.PublisherDomainLocalConverter{Dom: PublisherChan}).Init(ctx))
    // 通常のConverter
    a.routines = append(a.routines, (&infra.BoardGeneratorDomainLocalConverter{Dom: BoardGeneratorChan}).Init(ctx))
  // buffers
    a.routines = append(a.routines, (&infra.BoardGeneratorCodomainSingleBuffer{Cod: BoardGeneratorChan}).Init(ctx))
    a.routines = append(a.routines, (&infra.PublisherCodomainSingleBuffer{Cod: PublisherChan}).Init(ctx))
  return a
}
func (a *Adapter) Start(stopch <-chan struct{}) error {
  var err error
  a.finishInitialize()
  err = a.rdy_eg.Wait()
  if err != nil {
    return err
  }
  // grpc_serverにサービスが全てRegisterされていることを想定している。
  if a.grpc_rdy != nil {
    close(a.grpc_rdy)
  }
  wg := &sync.WaitGroup{}
  for _, r := range a.routines {
    wg.Add(1)
    routine := r
    name := reflect.TypeOf(routine).String()
    banner.Start(name)
    go func() {
      routine.Start()
      wg.Done()
      banner.Finish(name)
    }()
  }
  a.logger.Info("Category System is ready.")
  select {
  case <- stopch:
    a.logger.Info("stopch is caught.")
    for _, f := range a.firsts {
      f.Close()
    }
  case err := <- ErrChan:
    // 正常にarrowsが終了した場合EOF
    if err == adapter.EOF {
      a.logger.Info("ErrChan caught adapter.EOF. Now Normally closing systems.", zap.String("error", err.Error()))
      for _, f := range a.firsts {
        f.Close()
      }
    } else if err != io.EOF {
      a.logger.Error("ErrChan caught bad error.", zap.String("error", err.Error()))
      for _, f := range a.firsts {
        f.Close()
      }
    }
  }
  lastEg := &errgroup.Group{}
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  lastEg.Go(func()error{
    var merr error
    for {
      select {
      case <- ctx.Done():
        a.logger.Info("Error store system is closed.")
        return merr
      case err := <- ErrChan:
        // 正常にarrowsが終了した場合EOF
        if err != io.EOF && err != adapter.EOF {
          a.logger.Error("ErrChan is caught.", zap.String("error", err.Error()))
          merr = multierr.Append(merr, err)
        }
      }
    }
  })
  a.logger.Info("now waiting for arrow systems.")
  wg.Wait()
  a.cancel()
  a.logger.Info("now waiting for eg systems.")
  err = a.eg.Wait()
  if err != nil {
    a.logger.Error("got error", zap.String("error", err.Error()))
    return err
  }
  cancel()
  err = lastEg.Wait()
  if err != nil {
    a.logger.Error("got error", zap.String("error", err.Error()))
    return err
  }
  a.logger.Info("category system is closed.")
  return s_adapter.Shutdowned{}
}
