package monolith

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"testing"
)

func TestCreateApp(t *testing.T) {
	err := fx.ValidateApp(CreateApp())
	require.NoError(t, err)
}
