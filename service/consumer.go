package service

import (
	"sync"
)

var (
	instanceService serviceBD
	onceService     sync.Once
)

type serviceBD interface {
	Consumer(topic string, consumer string)
}

type service struct{}

func GetInstanceServiceBD() serviceBD {
	onceService.Do(func() {
		instanceService = &service{}
	})
	return instanceService
}

func (s *service) Consumer(topic string, consumer string) {

	/*fmt.Println("Consumindo topic : [", topic, "]")
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/db", os.Getenv("USER"), os.Getenv("PASS"), os.Getenv("ENDPOINT")))
	if err = db.Ping(); err != nil {
		msg := "error: " + err.Error()
		log.Fatal(msg)
	}
	defer db.Close()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: consumer,
		Topic:   topic,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		tx := db.MustBegin()
		tx.MustExec(insertOffsets, m.Offset, m.Partition, m.Key, m.Value)
		tx.Commit()
		fmt.Printf("Mensagem de [%s] transferida\n", topic)
	}

	r.Close()*/

}
