package dependencies

import (
	"observer/pkg/hostcheck"
	"observer/pkg/kafka_adaptor"
	"sync"
)

type DepsInterface interface {
	SendToKafka(data hostcheck.SingleResult) (bool, string)
	GetRequestor() hostcheck.Requestor
}
type Dependencies struct {
	kafka *kafka_adaptor.KafkaAdaptor
}

func (dependencies Dependencies) SendToKafka(data hostcheck.SingleResult) (bool, string) {
	dependencies.kafka.Keep(data)
	return true, ""
}

func (dependencies Dependencies) GetRequestor() hostcheck.Requestor {
	return hostcheck.BaseRequestor{}
}

var (
	dependencies *Dependencies
	once         sync.Once
)

func GetDeps() *Dependencies {
	once.Do(func() {
		dependencies = &Dependencies{
			kafka: kafka_adaptor.MakeKafkaAdaptor(),
		}
	})
	return dependencies
}
