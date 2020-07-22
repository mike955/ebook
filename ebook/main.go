package main

import (
	"context"
	"ebook/ebook/conf"
	"ebook/ebook/routers"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	
	gin.SetMode(conf.RUN_MODE)
	engine := gin.New()
	
	// 性能分析 - 正式环境不要使用！！！
	pprof.Register(engine)
	routers.InitRouter(engine)
	endPoint := fmt.Sprintf(":%d", conf.EBOOK_PORT)
	maxHeaderBytes := 1 << 20
	
	server := &http.Server{
		Addr:           endPoint,
		Handler:        engine,
		//ReadTimeout:    conf.READTIMEOUT,
		//WriteTimeout:   conf.WRITETIMEOUT,
		MaxHeaderBytes: maxHeaderBytes,
	}
	
	fmt.Println("|-----------------------------------|")
	fmt.Println("|            Ebook             |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("|    Port" + endPoint + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "        |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")
	
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()
	
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}