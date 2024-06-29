package pkg

import (
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int   `json:"limit"`
	Page       int   `json:"page"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// GetOffset calculates the offset based on the current page and limit.
//
// No parameters.
// Returns an integer representing the offset.
func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

// GetLimit returns the limit value of the Pagination struct.
//
// If the limit value is 0, it is set to 10.
// The function does not take any parameters.
// It returns an integer representing the limit value.
func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

// GetPage returns the current page value of the Pagination struct.
//
// If the page value is 0, it is set to 1.
// The function does not take any parameters.
// It returns an integer representing the page value.
func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

// Paginate generates pagination for a given value using the provided pagination and database objects.
//
// Parameters:
// - value: The value to paginate.
// - pagination: The pagination object containing the pagination parameters.
// - db: The database object used to perform the pagination.
//
// Returns:
// - A function that takes a database object and applies the pagination offset and limit.
func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var total int64
	db.Model(value).Count(&total)

	pagination.Total = total
	totalPages := int(math.Ceil(float64(total) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
