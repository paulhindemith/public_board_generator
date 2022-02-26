# public_board_generator
##### 注意
  このリポジトリはプライベートリポジトリboard_generatorのコピーです。
  使用できません。[ブログ](https://taizokaneko.info)からアクセスされていることを想定しています。
## 概要
取引所から価格を1秒ごとにwatchしてgRPCでpushするサービスです。
## ディレクトリの説明
hackはスクリプトが入っており、grpcの生成やkoによってkubernetesのyamlを生成しています。<br>
configはyamlの原本が入っています。<br>
arrows、infraは特殊ですが、arrowsで書いた内容をアプリケーションとして動かすためにinfraでコードを自動生成しています。<br>
システムの仕組み上、アプリケーション内部は点と矢印による論理しかありませんので、それをそのまま描き下せるような仕組みを導入しています。arrowsディレクトリの各ファイルのarrow()が矢印の部分で、型が点です。<br>
```go
// *adapter.Coterminalを*Boardに変換する処理をこのarrowで行っています。
// *adapter.Coterminal -> *Board
func (g *Generator) arrow(e *adapter.Coterminal) *Board {
  // 変換する処理
  return nil
}
```
```go
// *board_generator.Boardは上記の*Boardと同じです。上記の結果を受け取ってgRPCで送信する処理を行います。
// *board_generator.Board -> *Empty
func (g *Publisher) arrow(b *board_generator.Board) *Empty {
  // 送信する処理
  return nil
}
```
