package probe_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/soroworks/soroprobe/internal/probe"
)

func TestCostChangeJSONRoundTrip(t *testing.T) {
	t.Parallel()

	percent := 12.5
	original := probe.CostChange{
		Name:     "cpu instructions",
		Baseline: 1000,
		Current:  1125,
		Delta:    125,
		Percent:  &percent,
	}

	data, err := json.Marshal(original)
	require.NoError(t, err, "marshaling CostChange should succeed")

	var unmarshaled probe.CostChange
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err, "unmarshaling CostChange should succeed")

	assert.Equal(t, original.Name, unmarshaled.Name)
	assert.Equal(t, original.Baseline, unmarshaled.Baseline)
	assert.Equal(t, original.Current, unmarshaled.Current)
	assert.Equal(t, original.Delta, unmarshaled.Delta)
	require.NotNil(t, unmarshaled.Percent)
	assert.Equal(t, *original.Percent, *unmarshaled.Percent)
}

func TestCostComparisonJSONRoundTrip(t *testing.T) {
	t.Parallel()

	percent := 15.0
	maxIncrease := 20.0
	original := probe.CostComparison{
		Figures: []probe.CostChange{
			{
				Name:     "cpu instructions",
				Baseline: 100,
				Current:  115,
				Delta:    15,
				Percent:  &percent,
			},
		},
		MaxIncreasePercent: &maxIncrease,
		Exceeded:           false,
	}

	data, err := json.Marshal(original)
	require.NoError(t, err, "marshaling CostComparison should succeed")

	var unmarshaled probe.CostComparison
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err, "unmarshaling CostComparison should succeed")

	require.Len(t, unmarshaled.Figures, 1)
	assert.Equal(t, original.Figures[0].Name, unmarshaled.Figures[0].Name)
	assert.Equal(t, original.Figures[0].Baseline, unmarshaled.Figures[0].Baseline)
	assert.Equal(t, original.Figures[0].Current, unmarshaled.Figures[0].Current)
	assert.Equal(t, original.Figures[0].Delta, unmarshaled.Figures[0].Delta)
	require.NotNil(t, unmarshaled.Figures[0].Percent)
	assert.Equal(t, *original.Figures[0].Percent, *unmarshaled.Figures[0].Percent)

	require.NotNil(t, unmarshaled.MaxIncreasePercent)
	assert.Equal(t, *original.MaxIncreasePercent, *unmarshaled.MaxIncreasePercent)
	assert.Equal(t, original.Exceeded, unmarshaled.Exceeded)
}
