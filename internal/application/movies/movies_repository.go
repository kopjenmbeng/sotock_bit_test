package movies

import (
	"context"

	"github.com/kopjenmbeng/sotock_bit_test/internal/dto"
)

type IMoviesRepository interface {
	Search(ctx context.Context, search string, page int) (result *dto.Searching, status int, err error)
	GetDetail(ctx context.Context, id string) (result *dto.DetailMovie, status int, err error)
}

type MoviesRepository struct {
}

func NewMovieRepository() IMoviesRepository {
	return &MoviesRepository{}
}

func (repo *MoviesRepository) Search(ctx context.Context, search string, page int) (result *dto.Searching, status int, err error){
	return
}
func (repo *MoviesRepository) GetDetail(ctx context.Context, id string) (result *dto.DetailMovie, status int, err error){
	return
}
