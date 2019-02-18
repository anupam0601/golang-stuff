package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/anupam0601/golang-stuff/ints-perf-automation/filescreatormicro"
	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	errChan := make(chan error)

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	var svc filescreatormicro.FilesCreatorService
	svc = filescreatormicro.BasicFilesCreator{}
	svc = filescreatormicro.LoggingMiddleware(logger)(svc)
	endpoint := filescreatormicro.Endpoints{
		FilesCreateEndpoint: filescreatormicro.MakeFilesCreateEndpoint(svc),
	}


	r := filescreatormicro.MakeHttpHandler(ctx, endpoint, logger)

	// HTTP transport
	go func() {
		fmt.Println("Starting server at port 8080")
		handler := r
		errChan <- http.ListenAndServe(":8080", handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println(<-errChan)

}
