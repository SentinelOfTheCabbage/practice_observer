package kafka_adaptor

import (
	"context"
	"sync"
)

type kafkaAdaptor interface {
	Keep(data any) bool
}
type KafkaAdaptor struct {
}

var (
	kafka_communicator *KafkaAdaptor
	once               sync.Once
)

func MakeKafkaAdaptor() *KafkaAdaptor {
	once.Do(func() {
		kafka_communicator = &KafkaAdaptor{}
	})
	return kafka_communicator
}

func (k *KafkaAdaptor) Keep(data any) bool {
	context.TODO()
	return false
}
