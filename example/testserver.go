package example

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Testserver struct {
}

func (t Testserver) GetTimestamp(ctx context.Context, empty *Empty) (*Timestamp, error) {
	timestamp := Timestamp{
		Time: timestamppb.Now(),
	}

	return &timestamp, nil
}

func (t Testserver) Hello(ctx context.Context, test *Test) (*TestResponse, error) {
	fmt.Println("server: ", test)
	response := TestResponse{
		Msg: "alles angekommen",
	}
	return &response, nil
}

func (t Testserver) mustEmbedUnimplementedEchoServer() {
	//TODO implement me
	panic("implement me")
}
