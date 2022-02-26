// Code generated .* DO NOT EDIT.

package infra
import(
  "context"
)
// <Local>
//         arrow2
//        --------->
// |_|_|_|
//  dom
//
// <GRPC>
// |_|_|_|
//  dom
//  |
//  |(grpc)
//  |     arrow1
//  V     ------->
// |_|_|_|
//  dom
//
// <Delay>
// |_|_|_|
//  dom
//  |
//  |(late for any duration)
//  |     arrow1
//  V     ------->
// |_|_|_|
//  dom
//

  type PublisherDomainLocalConverter struct {
  Dom chan *PublisherMessage
}
func (s *PublisherDomainLocalConverter) Init(ctx context.Context) *PublisherDomainLocalConverter {
  return s
}
func (s *PublisherDomainLocalConverter) Start() {
}
    
  
  type BoardGeneratorDomainLocalConverter struct {
  Dom chan *BoardGeneratorMessage
}
func (s *BoardGeneratorDomainLocalConverter) Init(ctx context.Context) *BoardGeneratorDomainLocalConverter {
  return s
}
func (s *BoardGeneratorDomainLocalConverter) Start() {
}
    