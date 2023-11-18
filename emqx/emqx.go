package emqx

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	emqxHost = "122.51.59.109"
	emqxPort = 31101
	userName = "admin"
	passWord = "admin111"
)

func GetEmqxConfig(clientID string) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", emqxHost, emqxPort))
	opts.SetClientID(clientID)
	opts.SetUsername(userName)
	opts.SetPassword(passWord)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return nil, token.Error()
	}

	return client, nil
}
