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

	pct := 12.5
	original := probe.CostChange{
		Name:     "cpu instructions",
		Baseline: 1000,
		Current:  1125,
		Delta:    125,
		Percent:  &pct,
	}

	data, err := json.Marshal(original)
	require.NoError(t, err)

	var decoded probe.CostChange
	err = json.Unmarshal(data, &decoded)
	require.NoError(t, err)

	assert.Equal(t, original.Name, decoded.Name)
	assert.Equal(t, original.Baseline, decoded.Baseline)
	assert.Equal(t, original.Current, decoded.Current)
	assert.Equal(t, original.Delta, decoded.Delta)
	require.NotNil(t, decoded.Percent)
	assert.Equal(t, *original.Percent, *decoded.Percent)
}

func TestCostComparisonJSONRoundTrip(t *testing.T) {
	t.Parallel()

	maxInc := 10.0
	original := probe.CostComparison{
		Figures: []probe.CostChange{
			{
				Name:     "cpu instructions",
				Baseline: 1000,
				Current:  1125,
				Delta:    125,
			},
		},
		MaxIncreasePercent: &maxInc,
		Exceeded:           true,
	}

	data, err := json.Marshal(original)
	require.NoError(t, err)

	var decoded probe.CostComparison
	err = json.Unmarshal(data, &decoded)
	require.NoError(t, err)

	assert.Equal(t, original.Exceeded, decoded.Exceeded)
	require.NotNil(t, decoded.MaxIncreasePercent)
	assert.Equal(t, *original.MaxIncreasePercent, *decoded.MaxIncreasePercent)
	require.Len(t, decoded.Figures, 1)
	assert.Equal(t, original.Figures[0].Name, decoded.Figures[0].Name)
	assert.Equal(t, original.Figures[0].Baseline, decoded.Figures[0].Baseline)
	assert.Equal(t, original.Figures[0].Current, decoded.Figures[0].Current)
	assert.Equal(t, original.Figures[0].Delta, decoded.Figures[0].Delta)
}
