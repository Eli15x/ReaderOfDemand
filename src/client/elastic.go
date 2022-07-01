package client

import (
	"sync"
	"github.com/labstack/gommon/log"
	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

var (
	instanceClient elasticClient
	onceService     sync.Once
)

type elasticClient interface {
	Connect() error
}

type client struct{}

func GetInstanceClient() elasticClient {
	onceService.Do(func() {
		instanceClient = &client{}
	})
	return instanceClient
}

func (c *client) Connect() error {

	cfg := elasticsearch.Config{
        CloudID: "My_deployment:dXMtY2VudHJhbDEuZ2NwLmNsb3VkLmVzLmlvOjQ0MyRiN2VhMWNiMzY5MmU0ZjlhYjA5MmI5Njg3ZTRiM2FjNiQ4NGZhNDFmZmUzNjA0ZjVkODcwMTQyYTAzMjc2ZGJhYg==",
        APIKey: "cjJwUXQ0RUJoVTN3V2d3ZVZlSGE6TkxZblJnSVJUbHU5czFrZmRVdW43UQ==",
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}

	log.Infof("[ErrorConsumer] error : %s \n", es, "")

	return nil
}
