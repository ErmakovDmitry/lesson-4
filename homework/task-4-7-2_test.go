package homework

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

const N = 10

func TestSelectMany(t *testing.T) {

	cases := map[string]struct {
		nChannels int
	}{
		"one": {
			nChannels: 1,
		},
		"two": {
			nChannels: 2,
		},
		"many": {
			nChannels: 5,
		},
		"a lot": {
			nChannels: 200,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			inputs := make([]chan int64, tc.nChannels)
			expected := make([]int64, 0)
			for n := 0; n < tc.nChannels; n++ {
				inputs[n] = make(chan int64, N)
				for i := 0; i < N; i++ {
					random := int64(rand.IntN(100))
					expected = append(expected, random)
					inputs[n] <- random
				}
				close(inputs[n])
			}

			resultCh := selectMany(inputs)
			results := make([]int64, 0)
			for i := 0; i < (N * tc.nChannels); i++ {
				results = append(results, <-resultCh)
			}

			slices.Sort(results)
			slices.Sort(expected)

			require.Equal(t, len(results), len(expected), fmt.Sprintf("len(results) %d != len(expected) %d", len(results), len(expected)))
			require.True(t, reflect.DeepEqual(results, expected), fmt.Sprintf("%v != %v", results, expected))
		})
	}
}
