package page

type Cursor[T any] struct {
	Next  *string `json:"next"`
	Items []T     `json:"items"`
}

func NewEmpty[T any]() Cursor[T] {
	return Cursor[T]{
		Items: []T{},
		Next:  nil,
	}
}

func New[T any](items []T, limit int, next func(item T) string) Cursor[T] {
	if items == nil {
		items = []T{}
	}

	if limit > 0 && len(items) > limit {
		cursor := next(items[limit])
		return Cursor[T]{
			Items: items[:limit],
			Next:  &cursor,
		}
	}

	return Cursor[T]{
		Items: items,
		Next:  nil,
	}
}
