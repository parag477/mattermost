// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	request "github.com/mattermost/mattermost/server/public/shared/request"
	mock "github.com/stretchr/testify/mock"
)

// MessageExportInterface is an autogenerated mock type for the MessageExportInterface type
type MessageExportInterface struct {
	mock.Mock
}

// RunExport provides a mock function with given fields: c, format, since, limit
func (_m *MessageExportInterface) RunExport(c *request.Context, format string, since int64, limit int) (int64, *model.AppError) {
	ret := _m.Called(c, format, since, limit)

	var r0 int64
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(*request.Context, string, int64, int) (int64, *model.AppError)); ok {
		return rf(c, format, since, limit)
	}
	if rf, ok := ret.Get(0).(func(*request.Context, string, int64, int) int64); ok {
		r0 = rf(c, format, since, limit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*request.Context, string, int64, int) *model.AppError); ok {
		r1 = rf(c, format, since, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// StartSynchronizeJob provides a mock function with given fields: c, exportFromTimestamp
func (_m *MessageExportInterface) StartSynchronizeJob(c *request.Context, exportFromTimestamp int64) (*model.Job, *model.AppError) {
	ret := _m.Called(c, exportFromTimestamp)

	var r0 *model.Job
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(*request.Context, int64) (*model.Job, *model.AppError)); ok {
		return rf(c, exportFromTimestamp)
	}
	if rf, ok := ret.Get(0).(func(*request.Context, int64) *model.Job); ok {
		r0 = rf(c, exportFromTimestamp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(*request.Context, int64) *model.AppError); ok {
		r1 = rf(c, exportFromTimestamp)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewMessageExportInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMessageExportInterface creates a new instance of MessageExportInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMessageExportInterface(t mockConstructorTestingTNewMessageExportInterface) *MessageExportInterface {
	mock := &MessageExportInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
