package products

import (
	"commandservice/domain/models/categories"
	"commandservice/errs"
	"github.com/google/uuid"
)

// 商品
type Product struct {
	id *ProductID
	name *ProductName
	price *ProductPrice
	category *categories.Category
}