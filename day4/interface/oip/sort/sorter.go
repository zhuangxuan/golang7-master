package sort

import "go-course/interface/oip/common"

type Sorter interface {
	Sort([]*common.Product) []*common.Product //传入一批商品，返回排序之后的商品
	Name() string
}
