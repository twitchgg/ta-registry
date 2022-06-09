package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	registry "ntsc.ac.cn/ta-registry/internal/server"
	ccmd "ntsc.ac.cn/ta-registry/pkg/cmd"
)

var envs struct {
	listener        string
	certPath        string
	localDBPath     string
	managerEndpoint string
	etcdEndpoints   string
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
	rootCmd.Flags().StringVar(&envs.localDBPath,
		"local-db-path", "", "local database path")
	rootCmd.Flags().StringVar(&envs.managerEndpoint,
		"mgr-endpoint", "", "remote api url")
	rootCmd.Flags().StringVar(&envs.etcdEndpoints,
		"etcd-endpoints", "localhost:2379", "etcd endpoint urls")
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
	ccmd.ValidateStringVar(&envs.localDBPath, "local_db_path", false)
	ccmd.ValidateStringVar(&envs.managerEndpoint, "mgr_endpoint", false)
	if err = ccmd.ValidateStringVar(&envs.etcdEndpoints, "etcd_endpoints", true); err != nil {
		logrus.WithField("prefix", "cmd.root").
			Fatalf("check boot var failed: %s", err.Error())
	}
	go func() {
		ccmd.RunWithSysSignal(nil)
	}()
}

func run(cmd *cobra.Command, args []string) {
	serv, err := registry.NewRegistryServer(&registry.RegistryServerConfig{
		Listener:        envs.listener,
		CertPath:        envs.certPath,
		EtcdEndpoints:   envs.etcdEndpoints,
		LocalDBPath:     envs.localDBPath,
		ManagerEndpoint: envs.managerEndpoint,
	})
	if err != nil {
		logrus.WithField("prefix", "cmd.root").
			Fatalf("create registry server failed: %s", err.Error())
	}
	logrus.WithField("prefix", "cmd.root").
		Fatalf("run registry server failed: %s", <-serv.Start())
}
