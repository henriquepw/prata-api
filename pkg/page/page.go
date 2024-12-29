package page

import (
	"math"
)

type PageMeta struct {
	TotalItems   int `json:"totalItems"`
	CurrentPage  int `json:"currentPage"`
	ItemsPerPage int `json:"itemsPerPage"`
	ItemCount    int `json:"itemCount"`
	TotalPages   int `json:"totalPages"`
}

type Page[T any] struct {
	Items []T      `json:"items"`
	Meta  PageMeta `json:"meta"`
}

type Cursor[T any] struct {
	Next  *string `json:"next"`
	Items []T     `json:"items"`
}

func New[T any](items []T, page, limit, totalItems int64) *Page[T] {
	meta := PageMeta{
		TotalItems:   int(totalItems),
		CurrentPage:  int(page),
		ItemsPerPage: int(limit),
		ItemCount:    len(items),
		TotalPages:   int(math.Ceil(float64(totalItems) / float64(limit))),
	}

	return &Page[T]{
		Items: items,
		Meta:  meta,
	}
}

func NewEmpty[T any](page, limit int64) *Page[T] {
	return &Page[T]{
		Items: []T{},
		Meta: PageMeta{
			TotalItems:   0,
			CurrentPage:  int(page),
			ItemsPerPage: int(limit),
			ItemCount:    0,
			TotalPages:   1,
		},
	}
}

func NewEmptyCursor[T any]() *Cursor[T] {
	return &Cursor[T]{
		Items: []T{},
		Next:  nil,
	}
}

func NewCursor[T any](items []T, limit int64, next func(item T) string) *Cursor[T] {
	if len(items) > int(limit) {
		cursor := next(items[limit])
		return &Cursor[T]{
			Items: items[:limit],
			Next:  &cursor,
		}
	}

	return &Cursor[T]{
		Items: items,
		Next:  nil,
	}
}

func Skip(page, limit int64) int64 {
	return (page - 1) * limit
}

type Queryable[T any] interface {
	// Initialize initializes the queryable
	Initialize(f ...map[string]T)
	// GetFilter returns the filter
	GetFilter() map[string]T
}

type PaginationQueryable[T any] interface {
	Queryable[T]
	// GetPage returns the page number
	GetPage() int64
	// GetLimit returns the limit
	GetLimit() int64
	// GetSortBy returns the sort by
	GetSortBy() string
	// GetOrderBy returns the order by
	GetOrderBy() string
}

type FilterQuery[T any] struct {
	Filter map[string]T // Filtros para a busca
}

func (q *FilterQuery[T]) Initialize(f ...map[string]T) {
	if q.Filter == nil {
		if len(f) > 0 {
			q.Filter = f[0]
		}
	}
}

func (q *FilterQuery[T]) GetFilter() map[string]T {
	return q.Filter
}

type PaginationQuery[T any] struct {
	Page    int64        `json:"page"`     // número da pagina
	Limit   int64        `json:"limit"`    // quantidade de itens por pagina
	SortBy  string       `json:"sort_by"`  // nome da propriedade
	OrderBy string       `json:"order_by"` // "asc" ou "desc"
	Filter  map[string]T // Filtros para a busca
}

func (q *PaginationQuery[T]) Initialize(f ...map[string]T) {
	if q.Page < 1 {
		q.Page = 1
	}
	if q.Limit < 1 {
		q.Limit = 10
	}
	if q.SortBy == "" {
		q.SortBy = "createdAt"
	}
	if q.OrderBy == "" {
		q.OrderBy = "ASC"
	}
	if q.Filter == nil {
		if len(f) > 0 {
			q.Filter = f[0]
		}
	}
}

func (q *PaginationQuery[T]) GetFilter() map[string]T {
	return q.Filter
}

func (q *PaginationQuery[T]) GetPage() int64 {
	return q.Page
}

func (q *PaginationQuery[T]) GetLimit() int64 {
	return q.Limit
}

func (q *PaginationQuery[T]) GetSortBy() string {
	return q.SortBy
}

func (q *PaginationQuery[T]) GetOrderBy() string {
	return q.OrderBy
}
