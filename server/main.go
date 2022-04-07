package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	"modbus_server/config"
	"modbus_server/counter"
	"modbus_server/server"
)

func init() {
	config.Setup("config.ini")
}

func main() {
	var wg sync.WaitGroup
	stopchan := make(chan bool, 1)
	kill := make(chan os.Signal, 1)
	signal.Notify(kill,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		wg.Wait()

		close(stopchan)
	}()

	wg.Add(2)

	//1
	go func() {
		defer wg.Done()

		server.ModbusServer()
		runtime.Gosched()
	}()

	//2
	go func() {
		defer wg.Done()

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		counter.Counter(*ticker, stopchan)
		runtime.Gosched()
	}()

	exit_chan := make(chan int)
	go func() {
		for {
			s := <-kill
			switch s {
			// kill -SIGHUP XXXX
			case syscall.SIGHUP:
				log.Println("close_cause [hungup]")
				exit_chan <- 3

			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				log.Println("close_cause [interupt]")
				exit_chan <- 2

			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				log.Println("close_cause [force_stop]")
				exit_chan <- 0

			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				log.Println("close_cause [stop and core dump]")
				exit_chan <- 0

			default:
				log.Println("close_cause [Unknown signal]")
				exit_chan <- 1
			}
		}
	}()

	code := <-exit_chan
	log.Println("< close main >")

	os.Exit(code)
}
