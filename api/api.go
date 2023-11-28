package api

import (
	"context"
	_ "wow-admin/api/controller"
	"wow-admin/api/router"
	"wow-admin/utils"
)

func Init(ctx context.Context, wait *utils.WaitGroup) {
	router.InitAndStartWebServer(ctx, true, wait)
}
