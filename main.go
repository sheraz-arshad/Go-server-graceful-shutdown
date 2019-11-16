package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"demo/routes"

	"github.com/gorilla/mux"
)

func main() {

	router := setupRoutes()
	done := setupChannel()
	server := startServer(router)

	<-done
	shutdownServer(server)
}

func setupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", routes.Home)
	router.HandleFunc("/about", routes.About)
	return router
}

func setupChannel() chan os.Signal {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP)
	return done
}

func startServer(r *mux.Router) *http.Server {
	server := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server close error %s\n", err)
		}
	}()

	log.Println("Server has started on the port 3000")

	return server
}

func shutdownServer(server *http.Server) {
	log.Println("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	log.Println("Server Exited Properly")
}
