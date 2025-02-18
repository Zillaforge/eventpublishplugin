package cmd

import (
	"eventpublishplugin/cmd/args"
	_ "eventpublishplugin/configs" //loading config
	cnt "eventpublishplugin/constants"
	modCom "eventpublishplugin/module/common"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	tkErr "pegasus-cloud.com/aes/toolkits/errors"
	"pegasus-cloud.com/aes/toolkits/mviper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "EventPublishPlugin",
		Short: "ASUS Enterprise EventPublishPlugin",
		Long:  "Describe about ASUS EventPublishPlugin",
	}
	globalFolder = cnt.GlobalConfigPath + "/" + cnt.GlobalConfigFilename
)

//Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(NewServeCmd(), NewVersionCmd())
	rootCmd.PersistentFlags().StringVarP(&args.CfgFileG, "config", "c", globalFolder, "config file (default is /etc/ASUS/eventpublishplugin.yaml)")
	rootCmd.PersistentFlags().StringVar(&args.RedisChannel, "redis-channel", "", "redis channel")
	rootCmd.PersistentFlags().StringVar(&args.RabbitMQExchange, "rabbitmq-exchange", "", "rabbitmq exchange")
	rootCmd.PersistentFlags().StringVar(&args.RabbitMQRoutingKey, "rabbitmq-routing-key", "", "rabbitmq routing key")
}

func initConfig() {
	mviper.SetConfigType("yaml")
	if args.CfgFileG != "" {
		mviper.SetConfigFile(args.CfgFileG)
	}

	if args.RedisChannel != "" {
		modCom.RedisChannel = args.RedisChannel
	} else if args.RabbitMQExchange != "" && args.RabbitMQRoutingKey != "" {
		modCom.RabbitMQExchange = args.RabbitMQExchange
		modCom.RabbitMQRoutingKey = args.RabbitMQRoutingKey
	} else {
		err := tkErr.New(cnt.ServerPluginProviderInfoMustBeSetErr)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := mviper.MergeInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig() Failed : %v\n", err)
		os.Exit(1)
	}
}
