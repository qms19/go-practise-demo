package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)
func main()  {
	g := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("qms19 %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	g.GET("/", func(c *gin.Context) {
		c.JSON(200,"test")
	})
	srv := &http.Server{
		Addr: ":8080" ,
		Handler: g ,
	}
	srv2 := &http.Server{
		Addr:              ":8081",
		Handler:           g ,
	}

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer  wg.Done()
		quit := make(chan os.Signal, 2)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<- quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx) ; err != nil {
			log.Fatalf("Shutdown server failed: %s\n", err)
		}
		if err := srv2.Shutdown(ctx) ; err != nil {
			log.Fatalf("Shutdown server failed: %s\n", err)
		}
		log.Println("Shutdown server ")
	}()

	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe() ; err != nil && !errors.Is(err, http.ErrServerClosed){
			log.Fatalf("listen: %s\n",err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := srv2.ListenAndServe() ; err != nil && !errors.Is(err, http.ErrServerClosed){
			log.Fatalf("listen: %s\n",err)
		}
	}()

	wg.Wait()
}

