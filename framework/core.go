package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router map[string]*Tree
	middlewares []ControllerHandler
}

func NewCore() *Core {
	router := map[string]*Tree{}
	router["GET"] = newTree()
	router["PUT"] = newTree()
	router["POST"] = newTree()
	router["DELETE"] = newTree()
	return &Core{router: router}
}

func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}

func (c *Core) Get(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

func (c *Core) FindRouteNodeByRequest(request *http.Request) *node {
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.root.matchNode(uri)
	}
	return nil
}


func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := NewContext(request, response)
	nodes := c.FindRouteNodeByRequest(request)
	if nodes == nil {
		ctx.SetStatus(404, "no found")
		return
	}

	ctx.SetHandlers(nodes.handlers)

	params := nodes.parseParamsFromEndNode(request.URL.Path)
	ctx.params = params

	if err := ctx.Next(); err != nil {
		ctx.SetStaus(500, "inner error")
		return
	}
}