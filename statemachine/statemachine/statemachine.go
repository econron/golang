package statemachine

import "fmt"

// inspired by https://qiita.com/shikuno_dev/items/1de5129bdf1f004e4525

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

type StateKind string
type EventKind string
// 状態
const (
	InStock StateKind = "in stock"
	OutOfStock StateKind = "out of stock"
)
// イベント
const (
	StockArrived EventKind = "stock has arrived"
	Bought EventKind = "bought"
)
// 状態を表すデータ
type StockState struct {
	state StateKind
}
func NewStockState(newState StateKind) *StockState {
	return &StockState{
		state: newState,
	}
}
// 状態遷移
func (s StockState) Transition(e EventKind) (*StockState, error) {
	switch s.state {
	case InStock:
		if e == Bought {
			return NewStockState(OutOfStock), nil
		}
	case OutOfStock:
		if e == StockArrived {
			return NewStockState(InStock), nil
		}
	}
	return nil, fmt.Errorf("invalid state and event. state:%s, event;%s", s.state, e)
}

func NewEventKind(event string) EventKind {
	return EventKind(event)
}
