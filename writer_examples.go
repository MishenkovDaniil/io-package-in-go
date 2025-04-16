package gopres

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"time"
)

func WriteToBuf() {
	var buf bytes.Buffer
	buf.Write([]byte("Hello"))       // запись байтов
	fmt.Fprintf(&buf, ", %s!", "Go") // форматированный вывод
}

func WriteToFile() {
	// Создаем файл (если существует, перезаписываем)
	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close() // Закрываем файл в конце

	// Записываем строку
	bytesWritten, err := file.WriteString("Hello, File!\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Записано %d байт.\n", bytesWritten)

	// Записываем байты
	data := []byte{72, 101, 108, 108, 111, 33} // "Hello!" в ASCII
	bytesWritten, err = file.Write(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Записано %d байт.\n", bytesWritten)

	// Используем fmt.Fprintf для форматированной записи
	fmt.Fprintf(file, "Сегодня: %s\n", "2023-10-05")
}

func WriteToTCP() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go func(c net.Conn) {
			defer c.Close()
			c.Write([]byte("Время сервера: " + time.Now().String() + "\n"))
		}(conn)
	}
}
