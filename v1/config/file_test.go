package config_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/RichardKnop/machinery/v1/config"
	"github.com/stretchr/testify/assert"
)

var configYAMLData = `---
broker: broker
default_queue: default_queue
result_backend: result_backend
results_expire_in: 123456
amqp:
  binding_key: binding_key
  exchange: exchange
  exchange_type: exchange_type
  prefetch_count: 123
  queue_binding_args:
    image-type: png
    x-match: any
`

func TestReadFromFile(t *testing.T) {
	data, err := config.ReadFromFile("testconfig.yml")
	if err != nil {
		t.Fatal(err)
	}

	if string(data) == configYAMLData && err == nil {
		return
	}

	var buffer bytes.Buffer
	buffer.WriteString(
		fmt.Sprintf("Expected value:\n%v\n", configYAMLData))
	buffer.WriteString(
		fmt.Sprintf("Actual value:\n%v\n", string(data)))
	t.Error(buffer.String())
}

func TestNewFromYaml(t *testing.T) {
	config.Reset()

	cnf := config.NewFromYaml("testconfig.yml", true, false)

	assert.Equal(t, "broker", cnf.Broker)
	assert.Equal(t, "default_queue", cnf.DefaultQueue)
	assert.Equal(t, "result_backend", cnf.ResultBackend)
	assert.Equal(t, 123456, cnf.ResultsExpireIn)
	assert.Equal(t, "exchange", cnf.AMQP.Exchange)
	assert.Equal(t, "exchange_type", cnf.AMQP.ExchangeType)
	assert.Equal(t, "binding_key", cnf.AMQP.BindingKey)
	assert.Equal(t, "any", cnf.AMQP.QueueBindingArgs["x-match"])
	assert.Equal(t, "png", cnf.AMQP.QueueBindingArgs["image-type"])
	assert.Equal(t, 123, cnf.AMQP.PrefetchCount)
}
