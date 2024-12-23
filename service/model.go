package service

import (
	authservice "app/cmd/core-service/service/auth"
	videoservice "app/cmd/core-service/service/video"
	"app/internal/entity"
	query "app/pkg/query/basic"
	rawquery "app/pkg/query/raw"
)

type Service struct {
	AuthService  authservice.AuthService
	VideoService videoservice.VideoService

	QueryCourse    query.QueryService[entity.Course]
	RawQueryCourse rawquery.QueryRawService[entity.Course]

	QueryChapter    query.QueryService[entity.Chapter]
	RawQueryChapter rawquery.QueryRawService[entity.Chapter]

	QueryLession    query.QueryService[entity.Lession]
	RawQueryLession rawquery.QueryRawService[entity.Lession]

	QueryVideoLession    query.QueryService[entity.VideoLession]
	RawQueryVideoLession rawquery.QueryRawService[entity.VideoLession]
}
