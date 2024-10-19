package server

import (
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	counter int64

	server *http.Server
	router *gin.Engine
}

func New() *Server {
	router := gin.Default()
	server := &Server {
		router: router,
		counter: int64(0),
	}

	router.GET("/", server.CounterHandler)
	router.GET("/health_checks", server.HealthHandler)

	return server
}

// 成功レスを返却するだけ
func (s *Server) HealthHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"success": true})
}

// リクエスト数を計算
func (s *Server) CounterHandler(ctx *gin.Context) {
	counter := atomic.AddInt64(&s.counter, 2)
	ctx.JSON(200, gin.H{"counter": counter})
}

// サーバ起動・引数アドレスをリッスン
func (s *Server) Start(address string) error {
	s.server = &http.Server{
		Addr: address,
		Handler: s.router,
		ReadTimeout: 10 * time.Second,
	}
	log.Printf("start server on %s", address)

	return s.server.ListenAndServe()
}

// サーバを停止
func (s *Server) Stop() error {
	log.Printf("stop server")
	if s.server != nil {
		return s.server.Close()
	}

	return nil
}