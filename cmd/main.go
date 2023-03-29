package main

import (
	"fmt"
	"golang-cities-information-service/pkg/file"
	"golang-cities-information-service/pkg/handler"
	"golang-cities-information-service/pkg/service"
	"golang-cities-information-service/pkg/storage"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	Repository := storage.NewMemStorage()
	Serv := service.NewCityService(Repository)
	Hand := handler.NewHandler(Serv)

	file.OpenFileCsv(Repository)
	fmt.Println("Открываем файл")
	go func() {

		log.Fatal(http.ListenAndServe(":8080", Hand.InitRoutes()))
	}()
	fmt.Println("Сервер запустился")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	file.CloseFileCsv(Repository)
	fmt.Println("Завершаем работу")
}
