Что выведет программа? Объяснить вывод программы.

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

Ответ:
```
[77 78 79]

b - slice, ссылающийся на базовый массив а
    с первого индекса до четвертого.

```
