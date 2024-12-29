package page_test

import (
	"testing"

	"github.com/henriquepw/pobrin-api/pkg/page"
	"github.com/stretchr/testify/assert"
)

type test struct {
	list       []string
	page       int
	limit      int
	totalItems int
	itemCount  int
	totalPages int
}

func TestPage(t *testing.T) {
	tests := []test{
		{[]string{"1", "2"}, 1, 2, 3, 2, 2},
		{[]string{"1", "2"}, 1, 4, 2, 2, 1},
		{[]string{"1", "2"}, 2, 100, 102, 2, 2},
	}

	for _, tt := range tests {
		p := page.New(
			tt.list,
			int64(tt.page),
			int64(tt.limit),
			int64(tt.totalItems),
		)

		assert.Equal(t, tt.list, p.Items)
		assert.Equal(t, tt.totalItems, p.Meta.TotalItems)
		assert.Equal(t, tt.page, p.Meta.CurrentPage)
		assert.Equal(t, tt.limit, p.Meta.ItemsPerPage)
		assert.Equal(t, tt.itemCount, p.Meta.ItemCount)
		assert.Equal(t, tt.totalPages, p.Meta.TotalPages)
	}
}

type queryable struct {
	Filter map[string]any
}

func TestQueryable(t *testing.T) {
	tests := []queryable{
		{Filter: map[string]any{"id": 1}},
		{Filter: map[string]any{"id": "12313"}},
		{Filter: map[string]any{"id": false}},
	}

	t.Run("GetFilter returns filter", func(t *testing.T) {
		for _, tt := range tests {
			f := page.FilterQuery[any]{
				Filter: tt.Filter,
			}

			assert.Equal(t, tt.Filter, f.GetFilter())
		}
	})

	t.Run("Initialize set filter attribute and GetFilter returns filter", func(t *testing.T) {
		for _, tt := range tests {
			f := page.FilterQuery[any]{}

			f.Initialize(tt.Filter)

			assert.Equal(t, tt.Filter, f.GetFilter())
		}
	})

	t.Run("Initialize set filter attribute and GetFilter returns filter", func(t *testing.T) {
		for _, tt := range tests {
			f := page.FilterQuery[any]{
				Filter: tt.Filter,
			}

			f.Initialize()

			assert.Equal(t, tt.Filter, f.GetFilter())
		}
	})

	t.Run("Initialize receives with no filter and when has called nothing happens with filter", func(t *testing.T) {
		for range tests {
			f := page.FilterQuery[any]{}

			f.Initialize()

			assert.Nil(t, f.GetFilter())
		}
	})
}

type paginationQuery struct {
	Filter map[string]any
}

func TestPaginationQueryable(t *testing.T) {
	tests := []paginationQuery{
		{Filter: map[string]any{"id": 1}},
		{Filter: map[string]any{"id": "12313"}},
		{Filter: map[string]any{"id": false}},
	}

	t.Run("Creates Pagination Queryable", func(t *testing.T) {
		for _, tt := range tests {
			f := page.PaginationQuery[any]{
				Filter:  tt.Filter,
				Page:    1,
				Limit:   20,
				SortBy:  "id",
				OrderBy: "ASC",
			}

			assert.Equal(t, tt.Filter, f.GetFilter())
			assert.Equal(t, int64(1), f.GetPage())
			assert.Equal(t, int64(20), f.GetLimit())
			assert.Equal(t, "id", f.GetSortBy())
			assert.Equal(t, "ASC", f.GetOrderBy())
		}
	})

	t.Run("Initialize set filter attribute and GetFilter returns filter", func(t *testing.T) {
		for _, tt := range tests {
			f := page.PaginationQuery[any]{}

			f.Initialize(tt.Filter)

			assert.Equal(t, tt.Filter, f.GetFilter())
			assert.Equal(t, int64(1), f.GetPage())
			assert.Equal(t, int64(10), f.GetLimit())
			assert.Equal(t, "createdAt", f.GetSortBy())
			assert.Equal(t, "ASC", f.GetOrderBy())
		}
	})

	t.Run("Initialize set filter attribute and GetFilter returns filter", func(t *testing.T) {
		for _, tt := range tests {
			f := page.PaginationQuery[any]{
				Filter:  tt.Filter,
				Page:    1,
				Limit:   10,
				SortBy:  "id",
				OrderBy: "ASC",
			}

			f.Initialize()

			assert.Equal(t, tt.Filter, f.GetFilter())
			assert.Equal(t, int64(1), f.GetPage())
			assert.Equal(t, int64(10), f.GetLimit())
			assert.Equal(t, "id", f.GetSortBy())
			assert.Equal(t, "ASC", f.GetOrderBy())
		}
	})

	t.Run("Initialize receives with no filter and when has called nothing happens with filter", func(t *testing.T) {
		for range tests {
			f := page.PaginationQuery[any]{}

			f.Initialize()

			assert.Nil(t, f.GetFilter())
		}
	})
}
