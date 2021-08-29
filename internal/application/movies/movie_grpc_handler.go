package movies

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kopjenmbeng/sotock_bit_test/internal/proto"
	"github.com/kopjenmbeng/sotock_bit_test/internal/utility"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type MoviesGrpcHandler struct {
	srv *grpc.Server

	repository IMoviesRepository
}

func NewMoviesGrpcHandler(log *log.Logger, srv *grpc.Server, dbr sqlx.QueryerContext, dbw *sqlx.DB) MoviesGrpcHandler {
	server := MoviesGrpcHandler{srv: srv, repository: NewMovieRepository(dbr, dbw)}
	proto.RegisterMoviesServer(server.srv, &server)
	log.Printf("register rpc movies is ready !")
	return server
}

func (movie *MoviesGrpcHandler) Search(ctx context.Context, in *proto.SearchRequestMessage) (*proto.SearchResponseMessage, error) {
	data, _, err := movie.repository.Search(ctx, in.Search, int(in.Page))
	if err != nil {
		return nil, utility.GrpcInternalServerError(err)
	}
	var list []*proto.Search
	for _, i := range data.Search {
		list = append(list, &proto.Search{Title: i.Title, Year: i.Year, ImdbID: i.ImdbID, Type: i.Type, Poster: i.Poster})
	}
	result := proto.SearchResponseMessage{Search: list}
	return &result, nil
}

func (movie *MoviesGrpcHandler) GetDetail(ctx context.Context, in *proto.DetailMovieRequestMessage) (*proto.DetailMovieRequestResponse, error) {
	data, _, err := movie.repository.GetDetail(ctx, in.Id)
	if err != nil {
		return nil, utility.GrpcInternalServerError(err)
	}
	var list_rated []*proto.Rating
	for _, i := range data.Ratings {
		list_rated = append(list_rated, &proto.Rating{Source: i.Source, Value: i.Source})
	}
	result := proto.DetailMovieRequestResponse{
		Title:      data.Title,
		Year:       data.Year,
		Rated:      data.Rated,
		Released:   data.Released,
		Runtime:    data.Runtime,
		Genre:      data.Genre,
		Director:   data.Director,
		Writer:     data.Writer,
		Actors:     data.Actors,
		Plot:       data.Plot,
		Language:   data.Language,
		Country:    data.Country,
		Awards:     data.Awards,
		Poster:     data.Poster,
		Ratings:    list_rated,
		Metascore:  data.Metascore,
		ImdbRating: data.ImdbRating,
		ImdbVotes:  data.ImdbVotes,
		ImdbID:     data.ImdbID,
		Type:       data.Type,
		DVD:        data.DVD,
		BoxOffice:  data.BoxOffice,
		Production: data.Production,
		Website:    data.Website,
		Response:   data.Response,
	}
	return &result, nil
}
