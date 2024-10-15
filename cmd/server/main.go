package main

import (
	"eToDo/routes"
	"log"
	"net/http"
)

func main() {
	/*
		Функция для запуска сервер
	*/
	port := ":8080"                              //Порт для подключения по localhost
	routerHome := routes.HomeRouterHandler       //Homepage
	routerProfile := routes.ProfileRouterHandler //Profile

	http.HandleFunc("/", routerHome)
	http.HandleFunc("/me", routerProfile)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
