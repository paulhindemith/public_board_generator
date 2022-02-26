// Code generated .* DO NOT EDIT.

package infra

import(
  "context"
)
// <Single>
// arrow1          arrow2
// ------->      --------->
//        |_|_|_|
//         cod and dom
//
// <Publish>
// arrow1                       arrow2.1
// ------->                   --------->
//        |_|_|_| ____> |_|_|_|
//         cod   \       dom   arrow2.2
//                 \          --------->
//                  \_> |_|_|_|
//                       dom
//
type BoardGeneratorCodomainSingleBuffer struct {
  Cod chan *BoardGeneratorMessage
  Doms []chan *BoardGeneratorMessage
}
func (b *BoardGeneratorCodomainSingleBuffer) Init(ctx context.Context) *BoardGeneratorCodomainSingleBuffer {
  return b
}
func (b *BoardGeneratorCodomainSingleBuffer) Start() {
  b.Doms = []chan *BoardGeneratorMessage{b.Cod}
}
type PublisherCodomainSingleBuffer struct {
  Cod chan *PublisherMessage
  Doms []chan *PublisherMessage
}
func (b *PublisherCodomainSingleBuffer) Init(ctx context.Context) *PublisherCodomainSingleBuffer {
  return b
}
func (b *PublisherCodomainSingleBuffer) Start() {
  b.Doms = []chan *PublisherMessage{b.Cod}
}
