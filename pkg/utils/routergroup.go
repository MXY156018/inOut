package utils

import (
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest"
)

// go-zero 路由组封装
//
//使用示例
type RouterGroup struct {
	engine      *rest.Server
	middlewares []rest.Middleware
	opts        []rest.RouteOption
	name        string
}

func NewRouterGroup(name string, engine *rest.Server) *RouterGroup {
	r := &RouterGroup{}
	r.name = name
	r.engine = engine
	r.name = "/" + strings.TrimLeft(name, "/")

	return r
}

func (l *RouterGroup) Group(name string) *RouterGroup {
	r := NewRouterGroup(l.name+"/"+strings.TrimLeft(name, "/"), l.engine)
	r.Use(l.middlewares...)
	r.RouterOption(l.opts...)
	return r
}

func (l *RouterGroup) Use(middleware ...rest.Middleware) *RouterGroup{
	l.middlewares = append(l.middlewares, middleware...)
	return l
}

func (l *RouterGroup) RouterOption(opts ...rest.RouteOption) *RouterGroup {
	l.opts = append(l.opts, opts...)
	return l
}

func (l *RouterGroup) combinePath(path string) string {
	return l.name + "/" + strings.TrimLeft(path, "/")
}

func (l *RouterGroup) Post(path string, handler http.HandlerFunc) *RouterGroup {
	if len(l.middlewares) > 0 {
		l.engine.AddRoutes(
			rest.WithMiddlewares(l.middlewares, rest.Route{
				Method:  http.MethodPost,
				Path:    l.combinePath(path),
				Handler: handler,
			}),
			l.opts...,
		)
	} else {
		l.engine.AddRoutes([]rest.Route{{
			Method:  http.MethodPost,
			Path:    l.combinePath(path),
			Handler: handler,
		}}, l.opts...)
	}
	return l
}

func (l *RouterGroup) Get(path string, handler http.HandlerFunc) *RouterGroup {
	if len(l.middlewares) > 0 {
		l.engine.AddRoutes(
			rest.WithMiddlewares(l.middlewares, rest.Route{
				Method:  http.MethodGet,
				Path:    l.combinePath(path),
				Handler: handler,
			}),
			l.opts...,
		)
	} else {
		l.engine.AddRoutes([]rest.Route{{
			Method:  http.MethodGet,
			Path:    l.combinePath(path),
			Handler: handler,
		}}, l.opts...)
	}
	return l
}

func (l *RouterGroup) Put(path string, handler http.HandlerFunc) *RouterGroup {
	if len(l.middlewares) > 0 {
		l.engine.AddRoutes(
			rest.WithMiddlewares(l.middlewares, rest.Route{
				Method:  http.MethodPut,
				Path:    l.combinePath(path),
				Handler: handler,
			}),
			l.opts...,
		)
	} else {
		l.engine.AddRoutes([]rest.Route{{
			Method:  http.MethodPut,
			Path:    l.combinePath(path),
			Handler: handler,
		}}, l.opts...)
	}
	return l
}

func (l *RouterGroup) Delete(path string, handler http.HandlerFunc) *RouterGroup {
	if len(l.middlewares) > 0 {
		l.engine.AddRoutes(
			rest.WithMiddlewares(l.middlewares, rest.Route{
				Method:  http.MethodDelete,
				Path:    l.combinePath(path),
				Handler: handler,
			}),
			l.opts...,
		)
	} else {
		l.engine.AddRoutes([]rest.Route{{
			Method:  http.MethodDelete,
			Path:    l.combinePath(path),
			Handler: handler,
		}}, l.opts...)
	}
	return l
}
