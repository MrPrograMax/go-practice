package mymap

import (
	"fmt"
	"sync"
)

type MyMap struct {
	dict map[string]int
	mx   sync.Mutex

	//Используем мьютекс для безопасного обращения к общему объекту
}

// NewMyMap - конструктор
func NewMyMap() *MyMap {
	return &MyMap{dict: make(map[string]int)}
}

// SetValue - добавляет пару в словарь
func (m *MyMap) SetValue(key string, value int) {
	m.mx.Lock()         // Блокируем мьютекс
	defer m.mx.Unlock() // Разблокируем

	m.dict[key] = value // Создаем пару в словаре
}

// PrintMyMap - вывод всех пар из словаря
func (m *MyMap) PrintMyMap() {
	for key, value := range m.dict {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
}
