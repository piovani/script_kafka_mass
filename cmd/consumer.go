package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"github.com/piovani/script_kafka_mass/infra/config"
	"github.com/spf13/cobra"
)

var (
	ConsumerCMD = &cobra.Command{
		Use:     "consumer",
		Short:   "consumer very messages in topic",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			NewConsumer().Execute()
		},
	}
)

type Consumer struct {
	ready chan bool
}

func NewConsumer() *Consumer {
	return &Consumer{
		ready: make(chan bool),
	}
}

func (c *Consumer) Execute() {
	configSarama := sarama.NewConfig()
	configSarama.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewConsumerGroup([]string{config.Env.Brokers}, "script-consumer", configSarama)
	c.checkErr(err)

	for {
		err = client.Consume(context.TODO(), []string{config.Env.Topic}, c)
		c.checkErr(err)

		c.ready = make(chan bool)
	}

	<-c.ready
	err = client.Close()
	c.checkErr(err)
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(c.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message := <-claim.Messages():
			log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
			session.MarkMessage(message, "")

		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/Shopify/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}

func (c *Consumer) checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
