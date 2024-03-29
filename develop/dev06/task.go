package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fields := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	if *fields == "" {
		fmt.Println("you must specify a list of fields")
		os.Exit(-1)
	}

	fieldsInt := make([]int, 0)
	data := make([]string, 0)

	for _, field := range strings.Split(*fields, ",") {
		n, err := strconv.Atoi(field)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		fieldsInt = append(fieldsInt, n)
	}

	switch {
	case len(flag.Args()) > 0:
		for _, filePath := range flag.Args() {
			file, err := os.OpenFile(filePath, os.O_RDONLY, 0)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			defer file.Close()

			data = cut(data, file, fieldsInt, *delimiter, *separated)
		}

	default:
		data = cut(data, os.Stdin, fieldsInt, *delimiter, *separated)
	}

	fmt.Println(strings.Join(data, "\n"))
}

func cut(sliceRes []string, src io.Reader, fieldsInt []int, delimiter string, separated bool) []string {
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if separated && !strings.Contains(line, delimiter) {
			continue
		}

		separated := strings.Split(line, delimiter)
		res := make([]string, 0)
		for _, field := range fieldsInt {

			if field > len(separated) {
				break
			}
			res = append(res, separated[field-1])
		}

		sliceRes = append(sliceRes, strings.Join(res, "\t"))
	}
	return sliceRes
}
