package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"golang.org/x/net/context"

	"github.com/Tomoka64/janken_server/janken"
	"google.golang.org/grpc"
)

type taskServer struct{}

type Data struct {
	Win  int `json:"win"`
	Lose int `json:"lose"`
	Aiko int `json:"aiko"`
}

func main() {
	fmt.Println("listening to port 8888")
	srv := grpc.NewServer()
	var tasks taskServer
	jankenserve.RegisterTasksServer(srv, tasks)
	l, _ := net.Listen("tcp", ":8888")
	log.Fatal(srv.Serve(l))
}

func (taskServer) Jankening(ctx context.Context, janken *jankenserve.Janken) (*jankenserve.Result, error) {
	robot := robotPlay()
	result, err := RuleHandler(robot, janken.Janken)
	if err != nil {
		return nil, err
	}
	koma := Koma[int(robot)]
	ret := &jankenserve.Result{
		Result: result,
		Koma:   koma,
	}
	return ret, nil

}

var Koma = map[int]string{
	0: "Rock",
	1: "Paper",
	2: "Scissors",
}

func robotPlay() int64 {
	rand.Seed(time.Now().UTC().UnixNano())
	play := rand.Intn(len(Koma))
	return int64(play)
}

func RuleHandler(robot int64, player int64) (string, error) {
	switch {
	case robot == player:
		return "Aiko", nil
	case robot == 0:
		if player == 1 {
			return "you win", nil
		}
		if player == 2 {
			return "you lost", nil
		}
	case robot == 1:
		if player == 0 {
			return "you lost", nil
		}
		if player == 2 {
			return "you win", nil
		}
	case robot == 2:
		if player == 0 {
			return "you win", nil
		}
		if player == 1 {
			return "you lost", nil
		}
	}
	return "", fmt.Errorf("Not determinable")
}
