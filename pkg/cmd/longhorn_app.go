package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func makeInstallLongHorn() *cobra.Command {
	var longHorn = &cobra.Command{
		Use:          "longhorn",
		Short:        "Install Lognhorn",
		Long:         "Install Longhorn for providing a distributed block storage system for Kubernetes",
		Example:      "k3sup app install longhorn",
		SilenceUsage: true,
	}

	longHorn.Flags().StringP("namespace", "n", "longhorn-system", "The namespace to install longhorn to")

	longHorn.RunE = func(command *cobra.Command, args []string) error {
		kubeConfigPath := getDefaultKubeconfig()

		if command.Flags().Changed("kubeconfig") {
			kubeConfigPath, _ = command.Flags().GetString("kubeconfig")
		}

		fmt.Printf("Using kubeconfig: %s\n", kubeConfigPath)

		namespace, _ := command.Flags().GetString("namespace")

		if namespace != "longhorn-system" {
			return fmt.Errorf("To override the longhorn-system namespace, install longhorn via helm manually")
		}

		// Add longhorn with kubectl apply -f https://raw.githubusercontent.com/longhorn/longhorn/master/deploy/longhorn.yaml
		res, err := kubectlTask("apply", "-f", "https://raw.githubusercontent.com/longhorn/longhorn/master/deploy/longhorn.yaml")
		if err != nil {
			return err
		}

		if len(res.Stderr) > 0 {
			return fmt.Errorf("Error installing longhorn: %s", res.Stderr)
		}

		fmt.Println(`=======================================================================
= longhorn has been installed.                                        =
=======================================================================

# Check pod status with
kubectl -n longhorn-system get pods

# Get started with Longhorn here
# https://github.com/longhorn/longhorn

Thank you for using k3sup!`)

		return nil
	}

	return longHorn
}
