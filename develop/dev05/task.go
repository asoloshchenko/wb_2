package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "количество строк")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNum := flag.Bool("n", false, "печатать номер строки")

	flag.Parse()
	//fmt.Println(*after, *before, *context, *count, *ignoreCase, *invert, *fixed, *lineNum)

	data := make([]string, 0)

	switch {
	case len(flag.Args()) < 1:
		fmt.Println("No arguments")
		os.Exit(-1)

	case len(flag.Args()) > 1:
		for _, filePath := range flag.Args()[1:] {
			file, err := os.OpenFile(filePath, os.O_RDONLY, 0)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				data = append(data, scanner.Text())
			}
		}

	default:
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			data = append(data, scanner.Text())
		}
	}

	fmt.Println()

	pattern := flag.Args()[0]
	if *fixed {
		pattern = "\\Q" + pattern + "\\E"
	}

	if *ignoreCase {
		pattern = "(?i)" + pattern
	}

	rexp, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if *context > 0 {
		*after = *context
		*before = *context
	}

	matched := make(map[int]struct{})

	for i, line := range data {
		if (*invert && !rexp.MatchString(line)) || (!*invert && rexp.MatchString(line)) {
			addLine(i, *before, *after, matched, len(data))
		}
	}

	if *count {
		fmt.Printf("Count: %d\n", len(matched))
		return
	}

	for k := range matched {

		if *lineNum {
			fmt.Printf("%d:\t%s\n", k+1, data[k])
		} else {
			fmt.Printf("%s\n", data[k])
		}
	}

}

func addLine(ind int, before int, after int, m map[int]struct{}, lenData int) {

	for j := ind - before; j <= ind+after; j++ {
		if j < 0 || j >= lenData {
			continue
		}
		if _, ok := m[j]; ok {
			continue
		}

		m[j] = struct{}{}
	}

}
