package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ProducerCMD = &cobra.Command{
		Use:     "producer",
		Short:   "producer very messages in topic",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			NewProducer().Execute()
		},
	}
)

type Producer struct{}

func NewProducer() *Producer {
	return &Producer{}
}

func (p *Producer) Execute() {
	fmt.Println("AQUI PRODUCER")
}
