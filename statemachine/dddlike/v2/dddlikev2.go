package dddlike

import (
	"fmt"
	"errors"
)

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
	count int
}

const (
	INSTOCK StockStatus = "IN STOCK"
	OUTOFSTOCK StockStatus = "OUT OF STOCK"
)

func NewBook(name string) *Book {
	return &Book{
		name: name,
		count: 0,
	}
}

func (b *Book) Status() StockStatus {
	if b.count > 0 {
		return INSTOCK
	}
	return OUTOFSTOCK
}

func (b *Book) Arrived(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("amount is invalid. amount: %d", amount)
	}
	b.count += amount
	return nil
}

func (b *Book) Sold() error {
	if b.count == 0 {
		return errors.New("this book has already been out of stock.")
	}
	b.count-=1
	return nil
}