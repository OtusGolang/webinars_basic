background-image: url(../img/title.svg)

---

background-image: url(../img/rec.svg)

---
background-image: url(../img/topic.svg)

.topic[Знакомство и начало работы с Go]
.tutor[Алексей Семушкин]
.tutor_desc[Software Engineer at Semrush]

---

background-image: url(../img/rules.svg)

---

# О чем будем говорить:

- Состав курса
- Что такое Go
- Что потребуется для прохождения

---

# Давайте познакомимся

- ## Откуда вы?
- ## Был ли опыт в IT?
- ## Почему выбрали именно Go?

---

## Что нам предстоит на этом [курсе](https://otus.ru/lessons/go-basic/)?

---

# Знакомьтесь, язык Go

- Разработан в 2009 году корпорацией Google;
- Полное название — Golang (Google language);
- Создавался как современная альтернатива C и C++

---

# Создатели языка

.center-image[
![](img/TompsonKen.jpg)
]
</br>

Кен Томпсон — пионер компьютерной науки, один из создателей языка программирования C и операционной системы UNIX.

---

# Создатели языка

.center-image.threefour-image[
![](img/RobPike.jpeg)
]
</br>

Роб Пайк — разработчик операционных систем и языков программирования в Bell Labs. Один из создателей кодировки UTF-8.

---

# Какие проблемы призван решить Go?

- Медленная сборка программ;
- Неконтролируемые зависимости;
- Трудности с чтением чужого кода;
- Высокие трудозатраты при переходе на новую версию языка;
- Сложности разработки инструментария

---

# Что нам потребуется

- ОС Linux или Mac OS
- Текстовый редактор с поддержкой плагинов
- Дистрибутив Go
- Git

---

# Операционная система

- [Ubuntu](https://ubuntu.com/download/desktop)
- [Virtual Box](https://www.virtualbox.org/)
- [Гайд](https://ithowto.ru/ustanovka-ubuntu-2004-virtualbox.html)

---

# Редактор

## Почему (желательно) не Goland?

---

# Редактор

## Потому что Goland слишком все упрощает :)

---

# Редактор

- [VSCode](https://code.visualstudio.com/download)
- [Eclipse](https://www.eclipse.org/downloads/)
- [vim](https://www.vim.org/download.php) - для тех кто привык играть на сложности "Кошмар"

</br>
Рекомендую VSCode, [гайд](https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code) по настройке под Go.

---

# Первые шаги c Go

https://go.dev/tour

---

# Установка Go

<b>Getting Started</b></br>
https://go.dev/doc/install<br><br>

<b>Downloads</b><br>
https://go.dev/dl/<br><br>

Проще всего через `apt-get`
```
sudo apt-get update

sudo apt-get install golang
```

Или можно скачать с официального сайта
```
wget https://go.dev/dl/go1.20.5.linux-amd64.tar.gz

sudo tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz

sudo ln -s /usr/local/go/bin/go /usr/bin/go
```

Готово!

---

# Несколько версий Golang

https://go.dev/doc/manage-install

---

# GOROOT

`GOROOT` — переменная, которая указывает где лежит ваш дистрибутив Go, т.е.
компилятор, утилиты и стандартная библиотека. В новых версия Go (&gt; 1.0) утилиты сами определяют расположение Go.

<br><br>

Однако, вы можете узнать `GOROOT`

```
$ go env | grep ROOT
GOROOT="/usr/local/go"
```

И можете посмотреть исходный код Go =)
```
vim /usr/local/go/src/runtime/slice.go
```

---

# GOPATH

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

# Playground

https://go.dev/play/p/UCPdVNrl0-P

---

# Git

Git — распределённая система контроля версий, которая даёт возможность разработчикам отслеживать изменения в файлах и работать над одним проектом совместно с коллегами. 

</br>

GitHub — сервис онлайн-хостинга репозиториев, обладающий всеми функциями распределённого контроля версий и функциональностью управления исходным кодом

---

# Git

Ваши домашние задания лежат [здесь](https://github.com/OtusGolang/home_work_basic)
Гайд по сдаче [здесь](https://github.com/OtusGolang/home_work_basic). Он пригодится вам позже

---


background-image: url(../img/questions.svg)

---

background-image: url(../img/poll.svg)

---

background-image: url(../img/next_webinar.svg)
.announce_date[7 августа]
.announce_topic[Основы вычислительной техники]

---
background-image: url(../img/thanks.svg)

.tutor[Алексей Семушкин]
.tutor_desc[Software Engineer at Semrush]
