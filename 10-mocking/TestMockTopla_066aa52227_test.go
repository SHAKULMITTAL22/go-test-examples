package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) MockTopla(nums []int) (int, error) {
	args := m.Called(nums)
	return args.Int(0), args.Error(1)
}

func Matematik(repo *MockRepository) *MockRepository {
	return repo
}

func TestMockTopla(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("MockTopla", []int{2, 3}).Return(5, nil)

	testMatematik := Matematik(mockRepo)

	sonuc, err := testMatematik.MockTopla([]int{2, 3})

	mockRepo.AssertExpectations(t)

	assert.Equal(t, 5, sonuc)
	assert.Nil(t, err)

	mockRepo.On("MockTopla", []int{2, 2}).Return(4, nil)
	sonuc, err = testMatematik.MockTopla([]int{2, 2})
	assert.Equal(t, 4, sonuc)
	assert.Nil(t, err)

	mockRepo.On("MockTopla", []int{1, 2}).Return(0, errors.New("some error"))
	sonuc, err = testMatematik.MockTopla([]int{1, 2})
	assert.Equal(t, 0, sonuc)
	assert.NotNil(t, err)
}
