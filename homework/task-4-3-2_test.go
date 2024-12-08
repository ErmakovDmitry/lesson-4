package homework

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const channelLength = 10
const writeNumber = 500

// go test lesson-3/homework -run TestSumChannels
func TestSumChannels(t *testing.T) {
	cases := map[string]struct {
		nChannels int
		sum       int64
	}{
		"empty": {
			nChannels: 0,
			sum:       0,
		},
		"one": {
			nChannels: 1,
			sum:       writeNumber * channelLength,
		},
		"many": {
			nChannels: 5,
			sum:       writeNumber * channelLength * 5,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			inputs := make([]chan int64, tc.nChannels)
			for n := 0; n < tc.nChannels; n++ {
				inputs[n] = make(chan int64, 10)
				for i := 0; i < 10; i++ {
					inputs[n] <- writeNumber
				}
				close(inputs[n])
			}
			result := SumChannels(inputs)
			require.Equal(t, tc.sum, result)
		})
	}
}
