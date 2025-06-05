package utils

import (
	"math"

	"gorm.io/gorm"
)

// Paginate 通用的分页方法
func Paginate(pageNum int, pageSize int, db *gorm.DB) *gorm.DB {
	// 确保页码和每页数量是有效的
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100 // 限制最大每页数量
	}

	offset := (pageNum - 1) * pageSize
	return db.Offset(offset).Limit(pageSize)
}

func TotalPage(total int64, pageSize int) int {
	return int(math.Ceil(float64(total) / float64(pageSize)))
}
