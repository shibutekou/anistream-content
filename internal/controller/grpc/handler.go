package grpc

import (
	"context"
	"errors"
	"github.com/vgekko/anistream-content/internal/controller/grpc/pb"
	"github.com/vgekko/anistream-content/internal/entity"
	"github.com/vgekko/anistream-content/internal/usecase"
	"github.com/vgekko/anistream-content/pkg/apperror"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log/slog"
)

type ContentServer struct {
	uc  usecase.ContentUseCase
	log *slog.Logger

	pb.UnimplementedContentServiceServer
}

func NewContentServerGrpc(gserver *grpc.Server, uc usecase.ContentUseCase, log *slog.Logger) {
	cs := ContentServer{uc, log, pb.UnimplementedContentServiceServer{}}

	pb.RegisterContentServiceServer(gserver, &cs)

	reflection.Register(gserver)
}

type errorMap struct {
	err  error      // domain error
	code codes.Code // grpc status code
}

func (s *ContentServer) GetTitleContent(ctx context.Context, in *pb.GetTitleContentRequest) (*pb.GetTitleContentReply, error) {
	possibleErrors := []errorMap{
		{apperror.ErrTitleNotFound, codes.NotFound},
		{apperror.ErrNoSearchParams, codes.InvalidArgument},
		{apperror.ErrMissingOrInvalidToken, codes.Internal},
		{apperror.ErrUnknown, codes.Unknown},
	}

	content, err := s.uc.Search(entity.TitleFilter{Opt: in.Filter.Opt, Val: in.Filter.Val})
	for _, pErr := range possibleErrors {
		if errors.Is(errors.Unwrap(err), pErr.err) {
			s.log.Error(err.Error())
			return nil, status.Error(pErr.code, err.Error())
		}
	}

	reply := pb.GetTitleContentReply{}
	reply.Content = transformTitleContent(content)

	return &reply, status.New(codes.OK, "success").Err()
}

func transformTitleContent(content entity.Content) *pb.Content {
	transformedContent := new(pb.Content)

	for i := range content.Titles {
		tmp := pb.Title{
			Link:             content.Titles[i].Link,
			Title:            content.Titles[i].Title,
			TitleOrig:        content.Titles[i].TitleOrig,
			OtherTitle:       content.Titles[i].OtherTitle,
			Year:             content.Titles[i].Year,
			KinopoiskID:      content.Titles[i].KinopoiskID,
			ShikimoriID:      content.Titles[i].ShikimoriID,
			ImdbID:           content.Titles[i].IMDbID,
			AnimeStatus:      content.Titles[i].AnimeStatus,
			AnimeDescription: content.Titles[i].AnimeDescription,
			PosterURL:        content.Titles[i].PosterURL,
			Duration:         content.Titles[i].Duration,
			KinopoiskRating:  content.Titles[i].KinopoiskRating,
			ImdbRating:       content.Titles[i].IMDbRating,
			ShikimoriRating:  content.Titles[i].ShikimoriRating,
			PremiereWorld:    content.Titles[i].PremiereWorld,
			EpisodesTotal:    content.Titles[i].EpisodesTotal,
			Writers:          content.Titles[i].Writers,
			Screenshots:      content.Titles[i].Screenshots,
		}

		transformedContent.Titles = append(transformedContent.Titles, &tmp)
	}

	transformedContent.Total = content.Total

	return transformedContent
}
