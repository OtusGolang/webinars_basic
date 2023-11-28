background-image: url(../img/title.svg)

---

background-image: url(../img/rec.svg)

---
background-image: url(../img/topic.svg)

.topic[Работа с базами данных Часть 2]
.tutor[Родион Ступников]
.tutor_desc[Software Engineer]

---

background-image: url(../img/rules.svg)

---

# План занятия

.big-list[
* Подключение к СУБД и настройка пула подключений
* Выполнение запросов и получение результатов
* Стандартные интерфейсы sql.DB, sql.Rows, sql.Tx
* Использование транзакций
* SQL инъекции и борьба с ними
]

---

# Подключение к PostgreSQL из Go

Создание подключения (простой вариант):

```go
import "database/sql"
import _ "github.com/jackc/pgx/stdlib"

dsn := "..."
db, err := sql.Open("pgx", dsn)  // *sql.DB
if err != nil {
  log.Fatalf("failed to load driver: %v", err)
}
// создан пул соединений
```

Использование подключения:

```go
err := db.PingContext(ctx)
if err != nil {
  return fmt.Errorf("failed to connect to db: %w", err)
}
// работаем с db
```

http://go-database-sql.org/importing.html <br>
http://go-database-sql.org/accessing.html

---

# DataSourceName

DSN - строка подключения к базе, содержит все необходимые опции.
<br>
Синтаксис DSN зависит от используемой базы данных и драйвера.
<br><br>

Например для PostgreSQL:
```
"postgres://myuser:mypass@localhost:5432/mydb?sslmode=verify-full"
```
Или
```
"user=myuser dbname=mydb sslmode=verify-full password=mypass"
```
* `host` - Сервер базы данных или путь к UNIX-сокету (по-умолчанию localhost)
* `port` - Порт базы данных (по-умолчанию 5432)
* `dbname` - Имя базы данных
* `user` - Пользователь в СУБД (по умолчанию - пользователь OS)
* `password` - Пароль пользователя

Подробнее: [https://godoc.org/github.com/lib/pq](https://godoc.org/github.com/lib/pq)
---

# Пул соединений

`sql.DB` - это пул соединений с базой данных. Соединения будут открываться по мере необходимости.
<br><br>
`sql.DB` - безопасен для конкурентного использования (так же как `http.Client`)
<br><br>

Настройки пула:
```go
// Макс. число открытых соединений от этого процесса
db.SetMaxOpenConns(n int)

// Макс. число открытых неиспользуемых соединений
db.SetMaxIdleConns(n int)

// Макс. время жизни одного подключения
db.SetConnMaxLifetime(d time.Duration)
```

http://go-database-sql.org/connection-pool.html


---
# Пул соединений

.full-image[
![](img/no_pool.png)
]


---

# Пул соединений

.full-image[
![](img/with_pool.png)
]



---

# Выполнение запросов

```go
query := `insert into events(owner, title, descr, start_date, end_date) 
  values($1, $2, $3, $4, $5)`

result, err := db.ExecContext(ctx, query,
  42, "new year", "watch the irony of fate", "2019-12-31", "2019-12-31"
) // sql.Result
if err != nil {
  // обработать ошибку
}

// Авто-генерируемый ID (SERIAL)
eventId, err := result.LastInsertId()      // int64

// Количество измененных строк
rowsAffected, err := result.RowsAffected() // int64
```

http://go-database-sql.org/retrieving.html

---

# Получение результатов

```go
query := `
  select id, title, descr
  from events
  where owner = $1 and start_date = $2
`
rows, err := db.QueryContext(ctx, query, owner, date)
if err != nil {
  // ошибка при выполнении запроса
}
defer rows.Close()

for rows.Next() {
  var id int64
  var title, descr string
  if err := rows.Scan(&id, &title, &descr); err != nil {
    // ошибка сканирования
  }
  // обрабатываем строку
  fmt.Printf("%d %s %s\n", id, title, descr)
}
if err := rows.Err(); err != nil {
	// ошибка при получении результатов
}
```

---

# Объект sql.Rows

```go
// возвращает имена колонок в выборке
rows.Columns() ([]string, error)

// возвращает типы колонок в выборке
rows.ColumnTypes() ([]*ColumnType, error)

// переходит к следующей строке или возвращает false
rows.Next() bool

// заполняет переменные из текущей строки
rows.Scan(dest ...interface{}) error

// закрывает объект Rows
rows.Close()

// возвращает ошибку, встреченную при итерации
rows.Err() error
```

---

# Получение одной строки

```go
query := "select * from events where id = $1"

row := db.QueryRowContext(ctx, query, id)

var id int64
var title, descr string

err := row.Scan(&id, &title, &descr)

if err == sql.ErrNoRows {
  // строки не найдено
} else if err != nil {
  // "настоящая" ошибка
}
```

---

# PreparedStatements

*PreparedStatement* - это заранее разобранный запрос, который можно выполнять повторно.
*PreparedStatement* - временный объект, который *создается в СУБД* и живет в рамках сессии, или пока не будет закрыт.

```go
// создаем подготовленный запрос
stmt, err := db.Prepare("delete from events where id = $1") // *sql.Stmt
if err != nil {
	log.Fatal(err)
}

// освобождаем ресурсы в СУБД
defer stmt.Close()

// многократно выполняем запрос
for _, id := range ids {
  _, err := stmt.Exec(id)
  if err != nil {
    log.Fatal(err)
  }
}
```

http://go-database-sql.org/prepared.html

---

# Работа с соединением

`*sql.DB` - это пул соединений. Даже последовательные запросы могут использовать *разные* соединения с базой.
<br><br>
Если нужно получить одно конкретное соединение, то
```go
conn, err := db.Conn(ctx)  // *sql.Conn

// вернуть соединение в pool
defer conn.Close()

// далее - обычная работа как с *sql.DB
err := conn.ExecContext(ctx, query1, arg1, arg2)

rows, err := conn.QueryContext(ctx, query2, arg1, arg2)
```

---

# Основные методы

Определены у `*sql.DB`, `*sql.Conn`, `*sql.Tx`, `*sql.Stmt`:
```go
// изменение данных
ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error)

// получение данных (select)
QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error)

// получение одной строки
QueryRowContext(ctx context.Context, query string, args ...interface{}) *Row
```
<br><br>
Внимание, ошибка:
```go
_, err := db.QueryContext(ctx, "delete from events where id = $1", 42)
```
---

# NULL

В SQL базах любая колонка может быть объявлена к NULL / NOT NULL.
NULL - это не 0 и не пустая строка, это отсутствие значения.
```sql
create table users (
  id          serial primary key,
  name        text not null,
  age         int null
);
```

Для обработки `NULL` в Go предлагается использовать специальные типы:
```go
var id, realAge int64
var name string
var age sql.NullInt64
err := db.QueryRowContext(ctx, "select * from users where id = 1").Scan(&id, &name, &age)

if age.Valid {
  realAge = age.Int64
} else {
	// обработка на ваше усмотрение
}
```

---


# Проблемы database/sql

.big-list[
* placeholder зависят от базы: (`$1` в Postgres, `?` в MySQL, `:name` в Oracle)
* Есть только базовые типы, но нет, например `sql.NullDate`
* `rows.Scan(arg1, arg2, arg3)` - неудобен, нужно помнить порядок и типы колонок.
* Нет возможности `rows.StructScan(&event)`
  ]

---

# Расширение sqlx

`jmoiron/sqlx` - обертка, прозрачно расширяющая стандартную библиотеку `database/sql`.
<br><br>
* `sqlx.DB` - обертка над `*sql.DB`
* `sqlx.Tx` - обертка над `*sql.Tx`
* `sqlx.Stmt` - обертка над `*sql.Stmt`
* `sqlx.NamedStmt` - PreparedStatement с поддержкой именованых параметров

Подключение `jmoiron/sqlx`:

```go
import "github.com/jmoiron/sqlx"

db, err := sqlx.Open("pgx", dsn)  // *sqlx.DB

rows, err := db.QueryContext("select * from events") // *sqlx.Rows

...
```
---

# sqlx: именованные placeholder'ы

Можно передавать параметры запроса в виде словаря:
```go
sql := "select * from events where owner = :owner and start_date = :start"
rows, err := db.NamedQueryContext(ctx, sql, map[string]interface{}{
  "owner": 42,
  "start": "2019-12-31",
})
```

Или структуры:
```go
type QueryArgs{
  Owner int64
  Start string
}
sql := "select * from events where owner = :owner and start_date = :start"
rows, err := db.NamedQueryContext(ctx, sql, QueryArgs{
  Owner: 42,
  Start: "2019-12-31",
})
```
---

# sqlx: сканирование

Можно сканировать результаты в словарь:
```go
sql := "select * from events where start_date > $1"

rows, err := db.QueryContext(ctx, sql, "2020-01-01") // *sqlx.Rows

for rows.Next() {
  results := make(map[string]interface{})
  err := rows.MapScan(results)
  if err != nil {
    log.Fatal(err)
  }
  // обрабатываем result
}
```
---

# sqlx: сканирование

Можно сканировать результаты структуру:
```
type Event {
  Id          int64
  Title       string
  Description string   `db:"descr"`
}

sql := "select * from events where start_date > $1"

rows, err := db.NamedQueryContext(ctx, sql, "2020-01-01") // *sqlx.Rows

events := make([]Event)

for rows.Next() {
  var event Event
  err := rows.StructScan(&event)
  if err != nil {
    log.Fatal(err)
  }
  events = append(events, event)
}
```
---

# Драйверы для Postgres

* Лучший драйвер на текущий момент: [https://github.com/jackc/pgx](https://github.com/jackc/pgx)
* Другой часто используемый драйвер (менее производительный): [https://github.com/lib/pq](https://github.com/lib/pq)

---

# Миграции

* [https://github.com/pressly/goose](https://github.com/pressly/goose) - можно использовать как cli-тулзу и как библиотеку
* [https://flywaydb.org/](https://flywaydb.org/) - пожалуй, самая популярная штука для миграций

*Protip*: flyway можно запускать из докер-контейнера перед запуском основного приложения, см. https://github.com/flyway/flyway-docker

---

# ORM

* [https://gorm.io/](https://gorm.io/) - использует пустые интерфейсы :(
* [https://github.com/go-reform/reform](https://github.com/go-reform/reform) - использует кодогенерацию, но разработка немного заброшена

---

# Другие ресурсы для изучения

.big-list[
* [ru] [https://habr.com/ru/company/oleg-bunin/blog/461935/](Статья и видео о тонкостях работы с Postgres в Go)
* [en] [http://go-database-sql.org/index.html](Полезная документация по работе с SQL из Go)
* [en] [https://golang.org/pkg/database/sql](Документация к database/sql)
* [en] [https://jmoiron.github.io/sqlx](sqlx)
  ]

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
