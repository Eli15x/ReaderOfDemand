package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	service "github.com/Eli15x/ReaderOfDemand/service"
	_ "github.com/segmentio/kafka-go"
)

const (
	insertOffsets = `insert into db.t1 (_offset, _partition, _key, _value) values (?, ?, ?, ?)`
)

func main() {

	topics := strings.Split(os.Getenv("TOPICS"), ",")

	for _, v := range topics {
		c := fmt.Sprintf("consumer-%s", v)
		service.GetInstanceServiceBD().Consumer(v, c)
	}

	time.Sleep(300 * time.Second)
	fmt.Println("Encerrando threads")

}
