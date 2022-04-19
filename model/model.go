package model

import (
	"errors"
)

type Book struct {
	Id int
	Name string
	Author string
}

type BookShelf struct {
	Books []Book
}




func (b *BookShelf)AddBookOnShelf(book Book) error{
	if b.GetBook(book.Id) == nil{
		b.Books = append(b.Books, book)
		return nil
	}else {
		return errors.New("This book is already on the shelf")
	}
}

func (b *BookShelf)GetBook(id int) *Book {
	for _,element := range b.Books{
		if id == element.Id{
			return &element
		}
	}
	return nil
}

func (b *BookShelf)FindBookByName(name string) *Book{
	for _,element := range b.Books{
		if name == element.Name{
			return &element
		}
	}
	return nil
}

func (b *BookShelf)DelBook(id int) error {
	for i,element := range b.Books{
		if id == element.Id{
			b.Books[i] = b.Books[len(b.Books)-1]
			b.Books[len(b.Books)-1] = Book{}
			b.Books = b.Books[:len(b.Books)-1]
			return nil
		}
	}
	return errors.New("The book is missing from the shelf")
}

