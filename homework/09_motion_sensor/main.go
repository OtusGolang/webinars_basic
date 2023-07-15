package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Начинаем чтение данных
	dataChan := readDataRoutine()

	// Запускаем горутину обработки данных
	procChan := processDataRoutine(10, dataChan)

	// Получаем обработанные значения из канала и выводим их на экран
	for processedData := range procChan {
		fmt.Println("Обработанные данные:", processedData)
	}
}

func readDataRoutine() chan int {
	// Создаем канал для передачи данных сенсора
	dataChan := make(chan int)

	go func() {
		// Закрываем канал для читателей
		defer close(dataChan)

		// Шлем данные в течение минуты
		timer := time.NewTimer(time.Minute)

		for {
			select {
			case <-timer.C:
				return
			default:
				// Генерируем случайное значение сенсора движения
				data := rand.Intn(100)

				// Отправляем значение в канал
				dataChan <- data
				// Пауза между чтениями данных
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	return dataChan
}

func processDataRoutine(slidingInterval int, dataChan <-chan int) chan int {
	// Создаем канал для передачи обработанных данных
	procChan := make(chan int)

	go func() {
		// Закрываем канал для читателей
		defer close(procChan)

		var sum int
		var count int

		for data := range dataChan {
			// Обработка данных: суммирование и подсчет количества значений
			sum += data
			count++

			if count == slidingInterval {
				// Вычисление среднего значения
				average := float64(sum) / float64(count)

				// Отправляем обработанное значение в канал
				procChan <- int(average)

				// Обнуляем вычитанные значния
				sum = 0
				count = 0
			}
		}
	}()

	return procChan
}
