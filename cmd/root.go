/*
Copyright © 2023 Award Malisi
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobraScratch",
	Short: "A brief description of your application ROOT",
	Long:  `A longer description that spans multiple lines and likely `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobraScratch.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	commandCmd.Flags().String("localFlag", "St", "a local string flag")
	commandCmd.PersistentFlags().Bool("persistentFlag", false,
		"a persistent boolean flag")

}

//Need Local Flag and global flag info.
