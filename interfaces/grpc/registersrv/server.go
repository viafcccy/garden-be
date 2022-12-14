package registersrv

import (
	"google.golang.org/grpc"

	service "github.com/viafcccy/garden-be/application/service/user"
	"github.com/viafcccy/garden-be/interfaces/grpc/proto"
	"github.com/viafcccy/garden-be/interfaces/grpc/server/user"
)

func RegisterSrv(server *grpc.Server) {
	proto.RegisterUserServer(server, &user.UserServer{
		UserSrv: service.NewUserService(),
	})
}
