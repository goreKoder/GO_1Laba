package main //go run LAB_1/lab_1.go

import (
	"fmt"
	"time"
)

func main() {
	// Получаем текущее время
	currentTime := time.Now()

	// Форматируем время и дату
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	// Выводим результат
	fmt.Println("Текущая дата и время:", formattedTime)
}
