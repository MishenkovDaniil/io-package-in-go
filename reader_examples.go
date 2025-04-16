package gopres

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func ReadFromString() int {
	reader := strings.NewReader("Hello, io package!")
	data := make([]byte, 5)
	n, _ := reader.Read(data) // n = 5, data = "Hello"
	return n
}

func ReadFromFile() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close() // Закрываем файл при завершении

	// Читаем данные по частям (буфер 256 байт)
	buffer := make([]byte, 256)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("\nКонец файла.")
				break
			}
			panic(err)
		}
		fmt.Printf("Прочитано %d байт: %q\n", n, buffer[:n])
	}
}

func ReadFromTCP() {
	// Подключаемся к серверу
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close() // Закрываем соединение

	// Отправляем HTTP-запрос
	_, err = conn.Write([]byte("GET / HTTP/1.1\r\nHost: example.com\r\n\r\n"))
	if err != nil {
		panic(err)
	}

	// Читаем ответ от сервера
	response, err := io.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Ответ сервера (%d байт):\n%s\n", len(response), response)
}
