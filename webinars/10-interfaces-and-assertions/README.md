background-image: url(../img/title.svg)

---

background-image: url(../img/rec.svg)

---
background-image: url(../img/topic.svg)

.topic[Интерфейсы и утверждение типов]
.tutor[Алексей Романовский]
.tutor_desc[Разработчик в Resolver Inc.]

---

background-image: url(../img/rules.svg)

---

# Цели занятия

* знать что такое интерфейс;
* определять и реализовывать интерфейсы.

---
# Краткое содержание

* duck typing;
* композиция интерфейсов;
* пустой интерфейс (interface{}).
* утверждение типов (type assertion);
* type switch;

---

# Интерфейсы: определение

**Интерфейс** — набор методов, которые надо реализовать, чтобы удовлетворить интерфейсу. Ключевое слово: `interface`.

```go
type Stringer interface { // фреймворк Go, пакет fmt
    String() string
}

type Shape interface {
    Area() float64
    Perimeter() float64
}
```

- Одному интерфейсу могут соответствовать много типов
- Тип может реализовать несколько интерфейсов

---

# Интерфейсы и типы

Интерфейс можно использовать в качестве типа при объявлении переменной.

Переменная **типа интерфейс** может содержать только значение типа, реализующего этот интерфейс.

```go
var s Stringer // статический тип
s = time.Time{} // динамический тип
```
<br>

### Ссылки:
* https://go.dev/doc/effective_go#interfaces_and_types

---

# Интерфейсы и типы

<br>Значение типа интерфейс состоит из динамического типа и значения.
<br>Мы можем их смотреть при помощи %v и %T

```go
type Temp int

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " °C"
}


func main() {
	var x fmt.Stringer
	x = Temp(24)
	fmt.Printf("%v %T\n", x, x) // 24 °C main.Temp
}
```
<br>

### Пример: https://goplay.space/#JjXQsIsXwac


---

# Duck typing: Интерфейсы реализуются неявно

Duck typing ('Утиная типизация', 'неявная типизация') - <br>
это подход к типизации, при котором совместимость типов определяется только <br>
**наличием у них определенных методов**

```go
type Duck interface {
    Talk() string
    Walk()
    Swim()
}

type Dog struct {
    name string
}

func (d Dog) Talk() string {
    return "AGGGRRRR"
}

func (d Dog) Walk() { }

func (d Dog) Swim() { }

```
<br>

### Пример: https://goplay.space/#GWYHjaDPnLG

---

# Интерфейсы реализуются неявно

В стандартной библиотеке Go есть интерфейс `Stringer`:
```go
type Stringer interface {
    String() string
}
```


```go
type MyVeryOwnStringer struct { s string}

func (s MyVeryOwnStringer) String() string {
    return "my string representation of MyVeryOwnStringer"
}


func main() {
    // my string representation of MyVeryOwnStringer{}
    fmt.Println(MyVeryOwnStringer{"hello"})
}
```
<br>

### Пример: https://goplay.space/#ppTH6Ya-fX5


---

# Композиция интерфейсов

Пример из пакета `io`:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

type ReadCloser interface {
    Reader
    Closer
}
```

---

# Интерфейсы: композиция

```go
import "fmt"

type Greeter interface {
     hello()
}

type Stranger interface {
    Bye() string
    Greeter
    fmt.Stringer
}
```

---

# Интерфейсы: имена методов

Имена методов не должны повторяться:

```go
type Hound interface {
    destroy()
    bark(int)
}

type Retriever interface {
    Hound
    bark() // duplicate method
}

```

```shell
./prog.go:6:2: duplicate method bark
```
<br>

### Пример: https://goplay.space/#DF4oPHXDGLP


---

# Пустой интерфейс (interface{}, any)

Пустой интерфейс не содержит методов:

```go
type any interface{}  // представлен в Go 1.18
```

```go
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
   ...
}
```

<br>

### Ссылки:
* https://tip.golang.org/doc/go1.18

---

Переменная с типом "пустой интерфейс" может содержать значение любого типа.

```go
	var x any // any == interface{}
	x = 5
	x = "hello"
	x = time.Now()
	x = map[string]int{"one": 1}
	x = struct{ name string }{name: "John"} // анонимный тип

    x = map[interface{}]interface{}{}
	x = map[any]any{} // более читаемо

```

<br>

### Пример: https://goplay.space/#pfgv-f2Ob_q

---

# interface{} is says nothing

Предпочитаете более специфичный тип пустому интерфейсу вашем коде<br>
_если возможно_


<br>
### Ссылки:
* https://go-proverbs.github.io/

---

# Интерфейсы: проверка типа (type assertion)

Есть способы проверить какой тип хранится в переменной-интерфейсе.

<br>
Выражение `x.(T)` проверяет, что интерфейс `x != nil` и конкретная часть `x` имеет тип `T`:

- если T не интерфейс, то проверяем, что динамический тип x это T
- если T интерфейс: то проверяем, что динамический тип x его реализует
---

# Интерфейсы: type assertion

Выражение `x.(T)` возвращает одно или два значения:
* T (+ panic)
* T, bool

```go
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	r, ok := i.(fmt.Stringer)
	fmt.Println(r, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
```

<br>

### Пример: https://goplay.space/#x-NbzVMZMUp

---

# Интерфейсы: type assertion

```go
    var i interface{} = "hello"

	f, ok := i.(float64) // 0 false
	fmt.Println(f, ok)

	f = i.(float64) // panic: interface conversion:
					// interface {} is string, not float64
	fmt.Println(f)
```

Проверка типа возможна только для интерфейса:

```go
	s := 5
    // Invalid type assertion: s.(int) (non-interface type int on left)
	i := s.(int)
```

---

# Интерфейсы: type switch

<br>
Мы можем объединить проверку нескольких типов в один `type switch`:

```go
func checkSignature(/* ... */, publicKey crypto.PublicKey) (err error) {
    // ...

    switch pub := publicKey.(type) {
    case *rsa.PublicKey:
        // ...
    case *ecdsa.PublicKey:
        // ...
    case ed25519.PublicKey:
        // ...
    }
    return ErrUnsupportedAlgorithm
}
```

### Пример: [src/crypto/x509/x509.go](https://github.com/golang/go/blob/283d8a3d53ac1c7e1d7e297497480bf0071b6300/src/crypto/x509/x509.go#L837)

---


# Интерфейсы: type switch


Как и в обычном `switch` мы можем объединять типы:

```go
    case *rsa.PublicKey, *ecdsa.PublicKey:
        // Do some work...
    }
```

и обрабатывать `default`:

```go
switch publicKey.(type) {
default:
    // No case for input type...
}
```


---

# Интерфейсы и производительность

Вызов метода через интерфейс немного медленнее, чем вызов метода напрямую.

Очень сильно зависит от платформы и версии го.
<br><br>
Не оптимизируйте заранее, если не уверены, что это нужно. 

Имеет смысл исследовать "горячие точки" - места где вызов делается миллионы раз - 
и только с оглядкой на реальные бенчмарки с нужной платформой и версией го.

```go
// прямой вызов
v := MyStruct{}
v.String()

// через интерфейс
var v fmt.Stringer = MyStruct{}
v.String()

```

<br>

### Ссылки:
* https://github.com/tserkov/go-interface-benchmark



---

background-image: url(../img/questions.svg)

---

background-image: url(../img/poll.svg)

---

background-image: url(../img/next_webinar.svg)
.announce_date[6 Сентября]
.announce_topic[Обработка ошибок и паник]

---
background-image: url(../img/thanks.svg)

.announce_date[Ссылка на опрос:]
.announce_topic[https://otus.ru/polls/70297/]
.tutor[Алексей Романовский]
.tutor_desc[Преподаватель]