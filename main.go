package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	
	"github.com/labstack/gommon/log"
	"github.com/Eli15x/ReaderOfDemand/src/client"
	service "github.com/Eli15x/ReaderOfDemand/src/service"
	_ "github.com/segmentio/kafka-go"
)

const (
	insertOffsets = `insert into db.t1 (_offset, _partition, _key, _value) values (?, ?, ?, ?)`
)

func main() {

	topics := strings.Split(os.Getenv("TOPICS"), ",")
	err := client.GetInstanceClient().Connect()

	if err != nil {
		log.Infof("[ErrorClientConnection] error : %s \n", err, "")
		return
	}

	

	for _, v := range topics {
		c := fmt.Sprintf("consumer-%s", v)
		err := service.GetInstanceServiceBD().Consumer(v, c)
		if err != nil {
			log.Infof("[ErrorConsumer] error : %s \n", err, "")
		}
	}

	time.Sleep(300 * time.Second)
	fmt.Println("Encerrando threads")

}
