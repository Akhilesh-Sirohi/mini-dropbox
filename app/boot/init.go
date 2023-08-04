package boot

import (
	"context"
	"github.com/mini-dropbox/app/config"
	"github.com/mini-dropbox/app/providers"
	"github.com/mini-dropbox/app/router"
)

func Init(ctx context.Context, env string) error {
	err := config.Init(env)
	if err != nil {
		return err
	}

	cfg := config.GetConfig()
	err = providers.DB.Setup(cfg.Db)
	if err != nil {
		return err
	}

	router.Initialize(&cfg)

	return nil
}
