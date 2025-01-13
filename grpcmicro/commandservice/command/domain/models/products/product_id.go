package products

import (
	"commandservice/errs"
	"fmt"
	"regexp"
	"unicode/utf8"
)

// 商品番号
type ProductID struct {
	value string
}

func (ins *ProductID) Value() string {
	return ins.value
}

// 同一性検証
func (one *ProductID) Equals(another *ProductID) bool {
	if one == another {
		return true
	}
	return one.value == another.value
}

// コンストラクタ
func NewProductID(value string) (*ProductID, error) {
	// フィールドの長さ
	const LENGTH = 36
	// UUIDの正規表現
	const REGEXP string = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	// 引数の文字チェック
	if utf8.RuneCountInString(value) != LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("商品番号は%d文字でなければなりません", LENGTH))
	}
	// 正規表現チェック
	if matched, err := regexp.MatchString(REGEXP, value); !matched {
		return nil, errs.NewDomainError("商品番号はUUID形式でなければなりません" + err.Error())
	}
	return &ProductID{value: value}, nil
}