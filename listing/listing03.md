Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)


func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
this comparison is never true (SA4023)
	main.go:14:9: the lhs of the comparison is the 1st return value of this function call
	main.go:8:6: _/root/wbschool_exam_L2/listing.Foo never returns a nil interface value

error - это не базовый тип, а интерфейс 
type error interface {
    Error() string
}
значит, мы сравниваем интерфейс, содержащий строку Error со значением nil

сам интерфейс - это структура iface
type iface struct {
	tab  *itab // ссылка на хранимый тип
	data unsafe.Pointer // ссылка на данные
}

интерфейс == nil, когда отсутствует и значение и тип (interface{})
```
