package test

import (
	"context"
	"testing"

	"github.com/kopjenmbeng/sotock_bit_test/internal/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestSearchMovie(t *testing.T) {
	var conn *grpc.ClientConn
	// conn, err := grpc.Dial("13.228.12.30:30304", grpc.WithInsecure()) // for staging server testing
	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	assert.Equal(t, nil, err)
	if err != nil {
		t.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewMoviesClient(conn)
	response, err := c.Search(context.Background(), &proto.SearchRequestMessage{Search: "batman", Page: 1})
	// if err!=nil{
	// 	t.Fatalf("error %s",err)
	// }
	assert.Equal(t, nil, err)
	t.Log(response)
}

func TestDetailMovie(t *testing.T) {
	var conn *grpc.ClientConn
	// conn, err := grpc.Dial("13.228.12.30:30304", grpc.WithInsecure()) // for staging server testing
	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	assert.Equal(t, nil, err)
	if err != nil {
		t.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewMoviesClient(conn)
	response, err := c.GetDetail(context.Background(), &proto.DetailMovieRequestMessage{Id: "tt0096895"})
	// if err!=nil{
	// 	t.Fatalf("error %s",err)
	// }
	assert.Equal(t, nil, err)
	t.Log(response)
}
