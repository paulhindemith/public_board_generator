package publisher

import(
  "log"
  "context"
  "go.uber.org/zap"
  "github.com/paulhindemith/pkg-strategist/logging"
  "github.com/paulhindemith/board_generator/arrows/board_generator"
  "github.com/paulhindemith/pkg-strategist/grpc_server"
  "github.com/paulhindemith/board_generator/arrows/publisher/board/types"
  "github.com/paulhindemith/board_generator/arrows/publisher/board"
  ctx_init "github.com/paulhindemith/category/pkg/initialize"
  ctx_rdy "github.com/paulhindemith/category/pkg/ready"
)

type publishing struct {
  conn types.Board_PublishBoardsServer
  stopch chan struct{}
}

type Server struct {
	types.UnimplementedBoardServer
  clients map[string]publishing
}

func (s *Server) PublishBoards(in *types.BoardsRequest, gbs types.Board_PublishBoardsServer) error {
  stopch := make(chan struct{}, 0)
  s.clients[in.Id] = publishing{
    conn: gbs,
    stopch: stopch,
  }
  <- stopch
  return nil
}

type Empty struct{
}
type Publisher struct {
  logger *zap.Logger
  server *Server
}
func NewArrow(ctx context.Context) interface{} {
  g := &Publisher{
    logger: logging.FromContext(ctx),
    server: &Server{
      clients: map[string]publishing{},
    },
  }
  initCtx := ctx_init.FromContext(ctx)
  rdy_eg := ctx_rdy.FromContext(ctx)
  rdy_eg.Go(func() error {
    <- initCtx.Done()
    s := grpc_server.FromContext(ctx)
    types.RegisterBoardServer(s, g.server)
    return nil
  })
  return g.arrow
}
func (g *Publisher) arrow(b *board_generator.Board) *Empty {
  for id, pub := range(g.server.clients) {
    br := board.DomainToBoardReply(b.Board)
    if b.GetName() == "gmo_fx_btc_jpy" {
      g.logger.Info("called")
      log.Println(br.GetName())
    }
    err := pub.conn.Send(br)
    if err != nil {
      g.logger.Info(err.Error())
      close(pub.stopch)
      delete(g.server.clients, id)
      return nil
    }
  }
  return nil
}
