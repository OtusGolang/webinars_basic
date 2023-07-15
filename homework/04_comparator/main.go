package main

import (
	"fmt"
)

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

func (b *Book) GetID() int {
	return b.id
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetYear() int {
	return b.year
}

func (b *Book) GetSize() int {
	return b.size
}

func (b *Book) GetRate() float64 {
	return b.rate
}

type CompareMode int

const (
	CompareByYear CompareMode = iota
	CompareBySize
	CompareByRate
)

type BookComparator struct {
	mode CompareMode
}

func (bc BookComparator) Compare(b1, b2 Book) bool {
	switch bc.mode {
	case CompareByYear:
		return b1.year == b2.year
	case CompareBySize:
		return b1.size == b2.size
	case CompareByRate:
		return b1.rate == b2.rate
	default:
		return false
	}
}

func NewComparator(mode CompareMode) BookComparator {
	return BookComparator{mode: mode}
}

func main() {
	book1 := Book{id: 1, title: "Book 1", author: "Author 1", year: 2020, size: 200, rate: 4.5}
	book2 := Book{id: 2, title: "Book 2", author: "Author 2", year: 2020, size: 250, rate: 4.8}

	// Пример использования методов для установки и получения полей структуры
	book1.SetTitle("New Title")
	fmt.Println(book1.GetTitle()) // Вывод: New Title

	// Пример использования структуры BookComparator
	yearComparator := NewComparator(CompareByYear)
	fmt.Println(yearComparator.Compare(book1, book2)) // Вывод: true

	sizeComparator := NewComparator(CompareBySize)
	fmt.Println(sizeComparator.Compare(book1, book2)) // Вывод: false

	rateComparator := NewComparator(CompareByRate)
	fmt.Println(rateComparator.Compare(book1, book2)) // Вывод: false
}
