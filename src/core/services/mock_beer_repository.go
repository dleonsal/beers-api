// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package services

import (
	entities "github.com/dleonsal/beers-api/src/core/domain/entities"
	errors "github.com/dleonsal/beers-api/src/errors"

	mock "github.com/stretchr/testify/mock"
)

// MockBeerRepository is an autogenerated mock type for the BeerRepository type
type MockBeerRepository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: beerID
func (_m *MockBeerRepository) GetByID(beerID int64) (*entities.Beer, *errors.RestError) {
	ret := _m.Called(beerID)

	var r0 *entities.Beer
	if rf, ok := ret.Get(0).(func(int64) *entities.Beer); ok {
		r0 = rf(beerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Beer)
		}
	}

	var r1 *errors.RestError
	if rf, ok := ret.Get(1).(func(int64) *errors.RestError); ok {
		r1 = rf(beerID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.RestError)
		}
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *MockBeerRepository) List() ([]entities.Beer, *errors.RestError) {
	ret := _m.Called()

	var r0 []entities.Beer
	if rf, ok := ret.Get(0).(func() []entities.Beer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Beer)
		}
	}

	var r1 *errors.RestError
	if rf, ok := ret.Get(1).(func() *errors.RestError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.RestError)
		}
	}

	return r0, r1
}

// Save provides a mock function with given fields: beer
func (_m *MockBeerRepository) Save(beer entities.Beer) *errors.RestError {
	ret := _m.Called(beer)

	var r0 *errors.RestError
	if rf, ok := ret.Get(0).(func(entities.Beer) *errors.RestError); ok {
		r0 = rf(beer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.RestError)
		}
	}

	return r0
}