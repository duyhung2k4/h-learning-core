package service

import (
	authservice "app/cmd/core-service/service/auth"
	videoservice "app/cmd/core-service/service/video"
	"app/internal/connection"
	"app/internal/entity"
	query "app/pkg/query/basic"
	rawquery "app/pkg/query/raw"
)

func Register() Service {
	return Service{
		AuthService:  authservice.Register(),
		VideoService: videoservice.Register(),

		QueryCourse:    query.Register[entity.Course](),
		RawQueryCourse: rawquery.Register[entity.Course](),

		QueryChapter:    query.Register[entity.Chapter](),
		RawQueryChapter: rawquery.Register[entity.Chapter](),

		QueryLession:    query.Register[entity.Lession](),
		RawQueryLession: rawquery.Register[entity.Lession](),

		QueryVideoLession:    query.Register[entity.VideoLession](),
		RawQueryVideoLession: rawquery.Register[entity.VideoLession](),

		GrpcClientQuizz: connection.GetGrpcClientQuizz(),
	}
}
