package middleware

import "coreDemo/framework"

func Recovery() framework.ControllerHandler {
	return func(c *framework.Context) error {
		// 捕获 c.Next 过程的 panic
		defer func() {
			if err := recover(); err != nil {
				c.Json(500, err)
			}
		}()
		c.Next()
		return nil
	}
}
