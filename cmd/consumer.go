package cmd

import (
	"fmt"

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

type Consumer struct{}

func NewConsumer() *Consumer {
	return &Consumer{}
}

func (C *Consumer) Execute() {
	fmt.Println("AQUI CONSUMER")
}
