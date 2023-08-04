package api

import (
	"context"
	"github.com/mini-dropbox/app/boot"
	"github.com/mini-dropbox/app/common"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	environ, err := common.GetEnv()
	if err != nil {
		log.Fatal("read from env failed", "error", err)
	}

	if err := boot.Init(ctx, environ); err != nil {
		log.Fatal("application boot failed", "error", err)
	}
}
