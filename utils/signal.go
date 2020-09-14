package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
