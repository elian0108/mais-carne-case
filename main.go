package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	apiURL  = "https://api.maiscarne.com/pagamentos/health"
	fileLog = "health_check.log"
)

func main() {

	file, err := os.OpenFile(fileLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	logger := log.New(file, "API_LOG: ", log.Ldate|log.Ltime|log.Lmicroseconds)
	inicio := time.Now()
	response, err := http.Get(apiURL)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	tempo := time.Since(inicio).Milliseconds()
	status := response.StatusCode

	if status != 200 || tempo > 500 {
		logger.Printf("URL=%s STATUS=%d TEMPO=%d\n", apiURL, status, tempo)
	}
}
