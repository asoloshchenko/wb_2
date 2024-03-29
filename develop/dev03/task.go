package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	columnSort := flag.Int("k", 0, "указание колонки для сортировки")
	num := flag.Bool("n", false, "сортировать по числовому значению")
	isReverse := flag.Bool("r", false, "reverse")
	isUnique := flag.Bool("u", false, "не выводить повторяющиеся строки")
	// flag.Int("M", 0, "M")
	// flag.Int("b", 0, "b")
	// flag.Int("c", 0, "c")
	// flag.Int("h", 0, "h")

	flag.Parse()

	lines := make([][]string, 0)
	if len(flag.Args()) == 0 {
		fmt.Println("No arguments")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lines = append(lines, []string{scanner.Text()})
		}

	}

	for _, filePath := range flag.Args() {
		file, err := os.OpenFile(filePath, os.O_RDONLY, 0)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		reader := bufio.NewScanner(file)
		reader.Split(bufio.ScanLines)

		for reader.Scan() {
			lines = append(lines, []string{reader.Text()})
		}

	}

	if *columnSort > 0 {
		for i := 0; i < len(lines); i++ {
			lines[i] = strings.Split(lines[i][0], " ") // []string{[i])
		}
	} else {
		*columnSort = 1
	}

	if *num {
		sort.Slice(lines, func(i, j int) bool {
			num1, _ := strconv.Atoi(lines[i][*columnSort-1])
			num2, _ := strconv.Atoi(lines[j][*columnSort-1])
			return num1 < num2
		})
	} else {
		sort.Slice(lines, func(i, j int) bool {
			return lines[i][*columnSort-1] < lines[j][*columnSort-1]
		})
	}

	res := make([]string, len(lines))
	for i := range lines {
		res[i] = strings.Join(lines[i], " ")
	}

	if *isUnique {
		seen := make(map[string]struct{})
		j := 0
		for _, v := range res {
			if _, ok := seen[v]; ok {
				continue
			}
			seen[v] = struct{}{}
			res[j] = v
			j++
		}
		res = res[:j]
	}

	if *isReverse {
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}
	fmt.Println()
	for _, l := range res {
		fmt.Println(l)
	}

	//fmt.Println(lines)

}
