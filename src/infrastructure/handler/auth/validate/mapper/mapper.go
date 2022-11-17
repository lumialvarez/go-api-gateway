package mapper

import (
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/auth/validate/contract"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
)

type Mapper struct {
}

func (m Mapper) ToServiceRequest(dto contract.ValidateAuthRequest) pb.ValidateRequest {
	return pb.ValidateRequest{Token: dto.Token}
}

func (m Mapper) ToDTOResponse(dto *pb.ValidateResponse) contract.ValidateAuthResponse {
	return contract.ValidateAuthResponse{
		UserId:   dto.UserId,
		UserName: dto.UserName,
		Role:     dto.Role,
	}
}
