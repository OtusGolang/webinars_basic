background-image: url(../img/title.svg)

---

background-image: url(../img/rec.svg)

---
background-image: url(../img/topic.svg)

.topic[Инструментарий языка и организация проекта на Go]
.tutor[Алексей Семушкин]
.tutor_desc[Software Engineer at Semrush]

---

background-image: url(../img/rules.svg)

---

# О чем будем говорить

- Go SDK
- Компиляция и кросс-компиляция
- Пакеты и модули в Go
- Организация проекта
- Форматирование и линтинг кода
- Спецификация Go

---

# Компиляция

### Компиляция — сборка программы, включающая:

- трансляцию всех модулей программы, написанных на одном или нескольких исходных языках программирования в эквивалентные программные модули на низкоуровневом языке, близком машинному коду;
- последующую сборку (линковку) исполняемой машинной программы, в том числе вставку в программу кода всех функций, импортируемых из статических библиотек.


Интересный факт: компилятор Go написан на самом Go (c версии 1.5)

---

# Этапы компиляции в Go

- Лексический и синтаксический разбор
- Построение АСД
- Подготовка промежуточного представления
- Упрощение синтаксиса
- Генерация SSA (низкоуровневое представление)
- Генерация машинного кода
- Линковка

</br>
https://go.dev/src/cmd/compile/README

---

# Компиляция

```
package main

import "fmt"

func main() {
	fmt.Println("Hello!")
}
```

```
$ go build -o prog prog.go

$ file prog
prog: Mach-O 64-bit executable x86_64

$ ./prog
Hello!
```

```
$ go run prog.go
Hello!
```

---

# Кросс-компиляция

Go позволяет легко собирать программы для других архитектур и операционных систем.<br><br>
Для этого при сборке нужно переопределить переменные `GOARCH` и `GOOS`:

```
$ GOOS=windows go build -o /tmp/prog prog.go

$ file /tmp/prog
prog: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows

$ GOARCH=386 GOOS=darwin go build -o /tmp/prog prog.go

$ file /tmp/prog
prog: Mach-O i386 executable
```

Возможные значения `GOOS` и `GOARCH`
- `go tool dist list`
- [https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)

---

# One-shot запуск

Запустить файл "как скрипт".

```
go run main.go
```

Удобно для проверки кода и синтаксиса.<br><br>

Так же можно использовать Go PlayGround: https://go.dev/play/p/Fz3j-hbcocv

---

# Программа на Go

```
package main // Имя текущего пакета

// Импорты других пакетов
import "fmt"

// Неявная инициализация пакета
func init() {
    fmt.Println("Hello from init!")
}

// Функция main как точка входа
func main() {
	foo()
}

func foo() {
	fmt.Println("Foo!")
}
```

---

# Еще раз о GOPATH

`GOPATH` — переменная окружения, показывает где лежит ваше дерево исходников.

<br><br>
По умолчанию `$HOME/go`. Можно изменить, например, добавив в `.bashrc`
```
export GOPATH=/path/your/go/projects
```
<br>
```
$ tree -d -L 1 /path/your/go/projects
/path/your/go/projects
├── bin
├── pkg
└── src
```

---

# Go Modules

Начиная с Go 1.11 появилась поддержка модулей — системы версионирования и зависимостей, а также разработки вне `GOPATH`.
<br><br>
Стандартные команды (`go get`, `go install`, `go test` и т.д.) работают по-разному внутри модуля и внутри `GOPATH`.
<br><br>
Модуль — любая директория вне `GOPATH`, содержащая файл `go.mod`

---

# Создание Go модуля

* (Опционально) создайте и склонируйте (в любое место) репозиторий с проектом
```
git clone https://github.com/user/otus-go.git /home/user/otus-go
```
* Создайте внутри репозитория нужные вам директории
```
mkdir /home/user/otus-go/hw-1
```
* Зайдите в директорию и инициализируйте Go модуль
```
cd /home/user/otus-go/hw-1
go mod init github.com/user/otus-go/hw-1
```

Теперь `/home/user/otus-go/hw-1` — это Go модуль.

<br>
https://blog.golang.org/using-go-modules

---

# Добавление зависимостей

Внутри модуля, вы можете добавить зависимость от пакета с помощью
```
$ go get github.com/beevik/ntp
go: finding golang.org/x/net latest
```

```
$ cat go.mod
module github.com/mialinx/foobar

go 1.20

require (
	github.com/beevik/ntp v0.2.0 // indirect
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297 // indirect
)
```

<br><br>

Внимание: в момент добавления зависимостей их версии фиксируются в `go.sum`.

---

# Авто-добавление

Также можно просто редактировать код
```
package main

import (
	"fmt"

	"github.com/go-loremipsum/loremipsum"
)

func main() {
	fmt.Println(loremipsum.New().Word())
}
```
А потом запустить
```
$ go mod tidy
```
Это добавит новые и удалит неиспользуемые зависимости.

---

# Базовые команды

`go get -d` — скачивает пакеты из Git репозиториев в `$GOPATH/src`.
<br><br>
`go install` собирает и устанавливает указанные пакеты в `$GOPATH/pkg` и `$GOBIN` (по умолчанию `$GOPATH/bin`).
<br><br>
`go get` (без флажка `-d`) — так же вызовет `install`.
<br><br>
`go run prog.go` — сборка и запуск программы.
<br><br><br>
### Многоточия

`go get github.com/golang/protobuf/...` — многоточие тут означает
"и все дочерние пакеты".
<br>
Это необходимо если пакет сложный, и содержит подпакеты.
<br>
Для простых достаточно `go get github.com/beevik/ntp`

---

# Go Workspaces

### Workspaces появились в Go 1.18 и позволяют одновременно работать с несколькими модулями. 

- `go work init`
- `go work use`
- `go work replace`

https://go.dev/blog/get-familiar-with-workspaces

---

# Какие еще инструменты предоставляет Go?

- `go vet` 
- `go doc` 
- `go tool trace`
- `go tool pprof`
- `go tool cover`

Больше можно посмотреть через `go help tool`

---

# Получение справки

```
$ go help
Go is a tool for managing Go source code.

Usage:

	go <command> [arguments]

The commands are:
	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
...

$ go help build
usage: go build [-o output] [-i] [build flags] [packages]

Build compiles the packages named by the import paths,
along with their dependencies, but it does not install the results.
...
```

---

# Структура проекта

### Go проект состоит из пакетов

### Официальных требований к структуре пакетов - нет
Однако довольно широко распростанена вот [эта](https://github.com/golang-standards/project-layout) версия от комьюнити

---
# Структура проекта

- `/cmd` - содержит входные точки в проект
- `/pkg` - содержит код, который может использоваться другими проектами (общие утилиты, API)
- `/internal` - содержит внутреннюю реализацию проекта. Не может быть импортирован.
- `go.mod`
- `go.sum`

---

# Best practice

- Лучше много небольших файлов чем один большой
- Удобство навигации
- Располагайте типы в начале файла
- Организовывайте типы по их функциональному назначению
- Используйте Godoc
- Не пишите бизнес-логику в `main.go`

---

# Конвенции именования

- Lowercase имена пакетов
- Старайтесь избегать часто используемые имена (напр. util)
- Используйте единственное число

https://go.dev/doc/effective_go#names

---

# Форматирование кода

<br>
В Go нет style guide, зато есть `go fmt path/to/code.go`
<br><br>
.left-code[
Было:
```
package main
import "fmt"



const msg = "%d students in chat\n"
type Student struct{
	Name string
	Age int
}
func main() {
	for i:=99;i>0;i-- {
		fmt.Printf(msg, i)
		if i<10{
			break
		}

	}
}

```
]
.right-code[
Стало:
```
package main

import "fmt"

const msg = "%d students in chat\n"

type Student struct {
	Name string
	Age  int
}

func main() {
	for i := 99; i > 0; i-- {
		fmt.Printf(msg, i)
		if i < 10 {
			break
		}

	}
}
```
]


---

# Форматирование кода: advanced

```
$ cd /tmp
$ go get mvdan.cc/gofumpt
```

```
$ gofumpt -l -w .
```

---

# Обновление и сортировка импортов

```
$ go get golang.org/x/tools/cmd/goimports
$ goimports -local my/module/name -w path/to/code.go
```

```
import (
	"strings"
)

func main() {
	fmt.Println(loremipsum.New().Word())
}
```

```
import (
	"fmt"

	"github.com/go-loremipsum/loremipsum"
)

func main() {
	fmt.Println(loremipsum.New().Word())
}
```


---

# Обновление и сортировка импортов: advanced

```
$ cd /tmp
$ go get github.com/daixiang0/gci
```

```
$ gci -w -local github.com/Antonboom/hw01 .
```

---

# Линтеры

Линтер — программа, анализирующая код и сообщающая о потенциальных проблемах.
<br><br>
`go vet` — встроенный линтер
```
$ go vet ./run.go
# command-line-arguments
./run.go:14:3: Printf call needs 1 arg but has 2 args

$ echo $?
2
```

`golint` — популярный сторонний линтер
```
$ go get -u golang.org/x/lint/golint

$ ~/go/bin/golint -set_exit_status ./run.go
run.go:7:6: exported type Student should have comment or be unexported
Found 1 lint suggestions; failing.

$ echo $?
1
```

---

# Металинтеры

Металинтеры — обертка, запускающая несколько линтеров за один проход.
<br><br>

https://github.com/golangci/golangci-lint/

```
$ ~/go/bin/golangci-lint run ./run.go
run.go:14:3: printf: Printf call needs 1 arg but has 2 args (govet)
		fmt.Printf(msg, i, i)
		^
run.go:7:6: `Student` is unused (deadcode)
type Student struct {
     ^

$ echo $?
1
```

---

# Работа с golangici-lint

https://gist.github.com/kulti/25f9243939e699428f7e14c5a3c8c32c

https://www.youtube.com/watch?v=QnG8z-JWfEY - Демо урок

---

# Спецификация Go

https://go.dev/ref/spec


---

# Еще полезные ссылки

- https://go.dev/doc/
- https://go.dev/doc/effective_go
- https://gist.github.com/adamveld12/c0d9f0d5f0e1fba1e551
- https://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/

---

background-image: url(../img/questions.svg)

---

background-image: url(../img/poll.svg)

---

background-image: url(../img/next_webinar.svg)
.announce_date[21 августа]
.announce_topic[Переменные и элементарные типы данных]

---
background-image: url(../img/thanks.svg)

.tutor[Алексей Семушкин]
.tutor_desc[Software Engineer at Semrush]
