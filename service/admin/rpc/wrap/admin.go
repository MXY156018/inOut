package wrap

import (
	"fmt"
	"mall-admin/pkg"

	"mall-admin/rpc/admin"
	"mall-admin/rpc/internal/server"
	"mall-admin/rpc/internal/svc"

	"go.uber.org/zap"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start(c pkg.RpcConfig, log *zap.Logger) error {
	ctx := svc.NewServiceContext(c, log)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		admin.RegisterAuthServer(grpcServer, server.NewAuthServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	fmt.Printf("Starting admin_rpc at %s...\n", c.ListenOn)
	s.Start()
	return nil
}
