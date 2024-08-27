// Code generated by mockery. DO NOT EDIT.

package extractor

import (
	context "context"

	extractor "github.com/ilfey/hikilist-go/pkg/parser/extractor"
	mock "github.com/stretchr/testify/mock"

	models "github.com/ilfey/hikilist-go/pkg/parser/models"
)

// AnimeExtractor is an autogenerated mock type for the AnimeExtractor type
type AnimeExtractor struct {
	mock.Mock
}

type AnimeExtractor_Expecter struct {
	mock *mock.Mock
}

func (_m *AnimeExtractor) EXPECT() *AnimeExtractor_Expecter {
	return &AnimeExtractor_Expecter{mock: &_m.Mock}
}

// FetchById provides a mock function with given fields: id
func (_m *AnimeExtractor) FetchById(id uint) (*models.AnimeDetailModel, error) {
	ret := _m.Called(id)

	var r0 *models.AnimeDetailModel
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*models.AnimeDetailModel, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *models.AnimeDetailModel); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AnimeDetailModel)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnimeExtractor_FetchById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchById'
type AnimeExtractor_FetchById_Call struct {
	*mock.Call
}

// FetchById is a helper method to define mock.On call
//   - id uint
func (_e *AnimeExtractor_Expecter) FetchById(id interface{}) *AnimeExtractor_FetchById_Call {
	return &AnimeExtractor_FetchById_Call{Call: _e.mock.On("FetchById", id)}
}

func (_c *AnimeExtractor_FetchById_Call) Run(run func(id uint)) *AnimeExtractor_FetchById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *AnimeExtractor_FetchById_Call) Return(_a0 *models.AnimeDetailModel, _a1 error) *AnimeExtractor_FetchById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AnimeExtractor_FetchById_Call) RunAndReturn(run func(uint) (*models.AnimeDetailModel, error)) *AnimeExtractor_FetchById_Call {
	_c.Call.Return(run)
	return _c
}

// FetchList provides a mock function with given fields: page
func (_m *AnimeExtractor) FetchList(page uint) (*models.AnimeListModel, error) {
	ret := _m.Called(page)

	var r0 *models.AnimeListModel
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*models.AnimeListModel, error)); ok {
		return rf(page)
	}
	if rf, ok := ret.Get(0).(func(uint) *models.AnimeListModel); ok {
		r0 = rf(page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AnimeListModel)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnimeExtractor_FetchList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchList'
type AnimeExtractor_FetchList_Call struct {
	*mock.Call
}

// FetchList is a helper method to define mock.On call
//   - page uint
func (_e *AnimeExtractor_Expecter) FetchList(page interface{}) *AnimeExtractor_FetchList_Call {
	return &AnimeExtractor_FetchList_Call{Call: _e.mock.On("FetchList", page)}
}

func (_c *AnimeExtractor_FetchList_Call) Run(run func(page uint)) *AnimeExtractor_FetchList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *AnimeExtractor_FetchList_Call) Return(_a0 *models.AnimeListModel, _a1 error) *AnimeExtractor_FetchList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AnimeExtractor_FetchList_Call) RunAndReturn(run func(uint) (*models.AnimeListModel, error)) *AnimeExtractor_FetchList_Call {
	_c.Call.Return(run)
	return _c
}

// OnFetchError provides a mock function with given fields: err
func (_m *AnimeExtractor) OnFetchError(err error) extractor.Action {
	ret := _m.Called(err)

	var r0 extractor.Action
	if rf, ok := ret.Get(0).(func(error) extractor.Action); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Get(0).(extractor.Action)
	}

	return r0
}

// AnimeExtractor_OnFetchError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnFetchError'
type AnimeExtractor_OnFetchError_Call struct {
	*mock.Call
}

// OnFetchError is a helper method to define mock.On call
//   - err error
func (_e *AnimeExtractor_Expecter) OnFetchError(err interface{}) *AnimeExtractor_OnFetchError_Call {
	return &AnimeExtractor_OnFetchError_Call{Call: _e.mock.On("OnFetchError", err)}
}

func (_c *AnimeExtractor_OnFetchError_Call) Run(run func(err error)) *AnimeExtractor_OnFetchError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *AnimeExtractor_OnFetchError_Call) Return(_a0 extractor.Action) *AnimeExtractor_OnFetchError_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AnimeExtractor_OnFetchError_Call) RunAndReturn(run func(error) extractor.Action) *AnimeExtractor_OnFetchError_Call {
	_c.Call.Return(run)
	return _c
}

// OnResolveError provides a mock function with given fields: err
func (_m *AnimeExtractor) OnResolveError(err error) extractor.Action {
	ret := _m.Called(err)

	var r0 extractor.Action
	if rf, ok := ret.Get(0).(func(error) extractor.Action); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Get(0).(extractor.Action)
	}

	return r0
}

// AnimeExtractor_OnResolveError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnResolveError'
type AnimeExtractor_OnResolveError_Call struct {
	*mock.Call
}

// OnResolveError is a helper method to define mock.On call
//   - err error
func (_e *AnimeExtractor_Expecter) OnResolveError(err interface{}) *AnimeExtractor_OnResolveError_Call {
	return &AnimeExtractor_OnResolveError_Call{Call: _e.mock.On("OnResolveError", err)}
}

func (_c *AnimeExtractor_OnResolveError_Call) Run(run func(err error)) *AnimeExtractor_OnResolveError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *AnimeExtractor_OnResolveError_Call) Return(_a0 extractor.Action) *AnimeExtractor_OnResolveError_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AnimeExtractor_OnResolveError_Call) RunAndReturn(run func(error) extractor.Action) *AnimeExtractor_OnResolveError_Call {
	_c.Call.Return(run)
	return _c
}

// Resolve provides a mock function with given fields: ctx, detailModel
func (_m *AnimeExtractor) Resolve(ctx context.Context, detailModel *models.AnimeDetailModel) error {
	ret := _m.Called(ctx, detailModel)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.AnimeDetailModel) error); ok {
		r0 = rf(ctx, detailModel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AnimeExtractor_Resolve_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Resolve'
type AnimeExtractor_Resolve_Call struct {
	*mock.Call
}

// Resolve is a helper method to define mock.On call
//   - ctx context.Context
//   - detailModel *models.AnimeDetailModel
func (_e *AnimeExtractor_Expecter) Resolve(ctx interface{}, detailModel interface{}) *AnimeExtractor_Resolve_Call {
	return &AnimeExtractor_Resolve_Call{Call: _e.mock.On("Resolve", ctx, detailModel)}
}

func (_c *AnimeExtractor_Resolve_Call) Run(run func(ctx context.Context, detailModel *models.AnimeDetailModel)) *AnimeExtractor_Resolve_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.AnimeDetailModel))
	})
	return _c
}

func (_c *AnimeExtractor_Resolve_Call) Return(_a0 error) *AnimeExtractor_Resolve_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AnimeExtractor_Resolve_Call) RunAndReturn(run func(context.Context, *models.AnimeDetailModel) error) *AnimeExtractor_Resolve_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewAnimeExtractor interface {
	mock.TestingT
	Cleanup(func())
}

// NewAnimeExtractor creates a new instance of AnimeExtractor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAnimeExtractor(t mockConstructorTestingTNewAnimeExtractor) *AnimeExtractor {
	mock := &AnimeExtractor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}