package main

import (
 "context"
 "fmt"
 "log"
 "os"
 "os/signal"
 "syscall"
 "time"

 "http-go/routes"
 "http-go/server"
)

func main() {

 /* Register routes */
 router := routes.RegisterRoutes()

 s := server.NewServer(8080, router)

 // Set up graceful shutdown
 stop := make(chan os.Signal, 1)
 signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

 go func() {
  if err := s.Run(); err != nil {
   log.Fatalf("Error starting server: %v", err)
  }
 }()

 fmt.Println("Server is running...")

 //Interrupt signal
 <-stop

 fmt.Println("Shutting down server...")

 ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
 defer cancel()

 if err := s.Shutdown(ctx); err != nil {
  log.Fatalf("Server forced to shutdown: %v", err)
 }

 fmt.Println("Server gracefully stopped")
}