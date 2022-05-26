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
Горутины поочередно печатают числа благодаря тому что после каждой печати срабатывает задержка.
В дальнйшем бесконечно печатаются нули, т.к select в таком виде не прекращает работу в случае закрытия канала, 
а продолжает записывать значения по умолчанию.


Вариант решения
for {
    select {
    case x, ok := <-ch:
        fmt.Println("ch1", x, ok)
        if !ok {
            ch = nil
        }
    case x, ok := <-ch2:
        fmt.Println("ch2", x, ok)
        if !ok {
            ch2 = nil
        }
    }

    if ch == nil && ch2 == nil {
        break
    }
}

```
