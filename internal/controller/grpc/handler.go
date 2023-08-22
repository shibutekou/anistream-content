package grpc

import (
	"context"
	"github.com/vgekko/anistream-content/internal/controller/grpc/pb"
	"github.com/vgekko/anistream-content/internal/entity"
	"github.com/vgekko/anistream-content/internal/usecase"
	"github.com/vgekko/anistream-content/pkg/logger/sl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

func (s *ContentServer) GetTitleContent(ctx context.Context, in *pb.GetTitleContentRequest) (*pb.GetTitleContentReply, error) {
	titleInfos, err := s.uc.Search(entity.TitleFilter{Opt: in.Filter.Opt, Val: in.Filter.Val})
	if err != nil {
		s.log.Error("grpc: GetTitleContent: ", sl.Err(err))
		return nil, err
	}

	reply := pb.GetTitleContentReply{}
	reply.TitleContent = transformTitleContent(titleInfos)

	return &reply, nil
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
