background-image: url(../img/title.svg)

---

background-image: url(../img/rec.svg)

---
background-image: url(../img/topic.svg)

.topic[Рекурсия и динамическое программирование]
.tutor[Родион Ступников]
.tutor_desc[Software Engineer]

---

background-image: url(../img/rules.svg)

---
# Цели занятия

* Понять принцип работы рекурсии
* Разобраться в плюсах и минусах рекурсивных алгоритмов
* Разобраться в плюсах и минусах динамического программирования

---

# Рекурсия

.center-image[
![](img/gopher_recursion.png)
]

---

# Рекурсия

* В программировании рекурсия — вызов функции (процедуры) из неё же самой, 
непосредственно (простая рекурсия) или через другие функции (сложная или косвенная рекурсия).
* Например, функция A вызывает функцию B, а функция B — функцию A.
* Рекурсивная программа позволяет описать повторяющееся или даже потенциально бесконечное вычисление,
причём без явных повторений частей программы и использования циклов.

---

# Рекурсия

Потребление памяти при использовании рекурсии

.center-image[
![](img/recursion_mem.png)
]

---

# Рекурсия: числа Фибоначчи

```go
// Функция для вычисления числа Фибоначчи с использованием рекурсии
func fibonacci(n int) int {
	if n <= 1 { // терминальное условие
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
```

```go
result := fibonacci(6)
```

Что со сложностью в O-нотации?

---
# Рекурсия: числа Фибоначчи

Полная версия кода:
https://go.dev/play/p/0L3_5nfKa0c

---

# Рекурсия: числа Фибоначчи

На изображении показано дерево рекурсивных вызовов для вычисления числа Фибоначчи F(4).

.center-image[
![](img/recursion_fib.png)
]

---

# Рекурсия: числа Фибоначчи

* N - номер числа Фибоначчи
* Сложность рекурсивного вычисления: O(2^N)
* Это происходит из-за того, что функция вызывает себя дважды для каждого числа Фибоначчи
и создает рекурсивное дерево с экспоненциальной глубиной.
* А вот память: O(N)

---
# Рекурсия

* Часто позволяет сделать код короче и выразительнее.
* Требует дополнительных расходов памяти под каждый рекурсивный вызов функции.

---

# Хвостовая рекурсия

Обычный код без оптимизации
```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}
```

```go
// 5! = 5 * 4 * 3 * 2 * 1
result := factorial(5)
```

https://go.dev/play/p/S5bvbTquiY6

---
# Хвостовая рекурсия

Оптимизированный код

```go

func factorialTailRec(n, accumulator int) int {
	if n <= 1 {
		return accumulator
	}
	return factorialTailRec(n-1, n*accumulator)
}

func factorial(n int) int {
	return factorialTailRec(n, 1)
}
```
https://go.dev/play/p/ws0kBttPZnE

---
# Хвостовая рекурсия

Под капотом компилятор преобразует ее примерное в такую функцию:

```go
func factorial(n, accumulator int) int {
    for {
        if n <= 1 {
            return accumulator
        }
        accumulator *= n
        n--
    }
}

```
https://go.dev/play/p/WWOAYKsUCIf

---

# Динамическое программирование

.center-image[
![](img/dynamic_programming.png)
]

---
# Итоги

---

background-image: url(../img/questions.svg)

---

background-image: url(../img/poll.svg)

---

background-image: url(../img/next_webinar.svg)
.announce_date[1 января]
.announce_topic[Тема следующего вебинара]

---
background-image: url(../img/thanks.svg)

.tutor[Лектор]
.tutor_desc[Должность]
