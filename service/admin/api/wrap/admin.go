package wrap

import (
	"fmt"
	"mall-admin/api/internal/handler"
	"mall-admin/api/internal/svc"
	"mall-admin/pkg"
	"mall-pkg/api"

	"github.com/zeromicro/go-zero/rest"
	"go.uber.org/zap"
)

func Start(c pkg.ApiConfig, log *zap.Logger) error {
	ctx := svc.NewServiceContext(c, log)
	corsmd := api.NewCorsMiddleware()
	cors := rest.WithNotAllowedHandler(corsmd)
	server := rest.MustNewServer(c.RestConf, cors)
	server.Use(corsmd.Handle)

	err := ctx.Init()
	if err != nil {
		return err
	}
	handler.Register(server, ctx)

	defer server.Stop()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
	return nil
}
