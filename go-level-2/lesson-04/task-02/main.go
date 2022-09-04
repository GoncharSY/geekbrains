package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

type contextKey string

const key contextKey = "goroutine_number"
const count int = 3

func main() {
	var ctx context.Context
	var stop context.CancelFunc
	var pool = make(chan any, count)

	// Создадим основной контекст.
	ctx, stop = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	// Запустим обработчики.
	for i := 1; i <= count; i++ {
		ctx := context.WithValue(ctx, key, i)
		go startWork(ctx, pool)
	}

	// Ждем сигнал к завершению.
	waitSignal(ctx)

	// Ждем завершения горутин.
	stopWorks(stop, pool)

	// Завершаем основную программу.
	fmt.Println("Main: stopped")
}

// Начать обработку.
func startWork(ctx context.Context, stop chan<- any) {
	var num = ctx.Value(key).(int)

	defer func() {
		stop <- struct{}{}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Goroutine-%v: stopped\n", num)
			return
		default:
			fmt.Printf("Goroutine-%v: working\n", num)
			time.Sleep(time.Duration(num) * time.Second)
		}
	}
}

// Ожидать сигнала.
func waitSignal(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("Signal: received")
}

// Прекратить обработку.
func stopWorks(stop context.CancelFunc, pool <-chan any) {
	stop()

	for i := 0; i < count; i++ {
		<-pool
	}
}
