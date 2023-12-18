package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
	"wow-admin/api/middleware"
	"wow-admin/config"
	"wow-admin/global"
	"wow-admin/utils"
)

type DefaultHttpHandler struct {
	router     *gin.Engine
	isDebug    bool
	fileServer http.Handler
}

func (h *DefaultHttpHandler) Init(router *gin.Engine, debug bool) {
	h.router = router
	h.isDebug = debug
}

func (h *DefaultHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/debug/pprof") && h.isDebug {
		http.DefaultServeMux.ServeHTTP(w, r)
	} else {
		h.router.ServeHTTP(w, r)
	}
}

type WebServer interface {
	Start()
	Stop()
}

type defaultServer struct {
	httpServer *http.Server
	ctx        context.Context
}

var _defaultWebServer WebServer
var _defaultWebRouter = gin.New()

func init() {
	// _defaultWebRouter.Use(middleware.RouterLogger(), gin.Recovery())
	_defaultWebRouter.Use(middleware.RouterLogger())
	_defaultWebRouter.Use(middleware.ErrorRecovery(false)) // 自定义错误处理中间件，不使用gin.Recovery()
	_defaultWebRouter.Use(middleware.Cors())
	_defaultWebRouter.Use(middleware.AuthLogin())
}

func (s *defaultServer) Start() {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Logger.Error(global.SERVER_NAME+" web server started fail", zap.Error(err))
		}
	}()
}

func (s *defaultServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := s.httpServer.Shutdown(ctx); err != nil {
		utils.Logger.Error(global.SERVER_NAME+"server Shutdown ", zap.Error(err))
	}

	utils.Logger.Info(global.SERVER_NAME + " server exiting .....")
}

func InitAndStartWebServer(ctx context.Context, debug bool, wait *utils.WaitGroup) {
	gin.SetMode(gin.ReleaseMode)
	hHandler := &DefaultHttpHandler{}
	hHandler.Init(_defaultWebRouter, debug)

	backPort := config.Cfg.Server.BackPort
	server := &http.Server{
		Addr:         backPort,
		Handler:      hHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	_defaultWebServer = &defaultServer{
		httpServer: server,
		ctx:        ctx,
	}
	_defaultWebServer.Start()
	utils.Logger.Info(global.SERVER_NAME + " server started at http://localhost" + backPort)

	wait.Wrap(func() {
		select {
		case <-ctx.Done():
			_defaultWebServer.Stop()
		}
	})
}

func GetWebRouter() *gin.Engine {
	return _defaultWebRouter
}

func GetWebRouterGroup() *gin.RouterGroup {
	return _defaultWebRouter.Group("/api")
}
