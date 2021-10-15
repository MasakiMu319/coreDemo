package main

import "coreDemo/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
