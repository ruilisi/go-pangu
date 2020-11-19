package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// gorm.Model 的定义 带有uuid 和 创建/修改时间
type Model struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//搜索结果结构,用来返回错误内容
type SearchResult struct {
	Error  error
	Status int
}

const (
	ERROR     = -1
	FOUND     = 0
	NOT_FOUND = 1
)

func Result(err error) SearchResult {
	var status int
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = NOT_FOUND
		} else {
			status = ERROR
		}
	}
	return SearchResult{Error: err, Status: status}
}

func (r *SearchResult) Err() error {
	return r.Error
}

func (r *SearchResult) Found() bool {
	return r.Status == FOUND
}

func (r *SearchResult) NotFound() bool {
	return r.Status == NOT_FOUND
}

func (r *SearchResult) DBError() bool {
	return r.Status == ERROR
}
