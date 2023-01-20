/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	postrequest "test-next-header/postrequest"
	randomstring "test-next-header/randomstring"

	"github.com/spf13/cobra"
)

var stringSize string
var request bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "test-next-header",
	Long: `
		Generate a string of a given size in bytes
		Send a post request to the next-app with the generated string in the cookie header

		Usage: tnh --stringSize <int>

		Example: tnh --stringSize 17000
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	cookieHeader := randomstring.GenerateRandomString(stringSize)

	if request == true {
		response := postrequest.DoPostRequest(cookieHeader)
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(response.Status)
		fmt.Println(string(body))
	} else {
		fmt.Println(cookieHeader)
	}

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.test-next-header.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.Flags().BoolVarP(&request, "request", "r", false, "Send post request to nextjs-app")
	rootCmd.Flags().StringVarP(&stringSize, "stringSize", "s", "", "Desired size of string (bytes)")
}
