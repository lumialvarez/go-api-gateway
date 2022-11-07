package authentication

import (
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	"google.golang.org/grpc"
	"log"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) *ServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		log.Fatalln("Could not connect to Authorization Service:", err)
	}
	log.Print("Connected to Authorization Service successfully")

	serviceClient := &ServiceClient{
		Client: pb.NewAuthServiceClient(cc),
	}

	return serviceClient
}
