package main

import (
"context"
"log"
"os"
"os/signal"

"github.com/matthausen/thirdfort/service"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("System call:%+v", oscall)
		cancel()
	}()

	if err := service.GracefullyShutDown(ctx); err != nil {
		log.Printf("Failed to serve:+%v\n", err)
	}
}