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

	var restored probe.CostChange
	err = json.Unmarshal(data, &restored)
	require.NoError(t, err)

	assert.Equal(t, original.Name, restored.Name)
	assert.Equal(t, original.Baseline, restored.Baseline)
	assert.Equal(t, original.Current, restored.Current)
	assert.Equal(t, original.Delta, restored.Delta)
	require.NotNil(t, restored.Percent)
	assert.Equal(t, *original.Percent, *restored.Percent)
}

func TestCostComparisonJSONRoundTrip(t *testing.T) {
	t.Parallel()

	pct := 15.0
	limit := 20.0
	original := probe.CostComparison{
		Figures: []probe.CostChange{
			{
				Name:     "cpu instructions",
				Baseline: 1000,
				Current:  1150,
				Delta:    150,
				Percent:  &pct,
			},
		},
		MaxIncreasePercent: &limit,
		Exceeded:           false,
	}

	data, err := json.Marshal(original)
	require.NoError(t, err)

	var restored probe.CostComparison
	err = json.Unmarshal(data, &restored)
	require.NoError(t, err)

	assert.Equal(t, len(original.Figures), len(restored.Figures))
	assert.Equal(t, original.Figures[0].Name, restored.Figures[0].Name)
	assert.Equal(t, original.Figures[0].Baseline, restored.Figures[0].Baseline)
	assert.Equal(t, original.Figures[0].Current, restored.Figures[0].Current)
	assert.Equal(t, original.Figures[0].Delta, restored.Figures[0].Delta)
	require.NotNil(t, restored.Figures[0].Percent)
	assert.Equal(t, *original.Figures[0].Percent, *restored.Figures[0].Percent)
	require.NotNil(t, restored.MaxIncreasePercent)
	assert.Equal(t, *original.MaxIncreasePercent, *restored.MaxIncreasePercent)
	assert.Equal(t, original.Exceeded, restored.Exceeded)
}
