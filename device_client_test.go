package goadb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zach-klippenstein/goadb/wire"
)

func TestGetAttribute(t *testing.T) {
	s := &MockServer{
		Status:   wire.StatusSuccess,
		Messages: []string{"value"},
	}
	client := &DeviceClient{s, deviceWithSerial("serial")}

	v, err := client.getAttribute("attr")
	assert.Equal(t, "host-serial:serial:attr", s.Requests[0])
	assert.NoError(t, err)
	assert.Equal(t, "value", v)
}

func TestRunCommandNoArgs(t *testing.T) {
	s := &MockServer{
		Status:   wire.StatusSuccess,
		Messages: []string{"output"},
	}
	client := &DeviceClient{s, anyDevice()}

	v, err := client.RunCommand("cmd")
	assert.Equal(t, "host:transport-any", s.Requests[0])
	assert.Equal(t, "shell:cmd", s.Requests[1])
	assert.NoError(t, err)
	assert.Equal(t, "output", v)
}

func TestPrepareCommandLineNoArgs(t *testing.T) {
	result, err := prepareCommandLine("cmd")
	assert.NoError(t, err)
	assert.Equal(t, "cmd", result)
}

func TestPrepareCommandLineEmptyCommand(t *testing.T) {
	_, err := prepareCommandLine("")
	assert.EqualError(t, err, "command cannot be empty")
}

func TestPrepareCommandLineBlankCommand(t *testing.T) {
	_, err := prepareCommandLine("  ")
	assert.EqualError(t, err, "command cannot be empty")
}

func TestPrepareCommandLineCleanArgs(t *testing.T) {
	result, err := prepareCommandLine("cmd", "arg1", "arg2")
	assert.NoError(t, err)
	assert.Equal(t, "cmd arg1 arg2", result)
}

func TestPrepareCommandLineArgWithWhitespaceQuotes(t *testing.T) {
	result, err := prepareCommandLine("cmd", "arg with spaces")
	assert.NoError(t, err)
	assert.Equal(t, "cmd \"arg with spaces\"", result)
}

func TestPrepareCommandLineArgWithDoubleQuoteFails(t *testing.T) {
	_, err := prepareCommandLine("cmd", "quoted\"arg")
	assert.EqualError(t, err, "arg at index 0 contains an invalid double quote: quoted\"arg")
}
