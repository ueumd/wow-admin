package api

import (
	"context"
	_ "wow-admin/api/controller"
	"wow-admin/api/router"
	"wow-admin/utils"
)

func Init(ctx context.Context, address string, wait *utils.WaitGroup) {
	router.InitAndStartWebServer(address, ctx, true, wait)
}
