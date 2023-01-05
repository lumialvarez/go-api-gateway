package mapper

import "github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/handler/grpc/profile/pb"

type Mapper struct {
}

func (m Mapper) ToListRequest(id int64, language string) pb.ListRequest {
	var request pb.ListRequest
	if id > 0 {
		request.ProfileId = &id
	}
	if len(language) > 0 {
		request.ProfileLanguage = &language
	}
	return request
}
