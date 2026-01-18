package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Version is set at build time
	Version = "dev"

	// Flags
	serviceName string
	environment string
	verbose     bool
	dryRun      bool
)

var rootCmd = &cobra.Command{
	Use:   "sreq",
	Short: "Service-aware API client with automatic credential resolution",
	Long: `sreq eliminates the overhead of manually fetching credentials from
multiple sources when testing APIs. Just specify the service name
and environment — sreq handles the rest.

Example:
  sreq GET /api/v1/users -s auth-service -e dev
  sreq POST /api/v1/users -s auth-service -e prod -d '{"name":"test"}'`,
	Version: Version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&serviceName, "service", "s", "", "Service name")
	rootCmd.PersistentFlags().StringVarP(&environment, "env", "e", "", "Environment (dev/staging/prod)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Show what would be sent without executing")

	// Add subcommands
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(serviceCmd)
	rootCmd.AddCommand(envCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(requestCmd)
}

// initCmd initializes the configuration
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize sreq configuration",
	Long:  `Creates the default configuration files in ~/.sreq/`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing sreq configuration...")
		// TODO: Implement config initialization
		fmt.Println("Created ~/.sreq/config.yaml")
		fmt.Println("Created ~/.sreq/services.yaml")
		fmt.Println("\nEdit these files to configure your providers and services.")
	},
}

// serviceCmd manages services
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Manage services",
	Long:  `Add, list, or remove service configurations.`,
}

var serviceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured services",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Configured services:")
		// TODO: List services from config
		fmt.Println("  (no services configured yet)")
	},
}

var serviceAddCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add a new service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Adding service: %s\n", name)
		// TODO: Implement service addition
	},
}

func init() {
	serviceCmd.AddCommand(serviceListCmd)
	serviceCmd.AddCommand(serviceAddCmd)
}

// envCmd manages environments
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Manage environments",
	Long:  `List environments or switch the default environment.`,
}

var envListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available environments",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available environments:")
		fmt.Println("  - dev")
		fmt.Println("  - staging")
		fmt.Println("  - prod")
		// TODO: Read from config
	},
}

var envSwitchCmd = &cobra.Command{
	Use:   "switch [env]",
	Short: "Switch default environment",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		env := args[0]
		fmt.Printf("Switched default environment to: %s\n", env)
		// TODO: Implement environment switching
	},
}

func init() {
	envCmd.AddCommand(envListCmd)
	envCmd.AddCommand(envSwitchCmd)
}

// configCmd shows configuration
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current configuration:")
		// TODO: Show config
		configPath := os.Getenv("SREQ_CONFIG")
		if configPath == "" {
			configPath = "~/.sreq/config.yaml"
		}
		fmt.Printf("  Config file: %s\n", configPath)
	},
}

func init() {
	configCmd.AddCommand(configShowCmd)
}

// requestCmd handles HTTP requests
var requestCmd = &cobra.Command{
	Use:   "[METHOD] [path]",
	Short: "Make an HTTP request",
	Long: `Make an HTTP request to a service endpoint.

Examples:
  sreq GET /api/v1/users -s auth-service -e dev
  sreq POST /api/v1/users -s auth-service -d '{"name":"test"}'`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		method := args[0]
		path := args[1]

		if serviceName == "" {
			fmt.Println("Error: --service (-s) is required")
			os.Exit(1)
		}

		if environment == "" {
			environment = "dev" // default
		}

		if verbose {
			fmt.Printf("Service: %s\n", serviceName)
			fmt.Printf("Environment: %s\n", environment)
			fmt.Printf("Method: %s\n", method)
			fmt.Printf("Path: %s\n", path)
			fmt.Println("---")
		}

		if dryRun {
			fmt.Println("[DRY RUN] Would execute:")
			fmt.Printf("  %s https://<resolved-url>%s\n", method, path)
			return
		}

		// TODO: Implement actual request logic
		fmt.Printf("Making %s request to %s on %s (%s)...\n", method, path, serviceName, environment)
	},
}

var requestData string
var requestHeaders []string

func init() {
	requestCmd.Flags().StringVarP(&requestData, "data", "d", "", "Request body (or @filename)")
	requestCmd.Flags().StringArrayVarP(&requestHeaders, "header", "H", nil, "Add header (repeatable)")
}
