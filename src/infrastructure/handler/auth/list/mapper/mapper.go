package mapper

import (
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/auth/list/contract"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
)

type Mapper struct {
}

func (m Mapper) ToServiceRequest(id int64, userName string) pb.ListRequest {
	var request pb.ListRequest
	if id > 0 {
		request.UserId = &id
	}
	if len(userName) > 0 {
		request.UserName = &userName
	}
	return request
}

func (m Mapper) ToDTOResponse(dto *pb.ListResponse) contract.ListAuthResponse {
	var list []contract.User
	for _, dtoUser := range dto.GetUsers() {
		user := contract.User{
			UserId:   dtoUser.GetUserId(),
			Name:     dtoUser.GetName(),
			UserName: dtoUser.GetUserName(),
			Email:    dtoUser.GetEmail(),
			Role:     dtoUser.GetRole(),
			Status:   dtoUser.GetStatus(),
		}
		list = append(list, user)
	}
	return contract.ListAuthResponse{Users: list}
}
