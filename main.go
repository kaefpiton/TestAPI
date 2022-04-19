package main

import (
	model "Server/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)


//Возвращаем сервер с заданными параметрами
//TODO почитать поподробнее про параметры
func createServer() http.Server {

	// создадим handler для сервера
	handler := http.NewServeMux()

	handler.HandleFunc("/hello", Logger(HelloHandler) )
	handler.HandleFunc("/greater/", GreaterHandler)

	handler.HandleFunc("/books/", BooksHandler)


	return http.Server{
		Addr: "127.0.0.1:8080",
		Handler: handler,
		ReadTimeout: 10 * time.Second,//время на обработку запроса клиента
		WriteTimeout: 10 * time.Second,//время на формирование ответа сервера
		ReadHeaderTimeout: 10 * time.Second,//ToDo уточнить
		MaxHeaderBytes: 1 << 20, // максимальный размер заголовка (128 кбайт)
	}
}


func main()  {

	//Объявление книг
	Firstbook := model.Book{Id: 1, Name: "first", Author: "author 1"}
	Secondbook := model.Book{Id: 2, Name: "second", Author: "author 2"}
	Thirdbook := model.Book{Id: 3, Name: "third", Author: "author 3"}

	//Объявление полки
	CurShelf := model.BookShelf{}

	CurShelf.AddBookOnShelf(Firstbook)
	CurShelf.AddBookOnShelf(Secondbook)
	CurShelf.AddBookOnShelf(Thirdbook)

	CurBook := CurShelf.GetBook(1)

	fmt.Printf("текущая книга: \nНазвание => %v\nАвтор => %v ", CurBook.Name, CurBook.Author)

	server:= createServer()
	log.Fatal(server.ListenAndServe())
}






type Response struct {
	Message string
	Error string
}

type FashionResponse struct {
	Message string `json:"Message"`
	Error string `json:"Error,omitempty"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request)  {
	//формируем json для ответа
	response := FashionResponse{
		Message: "Hello",
	}
	responsejson,_ := json.Marshal(response)

	//указываем в заголовке, что будем отправлять json
	w.Header().Set("Content-Type", "application/json")
	//отправляем успешный статус
	w.WriteHeader(http.StatusOK)
	//отправляем запрос
	w.Write(responsejson)
}
//функция-обработчик приветствия
func GreaterHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	name := strings.Replace(r.URL.Path,"/greater/", "",1)

	response := FashionResponse{
		Message: "Hello, " + name,
	}
	responseJson,_ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
//Middleware (простой логгер для handler)
//Todo посмотреть, что можно еще вывести
func Logger(next http.HandlerFunc)  http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: \nURI => %v \nMethod => %v \nHeader=> %v\n", r.URL,r.Method,r.Header)
		next.ServeHTTP(w,r)
	}
}




//РАБОТА С КНИГАМИ (CRUD)

//структура книг


//функция-обработчик для book
func BooksHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	switch  r.Method{
		case "GET":{

			response := FashionResponse{
				Message: "GET BOOK",
			}
			responseJson,_ := json.Marshal(response)

			w.WriteHeader(http.StatusOK)
			w.Write(responseJson)
		}

		case "POST":{

			AddBook()
			response := FashionResponse{
				Message: "ADD BOOK",
			}
			responseJson,_ := json.Marshal(response)

			w.WriteHeader(http.StatusOK)
			w.Write(responseJson)
		}

	}
}

func AddBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	rspstring := strings.Split(r.URL.Path, "/")

	for _,val := range  rspstring{
		fmt.Print(val)
	}
}

