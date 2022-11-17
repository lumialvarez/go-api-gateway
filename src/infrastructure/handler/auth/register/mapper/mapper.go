package mapper

import (
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/auth/register/contract"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
)

type Mapper struct {
}

func (m Mapper) ToServiceRequest(dto contract.RegisterAuthRequest) pb.RegisterRequest {
	return pb.RegisterRequest{
		Name:     dto.Name,
		UserName: dto.UserName,
		Email:    dto.Email,
		Password: dto.Password,
		Role:     dto.Role,
	}
}

func (m Mapper) ToDTOResponse(dto *pb.RegisterResponse) contract.RegisterAuthResponse {
	return contract.RegisterAuthResponse{
		UserId: dto.UserId,
	}
}
