package products

import (
	"commandservice/errs"
	"fmt"
	"unicode/utf8"
)

// 商品名
type ProductName struct {
	value string
}

func (ins *ProductName) Value() string {
	return ins.value
}

func (one *ProductName) Equals(another *ProductName) bool {
	return one.value == another.value
}

// コンストラクタ
func NewProductName(value string) (*ProductName, error) {
	const MIN_LENGTH = 5
	const MAX_LENGTH = 30
	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("商品名は%d文字以上%d文字以下でなければなりません", MIN_LENGTH, MAX_LENGTH))
	}
	return &ProductName{value: value}, nil
}