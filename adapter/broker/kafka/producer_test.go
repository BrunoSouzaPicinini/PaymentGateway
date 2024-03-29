package kafka

import (
	"github.com/BrunoSouzaPicinini/imersao5-gateway/adapter/presenter/transaction"
	"github.com/BrunoSouzaPicinini/imersao5-gateway/domain/entity"
	"github.com/BrunoSouzaPicinini/imersao5-gateway/usecase/process_transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProducerPublish(t *testing.T) {
	expectedOutput := process_transaction.TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you do not have limit for this transaction",
	}

	//outputJson, _ := json.Marshal(expectedOutuput)

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}
	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())
	err := producer.Publish(expectedOutput, []byte("1"), "test")
	assert.Nil(t, err)
}
