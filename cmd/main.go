package main

import (
	"awesomeProject/pkg/file"
	"awesomeProject/pkg/handler"
	"awesomeProject/pkg/service"
	"awesomeProject/pkg/storage"
	"fmt"
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
