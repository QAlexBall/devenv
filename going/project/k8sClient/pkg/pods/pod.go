package pods

import (
	"context"
	"fmt"
	"k8sClient/pkg"
	"os"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var PodCmd = &cobra.Command{
	Use:   "pod",
	Short: "Pod operations",
	Long:  `Pod operations: get, list, delete`,
}

var namespace string

var list = &cobra.Command{
	Use:   "list",
	Short: "List pods",
	Run: func(cmd *cobra.Command, args []string) {
		clientset, err := pkg.GetKubernetesClient()
		if err != nil {
			fmt.Println("Error getting kubernetes client: ", err)
			os.Exit(1)
		}

		pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			fmt.Println("Error listing pods: ", err)
			os.Exit(1)
		}

		for _, pod := range pods.Items {
			fmt.Println(pod.Name)
		}
	},
}

var search = &cobra.Command{
	Use:   "search",
	Short: "Search pods",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Searching pods")
	},
}

func init() {
	PodCmd.AddCommand(list)
	list.Flags().StringVarP(&namespace, "namespace", "n", "default", "Namespace")
	PodCmd.AddCommand(search)
}
