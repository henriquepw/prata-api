package web

type Middleware func(HandlerFn) HandlerFn

func ApplyMiddlewares(mw []Middleware, handler HandlerFn) HandlerFn {
	for i := len(mw) - 1; i >= 0; i-- {
		h := mw[i]
		if h != nil {
			handler = h(handler)
		}
	}

	return handler
}
