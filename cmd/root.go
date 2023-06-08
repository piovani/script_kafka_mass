package cmd

import (
	"log"

	"github.com/piovani/script_kafka_mass/infra/config"
	"github.com/spf13/cobra"
)

func Execute() {
	cmd := &cobra.Command{
		Use:     "scritp_kafka_mass",
		Version: "1.0.0",
	}

	cmd.AddCommand(
		ProducerCMD,
		ConsumerCMD,
	)

	CheckFatal(config.InitConfig())
	CheckFatal(cmd.Execute())
}

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
