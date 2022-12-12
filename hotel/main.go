package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	//defer cancel()

	// go func() {
	// 	time.Sleep(time.Second * 4)
	// 	cancel()
	// }()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Tempo excedido para reservar o quarto")
	case <-time.After(time.Second * 5):
		fmt.Println("Quarto reservado com sucesso")
	}
}
