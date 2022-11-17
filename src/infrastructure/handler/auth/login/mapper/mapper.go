package mapper

import (
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/auth/login/contract"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
)

type Mapper struct {
}

func (m Mapper) ToServiceRequest(dto contract.LoginAuthRequest) pb.LoginRequest {
	return pb.LoginRequest{
		UserName: dto.UserName,
		Password: dto.Password,
	}
}

func (m Mapper) ToDTOResponse(dto pb.LoginResponse) contract.LoginAuthResponse {
	return contract.LoginAuthResponse{
		Token:    dto.Token,
		UserId:   dto.UserId,
		UserName: dto.UserName,
		Role:     dto.Role,
	}
}
