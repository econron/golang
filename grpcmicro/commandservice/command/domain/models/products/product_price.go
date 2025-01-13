package products

import (
	"commandservice/errs"
	"fmt"
)

// 商品単価
type ProductPrice struct {
	value uint32
}

func (ins *ProductPrice) Value() uint32 {
	return ins.value
}

// コンストラクタ
func NewProductPrice(value uint32) (*ProductPrice, error) {
	const MIN_VALUE = 50
	const MAX_VALUE = 1000000
	if value < MIN_VALUE || value > MAX_VALUE {
		return nil, errs.NewDomainError(fmt.Sprintf("商品単価は%d以上%d以下である必要があります", MIN_VALUE, MAX_VALUE))
	}
	return &ProductPrice{
		value: value,
	}, nil
}