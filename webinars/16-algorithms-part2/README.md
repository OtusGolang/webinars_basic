background-image: url(../img/title.svg)

---

background-image: url(../img/rec.svg)

---
background-image: url(../img/topic.svg)

.topic[Алгоритмы и структуры данных. Часть 2]
.tutor[Родион Ступников]
.tutor_desc[Software Engineer]

---

background-image: url(../img/rules.svg)

---
# Цели занятия

* изучить структуры данных такие как: множество, дерево (в т.ч. бинарное), граф, хеш таблица
* изучить эффективные алгоритмы поиска
* изучить случаи, когда приведенные выше структуры данных могут оказаться полезны

---
# Хеш таблица

В языке Go уже есть стандартная структура данных map, которая позволяет за константное 
время находить произвольный элемент по его ключу.

Map гарантирует, что все ключи внутри нее уникальны.



.center-image[
![](img/map.png)
]

---
# Хеш таблица

Значением в хеш-таблице могут быть любые типы

А вот ключами: любые не ссылочные типы

```go
map[int][]int //ok
map[bool]string // ok
map[string]map[string]string //ok, может хранить внутри сложные объекты
map[[10]int64]float64 //ok
map[[]int]float64 // а вот так нельзя, []int - ссылочный тип
map[SomeStruct]*string //ok, если внутри нет ссылочных типов
map[interface{}]string //ok,  в интерфейс можно передать все, что не содержит ссылочные типы
map[[10]int64]interface{} //ok 
map[*int]bool - // можно, можно из *int достать номер адреса
map[map[int]int]int // так нельзя, сама map - тоже ссылочный тип
map[interface{}]interface{} // ok

type s struct {A *int}
map[s]int{} // тоже ok
```

https://go.dev/play/p/VfiQDGlBQa2

---
# Хеш таблица

Map - это ссылочный тип данных (как слайс). 
Их всегда нужно инициализировать перед использованием!

Ниже показаны способы создания map:

* С выделением емкости под будущие значения
```go
m := make(map[string]int, 10)
```

* С заданными значениями:
```go
m := map[string]int{
	"виталий": 31,
	"артем": 20,
	"петр": 25,
}
```

Еще один пример:
https://go.dev/tour/moretypes/20

---
# Хеш таблица

Чтение по ключу из хеш-таблицы: сложность O(1)

```go
val := m["артем"]
```
Если значение не существует в map, будет возвращено Zero Value.

Мы также можем вторым аргументом вычитать булевый флаг isExist.
```go
val, isExist := m["артем"]
```
В основном это нужно в сценариях, когда мы хотим узнать, 
действительно ли там записан ключ с ZeroValue,  или такого ключа там нет?

---
# Хеш таблица

Запись по ключу в хеш таблицу: сложность O(1)
```go
m["артем"] = 10
```

---
# Хеш таблица

С помощью оператора delete можно удалять значения из хеш-таблицы. 
```go
delete(m, "artem")
```
Как думаете, какая сложность у операции в O-нотации?

Пример внесения изменений в map:
https://go.dev/tour/moretypes/22

---
# Хеш таблица

Перебор всех ключей и значений в map

```go
for key, value := range m {
	fmt.Println(key, value)
}
```
Обратите внимание, что ключи и значения будут перебраны в случайном порядке.

https://go.dev/play/p/_8Yj1LON1ne

---
# Хеш таблица

Как она работает внутри:

.center-image[
![](img/map_detail.png)
]

---
# Хеш таблица

Больше можно узнать тут:
https://habr.com/ru/articles/457728/

---
# Множество

Множество - это как правило набор уникальных значений. 
В Go множество как правило реализуется через хеш-таблицу.

.

.center-image[
![](img/set.png)
]

---

# Множество

Реализация

```go
type Set struct {
	data map[int]struct{}
}

// Создать новое множество
func NewSet() *Set {
	s := &Set{
		data: make(map[int]struct{}),
	}
	return s
}

// Добавить элемент в множество
func (s *Set) Append(item int) {
    s.data[item] = struct{}{}
}

```

---
# Множество

Множествами удобно моделировать объекты из реального мира.

Допустим, у нас есть студенты, которые могут относиться как к множеству, записаны на курс Go,
так и к множеству, которые записаны на курс по Алгоритмам.

---

# Множество

Операция объединения:

.center-image[
![](img/set_add.png)
]

Что будет со сложностью?
Сколько нужно памяти?

---

# Множество

Операция объединения:

```go
func (s *Set) Add(other *Set) *Set {
    result := NewSet()

    // Добавляем элементы из текущего множества
    for item := range s.data {
        result.Append(item)
    }

    // Добавляем элементы из другого множества
    for item := range other.data {
        result.Append(item)
    }

    return result
}
```

Что будет со сложностью?
Сколько нужно памяти?

---
# Множество

Операция пересечения:

.center-image[
![](img/set_intersect.png)
]

Что будет со сложностью?
Сколько нужно памяти?


---
# Множество

Операция пересечения:

```go
// Пересечение множеств
func (s *Set) Intersect(other *Set) *Set {
	result := NewSet()
	for item := range s.data {
		if other.Contains(item) {
			result.Append(item)
		}
	}
	return result
}

// Проверить, содержит ли множество элемент
func (s *Set) Contains(item int) bool {
    _, found := s.data[item]
    return found
}
```
Что будет со сложностью?
Сколько нужно памяти? 
Как можно оптимизировать?

---
# Множество

Потребление памяти при операции пересечения: O(min(n, m)) ~= O(N)

---
# Множество
Улучшенная операция пересечения

```go
func (s *Set) Intersect(other *Set) *Set {
	result := NewSet()
	var smaller, larger *Set

	// Определяем множество меньшего размера
	if len(s.data) < len(other.data) {
		smaller = s
		larger = other
	} else {
		smaller = other
		larger = s
	}

	for item := range smaller.data {
		if larger.Contains(item) {
			result.Append(item)
		}
	}
	return result
}
```

Теперь сложность: O(min(N,M)) ~= O(N)

Потребление памяти не изменилось: O(min(N,M))

---
# Множество

Операция вычитания:

.center-image[
![](img/set_sub.png)
]

---
# Множество

Операция вычитания:

```go

func (s *Set) Sub(other *Set) *Set {
    result := NewSet()
    for item := range s.data {
        if !other.Contains(item) {
            result.Add(item)
		}
    }
    return result
}
```

Что будет со сложностью?
А сколько доп-памяти потребуется?

---
# Множество

Полная версия кода:

https://go.dev/play/p/ihOzLkvHUbD

---
# Множество

Итоги:
* Операция добавления - O(1)
* Операция объединения - O(N+M) ~= O(N)
* Операция пересечения - O(min(N,M)), Память: O(min(N,M))
* Операция вычитания - O(M), Память: O(N-M)

---
# Бинарное Дерево

Бинарное дерево - это
иерархическая структура данных, в которой каждый узел
имеет не более двух потомков(детей).
Как правило, первый называется родительским узлом,
а дети называются левым и правым наследниками.
Дерево является частным случаем графа.

.center-image[
![](img/bin_tree.png)
]

---
# Бинарное Дерево

Описание структуры:

```go
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}
```

---
# Бинарное Дерево

Операция поиска:
```go
// Метод для поиска значения в бинарном дереве
func (bt *BinaryTree) Search(value int) bool {
	currentNode := bt.Root
	for currentNode != nil {
		if value == currentNode.Value {
			return true
		} else if value < currentNode.Value {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
	return false
}
```

---
# Бинарное Дерево

Операция вставки:
```go
func (bt *BinaryTree) Insert(value int) {
    newNode := &TreeNode{Value: value}
    if bt.Root == nil {
        bt.Root = newNode
        return
    }
    currentNode := bt.Root
    for {
        if value < currentNode.Value {
            if currentNode.Left == nil {
                currentNode.Left = newNode
                return
            }
            currentNode = currentNode.Left
        } else {
            if currentNode.Right == nil {
                currentNode.Right = newNode
                return
            }
            currentNode = currentNode.Right
        }
    }
}
```

---
# Бинарное Дерево

Получение отсортированного по возрастанию массива:

```go
// Метод для обхода дерева в порядке "in-order"
func (bt *BinaryTree) ToSortedSlice() []int {
	stack := []*TreeNode{}
	currentNode := bt.Root
	var restult []int

	for currentNode != nil || len(stack) > 0 {
		for currentNode != nil {
			stack = append(stack, currentNode)
			currentNode = currentNode.Left
		}
		currentNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		restult = append(restult, currentNode.Value)
		currentNode = currentNode.Right
	}
	return restult
}
```

---
# Бинарное Дерево

Полная версия кода:
https://go.dev/play/p/yG2f0O0CY0F

---
# Несбалансированное бинарное дерево

Проблема несбалансированного дерева:

.center-image[
![](img/bin_tree_unbalanced.png)
]

---
# Несбалансированное бинарное дерево

Поиск в таком дереве по эффективности не будет отличаться от поиска в связном списке.

То есть по сложности он составит O(N).
---
# Несбалансированное бинарное дерево

Есть специальные виды деревьев, которые умеют поддерживать сбалансированные состояния:
* Красно-черное дерево: https://ru.wikipedia.org/wiki/%D0%9A%D1%80%D0%B0%D1%81%D0%BD%D0%BE-%D1%87%D1%91%D1%80%D0%BD%D0%BE%D0%B5_%D0%B4%D0%B5%D1%80%D0%B5%D0%B2%D0%BE
* AVL-дерево: https://ru.wikipedia.org/wiki/%D0%90%D0%92%D0%9B-%D0%B4%D0%B5%D1%80%D0%B5%D0%B2%D0%BE

---
# Дерево
Итого:
* Вставка элемента при сбалансированности: O(log N) 
* Поиск элемента при сбалансированности: O(log N)
* Получение отсортированного массива: O(N)
* Так же присутствует возможность искать по диапазонам значений за O(log N) при сбалансированности

---
# Бинарный поиск

Этот алгоритм позволяет найти элемент в отсортированном массиве за O(log N). 
Будет в домашнем задании.
.center-image[
![](img/sorted_arr.png)
]

---

# Бинарный поиск

.center-image[
![](img/bin_search.png)
]

---

# Граф
Граф — математическая абстракция реальной системы любой природы,
объекты которой обладают парными связями. 

Граф как математический объект есть совокупность двух множеств — 
множества самих объектов, называемого множеством вершин, 
и множества их парных связей, называемого множеством рёбер.

Элемент множества рёбер есть пара элементов множества вершин.

.center-image[
![](img/graph.png)
]

---
# Граф

Графы в реальной жизни:
* Карту дорог можно представить в виде графа
* Социальные связи людей в соцсетях тоже можно представить в виде графов

---
# Граф

Графы в Go можно представить следующим образом:

```go
type Node struct {
	Val       interface{}
	Vertices  []int
}

type Graph struct {
	Nodes map[int]Node
}

// Функция для добавления вершины в граф
func (g *Graph) AddNode(nodeID int, val interface{}) {
	newNode := Node{Val: val}
	g.Nodes[nodeID] = newNode
}
```

---
# Граф

Добавление ребра(связи) между вершинами:

```go
// Функция для добавления ребра между вершинами
func (g *Graph) AddEdge(nodeID1, nodeID2 int) {
	node1, exists1 := g.Nodes[nodeID1]
	node2, exists2 := g.Nodes[nodeID2]

	if !exists1 || !exists2 {
		fmt.Println("Одна из вершин не существует")
		return
	}

	node1.Vertices = append(node1.Vertices, nodeID2)
	node2.Vertices = append(node2.Vertices, nodeID1)

	g.Nodes[nodeID1] = node1
	g.Nodes[nodeID2] = node2
}
```

---
# Граф
Алгоритм поиска в ширину систематически обходит все ребра графа
для «открытия» всех вершин, достижимых из
стартовой ноды, вычисляя при этом расстояние (минимальное количество рёбер) от
стартовой ноды до каждой достижимой вершины.

.center-image[
![](img/graph_bfs.png)
]

---
# Граф
* Подробнее об алгоритме поиска в ширину можно почитать здесь:
https://ru.wikipedia.org/wiki/%D0%9F%D0%BE%D0%B8%D1%81%D0%BA_%D0%B2_%D1%88%D0%B8%D1%80%D0%B8%D0%BD%D1%83

---
# Граф

* Код простого графа здесь с алгоритмом поиска в ширину:
https://go.dev/play/p/rMQPZovZUvo

---
# Граф

Сложность алгоритма BFS: `O(V+E)`, 

где V - вершины, а E - ребра

---
# Граф со взвешенными ребрами
У ребер в графе могут быть веса, что может означать, к примеру, стоимость пути.
Чтобы находить в таком графе наиболее дешевый путь, пригодится алгоритм Дейкстры. 
Этот алгоритм находит кратчайшие пути от одной из вершин графа до всех остальных.

.center-image[
![](img/graph_deykstra.png)
]

---
# Граф со взвешенными ребрами

Подробнее о нем написано тут:
https://ru.wikipedia.org/wiki/%D0%90%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC_%D0%94%D0%B5%D0%B9%D0%BA%D1%81%D1%82%D1%80%D1%8B

---
# Граф со взвешенными ребрами

Реализация алгоритма Дейкстры:
https://go.dev/play/p/UGbFjjvrI1a

Как думаете, что со сложностью этого алгоритма?

---
# Граф со взвешенными ребрами

Алгоритм Дейкстры:
* Сложность обычного графа с использованием приоритетных очередей: `O((V + E) * log V)`
* Сложность обхода плотно связанного графа приближается к `O(V^2 * log V)`
* Сложность обхода разреженного графа с использованием приоритетных очередей: `O(V * max(V->E)) ~= O(V)`, где `V->E` - количество ребер у вершины
* Обсчитывает сложность пути к абсолютно всем точкам, что в некоторых задачах может быть избыточно.
* Для поиска эффективного пути к конкретной точке лучше использовать A* (А-star) - алгоритм.

---
# Опциональное задание

При желании, можете попробовать разобраться в алгоритме A*, 
и попробовать его реализовать для взвешенного графа.

---


# Граф

Итого:
* Используется для моделирования многих объектов реального мира
* Состоит из вершин (V) и ребер (E).
* Может быть ориентированным, так и не ориентированным.
* Существует больше множество разных алгоритмов обхода графа и поиска пути.

---
# Итоги

Разобрали сегодня:
* hash-map
* множества (set)
* бинарные деревья
* бинарный поиск в отсортированном массиве
* графы и алгоритмы обхода их


---

background-image: url(../img/questions.svg)

---

background-image: url(../img/poll.svg)

---

background-image: url(../img/next_webinar.svg)
.announce_date[16 октября]
.announce_topic[Рекурсия и динамическое программирование]

---
background-image: url(../img/thanks.svg)

.tutor[Родион Ступников]
.tutor_desc[Software Engineer]
