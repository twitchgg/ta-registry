package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var envs struct {
	listener string
	certPath string
}

var rootCmd = &cobra.Command{
	Use:    "ta-registry",
	Short:  "TA registry",
	PreRun: prerun,
	Run:    run,
}

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	formatter := new(prefixed.TextFormatter)
	logrus.SetFormatter(formatter)
	cobra.OnInitialize(func() {})
	viper.AutomaticEnv()
	viper.SetEnvPrefix("TA")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func prerun(cmd *cobra.Command, args []string) {

}

func run(cmd *cobra.Command, args []string) {

}
