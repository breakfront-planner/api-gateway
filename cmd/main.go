package main

import (
	"context"
	"os/signal"
	"syscall"
	//"github.com/breakfront-planner/api-gateway/internal/app"
)

func main() {
	_, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	/*

		application, err := app.New(ctx)
		if err != nil {
			log.Println("error while inject dependencies: ", err)
			os.Exit(1)
		}

		if err := application.Start(ctx); err != nil {
			log.Println("running the program: ", err)
			os.Exit(1)
		}

		closeCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := application.Close(closeCtx); err != nil {
			log.Println("closing the program: ", err)
		}
		log.Println("application stopped")
	*/
}
