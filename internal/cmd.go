package internal

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/itsfunky/funky/providers"
)

var (
	ldFlag  = "-X github.com/itsfunky/funky.%s=%s"
	command = "go %s -tags %s -ldflags \"%s\" *.go"
)

// CommandAction enum values.
var (
	BuildAction CommandAction = "build"
	RunAction   CommandAction = "run"
)

// CommandAction represents the available actions CreateCommand supports.
type CommandAction string

// GetLDFlags returns a list of LD Flags as a string.
func GetLDFlags(_ context.Context, provider providers.Provider, function string) string {
	flagsMap := map[string]string{
		"FunctionName": function,
		"Provider":     provider.Name,
	}

	flags := make([]string, 0, len(flagsMap))
	for k, v := range flagsMap {
		flags = append(flags, fmt.Sprintf(ldFlag, k, v))
	}

	return strings.Join(flags, " ")
}

// CreateCommand instantiates a command to be started or run later.
func CreateCommand(ctx context.Context, provider providers.Provider, function string, action CommandAction) *exec.Cmd {
	ldFlags := GetLDFlags(ctx, provider, function)
	cmdText := fmt.Sprintf(command, action, provider.Name, ldFlags)
	cmd := exec.CommandContext(ctx, "sh", "-c", cmdText)

	cmd.Dir = filepath.Join("functions", function)
	cmd.Env = append(
		os.Environ(),
		fmt.Sprintf("FUNKY_FUNCTION_NAME=%s", function),
	)

	return cmd
}
