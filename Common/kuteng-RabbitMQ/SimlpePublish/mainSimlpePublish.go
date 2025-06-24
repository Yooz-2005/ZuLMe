package SimlpePublish

import (
	"Common/kuteng-RabbitMQ/RabbitMQ"
	"encoding/json"
	"fmt"
)

func SimplePublish(data interface{}) error {
	mq := RabbitMQ.NewRabbitMQSimple("" +
		"Order")

	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	mq.PublishSimple(string(marshal))
	fmt.Println("发送成功！")
	return nil
}
