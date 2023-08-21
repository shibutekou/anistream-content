package grpc

import (
	"context"
	"github.com/vgekko/anistream-content/internal/entity"
	"github.com/vgekko/anistream-content/internal/usecase"
	"github.com/vgekko/anistream-content/pkg/logger/sl"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/vgekko/anistream-content/internal/controller/grpc/pb"
)

type ContentServer struct {
	i   usecase.InfoUseCase
	l   usecase.LinkUseCase
	log *slog.Logger

	pb.UnimplementedContentServiceServer
}

func NewContentServerGrpc(gserver *grpc.Server, i usecase.InfoUseCase, l usecase.LinkUseCase, log *slog.Logger) {
	cs := ContentServer{i, l, log, pb.UnimplementedContentServiceServer{}}

	pb.RegisterContentServiceServer(gserver, &cs)

	reflection.Register(gserver)
}

func (s *ContentServer) GetTitleInfo(ctx context.Context, in *pb.GetTitleInfoRequest) (*pb.GetTitleInfoReply, error) {
	titleInfos, err := s.i.Search(entity.TitleFilter{Option: in.Filter.Svc, Value: in.Filter.Val})
	if err != nil {
		s.log.Error("grpc: GetTitleInfo: ", sl.Err(err))
		return nil, err
	}

	reply := pb.GetTitleInfoReply{}
	reply.TitleInfo = transformTitleInfo(titleInfos)

	return &reply, nil
}

func (s *ContentServer) GetLink(ctx context.Context, in *pb.GetLinkRequest) (*pb.GetLinkReply, error) {
	link, err := s.l.Search(entity.TitleFilter{Option: in.Filter.Svc, Value: in.Filter.Val})
	if err != nil {
		s.log.Error("grpc: GetLink: ", sl.Err(err))
		return nil, err
	}

	replyLink := &pb.Link{Url: link.URL}

	return &pb.GetLinkReply{Link: replyLink}, nil
}

func transformTitleInfo(info []entity.TitleInfo) []*pb.TitleInfo {
	var result = make([]*pb.TitleInfo, 0, len(info))

	for i, _ := range info {
		tmp := pb.TitleInfo{
			Title:       info[i].Title,
			TitleOrig:   info[i].TitleOrig,
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
