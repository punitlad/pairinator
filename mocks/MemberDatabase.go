// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	models "pairinator/models"

	mock "github.com/stretchr/testify/mock"
)

// MemberDatabase is an autogenerated mock type for the MemberDatabase type
type MemberDatabase struct {
	mock.Mock
}

// Add provides a mock function with given fields: member
func (_m *MemberDatabase) Add(member models.Member) {
	_m.Called(member)
}

// GetAll provides a mock function with given fields:
func (_m *MemberDatabase) GetAll() []models.Member {
	ret := _m.Called()

	var r0 []models.Member
	if rf, ok := ret.Get(0).(func() []models.Member); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Member)
		}
	}

	return r0
}
