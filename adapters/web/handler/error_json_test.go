package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_errorJSON(t *testing.T) {
	msg := "error JSON"
	result := errorJSON(msg)
	require.Equal(t, string([]byte(`{"message":"error JSON"}`)), string(result))

}
