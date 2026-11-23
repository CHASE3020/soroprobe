package probe_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/soroworks/soroprobe/internal/probe"
)

func TestCostComparisonJSONRoundTrip(t *testing.T) {
	t.Parallel()

	percent := 15.5
	original := probe.CostComparison{
		Figures: []probe.CostChange{
			{
				Name:     "cpu instructions",
				Baseline: 1000,
				Current:  1200,
				Delta:    200,
				Percent:  &percent,
			},
		},
		MaxIncreasePercent: &percent,
		Exceeded:           true,
	}

	data, err := json.Marshal(original)
	require.NoError(t, err)

	var unmarshalled probe.CostComparison
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)

	assert.Equal(t, original, unmarshalled)
}
