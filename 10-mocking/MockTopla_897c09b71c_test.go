package tests

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepositoryTest struct {
	mock.Mock
}

func (mock *MockRepositoryTest) MockTopla(sayilar []int) (int, error) {
	args := mock.Called(sayilar)
	result := args.Get(0)

	return result.(int), args.Error(1)
}

func TestMockTopla_897c09b71c(t *testing.T) {
	mockRepo := new(MockRepositoryTest)

	mockRepo.On("MockTopla", []int{1, 2, 3}).Return(6, nil)

	result, err := mockRepo.MockTopla([]int{1, 2, 3})

	assert.NoError(t, err)
	assert.Equal(t, 6, result)
	t.Log("Success: Sum of integers is correct")

	mockRepo.On("MockTopla", []int{1, 2, 3}).Return(0, fmt.Errorf("some error"))

	result, err = mockRepo.MockTopla([]int{1, 2, 3})

	assert.Error(t, err)
	t.Log("Success: Error was returned as expected")

	assert.NotEqual(t, 6, result)
	t.Log("Success: Sum of integers is not as expected due to error")
}
