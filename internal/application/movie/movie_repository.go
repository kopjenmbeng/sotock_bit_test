package movie

import (
	"context"

	"github.com/kopjenmbeng/sotock_bit_test/internal/dto"
)

type IMovieRepository interface {
	Search(ctx context.Context,search string ,page int)(result *dto.Searching,status int,err error)
	GetDetail(ctx context.Context,id string)(result *dto.DetailMovie,status int,err error)
}

type MovieRepository struct{

}

func NewMovieRepository()IMovieRepository{
	return &MovieRepository{}
}

func (repo *MovieRepository) Search(ctx context.Context,search string ,page int)(result *dto.Searching,status int,err error)
func (repo *MovieRepository)	GetDetail(ctx context.Context,id string)(result *dto.DetailMovie,status int,err error)