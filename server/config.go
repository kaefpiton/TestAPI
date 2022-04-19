package server

import (
	"net/http"
	"time"
)

func ServerConfig() http.Server {

	return http.Server{
		Addr: "127.0.0.1:8080",
		ReadTimeout: 10 * time.Second,//время на обработку запроса клиента
		WriteTimeout: 10 * time.Second,//время на формирование ответа сервера
		ReadHeaderTimeout: 10 * time.Second,//ToDo уточнить
		MaxHeaderBytes: 1 << 20, // максимальный размер заголовка (128 кбайт)
	}

}
