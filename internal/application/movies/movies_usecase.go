package movies

type IMoviesUsecase interface{}

type MoviesUsecase struct{
	repository IMoviesRepository
}

func NewMoviesUsecase(repo IMoviesRepository)IMoviesUsecase{
	return &MoviesUsecase{repository: repo}
}