package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ID int
}

func (w *Worker) DoWork(counter *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	// Выполнение задачи работника
	time.Sleep(time.Second)
	fmt.Printf("Работник %d завершил задачу\n", w.ID)

	// Увеличение счетчика выполненных задач
	mutex.Lock()
	*counter++
	mutex.Unlock()
}

func main() {
	var counter int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	numWorkers := 5

	// Добавление работников в группу WaitGroup
	wg.Add(numWorkers)

	// Создание и запуск горутин работников
	for i := 1; i <= numWorkers; i++ {
		worker := &Worker{ID: i}
		go worker.DoWork(&counter, &wg, &mutex)
	}

	// Ожидание завершения всех работников
	wg.Wait()

	fmt.Println("Все работники завершили задачи")
	fmt.Printf("Общее количество выполненных задач: %d\n", counter)
}
