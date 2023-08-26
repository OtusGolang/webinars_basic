background-image: url(../img/title.svg)

---

background-image: url(../img/rec.svg)

---
background-image: url(../img/topic.svg)

.topic[Пользовательские типы данных]
.tutor[Илья Феоктистов]
.tutor_desc[Software Engineer at Agoda]

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

# Структуры

Структуры - фиксированный набор именованных переменных. <br>
Переменные размещаются рядом в памяти и обычно используются совместно.

```
struct{}  // Пустая структура, не занимает памяти

type User struct { // Структура с именованными полями
  Id      int64
  Name    string
  Age     int
  friends []int64  // Приватный элемент
}
```
https://golang.org/ref/spec#Struct_types

---

# Литералы структур

```
var u0 User                      // Zero Value для типа User

u1 := User{}                     // Zero Value для типа User

u2 := &User{}                    // То же, но указатель

u3 := User{1, "Vasya", 23, nil}  // По номерам полей

u4 := User{                      // По именам полей
  Id:       1,
  Name:     "Vasya",
  friends:  []int64{1, 2, 3},
}

```

---

# Анонимные типы и структуры

Анонимные типы задаются литералом, у такого типа нет имени.<br>
Типичный сценарий использования: когда структура нужна только внутри одной функции. 

```
var wordCounts []struct{w string; n int}
```

```
var resp struct {
    Ok        bool `json:"ok"`
    Total     int  `json:"total"`
    Documents []struct{
        Id    int    `json:"id"`
        Title string `json:"title"`
    } `json:"documents"`
}
json.Unmarshal(data, &resp)
fmt.Println(resp.Documents[0].Title)
```
<br>
https://go.dev/play/p/rE-DsbSFgN1

[//]: # (https://goplay.space/#rE-DsbSFgN1)


---

# Анонимные типы и структуры

```
testCases := []struct{
    name     string
    input    string
    expected int
    err      error  
} {
    name: "case1",
    input: "aaa",
    expected: 10,
    err: nil,
}
```

---

# Размер и выравнивание структур

https://go.dev/play/p/0WdB68TTmkj <br>

[//]: # (https://goplay.space/#0WdB68TTmkj <br>)

```
unsafe.Sizeof(1)   // 8 на моей машине
unsafe.Sizeof("A") // 16 (длина + указатель)

var x struct {
    a bool   // 1 (offset 0)
    c bool   // 1 (offset 1)
    b string // 16 (offset 8)
}

unsafe.Sizeof(x) // 24!
```
![img/aling.png](img/align.png)

https://github.com/dominikh/go-tools/tree/master/cmd/structlayout <br>
https://en.wikipedia.org/wiki/Data_structure_alignment

---


# Копирование указателей и структур

При присвоении переменных типа структура - данные копируются.
```
a := struct{x, y int}{0, 0}
b := a
a.x = 1
fmt.Println(b.x) // ?
```

При присвоении указателей - копируется только адрес данных.
```
a := new(struct{x, y int})
b := a
a.x = 1
fmt.Println(b.x) // ?
```

```
a := struct{x *int}{new(int)}
b := a
*a.x = 1
fmt.Println(b.x) // ?
```

---

# Экспортируемые и приватные элементы

Поля структур, начинающиеся со строчной буквы - **приватные**, они будут видны
только в том же пакете, где и структура. <br><br>
Поля, начинающиеся с заглавной - **публичные**, они будут видны везде.

```
type User struct {
  Id      int64
  Name    string   // Экспортируемое поле
  Age     int
  friends []int64  // Приватное поле
}
```

Не совсем очевидное следствие: пакеты стандартной библиотеки, например, `encoding/json` тоже не могут
работать с приватными полями :)<br><br>
Доступ к приватным элементам (на чтение!) все же можно получить с помощью пакета `reflect`.

<br>
https://goplay.space/#g9sldeRCgaO

---

# Функции-конструкторы

В Go принят подход Zero Value: постарайтесь сделать так, что бы
ваш тип работал без инициализации, как реализованы, например
```
var b strings.Builder
var wg sync.WaitGroup
```

Если ваш тип содержит словари, каналы или инициализация обязательна - скройте
ее от пользователя, создав функции-конструкторы:

```
func NewYourType() (*YourType) {
  // ...  
}
func NewYourTypeWithOption(option int) (*YourType) {
  // ...
}
```

https://goplay.space/#5lfGpAcfTyU

---

# nil receiver

```
type RateLimiter struct {
    ...
}

func (r *RateLimiter) Allow() bool {
    if r == nil {
        return true
    }
    return r.allow()
}
```


---


# Встроенные структуры

В Go есть возможность "встраивать" типы внутрь структур. <br>
При этом у элемента структуры НЕ задается имя.

```
type LinkStorage struct {
    sync.Mutex                  // Только тип!
    storage map[string]string   // Тип и имя
}
```

Обращение к элементам встроенных типов:
```
var storage LinkStorage
storage.Mutex.Lock()     // Имя типа используется 
storage.Mutex.Unlock()   // как имя элемента структуры
```

---

# Продвижение методов

При встраивании методы встроенных структур можно вызывать у ваших типов!

```
// Вместо
storage.Mutex.Lock()
// можно просто
storage.Lock()
```


---

# Но, это не наследование

```
type Base struct {}

func (b Base) Name() string {
    return "Base"
}

func (b Base) Say() {
    fmt.Println(b.Name())
}

type Child struct {
    Base
    Name string
}

func (c Child) Name() string {
    return "Child"
}

var c Child
c.Say() // Увы "Base" :(
```
https://goplay.space/#AOyLzYid61L


---

# Тэги элементов структуры

К элементам структуры можно добавлять метаинформацию - тэги. <br>
Тэг это просто литерал строки, но есть соглашение о структуре такой строки.

<br>
Например,
```
type User struct {
    Id      int64    `json:"-"`    // Игнорировать в encode/json
    Name    string   `json:"name"`
    Age     int      `json:"user_age" db:"how_old"`
    friends []int64 
}
```

Получить информацию о тэгах можно через `reflect`
```
var u User
ageField := reflect.TypeOf(u).FieldByName("Age")
jsonFieldName := ageField.Get("json")  // "user_age"
```

https://github.com/golang/go/wiki/Well-known-struct-tags

---

# Использование тэгов для JSON сериализации

Для работы с JSON используется пакет `encoding/json`

```
// Можно задать имя поля в JSON документе
Field int `json:"myName"`

// Не выводить в JSON поля у которых Zero Value
Author *User `json:"author,omitempty"`

// Использовать имя поля Author, но не выводить Zero Value
Author *User `json:"omitempty"`

// Игнорировать это поле при сериализации / десереализации
Field int `json:"-"`
```

---

# Использование тэгов для работы с СУБД

Зависит от пакета для работы с СУБД.<br>
Например, для `github.com/jmoiron/sqlx`
```
var user User
row := db.QueryRow("SELECT * FROM users WHERE id=?", 10)
err = row.Scan(&user)
```

Для ORM библиотеки GORM `github.com/jinzhu/gorm` фич намного больше
```
type User struct {
  gorm.Model
  Name         string
  Email        string  `gorm:"type:varchar(100);unique_index"`
  Role         string  `gorm:"size:255"` // set field size to 255
  MemberNumber *string `gorm:"unique;not null"` // set member number to unique and not null
  Num          int     `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
  Address      string  `gorm:"index:addr"` // create index with name `addr` for address
  IgnoreMe     int     `gorm:"-"` // ignore this field
}
```

---

# Ещё раз про пустые структуры

https://dave.cheney.net/2014/03/25/the-empty-struct

```
type Set map[int]struct{}
```

```
ch := make(chan struct{})
ch <- struct{}{}
```


<!--
# Небольшой тест

.left-text[
Проверим что мы узнали за этот урок
<br><br>
[https://forms.gle/xLLab1NXH9NLKJij8](https://forms.gle/xLLab1NXH9NLKJij8)
]

.right-image[
![](img/gopher9.png)
]-->

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!