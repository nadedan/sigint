package sigint

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var deferredFuncs []func()

var deferred chan func()
var signals chan os.Signal

func init() {
	deferredFuncs = make([]func(), 0)
	deferred = make(chan func(), 5)

	signals = make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go listen()
}

// Notify adds to the list of signals that will be captured.
//
// Calling Notify with any of the defaulted signals will have no effect.
//
// Defaults: syscall.SIGINT, syscall.SIGTERM
func Notify(sig ...os.Signal) {
	signal.Notify(signals, sig...)
}

// Defer a function for execution when any of the notified signals are caught
func Defer(df func()) {
	deferred <- df
}

// execute the deferred functions
func execute() {
	for _, f := range deferredFuncs {
		f()
	}
}

// listen for the signals and for new deferreds to be added
func listen() {
	for {
		select {
		case df := <-deferred:
			deferredFuncs = append(deferredFuncs, df)

		case s := <-signals:
			fmt.Printf("\nStopping program due to %s\n", s.String())
			execute()

			os.Exit(1)
		}
	}
}
