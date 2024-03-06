Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Сначала будут выведены числа от 1 до 8 включительно (порядок недетерминированный), затем бесконечно нули
```
При чтении из закрытого канала в переменную пишется значение по умолчанию (для `int` - `0`)
В функции `merge()` нет проверки на закрытые каналов, и она сама возвращаемый канал не закрывает.
Цикл в функции `main` так же будет длится бесконечно, так как канал `c` не будет закрыт.
