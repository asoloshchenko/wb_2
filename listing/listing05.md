Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```
Аналогично листингу 2
Переменная `err` имеет тип `*customError`, хоть она и ничем и не инициализированна, она не равна `nil`

Происходит это из-за того, что функция `test()` возвращает неинициализированный `*customError` вместо интерфейса