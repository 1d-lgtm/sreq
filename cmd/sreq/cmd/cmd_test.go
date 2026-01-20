package cmd

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRootCmd(t *testing.T) {
	// Test that root command exists and has correct properties
	if rootCmd.Use != "sreq" {
		t.Errorf("rootCmd.Use = %q, want %q", rootCmd.Use, "sreq")
	}

	if rootCmd.Short == "" {
		t.Error("rootCmd.Short should not be empty")
	}

	if rootCmd.Long == "" {
		t.Error("rootCmd.Long should not be empty")
	}
}

func TestRootCmd_PersistentFlags(t *testing.T) {
	flags := []struct {
		name      string
		shorthand string
	}{
		{"service", "s"},
		{"context", "c"},
		{"env", "e"},
		{"region", "r"},
		{"project", "p"},
		{"app", "a"},
		{"verbose", "v"},
		{"dry-run", ""},
	}

	for _, f := range flags {
		t.Run(f.name, func(t *testing.T) {
			flag := rootCmd.PersistentFlags().Lookup(f.name)
			if flag == nil {
				t.Errorf("Flag --%s not found", f.name)
				return
			}
			if f.shorthand != "" && flag.Shorthand != f.shorthand {
				t.Errorf("Flag shorthand = %q, want %q", flag.Shorthand, f.shorthand)
			}
		})
	}
}

func TestVersionCmd(t *testing.T) {
	if versionCmd.Use != "version" {
		t.Errorf("versionCmd.Use = %q, want %q", versionCmd.Use, "version")
	}

	if versionCmd.Short == "" {
		t.Error("versionCmd should have short description")
	}
}

func TestSubCommands(t *testing.T) {
	// Get all subcommands
	subcommands := rootCmd.Commands()

	// Check expected commands exist
	expectedCmds := []string{"version", "init", "run", "service", "auth", "config", "env", "history", "cache", "tui"}

	cmdMap := make(map[string]bool)
	for _, cmd := range subcommands {
		cmdMap[cmd.Use] = true
		// Also check commands with arguments like "run [method] [path]"
		parts := strings.Fields(cmd.Use)
		if len(parts) > 0 {
			cmdMap[parts[0]] = true
		}
	}

	for _, expected := range expectedCmds {
		if !cmdMap[expected] {
			t.Errorf("Expected subcommand %q not found", expected)
		}
	}
}

func TestInitCmd(t *testing.T) {
	cmd := findSubCommand("init")
	if cmd == nil {
		t.Fatal("init command not found")
	}

	if cmd.Short == "" {
		t.Error("init command should have short description")
	}
}

func TestRunCmd(t *testing.T) {
	cmd := findSubCommand("run")
	if cmd == nil {
		t.Fatal("run command not found")
	}

	// Check run command flags
	flags := []string{"header", "data", "timeout", "output"}
	for _, f := range flags {
		if flag := cmd.Flags().Lookup(f); flag == nil {
			t.Errorf("run command missing flag: %s", f)
		}
	}
}

func TestServiceCmd(t *testing.T) {
	cmd := findSubCommand("service")
	if cmd == nil {
		t.Fatal("service command not found")
	}

	// Check subcommands
	subCmds := cmd.Commands()
	subCmdMap := make(map[string]bool)
	for _, sub := range subCmds {
		parts := strings.Fields(sub.Use)
		if len(parts) > 0 {
			subCmdMap[parts[0]] = true
		}
	}

	expectedSubCmds := []string{"add", "list", "remove"}
	for _, expected := range expectedSubCmds {
		if !subCmdMap[expected] {
			t.Errorf("service subcommand %q not found", expected)
		}
	}
}

func TestAuthCmd(t *testing.T) {
	cmd := findSubCommand("auth")
	if cmd == nil {
		t.Fatal("auth command not found")
	}

	// Check subcommands
	subCmds := cmd.Commands()
	subCmdMap := make(map[string]bool)
	for _, sub := range subCmds {
		parts := strings.Fields(sub.Use)
		if len(parts) > 0 {
			subCmdMap[parts[0]] = true
		}
	}

	expectedSubCmds := []string{"consul", "aws"}
	for _, expected := range expectedSubCmds {
		if !subCmdMap[expected] {
			t.Errorf("auth subcommand %q not found", expected)
		}
	}
}

func TestConfigCmd(t *testing.T) {
	cmd := findSubCommand("config")
	if cmd == nil {
		t.Fatal("config command not found")
	}

	// Check subcommands
	subCmds := cmd.Commands()
	subCmdMap := make(map[string]bool)
	for _, sub := range subCmds {
		parts := strings.Fields(sub.Use)
		if len(parts) > 0 {
			subCmdMap[parts[0]] = true
		}
	}

	expectedSubCmds := []string{"show", "path", "test"}
	for _, expected := range expectedSubCmds {
		if !subCmdMap[expected] {
			t.Errorf("config subcommand %q not found", expected)
		}
	}
}

func TestHistoryCmd(t *testing.T) {
	cmd := findSubCommand("history")
	if cmd == nil {
		t.Fatal("history command not found")
	}

	// History is a single command with flags (not subcommands)
	// Check key flags
	flags := []string{"service", "env", "method", "all", "clear", "before", "curl", "httpie", "replay"}
	for _, f := range flags {
		if flag := cmd.Flags().Lookup(f); flag == nil {
			t.Errorf("history command missing flag: %s", f)
		}
	}
}

func TestCacheCmd(t *testing.T) {
	cmd := findSubCommand("cache")
	if cmd == nil {
		t.Fatal("cache command not found")
	}

	// Check subcommands
	subCmds := cmd.Commands()
	subCmdMap := make(map[string]bool)
	for _, sub := range subCmds {
		parts := strings.Fields(sub.Use)
		if len(parts) > 0 {
			subCmdMap[parts[0]] = true
		}
	}

	expectedSubCmds := []string{"clear", "status"}
	for _, expected := range expectedSubCmds {
		if !subCmdMap[expected] {
			t.Errorf("cache subcommand %q not found", expected)
		}
	}
}

func TestEnvCmd(t *testing.T) {
	cmd := findSubCommand("env")
	if cmd == nil {
		t.Fatal("env command not found")
	}

	// Check subcommands
	subCmds := cmd.Commands()
	subCmdMap := make(map[string]bool)
	for _, sub := range subCmds {
		parts := strings.Fields(sub.Use)
		if len(parts) > 0 {
			subCmdMap[parts[0]] = true
		}
	}

	expectedSubCmds := []string{"list", "switch", "current"}
	for _, expected := range expectedSubCmds {
		if !subCmdMap[expected] {
			t.Errorf("env subcommand %q not found", expected)
		}
	}
}

func TestTuiCmd(t *testing.T) {
	cmd := findSubCommand("tui")
	if cmd == nil {
		t.Fatal("tui command not found")
	}

	if cmd.Short == "" {
		t.Error("tui command should have short description")
	}
}

// Helper function to find a subcommand by name
func findSubCommand(name string) *cobra.Command {
	for _, cmd := range rootCmd.Commands() {
		parts := strings.Fields(cmd.Use)
		if len(parts) > 0 && parts[0] == name {
			return cmd
		}
	}
	return nil
}

func TestExecute(t *testing.T) {
	// Test that Execute function exists and doesn't panic
	// We can't fully test Execute without causing side effects
	// but we can verify it's callable
	_ = Execute // Just verify it compiles
}

func TestVersionVariable(t *testing.T) {
	// Version should have a default value
	if Version == "" {
		t.Error("Version should not be empty")
	}
}
