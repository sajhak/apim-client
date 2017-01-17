/*
 *
 */

package cmd

import (
	"fmt"
	"github.com/renstrom/dedent"
	"github.com/spf13/cobra"
)

var Verbose bool

var (
	rootLongDesc = dedent.Dedent(`
	WUM keeps WSO2 products up-to-date.
	`)
)

var RootCmd = &cobra.Command{
	Use:   "apim",
	Short: "apim client tool",
	Long:  rootLongDesc,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	//	RootCmd.AddCommand(InitCmd)
	//	RootCmd.AddCommand(versionCmd)
	//	if err := RootCmd.Execute(); err != nil {
	//		os.Exit(-1)
	//	}
}

func init() {
	fmt.Println("......Inside init...")
	//	cobra.OnInitialize(InitConfig)
	//	cobra.EnableCommandSorting = false
	//	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "enable verbose mode")
}
