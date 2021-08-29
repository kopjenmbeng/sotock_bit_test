package movies

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kopjenmbeng/goconf"
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

func (repo *MoviesRepository) Search(ctx context.Context, search string, page int) (result *dto.Searching, status int, err error) {
	uri := fmt.Sprintf("%s?apikey=%s&s=%s&page=%d", goconf.GetString("omdb.uri"), goconf.GetString("omdb.api_key"), search, page)
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(uri)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	status = http.StatusOK
	return
}
func (repo *MoviesRepository) GetDetail(ctx context.Context, id string) (result *dto.DetailMovie, status int, err error) {
	uri := fmt.Sprintf("%s?apikey=%s&s=%s&page=%d", goconf.GetString("omdb.uri"), goconf.GetString("omdb.api_key"), id)
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(uri)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	status = http.StatusOK
	return
}
