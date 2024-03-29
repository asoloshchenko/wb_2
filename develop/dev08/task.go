package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			err := scanner.Err()
			if err != nil {
				fmt.Println("Error reading input:", err)
			}
			break
		}
		input := scanner.Text()

		pipes := strings.Split(input, "|")
		fmt.Println(pipes)

		commands := make([]*exec.Cmd, len(pipes))
		for i, pipe := range pipes {
			args := strings.Fields(pipe)
			cmd := exec.Command(args[0], args[1:]...)
			commands[i] = cmd
			if i == 0 {
				cmd.Stdin = os.Stdin
			} else {
				cmd.Stdin, _ = commands[i-1].StdoutPipe()
			}
			if i == len(pipes)-1 {
				cmd.Stdout = os.Stdout
			}
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				exitErr, ok := err.(*exec.ExitError)
				if ok {
					fmt.Printf("error running command %d: %s", i, string(exitErr.Stderr))
				} else {
					fmt.Printf("error running command %d: %s", i, err.Error())
				}
			}
		}

		// args := strings.Fields(input)
		// fmt.Println(args)
		// //command := make([]string, 0)
		// start := 0
		// var prevCommand *exec.Cmd

		// for i, arg := range args {
		// 	switch arg {
		// 	case "&":
		// 		cmd := exec.Command(args[start], args[start+1:i]...)
		// 		err := cmd.Start()
		// 		if err != nil {
		// 			fmt.Println(err)
		// 		}
		// 		defer cmd.Wait()
		// 		prevCommand = cmd

		// 		start = i + 1

		// 	case "|":
		// 		cmd := exec.Command(args[start], args[start+1:i]...)
		// 		if prevCommand == nil {
		// 			cmd.Stdin = os.Stdin
		// 		}else{
		// 			cmd.Stdin = prevCommand.StdoutPipe()
		// 		}
		// 		cmd.Stdout = os.Stdout
		// 		cmd.Stderr = os.Stderr
		// 		err := cmd.Run()
		// 		if err != nil {
		// 			fmt.Println(err)
		// 		}
		// 		start = i + 1
		// 		prevCommand = cmd

		// 	default:
		// 		continue
		// 	}
		// }
		// if start < len(args) {
		// 	cmd := exec.Command(args[start], args[start+1:]...)
		// 	cmd.Stdin = os.Stdin
		// 	cmd.Stdout = os.Stdout
		// 	cmd.Run()
		// }
	}
}

// func cd(path string) error {
// 	return os.Chdir(path)
// }
// func pwd() (string, error) {
// 	return os.Getwd()
// }
// func echo(args ...string) (string, error) {
// 	return strings.Join(args, " "), nil
// }
// func kill(pid int) error {
// 	proc, err := os.FindProcess(pid)
// 	if err != nil {
// 		return err
// 	}
// 	if err = proc.Kill(); err != nil {
// 		return err
// 	}

// 	return nil
// }
// func ps() error {
// 	cmd := exec.Command("ps")

// }
// func netcat() {
// }

func RunString(command string) (string, error) {

	// TODO
	return "", nil
}
