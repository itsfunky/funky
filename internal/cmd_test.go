package internal

import (
	"context"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/itsfunky/funky/providers"
)

func TestGetLDFlags(t *testing.T) {
	p := providers.Provider{
		Name: "tester",
	}

	out := GetLDFlags(context.Background(), p, "test_func")
	expected := "-X github.com/itsfunky/funky.FunctionName=test_func -X github.com/itsfunky/funky.Provider=tester"

	assert.Equal(t, expected, out, "LD Flags did not match.")
}

func testCreateCommandSuite(t *testing.T, action CommandAction) {
	p := providers.Provider{
		Name: "tester",
	}

	cmd := CreateCommand(context.Background(), p, "foobar", action)
	assert.Equal(t, filepath.Join("functions", "foobar"), cmd.Dir, "Dir should point to the function's folder.")

	assert.Equal(t, 3, len(cmd.Args), "Command should only have 2 arguments")
	assert.Equal(t, "sh", cmd.Args[0], "Command path should be 'sh'.")
	assert.Equal(t, "-c", cmd.Args[1], "First command arg should be '-c'.")

	expectedLDFlags := GetLDFlags(context.Background(), p, "foobar")
	expectedCommand := "go " + string(action) + " -tags tester -ldflags \"" + expectedLDFlags + "\" *.go"
	assert.Equal(t, expectedCommand, cmd.Args[2], "Second command arg should be a proper 'go' command.")

	gotFunkyEnvs := map[string]string{}
	expectedFunkyEnvs := map[string]string{
		"FUNKY_FUNCTION_NAME": "foobar",
	}

	for _, v := range cmd.Env {
		if strings.HasPrefix(v, "FUNKY_") {
			kv := strings.SplitN(v, "=", 2)
			gotFunkyEnvs[kv[0]] = kv[1]
		}
	}

	assert.Equal(t, expectedFunkyEnvs, gotFunkyEnvs, "Env should include both Environment variables and Funky variables.")
}

func TestCreateCommandBuild(t *testing.T) {
	testCreateCommandSuite(t, BuildAction)
}

func TestCreateCommandRun(t *testing.T) {
	testCreateCommandSuite(t, RunAction)
}
