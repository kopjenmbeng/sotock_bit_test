package movies

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kopjenmbeng/goconf"
	"github.com/kopjenmbeng/sotock_bit_test/internal/dto"
)

type IMoviesRepository interface {
	Search(ctx context.Context, search string, page int) (result *dto.Searching, status int, err error)
	GetDetail(ctx context.Context, id string) (result *dto.DetailMovie, status int, err error)
}

type MoviesRepository struct {
	dbr sqlx.QueryerContext
	dbw *sqlx.DB
}

func NewMovieRepository(dbr sqlx.QueryerContext, dbw *sqlx.DB) IMoviesRepository {
	return &MoviesRepository{dbr: dbr,dbw: dbw}
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
	go repo.SaveLog(uuid.New().String(),"search",fmt.Sprintf("search=%s,page=%d",search,page),string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	status = http.StatusOK
	return
}

func(repo *MoviesRepository)SaveLog( id string,method string,request string,response string){
	query:=fmt.Sprintf(`
	INSERT INTO tbl_log
	(id,method,request,response,create_at)
	VALUES(?,?,?,?,?)
	`)
	now:=time.Now().UTC()
	_,err:=repo.dbw.Exec(query,&id,&method,&request,&response,&now)
	if err!=nil{
		fmt.Println(fmt.Errorf("error %s",err)) 
	}
}
func (repo *MoviesRepository) GetDetail(ctx context.Context, id string) (result *dto.DetailMovie, status int, err error) {
	uri := fmt.Sprintf("%s?apikey=%s&i=%s", goconf.GetString("omdb.uri"), goconf.GetString("omdb.api_key"), id)
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
	go repo.SaveLog(uuid.New().String(),"search",fmt.Sprintf("id=%s",id),string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	status = http.StatusOK
	return
}
