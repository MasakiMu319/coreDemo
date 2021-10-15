package main

import (
	"context"
	"coreDemo/framework"
	"fmt"
	"log"
	"time"
)

func FooControllerHandler(ctx *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	// durationCtx
	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), time.Duration(1 * time.Second))
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// 1. test panic is solved;
		// panic("panic test")

		// 2. test timeout situation;
		// time.Sleep(10 * time.Second)

		// 3. test if everything is ok;
		ctx.Json(200, "ok")

		finish <- struct{}{}
	}()
	select {
	case p := <- panicChan:
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		log.Println(p)
		ctx.Json(500, "panic")
	case <- finish:
		fmt.Println("finish")
	case <- durationCtx.Done():
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		ctx.Json(500, "time out")
		ctx.SetHasTimeout()
	}
	return nil
}