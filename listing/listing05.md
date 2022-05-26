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
Поскольку тип customError реализует метод Error(), то он удовлетворяет интерфейсу error. 
Функция test возвращает указатель типа customError со значением nil. Таким образом, поле 
структуры tab структуры iface не будет пустым и переменная err снова не равна nil. 

Вывод:
error
```
