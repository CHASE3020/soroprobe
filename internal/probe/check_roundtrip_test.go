package probe_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/soroworks/soroprobe/internal/probe"
)

func TestCheckJSONRoundTrip(t *testing.T) {
	t.Parallel()

	original := probe.Check{
		Name:    "instance_ttl",
		Outcome: probe.OutcomePass,
		Detail:  "instance is healthy with 100000 ledgers remaining",
	}

	data, err := json.Marshal(original)
	require.NoError(t, err)

	var decoded probe.Check
	err = json.Unmarshal(data, &decoded)
	require.NoError(t, err)

	assert.Equal(t, original, decoded)
}
