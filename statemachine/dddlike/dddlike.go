package dddlike

import "fmt"

// 集約ルートをベースにイベントで状態を変化するように書く
/*
本屋の在庫管理の状態遷移を想像する

イベント
- 納入された
- 棚に置いた
- お客さんが買う

状態
- 店舗在庫なし
- 店舗在庫あり
- 棚陳列済み

店舗在庫なし & 納入された -> 店舗在庫あり
店舗在庫あり & 棚に置いた -> 棚陳列済み
棚陳列済み & お客さんが買う -> 店舗在庫なし

店舗在庫あり & 棚に置いた -> 棚陳列済み　は店にも客にも不要そう？なので削除する。

最終
- 店舗在庫なし & 納入された -> 店舗在庫あり
- 店舗在庫あり & お客さんが買う -> 店舗在庫なし

*/
// 内部の値を変えるのはメソッドを通じてのみ
// newした時点でもう値を受け付けないようにする
type StockStatus string

type Book struct {
	name string
	status StockStatus
}

const (
	INSTOCK StockStatus = "IN STOCK"
	OUTOFSTOCK StockStatus = "OUT OF STOCK"
)

func NewBook(name string) *Book {
	return &Book{
		name: name,
		status: OUTOFSTOCK,
	}
}

// UI用
func (b *Book) Status() StockStatus {
	return b.status
}

// 本が到着した
func (b *Book) Arrived() error {
	if b.status == INSTOCK {
		return fmt.Errorf("invalid status: %s", b.status)
	}
	b.status = INSTOCK
	return nil
}

// 本が売れた
func (b *Book) Sold() error {
	if b.status == OUTOFSTOCK {
		return fmt.Errorf("invalid status: %s", b.status)
	}
	b.status = OUTOFSTOCK
	return nil
}
