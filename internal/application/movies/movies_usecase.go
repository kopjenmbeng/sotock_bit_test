package movies

import (
	"context"

	"github.com/kopjenmbeng/sotock_bit_test/internal/dto"
)

type IMoviesUsecase interface {
	Search(ctx context.Context, search string, page int) (result *dto.Searching, status int, err error)
	GetDetail(ctx context.Context, id string) (result *dto.DetailMovie, status int, err error)
}

type MoviesUsecase struct {
	repository IMoviesRepository
}

func NewMoviesUsecase(repo IMoviesRepository) IMoviesUsecase {
	return &MoviesUsecase{repository: repo}
}

func (use_case *MoviesUsecase)Search(ctx context.Context, search string, page int) (result *dto.Searching, status int, err error){
	result,status,err=use_case.repository.Search(ctx,search,page)
	return
}
func (use_case *MoviesUsecase)	GetDetail(ctx context.Context, id string) (result *dto.DetailMovie, status int, err error){
	result,status,err=use_case.repository.GetDetail(ctx,id)
	return
}
