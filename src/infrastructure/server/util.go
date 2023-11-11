package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/cmd/config"
)

// TODO: サーバ起動とgraceful shutdownを行う

func runWithGracefulShutdown(handler *http.ServeMux) {

	srv := &http.Server{
		Addr:    config.Config.Server.Port,
		Handler: handler,
	}

	go func() {
		log.Printf("Server listening on port %s\n", config.Config.Server.Port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server listen error: %s\n", err)
		}
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q, os.Interrupt)

	<-q

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %s\n", err)
	}

	log.Println("Server shutdown")
}
