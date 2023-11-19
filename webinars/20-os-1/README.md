background-image: url(../img/title.svg)

---

background-image: url(../img/rec.svg)

---
background-image: url(../img/topic.svg)

.topic[Работа с ОС часть 1]
.tutor[Кристина Ступникова]
.tutor_desc[Software Engineer at Tabby]

---

background-image: url(../img/rules.svg)

---

# На этом занятии мы научимся

.big-list[

* Использовать пакеты flags и pflag для обработки аргументов командной строки в Go.
* Управлять сигналами операционной системы и корректно обрабатывать их в приложениях на Go.
  ]

---

# Цель занятия

## Узнать, какие средства <br> взаимодействия с ОС есть в Go.

---

# Введение в командную строку

- ### Командная строка позволяет взаимодействовать с операционной системой и запускать приложения с различными параметрами

---

# Введение в командную строку

- ### Интерфейс командной строки (CLI) – инструмент для запуска и настройки программ.
- ### Аргументы командной строки – данные, передаваемые программе при её запуске.
- ### Флаги – специальные аргументы, используемые для модификации работы программы.

Примеры:

```
go run myapp.go --verbose --port=8080
```

---

# Почему важно уметь обрабатывать аргументы командной строки?

- ### Гибкость запуска: Позволяет задавать различные режимы работы программы.
- ### Автоматизация: Возможность интегрировать приложения в скрипты и системы сборки.
- ### Конфигурирование: Удобство изменения настроек приложения без изменения кода.

---

# Пакет flags в Go

- ### Пакет flags – стандартное средство для обработки аргументов командной строки в Go
- ### Позволяет задавать флаги и аргументы командной строки, а также получать их значения внутри программы

---

# Ключевые особенности пакета flags

- ### Простой и интуитивно понятный интерфейс
- ### Поддержка флагов с базовыми типами данных: string, int, bool и др.
- ### Автоматическое генерирование справки по флагам

---

# flag

https://golang.org/pkg/flag/

```
func main() {
	var msg string

	verbose := flag.Bool("verbose", false, "verbose output")
	flag.StringVar(&msg, "msg", "hello world", "message to print")

	flag.Parse()

	if *verbose {
		fmt.Println("you say:", msg)
	} else {
		fmt.Println(msg)
	}
}
```

---

# Введение в пакет pflag

- ### Пакет pflag – это альтернатива пакету flags с расширенными возможностями
- ### Позволяет задавать флаги с длинными и короткими именами, а также флаги без значений
- ### Поддерживает широкий спектр пользовательских типов данных

---

# Отличия pflag от стандартного пакета flags

- ### Длинные и короткие опции (-f и --flag)
- ### Возможность создания своих типов флагов
- ### Встроенная поддержка слайсов и других сложных структур данных.

---

# Примеры использования pflag

### pflag улучшает читаемость и гибкость обработки флагов

```
var verbose = pflag.BoolP("verbose", "v", false, "Вывести подробную информацию")
```

---

# pflag

https://github.com/spf13/pflag

```
func main() {
	var msg string
	var slice []int

	verbose := pflag.BoolP("verbose", "v", false, "verbose output")
	pflag.StringVarP(&msg, "msg", "m", "hello", "your msg")
	pflag.IntSliceVarP(&slice, "s", "s", []int{}, "your slice")

	pflag.Parse()

	if *verbose {
		fmt.Println("you say:", msg)
	} else {
		fmt.Println(msg)
	}

	fmt.Println("your slice: ", slice)
}
```

---

# pflag: флаги без значений

```
pflag.StringVar(&msg, "msg", "hello", "message to print")
pflag.Lookup("msg").NoOptDefVal = "bye"
```

<br><br>

.left-text[
|Флаг |Значение
|:---------------|:-------
|--port=9999 |ip=9999
|--port |ip=80
|[nothing]         |ip=8080
]

---

# Работа с сигналами ОС в Go

- ### Сигналы ОС – это уведомления отправляемые процессу операционной системой для уведомления о событиях или для управления исполнением
- ### Сигналы могут быть отправлены из командной строки или других процессов
- ### Сигналы могут быть обработаны программой для корректного завершения работы
- ### Примеры использования сигналов: остановка процесса (Ctrl+C), завершение работы и т.д

---

# Общие сигналы Unix

| Сигнал  | Значение | Описание                                                    |
|:--------|:---------|:------------------------------------------------------------|
| SIGHUP  | 1        | Завершение процесса                                         |
| SIGINT  | 2        | Завершение процесса (Ctrl+C)                                |
| SIGQUIT | 3        | Завершение процесса (Ctrl+\)                                |
| SIGKILL | 9        | Завершение процесса (не может быть перехвачен)              |
| SIGTERM | 15       | Завершение процесса (может быть перехвачен или игнорирован) |
| SIGSTOP | 17,19,23 | Остановка процесса (не может быть перехвачен)               |
| SIGCONT | 18,25,26 | Продолжить выполнение процесса (не может быть перехвачен)   |

---

# Сигналы

Сигналы - механизм OS, позволяющий посылать уведомления программе в особых ситуациях.
<br><br>

| Сигнал  | Поведение  | Применение                                             |
|:--------|:-----------|:-------------------------------------------------------|
| SIGINT  | Завершить  | `Ctrl+C` в консоли                                     |
| SIGKILL | Завершить  | `kill -9`, остановка зависших программ                 |
| SIGHUP  | Завершить  | Сигнал для переоткрытия логов и перечитывания конфига  |
| SIGUSR1 |            | На усмотрение пользователя                             |
| SIGUSR2 |            | На усмотрение пользователя                             |
| SIGPIPE | Завершить  | Отправляется при записи в закрытый файловый дескриптор |
| SIGSTOP | Остановить | При использовании отладчика                            |
| SIGCONT | Продолжить | При использовании отладчика                            |

<br><br>
Некоторые сигналы, например `SIGINT`, `SIGUSR1`, `SIGHUP`, можно игнорировать или установить обработчик.
<br><br>
Некоторые, например `SIGKILL`, обработать нельзя.
---

# Обработка сигналов в Go

### Go предоставляет пакет os/signal для перехвата и обработки сигналов операционной системы.

---

# Обработка сигналов

```
func main() {
c := make(chan os.Signal, 1)
signal.Notify(c, syscall.SIGINT, syscall.SIGKILL)
signal.Ignore(syscall.SIGTERM)

	for s := range c {
		fmt.Println("Got signal:", s)
	}
}
```

---

# Регистрация обработчиков сигналов в отдельной горутине

```
// Продолжение предыдущего примера
go func() {
  sig := <-sigs
  fmt.Println("Получен сигнал:", sig)
  // ... Обработка сигнала ...
}()
```

---

# Регистрация обработчиков сигналов в отдельной горутине

```
// ... Ваш код с предыдущих слайдов ...
go func() {
  for {
    sig := <-sigs
    switch sig {
    case syscall.SIGINT:
      fmt.Println("Прерывание: Ctrl+C")
      // ... Функция очистки ...
      os.Exit(0)
    case

```

---

background-image: url(../img/questions.svg)

---

background-image: url(../img/poll.svg)

---

background-image: url(../img/next_webinar.svg)
.announce_date[15 Ноября]
.announce_topic[Работа с ОС часть 2]

---
background-image: url(../img/thanks.svg)

.announce_date[Ссылка на опрос:]
.announce_topic[https://otus.ru/polls/70297/]

.tutor[Кристина Ступникова]
.tutor_desc[Software Engineer at Tabby]