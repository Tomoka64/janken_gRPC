package main

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/net/context"

	"github.com/Tomoka64/janken_server/janken"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "<usage> %s <num0~2>\n0: Rock\n1: Paper\n2: Scissors\n", os.Args[0])
		os.Exit(1)
	}
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect%v\n", err)
		os.Exit(1)
	}
	client := jankenserve.NewTasksClient(conn)
	num, _ := strconv.Atoi(os.Args[1])
	if num != 0 && num != 1 && num != 2 {
		fmt.Fprintln(os.Stderr, "put a number between 0 and 2\n0: Rock\n1: Paper\n2: Scissors")
		os.Exit(1)
	}
	err = Jankening(context.Background(), client, int64(num))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var Koma = map[int]string{
	0: "Rock",
	1: "Paper",
	2: "Scissors",
}

func Jankening(ctx context.Context, client jankenserve.TasksClient, in int64) error {
	result, err := client.Jankening(ctx, &jankenserve.Janken{Janken: in})
	if err != nil {
		return fmt.Errorf("could not add: %v", err)
	}
	fmt.Printf("robot: %s\nyou: %s\nresult: %s\n", result.Koma, Koma[int(in)], result.Result)
	return nil
}
