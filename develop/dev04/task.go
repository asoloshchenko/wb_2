package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func CheckIfStringsAreAnagram(source string, target string) bool {

	// Basic use case check both strings length
	if len(source) != len(target) {
		return false
	}

	// sort source & target arrays
	sourceArray := []rune(source)
	sort.Slice(sourceArray, func(i, j int) bool {
		return sourceArray[i] < sourceArray[j]
	})
	targetArray := []rune(target)
	sort.Slice(targetArray, func(i, j int) bool {
		return targetArray[i] < targetArray[j]
	})

	// Loop through the arrays and check character by character
	for i := 0; i < len(sourceArray); i++ {
		if sourceArray[i] != targetArray[i] {
			return false
		}
	}
	return true
}

func f(words []string) map[string][]string {
	res := make(map[string][]string)
Outer:
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
		for k := range res {
			if CheckIfStringsAreAnagram(words[i], k) && words[i] != k {
				res[k] = append(res[k], words[i])
				continue Outer
			}
		}
		res[words[i]] = make([]string, 0)
	}
	return res
}

func main() {

	words := []string{"Пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	res := f(words)

	fmt.Println(res)
}
