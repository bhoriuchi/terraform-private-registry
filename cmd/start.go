package cmd

import (
	"github.com/bhoriuchi/terraform-private-registry/server"
	"github.com/spf13/cobra"
)

var cmdStart = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewServer(server.Options{
			TLSCert: tlsCertFile,
			TLSKey:  tlsKeyFile,
			Addr:    serverAddr,
		})
		s.Start()
	},
}

var (
	tlsCertFile string
	tlsKeyFile  string
	serverAddr  string
)

func initCmdStart() {
	cmdStart.Flags().StringVarP(&tlsCertFile, "cert", "c", `cert.pem`, "Certificate file")
	cmdStart.Flags().StringVarP(&tlsKeyFile, "key", "k", `key.pem`, "Key file")
	cmdStart.Flags().StringVarP(&serverAddr, "addr", "a", ":8443", "Server address. Must be http")

	rootCmd.AddCommand(cmdStart)
}
