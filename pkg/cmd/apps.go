package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func MakeApps() *cobra.Command {
	var command = &cobra.Command{
		Use:          "app",
		Short:        "Manage Kubernetes apps",
		Long:         `Manage Kubernetes apps`,
		Example:      `  k3sup app install openfaas`,
		SilenceUsage: false,
	}

	var install = &cobra.Command{
		Use:          "install",
		Short:        "Install a Kubernetes app",
		Long:         `Install a Kubernetes app`,
		Example:      `  k3sup app install [APP]`,
		SilenceUsage: true,
	}

	install.Flags().String("kubeconfig", "kubeconfig", "Local path for your kubeconfig file")

	install.RunE = func(command *cobra.Command, args []string) error {

		if len(args) == 0 {
			fmt.Printf("You can install: %s\n", strings.TrimRight(strings.Join(getApps(), ", "), ", "))
			return nil
		}

		return nil
	}

	command.AddCommand(install)
	install.AddCommand(makeInstallOpenFaaS())
	install.AddCommand(makeInstallMetricsServer())
	install.AddCommand(makeInstallInletsOperator())
	install.AddCommand(makeInstallCertManager())
	install.AddCommand(makeInstallLongHorn())
	install.AddCommand(makeInstallOpenFaaSIngress())
	install.AddCommand(makeInstallNginx())

	return command
}

func getApps() []string {
	return []string{"openfaas", "nginx-ingress", "cert-manager", "openfaas-ingress", "inlets-operator", "metrics-server", "longhorn"}
}
