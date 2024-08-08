package utils

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

// go-zero 路由组封装
//
//使用示例
type Router struct {
	engine      *rest.Server
	middlewares []rest.Middleware
	opts        []rest.RouteOption
}

func NewRouter(engine *rest.Server) *Router {
	r := &Router{}
	r.engine = engine
	return r
}

func (l *Router) Use(middleware ...rest.Middleware) *Router {
	l.middlewares = append(l.middlewares, middleware...)
	return l
}

func (l *Router) RouterOption(opts ...rest.RouteOption) *Router {
	l.opts = append(l.opts, opts...)
	return l
}

func (l *Router) Post(path string, handler http.HandlerFunc) *Router {
	if len(l.middlewares) > 0 {
		l.engine.AddRoutes(
			rest.WithMiddlewares(l.middlewares, rest.Route{
				Method:  http.MethodPost,
				Path:    path,
				Handler: handler,
			}),
			l.opts...,
		)
	} else {
		l.engine.AddRoutes([]rest.Route{{
			Method:  http.MethodPost,
			Path:    path,
			Handler: handler,
		}}, l.opts...)
	}
	return l
}

func (l *Router) Get(path string, handler http.HandlerFunc) *Router {
	if len(l.middlewares) > 0 {
		l.engine.AddRoutes(
			rest.WithMiddlewares(l.middlewares, rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: handler,
			}),
			l.opts...,
		)
	} else {
		l.engine.AddRoutes([]rest.Route{{
			Method:  http.MethodGet,
			Path:    path,
			Handler: handler,
		}}, l.opts...)
	}
	return l
}

func (l *Router) Put(path string, handler http.HandlerFunc) *Router {
	if len(l.middlewares) > 0 {
		l.engine.AddRoutes(
			rest.WithMiddlewares(l.middlewares, rest.Route{
				Method:  http.MethodPut,
				Path:    path,
				Handler: handler,
			}),
			l.opts...,
		)
	} else {
		l.engine.AddRoutes([]rest.Route{{
			Method:  http.MethodPut,
			Path:    path,
			Handler: handler,
		}}, l.opts...)
	}
	return l
}

func (l *Router) Delete(path string, handler http.HandlerFunc) *Router {
	if len(l.middlewares) > 0 {
		l.engine.AddRoutes(
			rest.WithMiddlewares(l.middlewares, rest.Route{
				Method:  http.MethodDelete,
				Path:    path,
				Handler: handler,
			}),
			l.opts...,
		)
	} else {
		l.engine.AddRoutes([]rest.Route{{
			Method:  http.MethodDelete,
			Path:    path,
			Handler: handler,
		}}, l.opts...)
	}
	return l
}
