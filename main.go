package main

import (
	"github.com/goravel/framework/contracts/queue"
	"os"
	"os/signal"
	"syscall"

	"github.com/goravel/framework/facades"

	"goravel-blog/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start http server by facades.Route().
	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route Run error: %v", err)
		}
	}()

	// Listen for the OS signal
	go func() {
		<-quit
		if err := facades.Route().Shutdown(); err != nil {
			facades.Log().Errorf("Route Shutdown error: %v", err)
		}

		os.Exit(0)
	}()

	go func() {
		if err := facades.Queue().Worker().Run(); err != nil {
			facades.Log().Errorf("Queue run error: %v", err)
		}
	}()
	go func() {
		if err := facades.Queue().Worker(queue.Args{
			Connection: "database",
			Queue:      "default",
			Concurrent: 1,
			Tries:      1,
		}).Run(); err != nil {
			facades.Log().Errorf("Queue run error: %v", err)
		}
	}()

	go facades.Schedule().Run()

	select {}
}
