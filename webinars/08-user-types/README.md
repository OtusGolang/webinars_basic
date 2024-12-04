background-image: url(../img/title.svg)

---

background-image: url(../img/rec.svg)

---
background-image: url(../img/topic.svg)

.topic[Пользовательские типы данных]
.tutor[Алексей Романовский]
.tutor_desc[Разработчик в Resolver Inc.]

---

background-image: url(../img/rules.svg)

---
# Цели занятия

* научиться работать с пользовательскими типами данных;

---

# План занятия

* выравнивание структур;
* определение структур;
* инкапсуляция полей структуры;
* вложенные и анонимные структуры;

---
# Объявление пользовательских типов
```go
type <Name> <Declaration> 
```

`Declaration` может быть, например:
* Структурой - struct  
* Простым типом - string, int... (называется "alias")
* Интерфейсом - interface
* Функцией - func
* Коллекцией - array, slice, map
* Каналом - chan

См также:
* https://go.dev/ref/spec#Type_declarations

---

# Структуры

Структуры - фиксированный набор именованных переменных. <br>
Переменные размещаются рядом в памяти и обычно используются совместно.

```go
type User struct { // Структура с именованными полями
  Id      int64
  Name    string
  Age     int
  friends []int64  // Приватный элемент
}
```

```go
struct{}  // Пустая структура, не занимает памяти
```

### Пример: https://golang.org/ref/spec#Struct_types

---

# Методы

```go
type Mutex struct         { /* Mutex fields */ }
func (m *Mutex) Lock()    { /* Lock implementation */ }
func (m *Mutex) Unlock()  { /* Unlock implementation */ }
```

```go
m := Mutex{}
m.Lock()
```

---

# Литералы структур

```go
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

```go
var wordCounts []struct{w string; n int}
```

```go
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

### Пример: https://goplay.space/#rE-DsbSFgN1

---

# Анонимные типы и структуры

```go
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

# Экспортируемые и приватные элементы

Поля структур, начинающиеся со строчной буквы - **приватные**, они будут видны
только в том же пакете, где и структура. <br><br>
Поля, начинающиеся с заглавной - **публичные**, они будут видны везде.

```go
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

### Пример: https://go.dev/play/p/g9sldeRCgaO

---

# Встроенные структуры

В Go есть возможность "встраивать" типы внутрь структур. <br>
При этом у элемента структуры НЕ задается имя.

```go
type Human struct {
    name string
    age int
}
```

Тип Human встроен в тип Student:

```go
type Student struct {
    Human
    school string
}
```

---


# Размер и выравнивание структур


```go
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

### Пример: https://go.dev/play/p/f7hbieZyh_z

* https://github.com/dominikh/go-tools/tree/master/cmd/structlayout
* https://en.wikipedia.org/wiki/Data_structure_alignment
* https://dave.cheney.net/2015/10/09/padding-is-hard
* [Линтер fieldalignment](https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/fieldalignment)

---

# Копирование указателей и структур

При присвоении переменных типа структура - данные копируются.
```go
a := struct{x, y int}{0, 0}
b := a
a.x = 1
fmt.Println(b.x) // ?
```

При присвоении указателей - копируется только адрес данных.
```go
a := &struct{x, y int}{}
// a := new(struct{x, y int}) // так тоже можно
b := a
a.x = 1
fmt.Println(b.x) // ?
```

```go
a := struct{x *int}{new(int)}
b := a
*a.x = 1
fmt.Println(b.x) // ?
```

---

# Функции-конструкторы

В Go принят подход Zero Value: постарайтесь сделать так, что бы
ваш тип работал без инициализации, как реализованы, например

```go
var b strings.Builder
var wg sync.WaitGroup
```

Если ваш тип содержит словари, каналы или инициализация обязательна - скройте
ее от пользователя, создав функции-конструкторы:

```go
func NewYourType() (*YourType) {
  // ...  
}
func NewYourTypeWithOption(option int) (*YourType) {
  // ...
}
```

### Пример: https://go.dev/play/p/CAP9Q4wS4xX

---

# Тэги элементов структуры

К элементам структуры можно добавлять метаинформацию - тэги. <br>
Тэг это просто литерал строки, но есть соглашение о структуре такой строки.

<br>
Например,

```go
type User struct {
    Id      int64    `json:"-"`    // Игнорировать в encode/json
    Name    string   `json:"name"`
    Age     int      `json:"user_age" db:"how_old"`
    friends []int64 
}
```

Получить информацию о тэгах можно через `reflect`

```go
var u User
ageField := reflect.TypeOf(u).FieldByName("Age")
jsonFieldName := ageField.Get("json")  // "user_age"
```

### Ссылки:
* https://go.dev/wiki/Well-known-struct-tags

---

# Использование тэгов для JSON сериализации

Для работы с JSON используется пакет `encoding/json`

```go
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

```go
var user User
row := db.QueryRow("SELECT * FROM users WHERE id=?", 10)
err = row.Scan(&user)
```

Для ORM библиотеки GORM `github.com/jinzhu/gorm` фич намного больше

```go
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


```go
type Set map[int]struct{}
```

```go
ch := make(chan struct{})
ch <- struct{}{}
```

### Ссылки:
* https://dave.cheney.net/2014/03/25/the-empty-struct

---

background-image: url(../img/questions.svg)

---

background-image: url(../img/poll.svg)

---
background-image: url(../img/thanks.svg)

.tutor[Алексей Романовский]
.tutor_desc[Разработчик в Resolver Inc.]
