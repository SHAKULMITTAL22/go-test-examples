package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type islem struct{}

func (*islem) MockTopla(sayilar []int) (int, error) {
	toplam := 0
	for i := range sayilar {
		toplam = toplam + sayilar[i]
	}
	return toplam, nil
}

func TestMockTopla_fcac46a2e6(t *testing.T) {
	testCases := []struct {
		name     string
		sayilar  []int
		expected int
	}{
		{
			name:     "Adding positive numbers",
			sayilar:  []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "Adding negative numbers",
			sayilar:  []int{-1, -2, -3, -4, -5},
			expected: -15,
		},
		{
			name:     "Adding mix of positive and negative numbers",
			sayilar:  []int{-1, 2, -3, 4, -5},
			expected: -3,
		},
		{
			name:     "Adding zero",
			sayilar:  []int{0, 0, 0, 0, 0},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i := &islem{}
			result, err := i.MockTopla(tc.sayilar)
			assert.Equal(t, tc.expected, result)
			assert.NoError(t, err)
			t.Logf("Test case %s passed", tc.name)
		})
	}
}
