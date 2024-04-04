package main

import (
	"context"
	"log"

	"blogpost/internal/api/http"
	"blogpost/internal/metrics"
	"blogpost/internal/storage/inmem"
	"blogpost/internal/usecase/blog"
)

func main() {
	storage := inmem.New()

	usecase := blog.New(storage)

	ms := metrics.NewMetrics()

	server := http.NewServer(usecase, ms.HTTPAPI)

	ctx := context.Background()

	if err := server.Run(ctx); err != nil {
		log.Println("server error", err)
		return
	}
}
