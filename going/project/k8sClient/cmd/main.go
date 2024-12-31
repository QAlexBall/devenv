package main

import (
	"k8sClient/pkg/pods"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "k8sClient",
		Short: "A simple k8s client",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(pods.PodCmd)
	rootCmd.Execute()
}
