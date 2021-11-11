package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ueumd/logger"
	"net/http"
	"strings"
	"wow-admin/api/middleware"
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
	_defaultWebRouter.Use(middleware.RouterLogger(), gin.Recovery())
	_defaultWebRouter.Use(middleware.Cors())
	_defaultWebRouter.Use(middleware.AuthLogin())
}

func (s *defaultServer) Start() {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			logger.Fatalf(global.SERVER_NAME + " web server started fail:%v", err)
		}
	}()
}


func (s *defaultServer) Stop() {
	s.httpServer.Close()
	logger.Infoln(global.SERVER_NAME + " server stopped .....")
}


func InitAndStartWebServer(address string, ctx context.Context, debug bool, wait *utils.WaitGroup) {
	gin.SetMode(gin.ReleaseMode)
	hHandler := &DefaultHttpHandler{}
	hHandler.Init(_defaultWebRouter, debug)
	server := &http.Server{Addr: address, Handler: hHandler}

	_defaultWebServer = &defaultServer{
		httpServer: server,
		ctx:        ctx,
	}
	_defaultWebServer.Start()
	logger.InfoF(global.SERVER_NAME + " server started at %s ...", address)
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