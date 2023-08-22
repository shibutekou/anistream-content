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
	uc  usecase.InfoUseCase
	log *slog.Logger

	pb.UnimplementedContentServiceServer
}

func NewContentServerGrpc(gserver *grpc.Server, uc usecase.InfoUseCase, log *slog.Logger) {
	cs := ContentServer{uc, log, pb.UnimplementedContentServiceServer{}}

	pb.RegisterContentServiceServer(gserver, &cs)

	reflection.Register(gserver)
}

func (s *ContentServer) GetTitleInfo(ctx context.Context, in *pb.GetTitleInfoRequest) (*pb.GetTitleInfoReply, error) {
	titleInfos, err := s.uc.Search(entity.TitleFilter{Opt: in.Filter.Opt, Val: in.Filter.Val})
	if err != nil {
		s.log.Error("grpc: GetTitleInfo: ", sl.Err(err))
		return nil, err
	}

	reply := pb.GetTitleInfoReply{}
	reply.TitleInfo = transformTitleInfo(titleInfos)

	return &reply, nil
}

func transformTitleInfo(info []entity.TitleInfo) []*pb.TitleInfo {
	var result = make([]*pb.TitleInfo, 0, len(info))

	for i := range info {
		tmp := pb.TitleInfo{
			Link:        info[i].Link,
			Title:       info[i].Title,
			TitleOrig:   info[i].TitleOrig,
			OtherTitle:  info[i].OtherTitle,
			Year:        info[i].Year,
			KinopoiskID: info[i].KinopoiskID,
			ShikimoriID: info[i].ShikimoriID,
			ImdbID:      info[i].IMDbID,
			Screenshots: info[i].Screenshots,
		}

		result = append(result, &tmp)
	}

	return result
}
