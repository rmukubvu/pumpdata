package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/rmukubvu/pumpdata/controller"
	"github.com/rmukubvu/pumpdata/sensors"
	"github.com/rmukubvu/pumpdata/store"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var port = flag.Int("p", store.WebServerPort(), "port number")

func main() {
	flag.Parse()
	//r := handles.InitRouter()
	r := controller.InitRouter()
	//convert to port format
	sPort := fmt.Sprintf(":%d", *port)
	//create server
	srv := &http.Server{
		Addr: sPort,
		Handler: handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r),
	}
	//show on stdout
	log.Println("Server started. Press Ctrl-C to stop server")
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("Grpc Server started on port 7777")
	go func() {
		startGrpcServer()
	}()
	// Catch the Ctrl-C and SIGTERM from kill command
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//done
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	store.CloseDBConnection()
	log.Println("Server exiting")
	//log.Fatal(http.ListenAndServe(sPort, )
}

func startGrpcServer() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := sensors.Server{}
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	sensors.RegisterSensorServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
