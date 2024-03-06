Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1
```

Это происходит так, потому что `defer()` может читать и присваивать значения именованым возвращаемым значениям:

> 3) Deferred functions may read and assign
> to the returning function’s named return values.

В случае неименованных параметров, такого не происходит.

defer'ы записываются в стэк, и вызываются в порядке, обратном порядку их вызова (в случае вызова нескольких `defer()`)

Пример:
```go
 func c() (i int) {
    defer func() { fmt.Println("third") }()
    defer func() { fmt.Println("second") }()
    defer func() { fmt.Println("first") }()

    return 1
}
```
Выполнится в порядке:
```
i = 1 (return 1)
"first"
"second"
"third"
```