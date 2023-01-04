package mapper

import (
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
)

type Mapper struct {
}

func (m Mapper) ToListRequest(id int64, userName string) pb.ListRequest {
	var request pb.ListRequest
	if id > 0 {
		request.UserId = &id
	}
	if len(userName) > 0 {
		request.UserName = &userName
	}
	return request
}
