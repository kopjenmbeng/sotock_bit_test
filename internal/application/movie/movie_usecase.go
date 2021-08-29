package movie

type IMovieUsecase interface {
}

type MovieUsecase struct {
	repository IMovieRepository
}

func NewMovieUsecase(repo IMovieRepository) IMovieUsecase {
	return &MovieUsecase{}
}
