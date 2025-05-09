package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnInvalidPage_WhenCreateANewListOrders_ThenShouldReceiveAnError(t *testing.T) {
	listOrders := &ListOrders{Page: 0, Limit: 10}
	err := listOrders.IsValid()
	assert.ErrorIs(t, err, errInvalidPage)
}

func TestGivenAnInvalidLimit_WhenCreateANewListOrders_ThenShouldReceiveAnError(t *testing.T) {
	listOrders := &ListOrders{Page: 1, Limit: 0}
	err := listOrders.IsValid()
	assert.ErrorIs(t, err, errInvalidLimit)
}

func TestGivenValidParams_WhenCreateANewListOrders_ThenShouldBeValid(t *testing.T) {
	listOrders := &ListOrders{Page: 1, Limit: 10}
	err := listOrders.IsValid()
	assert.Nil(t, err)
	assert.Equal(t, 1, listOrders.Page)
	assert.Equal(t, 10, listOrders.Limit)
}

func TestGivenValidParams_WhenICallNewListOrdersFunc_ThenShouldReturnListOrders(t *testing.T) {
	listOrders, err := NewListOrders(1, 10)
	assert.Nil(t, err)
	assert.Equal(t, 1, listOrders.Page)
	assert.Equal(t, 10, listOrders.Limit)
}

func TestGivenInvalidParams_WhenICallNewListOrdersFunc_ThenShouldReturnError(t *testing.T) {
	listOrders, err := NewListOrders(0, -5)
	assert.Nil(t, listOrders)
	assert.Error(t, err)
}
