package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/OtusTeam/Go-Basic/open_lessons/go_http/service"
)

func main() {
	ctx := context.Background()

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill, syscall.SIGTERM)
	defer cancel()

	listenCfg := net.ListenConfig{}

	ln, err := listenCfg.Listen(ctx, "tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Handler: service.NewService(),
	}

	go func() {
		<-ctx.Done()

		log.Printf("shutting down due to context signal")

		if err := server.Shutdown(ctx); err != nil {
			log.Print(err)
		}
	}()

	log.Print(server.Serve(ln))
}
