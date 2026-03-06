package cmd

import (
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "error-ball",
	Short: "A CLI fortune cookie predicting the next error you might receive",
	Long:  "A CLI fortune cookie predicting the next error you might receive. Inspired by the 8-ball fortune teller.",
	Run: func(cmd *cobra.Command, args []string) {
		displayFortune()
	},
}

// fortunes contains an array of 40 popular and famous software errors
var fortunes = []string{
	"Error 404: Not Found",
	"Segmentation fault (core dumped)",
	"Blue Screen of Death (BSoD)",
	"System error 67 has occurred.",
	"General Protection Fault",
	"You don't have permission to access this resource.",
	"This program has performed an illegal operation and will be shut down.",
	"Abort, Retry, Fail?",
	"Kernel panic - not syncing: Fatal exception",
	"The server is down for maintenance",
	"Your device ran into a problem and needs to restart.",
	"Cannot open file: Permission denied",
	"OutOfMemoryError: Java heap space",
	"Device not ready",
	"No space left on device",
	"Error establishing a database connection",
	"Connection timed out",
	"The database is read-only",
	"Mount: block device is write-protected; mounting read-only",
	"Error: Failed to fetch",
	"Access denied",
	"File not found",
	"Network is unreachable",
	"SSL certificate has expired",
	"Port already in use",
	"Stack overflow",
	"Null pointer exception",
	"Division by zero",
	"Bad Request - Error 400",
	"Internal Server Error - Error 500",
	"Service Unavailable - Error 503",
	"Forbidden - Error 403",
	"DNS lookup failed",
	"Could not resolve hostname",
	"Memory access violation",
	"Buffer overflow detected",
	"Certificate verification failed",
	"Authentication failed",
	"Disk full",
	"Page fault in non-paged area",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// displayFortune selects and displays a random fortune
func displayFortune() {
	// Select a random fortune (rand/v2 is automatically seeded)
	randomIndex := rand.IntN(len(fortunes))
	selectedFortune := fortunes[randomIndex]

	// Display the fortune with some formatting
	fmt.Println("🥠 Your Error Fortune:")
	fmt.Printf("   %s\n", selectedFortune)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.first-action.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
