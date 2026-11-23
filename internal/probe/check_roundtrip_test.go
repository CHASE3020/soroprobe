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
		Detail:  "instance is healthy with plenty of ledgers remaining",
	}

	data, err := json.Marshal(original)
	require.NoError(t, err, "marshaling Check should succeed")

	var unmarshaled probe.Check
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err, "unmarshaling Check should succeed")

	assert.Equal(t, original, unmarshaled, "round-tripped Check should match original")
}
