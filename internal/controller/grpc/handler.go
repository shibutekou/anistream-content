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

	titleInfos, err := s.uc.Search(entity.TitleFilter{Opt: in.Filter.Opt, Val: in.Filter.Val})
	for _, pErr := range possibleErrors {
		if errors.Is(errors.Unwrap(err), pErr.err) {
			s.log.Error(err.Error())
			return nil, status.Error(pErr.code, err.Error())
		}
	}

	reply := pb.GetTitleContentReply{}
	reply.TitleContent = transformTitleContent(titleInfos)

	return &reply, status.New(codes.OK, "success").Err()
}

func transformTitleContent(titles []entity.TitleContent) []*pb.TitleContent {
	var result = make([]*pb.TitleContent, 0, len(titles))

	for i := range titles {
		tmp := pb.TitleContent{
			Link:             titles[i].Link,
			Title:            titles[i].Title,
			TitleOrig:        titles[i].TitleOrig,
			OtherTitle:       titles[i].OtherTitle,
			Year:             titles[i].Year,
			KinopoiskID:      titles[i].KinopoiskID,
			ShikimoriID:      titles[i].ShikimoriID,
			ImdbID:           titles[i].IMDbID,
			AnimeStatus:      titles[i].AnimeStatus,
			AnimeDescription: titles[i].AnimeDescription,
			PosterURL:        titles[i].PosterURL,
			Duration:         titles[i].Duration,
			KinopoiskRating:  titles[i].KinopoiskRating,
			ImdbRating:       titles[i].IMDbRating,
			ShikimoriRating:  titles[i].ShikimoriRating,
			PremiereWorld:    titles[i].PremiereWorld,
			EpisodesTotal:    titles[i].EpisodesTotal,
			Writers:          titles[i].Writers,
			Screenshots:      titles[i].Screenshots,
		}

		result = append(result, &tmp)
	}

	return result
}
