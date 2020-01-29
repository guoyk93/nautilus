package in_k8s

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetStatefulSetSequenceID(t *testing.T) {
	require.Equal(t, uint64(0), extractStatefulSetSequenceID("test"))
	require.Equal(t, uint64(1), extractStatefulSetSequenceID("test-0"))
	require.Equal(t, uint64(2), extractStatefulSetSequenceID("test-1"))
}
