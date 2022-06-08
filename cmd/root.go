package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	ccmd "ntsc.ac.cn/ta-registry/pkg/cmd"
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
	cobra.OnInitialize(func() {})
	viper.AutomaticEnv()
	viper.SetEnvPrefix("TA")
	rootCmd.Flags().StringVar(&ccmd.GlobalEnvs.LoggerLevel,
		"logger-level", "DEBUG", "logger level")
	rootCmd.Flags().StringVar(&envs.listener,
		"rpc-listener", "tcp://0.0.0.0:1358", "grpc listener url")
	rootCmd.Flags().StringVar(&envs.certPath,
		"cert-path", "/etc/ntsc/ta/registry/certs", "system certificates path")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func prerun(cmd *cobra.Command, args []string) {
	ccmd.InitGlobalVars()
	var err error
	if err = ccmd.ValidateStringVar(&envs.certPath, "cert_path", true); err != nil {
		logrus.WithField("prefix", "cmd.root").
			Fatalf("check boot var failed: %s", err.Error())
	}
	if err = ccmd.ValidateStringVar(&envs.listener, "rpc_listener", true); err != nil {
		logrus.WithField("prefix", "cmd.root").
			Fatalf("check boot var failed: %s", err.Error())
	}
}

func run(cmd *cobra.Command, args []string) {

}
