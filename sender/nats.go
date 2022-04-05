package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "pusher", stan.NatsURL("nats://localhost:4222"))
	if err != nil && err != io.EOF {
		log.Fatalln(err)
	} else {
		log.Println("Подключились к Nats-streaming")
	}

	var filename string
	for {
		fmt.Println("Введите имя и расширение файла для отправки:")
		fmt.Scanf("%s\n", &filename)
		dataFromFile, _ := ioutil.ReadFile(filename)
		sc.Publish("foo1", dataFromFile)
		println("Данные отправлены!")
	}
}
