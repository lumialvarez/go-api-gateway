package profile

import (
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/handler/grpc/profile/pb"
	"google.golang.org/grpc"
	"log"
)

type ServiceClient struct {
	Client pb.ProfileServiceClient
}

func InitServiceClient(c *config.Config) *ServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ProfileSvcUrl, grpc.WithInsecure())

	if err != nil {
		log.Print("Could not connect to Profile Service:", err)
	} else {
		log.Print("Connected to Profile Service successfully")
	}

	serviceClient := &ServiceClient{
		Client: pb.NewProfileServiceClient(cc),
	}

	return serviceClient
}
