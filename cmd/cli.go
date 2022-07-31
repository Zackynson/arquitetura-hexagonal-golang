/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/Zackynson/go-hexagonal/adapters/cli"
	"github.com/spf13/cobra"
)

var action string
var productId string
var productName string
var productPrice float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := cli.Run(productService, action, productId, productName, productPrice)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	// save args in vars
	cliCmd.Flags().StringVarP(&action, "action", "a", "", "Action to execute")
	cliCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")
	cliCmd.Flags().StringVarP(&productName, "name", "n", "", "Product Name")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "Product Price")
}
