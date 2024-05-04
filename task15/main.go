package main

import (
	"bytes"
	"fmt"
	"strings"
)

/*
К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить?
Приведите корректный пример реализации.

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}

Ответ:

Одной из главных проблем является утечка памяти,
т.к мы ссылаем на часть слайса

Важно помнить, что строки в Golang неизменяемы =>
нам нужен интсрумент для эффективного изменения,
хорошим вариантом является использование буфера
*/

var justString string

func createHugeString(stringLen int) string {
	var buf bytes.Buffer
	for i := 0; i < stringLen; i++ {
		buf.WriteString("s")
	}
	return buf.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
	fmt.Println(len(justString), justString)
}

func main() {
	someFunc()
}
