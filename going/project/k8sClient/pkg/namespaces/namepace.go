package namespaces

import (
	"context"
	"fmt"
	"k8sClient/pkg"
	"os"

	"github.com/olekukonko/tablewriter"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NamespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Namespace operations",
	Long:  `Namespace operations: get, list, delete`,
}

var list = &cobra.Command{
	Use:   "list",
	Short: "List namespaces",
	Run: func(cmd *cobra.Command, args []string) {
		clientset, err := pkg.GetKubernetesClient()
		if err != nil {
			fmt.Println("Error getting kubernetes client: ", err)
			os.Exit(1)
		}

		namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
		if err != nil {
			fmt.Println("Error listing namespaces: ", err)
			os.Exit(1)
		}
		table_print([]string{"Name"}, [][]string{
			[]string{"default"},
			[]string{"kube-system"},
		})
	},
}

func table_print(headers []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func init() {
	NamespaceCmd.AddCommand(list)
}
