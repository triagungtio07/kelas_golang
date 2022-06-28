package utils

import (
	"fmt"
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DEFAULT_LIMIT = 10

type (
	Total struct {
		Rows  int64 `json:"rows"`
		Pages int   `json:"pages"`
	}

	Paginator struct {
		Query      string      `json:"query" query:"query"`
		Limit      int         `json:"limit" query:"limit"`
		Page       int         `json:"page" query:"page"`
		Sorts      []string    `json:"sorts" query:"sorts"`
		Directions []string    `json:"directions" query:"directions"`
		Total      Total       `json:"total"`
		Data       interface{} `json:"data"`
	}
)

func NewPaginator(ctx *fiber.Ctx) *Paginator {
	pagination := Paginator{
		Total:      Total{},
		Sorts:      []string{},
		Directions: []string{},
	}

	ctx.QueryParser(&pagination)
	if pagination.Limit < DEFAULT_LIMIT {
		pagination.Limit = DEFAULT_LIMIT
	}

	if pagination.Page == 0 {
		pagination.Page = 1
	}

	return &pagination
}

func (p *Paginator) Paginate(value interface{}, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	db.Model(value).Count(&p.Total.Rows)

	p.Total.Pages = int(math.Ceil(float64(p.Total.Rows) / float64(p.Limit)))
	p.Data = value

	return func(db *gorm.DB) *gorm.DB {
		if len(p.Directions) == len(p.Sorts) {
			for k, v := range p.Sorts {
				direction := p.Directions[k]
				if direction != "asc" {
					direction = "desc"
				}

				db.Order(fmt.Sprintf("%s %s", v, direction))
			}
		}

		return db.Offset((p.Page - 1) * p.Limit).Limit(p.Limit)
	}
}
