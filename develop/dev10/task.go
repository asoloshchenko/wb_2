package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {

	timeoutStr := flag.String("timeout", "10s", "timeout")
	flag.Parse()
	fmt.Println(flag.Args())
	var host, port string
	if len(flag.Args()) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(-1)
	}
	host = flag.Arg(0)
	port = flag.Arg(1)

	fmt.Println(host, port, *timeoutStr)

	// Парсим таймаут
	timeout, err := time.ParseDuration(*timeoutStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	addr := host + ":" + port
	maxRetry := int(timeout / time.Second)

	conn, err := ConnectWithRetry(addr, timeout, maxRetry)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer conn.Close()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Connected to the server")
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			conn.Write([]byte(scanner.Text()))
		}
		if scanner.Err() != nil {
			fmt.Println(scanner.Err())
			os.Exit(-1)
		}

		done <- syscall.SIGINT

	}(ctx)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err == io.EOF {
					done <- syscall.SIGINT
					return
				}
				if err != nil {
					//fmt.Println("ERR read: ", err)
					done <- syscall.SIGINT
					return
				}
				fmt.Println(string(buf[:n]))
			}

		}
	}(ctx)

	<-done
	cancel()

}

func ConnectWithRetry(addr string, timeout time.Duration, maxRetry int) (net.Conn, error) {
	for i := 0; i < maxRetry; i++ {
		conn, err := net.DialTimeout("tcp", addr, timeout)
		if err == nil {
			return conn, nil
		}
		time.Sleep(time.Second)
	}

	return nil, errors.New("failed to connect to the server")
}
