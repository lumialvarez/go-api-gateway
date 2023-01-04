package mapper

import (
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/auth/update/contract"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
)

type Mapper struct {
}

func (m Mapper) ToServiceRequest(dto contract.UpdateAuthRequest) pb.UpdateRequest {
	user := pb.UpdateRequest_UserUpdate{
		UserId:   dto.User.UserId,
		Name:     dto.User.Name,
		UserName: dto.User.UserName,
		Email:    dto.User.Email,
		Password: dto.User.Password,
		Role:     dto.User.Role,
		Status:   dto.User.Status,
	}
	return pb.UpdateRequest{User: &user}
}
