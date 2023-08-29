background-image: url(../img/title.svg)

---

background-image: url(../img/rec.svg)

---
background-image: url(../img/topic.svg)

.topic[Управляющие конструкции языка и отладка программ]
.tutor[Хохлов Александр]
.tutor_desc[Архитектор платформенных решений ГК "Иннотех"]

---

background-image: url(../img/rules.svg)

---

# Цели занятия

- ### Научиться применять управляющие конструкции
- ### Уметь использовать отладчик

---


# Операторы ветвления if и else

Основной набор операторов для управления выбором выполнения программы


```
func isFileExist(_ string) bool {
	return false
}

func Example1() {
	exist := isFileExist("/etc/password")

	if exist {
		fmt.Printf("Exist\n")
	} else {
		fmt.Printf("Not exist\n")
	}
}
```


---

# Операторы ветвления else if


Проверяем дополнительные условия

```
func findUserBySNILS(_ string) (int, string, string) {
	return 30, "f", "Moscow"
}

func Example2() {
	age, sex, city := findUserBySNILS("000-000-000 00")

	if age <= 14 {
		fmt.Printf("Ребенок\n")
	} else if sex != "m" && city == "Moscow" {
		fmt.Printf("Москвичка\n")
	} else if sex == "m" {
		fmt.Printf("Мужчина\n")
	} else {
		fmt.Printf("Человек %\n", age)
	}
}
```

---
# Инициализация переменных в if

Можно объявить и инициализирорвать переменные в операторе if. В этом случае они будут доступны в основной и алльтернативных ветках выполнения (else/else if). Но за пределами оператора if такие переменные не будут доступны.

```
	if exist := isFileExist("/etc/password"); exist {
		fmt.Printf("Exist\n")
	} else {
		fmt.Printf("Not exist\n")
	}
```

---

# Операторы сравнения

- == - равны ли два значения
- != - не равны ли два значения
- < - является ли одно значение меньше другого
- \> - является ли одно значение больше другого
- <= - является ли одно значение меньше или равно другому значению
- \>= - является ли одно значение больше или равно другому значению

---

# Логические операторы 

- **&&** (x && y) является оператором **и**. Это верно, если оба выражения верны.
- **||** (x || y) является оператором **или**. Это верно, если хотя бы одно выражение является верным.
- **!** (! x) является оператором **нет**. Это верно, только если выражение является ложным.

---

# Оператор ветвления switch

Оператор switch позволяет упростить код в случае когда сравниваем множество значений с одной из переменных

```
	age, _, _ := exampleIf.FindUserBySNILS("000-000-000 00")

	switch age {
	case 30, 25:
		fmt.Println(30)
		fmt.Println(25)
	case 14:
		fmt.Println(14)
	case 5:
		fmt.Println(5)
	default:
		fmt.Printf("Age = %d\n", age)
	}
```

---

# Оператор ветвления switch

Оператор switch позволяет в проверках использовать другие переменные и даже выражения. 

```
func af() int {
	return 25
}

func Example3() {
	age, _, _ := exampleIf.FindUserBySNILS("000-000-000 00")

	a := af()

	switch age {
	case 30, 25:
		fmt.Println(30)
		fmt.Println(25)
	case 14:
		fmt.Println(14)
	case a * 2:
		fmt.Printf("Var Age = %d\n", age)
	default:
		fmt.Printf("Age = %d\n", age)
	}
}

```

---

# Оператор ветвления switch

При одинаковых значениях с использованием переменной сработает то что будет сравнено раньше

В случае константных/литеральных значений приложение не соберется при дубликатах. 

```
	const a = 25

	switch age {
	case 30, 25:
		fmt.Println(30)
		fmt.Println(25)
	case 14:
		fmt.Println(14)
	case a:
		fmt.Printf("Var Age = %d\n", age)
	default:
		fmt.Printf("Age = %d\n", age)
	}

```

---

# Оператор ветвления switch и fallthrough

Если необходимо продолжить выполнение в следующем case то следует вызывать fallthrough

```
	switch age {
	case 30, 25:
		fmt.Println(30)
		fmt.Println(25)
	case 14:
		fmt.Println(14)
		fallthrough
	case 5:
		fmt.Println(5)
	default:
		fmt.Printf("Age = %d\n", age)
	}
```

---
# Оператор ветвления switch

Исходную переменную можно не указывать, тогда в каждом case будет проверяться отдельные условия:

```
	t := time.Now()
	switch {
	case t.Second() < 15:
		fmt.Printf("<15 %v\n", t)
	case t.Second() < 30:
		fmt.Printf("<30 %v\n", t)
	case t.Second() < 45:
		fmt.Printf("<45 %v\n", t)
	default:
		fmt.Printf(">45 %v\n", t)
	}
```

---
# Оператор ветвления switch

В качестве сравнимого значения могут выступать и типы

```
	var s interface{}

	s = 12345.6789

	switch stype := s.(type) {
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	case float64:
		fmt.Println("float64")
	case float32:
		fmt.Println("float32")
	case int:
		fmt.Println("int")
	default:
		fmt.Printf("%T", stype)
	}
```

---
# Инициализация переменных в switch

Также как и с if в switch можно объявить и инициализирорвать переменные. В этом случае они будут доступны в операторе switch.

```
	switch age, _, _ := exampleIf.FindUserBySNILS("000-000-000 00"); age {
	case 30, 25:
		fmt.Println(30)
		fmt.Println(25)
	case 14:
		fmt.Println(14)
	case 5:
		fmt.Println(5)
	default:
		fmt.Printf("Age = %d\n", age)
	}
```


---
# Оператор цикла for

Циклы позволяют описать повторяющиеся команды. Можно применить несколько вариантов циклов. 

Первый классический вариант с инициализацией, проверкой условия выполнения цикла и действием перед следующей проверкой.

```
	for i := 1; i <= 3; i++ {
		fmt.Printf("i = %d\n", i)
	}
```

---
# Оператор цикла for

Второй вариант - оставляем только проверку условия

```
	i := 1
	for i <= 3 {
		fmt.Printf("i = %d\n", i)
		i++
	}
```

---
# Оператор цикла for

Третий вариант - бесконечный цикл с управление выхода из цикла в самом цикле

```
	i := 0
	for {
		i++

		if i%2 == 0 {
			continue
		} else if i > 3 {
			break
		}

		fmt.Printf("i = %d\n", i)
	}
```

- **break** - команда выхода из цикла
- **continue** - команда перехода к новой итерации цикла

---
# Управление с помощью break и continue


```
	for i := 1; i <= 3; i++ {
		fmt.Printf("i = %d\n", i)

		switch i {
		case 2:
			break
		default:
			fmt.Println("Ok")
		}
		fmt.Println(".")
	}
```

- **break** - команда которая может применяться и в for и в switch. Внимательно ее использовать
- **continue** - команда применяется только в циклах

---
# Оператор цикла for

Четвертый вариант - цикл range. Применяем для перебора значений в:
- массивы или слайсы
- строки
- словари
- каналы


```
	letters := []string{"a", "b", "c"}

	for i, letter := range letters {
		fmt.Printf("Index: %d Value:%s\n", i, letter)
	}
```

---
# Оператор цикла for


Перебор только значений
```
	for _, letter := range letters {
		fmt.Printf("Value: %s\n", letter)
	}
```

Перебор индексов
```
	for i := range letters {
		fmt.Printf("Index: %d\n", i)
	}
```

Перебор состава массива
```
	i := 0
	for range letters {
		fmt.Printf("Index: %d Value: %s\n", i, letters[i])
		i++
	}
```

---
# Итерация по строке

По байтам
```
for i := 0; i < len(s); i++ {
    b := s[i]
    // i строго последоваельно
    // b имеет тип byte, uint8
}
```

По рунам
```
for i, r := range s {
    // i может перепрыгивать значения 1,2,4,6,9...
    // r - имеет тип rune, int32
}
```

---
# Итерация по каналу


```
func Example12() {
	ch := make(chan string)

	go pushToChannel(ch)

	for val := range ch {
		fmt.Println(val)
	}
}
func pushToChannel(ch chan<- string) {
	ch <- "a"
	ch <- "b"
	ch <- "c"

	close(ch)
}

```

---
# Метки и работа с ними

Метки присутствуют в языке но их использование не рекомендуется тк это существенно снижает читаемость и сопровождаемость программ.

```
	fmt.Println("a")
	goto FINISH
	fmt.Println("b")

FINISH:
	fmt.Println("c")
```

---
# Метки и работа с ними

Метки должны быть доступны в той же функции откуда возможен вызов goto. Если метка не используется, то ее наличие считается ошибкой также как и не используемые переменные.

```
func Example14() {
	fmt.Println("a")
	goto FINISH
	fmt.Println("b")
}

func test() {
FINISH:
	fmt.Println("c")
}

```
---
# Метки когда можно использовать 1/2

В отдельных ситуациях переход по goto может все-таки быть более предпочтителен. Например стандратная библиотека и функция Gamma
https://go.dev/src/math/gamma.go


```
z := 1.0
	for x >= 3 {
		x = x - 1
		z = z * x
	}
	for x < 0 {
		if x > -1e-09 {
			goto small
		}
		z = z / x
		x = x + 1
	}
	for x < 2 {
		if x < 1e-09 {
			goto small
		}
		z = z / x
		x = x + 1
	}
```

---
# Метки когда можно использовать 1/2

```
	if x == 2 {
		return z
	}

	x = x - 2
	p = (((((x*_gamP[0]+_gamP[1])*x+_gamP[2])*x+_gamP[3])*x+_gamP[4])*x+_gamP[5])*x + _gamP[6]
	q = ((((((x*_gamQ[0]+_gamQ[1])*x+_gamQ[2])*x+_gamQ[3])*x+_gamQ[4])*x+_gamQ[5])*x+_gamQ[6])*x + _gamQ[7]
	return z * p / q

small:
	if x == 0 {
		return Inf(1)
	}
	return z / ((1 + Euler*x) * x)
```

---
# Отладка приложений

- Установка точек останова (breakpoint)
- Пошаговое выполнение с входом в функциию (Step Into) и обходом функции (Step Over)
- Выход из функции (Step Out)
- Добавление переменных в наблюдение (Watch) и редактирование значений

---

background-image: url(../img/questions.svg)

---

background-image: url(../img/poll.svg)

---

background-image: url(../img/next_webinar.svg)
.announce_topic[Пользовательские типы данных]

---
background-image: url(../img/thanks.svg)

.tutor[Хохлов Александр]
.tutor_desc[Архитектор платформенных решений ГК "Иннотех"]
