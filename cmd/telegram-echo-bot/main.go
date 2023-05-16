package main

import (
	"fmt"
	"net/http"

	"github.com/makropus/telegram-echo-bot/internal/echo"
)

func main() {
	fmt.Println("Hello, world.")
}

func YandexCFHandler(rw http.ResponseWriter, req *http.Request) {

	echo.HandleWebHookRequest(rw, req)
}
