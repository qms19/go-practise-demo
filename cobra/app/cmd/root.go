package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var configFile string

var rootCommand = &cobra.Command{
	Use: "app" ,
	Short: "app" ,
	Long: "app a" ,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("欢迎使用cobra。。。")
	},
}

func Execute()  {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init()  {
	cobra.OnInitialize(initConfig)
	rootCommand.PersistentFlags().StringVarP(&configFile,"config","f","","default is $HOME/app.yaml")
	rootCommand.PersistentFlags().StringP("author","a","","author")
	rootCommand.PersistentFlags().Bool("viper",true,"viper")

	viper.BindPFlag("useViper", rootCommand.PersistentFlags().Lookup("viper"))
	viper.BindPFlag("author",rootCommand.PersistentFlags().Lookup("author"))

	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")

}

type Config struct {
	App string
	Redis RedisConfig `mapstructure:"redis"`
}

type RedisConfig struct {
	Host string
	Port int
}

func initConfig()  {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName("app")

	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("config ",viper.ConfigFileUsed())
	}

	//fmt.Println(viper.Get("app"))
	//fmt.Println(viper.Get("useViper"))
	//fmt.Println(viper.Get("author"))
	//fmt.Println(viper.Get("redis.host"))

	var conf Config
	viper.Unmarshal(&conf)

	fmt.Println("tt---",conf.Redis)

}


