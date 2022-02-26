package board

import(
  "github.com/paulhindemith/scraper/pkg/domain"
  "github.com/paulhindemith/scraper/pkg/domain/any/board"
  "github.com/paulhindemith/board_generator/arrows/publisher/board/types"
)

func DomainToBoardReply(b domain.Board)*types.BoardReply {
  asks := make([]*types.Order, 0)
  bids := make([]*types.Order, 0)
  for _, o := range b.GetAsks().Orders() {
    asks = append(asks, &types.Order{
      Price: float64(o.GetPrice()),
      Size: o.GetSize(),
      F: int64(o.GetF()),
    })
  }
  for _, o := range b.GetBids().Orders() {
    bids = append(bids, &types.Order{
      Price: float64(o.GetPrice()),
      Size: o.GetSize(),
      F: int64(o.GetF()),
    })
  }
  return &types.BoardReply{
    Name: b.GetName(),
    Time: b.GetTime(),
    ResponseTime: b.GetResponseTime(),
    RequestTime: b.GetRequestTime(),
    Asks: asks,
    Bids: bids,
  }
}
func BoardReplyToDomain(br *types.BoardReply) domain.Board {
  askTypedOrders := make([]*domain.TypedOrder, 0)
  bidTypedOrders := make([]*domain.TypedOrder, 0)
  for _, o := range br.Asks {
    askTypedOrders = append(askTypedOrders, &domain.TypedOrder{
      Price: o.GetPrice(),
      Size: o.GetSize(),
    })
  }
  for _, o := range br.Bids {
    bidTypedOrders = append(bidTypedOrders, &domain.TypedOrder{
      Price: o.GetPrice(),
      Size: o.GetSize(),
    })
  }
  asks := domain.NewAsks(askTypedOrders)
  bids := domain.NewBids(bidTypedOrders)
  return &board.Board{
    Name: br.GetName(),
    Time: br.GetTime(),
    ResponseTime: br.GetResponseTime(),
    RequestTime: br.GetRequestTime(),
    Asks: asks,
    Bids: bids,
  }
}
