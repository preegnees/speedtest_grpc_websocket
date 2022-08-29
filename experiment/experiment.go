package main

import (
	"bufio"
	// "time"
	// "io"
	"log"
	"os"

	// "strings"
	"sync"
)

var PATH string = "C:\\Users\\secrr\\Desktop\\speedtest_grpc_websocket\\websocket\\files\\cli_file.txt"
var BUFFER []byte = make([]byte, 5)

// func main() {
// 	file, err := os.Open(PATH)
// 	if err != nil {
// 		log.Fatal("Ошикба при открытии, err:=", err)
// 	}
// 	defer file.Close()

// 	r := bufio.NewReader(file)

// 	n, err := r.Read(BUFFER)
// 	buf := BUFFER[:n]

// 	// _, _ := io.ReadFull(r, )

// 	log.Println("n: ", len(buf))
// 	log.Printf("buf: %s\n", buf)
// }
// можно узнать размер файла, узнать на сколько частей его делить,
// создать буфер с его размеро и указывать срезы, читая этот файл

func main() {
	f, err := os.Open(PATH)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, _ = f.Stat()
	// log.Println("stat:", s.Size(), "bytes")

	var wg sync.WaitGroup

	wg.Add(2)

	// go func(f *os.File) {
	// 	s, _ := f.Seek(1, 0)
	// 	time.Sleep(1 * time.Second)
	// 	log.Println("1, s:", s)
	// 	wg.Done()
	// }(f)
	// go func(f *os.File) {
	// 	s, _ := f.Seek(2, 0)
	// 	log.Println("2, s:", s)
	// 	wg.Done()
	// }(f)
	// wg.Wait()

	numb, _ := f.Seek(30, 0)
	log.Println(numb)
	r := bufio.NewReader(f)

	// log.Println("before n:", len(BUFFER))
	// log.Println("before buf:", string(BUFFER))
	
	n, _ := r.Read(BUFFER)
	buf := BUFFER[:n]
	
	log.Println("n:", len(buf))
	log.Printf("buf: %s", string(buf))

	// n, err := io.ReadFull(r, header[9:])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("n: ", n)
	// log.Println("header: ", string(header[:]))
	// _, err = reader.Discard(64) // discard the following 64 bytes

}
// должны создаваться несколько горутин, в них передается сдвиг в соответсвии с мапой, 
// там происходит установка сдвига и чтение, потом отправка

// еще нужно сделать пулл потоков, чтобы при чтении в память они не перегрузили систему

// еще можно сгенерировать несколько каналов, чтобы читать паралельно

// нужно учитывать 1024 при конфигурации!!!!!!!!!