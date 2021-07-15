package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/qms19/go-practise-demo/error/pkg/code"
	"github.com/qms19/go-practise-demo/error/pkg/core"
)

func main()  {
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
			core.WriteResponse(c,errors.WithCode(code.ErrSuccess,"qms"),"test")
	})
	g.Run()
}
