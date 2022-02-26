package board_generator

import(
  "log"
  "io"
  "time"
  "sync"
  "context"
  "go.uber.org/zap"
  "github.com/paulhindemith/scraper/pkg/limiter"
  "github.com/paulhindemith/pkg-strategist/logging"
  "github.com/paulhindemith/category/pkg/adapter"
  "github.com/paulhindemith/scraper/pkg/helper"
  "github.com/paulhindemith/scraper/pkg/domain"
  "github.com/paulhindemith/scraper/pkg/global"
  oanda "github.com/paulhindemith/scraper/pkg/domain/oanda/v1/prices"
  "github.com/paulhindemith/pasculli/pkg/mind"
  ctx_init "github.com/paulhindemith/category/pkg/initialize"
)
func init(){
  log.SetFlags(log.LstdFlags | log.Llongfile)
}

type Board struct{
  domain.Board
  Name mind.PriceName
  err error
}
func (b *Board) Err() error {
  return b.err
}
func (b *Board) GetName() mind.PriceName {
  return b.Name
}
func (b *Board) GetElement() interface{} {
  return b
}
func (b *Board) GraphvizTitle() string {
  return "PN × Time × Board"
}

type Generator struct {
  logger *zap.Logger
  config *Config
  stopch chan struct{}
  boards chan *Board
  initialized <- chan struct{}
}
func NewArrow(ctx context.Context) interface{} {
  g := &Generator{
    logger: logging.FromContext(ctx),
  }
  if initCtx := ctx_init.FromContext(ctx); initCtx != nil {
    g.initialized = initCtx.Done()
  }
  g.config = g.getConfig()
  oanda.Secret = domain.Secret(g.config.OandaSecret)
  return g.newProductionArrow(ctx)
}
func (g *Generator) newProductionArrow(ctx context.Context) BoardGenerator {
  g.stopch =  make(chan struct{}, 1)
  g.boards =  make(chan *Board, 10)
  g.logger.Info("board is defined")
  global.SetLogger(g.logger)
  for _, tgt := range g.config.Targets {
    boardHandler := g.handleBoardFunc(mind.PriceName(tgt.PriceName))
    limiter := limiter.EverySecondLimiter(tgt.PriceName, g.config.Wait)
    global.Add(helper.NewBoardGetter(tgt.PriceName, limiter, tgt.ReadLimit, tgt.WriteLimigt, boardHandler))
  }
  ctx, cancel := context.WithCancel(ctx)
  go func(){
    <- g.initialized
    global.Start(ctx)
    g.logger.Info("board is stoped")
    close(g.boards)
  }()
  go func(){
    <-g.stopch
    cancel()
  }()
  return g.arrow
}

func (g *Generator) handleBoardFunc(name mind.PriceName ) func(domain.Board) error {
  return func(b domain.Board) error {
    nb := &Board{
      Name: name,
      Board: b,
    }
    g.logger.Debug("got board.",
      zap.String("name", string(name)),
      zap.Any("asks", b.GetAsks().Len()),
      zap.Any("bids", b.GetBids().Len()))
    g.boards <- nb
    return nil
  }
}

var Ended = &Board{err: io.EOF}
func (g *Generator) arrow(e *adapter.Coterminal) *Board {
  if e.Err() == io.EOF {
    g.stop()
  }
  select {
  case <- time.After(1*time.Second):
  case b, ok := <- g.boards:
    if !ok {
      g.logger.Info("closed generator.")
      return Ended
    }
    return b
  }
  return nil
}
var closing sync.Once
func (g *Generator) stop() {
  closing.Do(func() {
    close(g.stopch)
  })
}

type BoardGenerator func(*adapter.Coterminal) *Board
func (bg BoardGenerator) Title() string {
  return "PN × Board"
}
