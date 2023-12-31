package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

	chanForResp := make(chan resp)
	go RPCCall(ctx, chanForResp)

	res := <-chanForResp
	fmt.Println(res.id, res.err)
}

type resp struct {
	id  int
	err error
}

func RPCCall(ctx context.Context, ch chan<- resp) {
	select {
	case <-ctx.Done():
		ch <- resp{
			id:  0,
			err: errors.New("timeout expired"),
		}
	case <-time.After(time.Second * time.Duration(rand.Intn(5))):
		ch <- resp{
			id: rand.Int(),
		}
	}
}
