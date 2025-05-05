package service

import (
	. "github.com/tattoo1880/testgoinit/config"
)
func Suckit(queueName string) {
	message ,err := MyRabbitMQ.Consume(queueName)
	if err != nil {
		panic(err)
	}
	for msg := range message {
		// 处理消息
		// 这里可以添加你的业务逻辑
		println("Received message:", string(msg.Body))
	}

}