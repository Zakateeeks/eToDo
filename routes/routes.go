package routes

import (
	"fmt"
	"net/http"
	"strings"
)

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	/*
		Функция разбирает параметры формы, если они присутствуют,
		и выводит их содержимое в консоль для отладки.
		Далее она выводит путь и схему запроса URL.
		И в завершение возвращает ответ "Homepage!" в HTTP-ответе.
	*/
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Fprintf(w, "Homepage!")
}

func ProfileRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Profile!")
}
