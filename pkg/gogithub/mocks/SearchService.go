// Code generated by mockery v1.0.0. DO NOT EDIT.

// If you want to rebuild this file, make mock-pkg

package mocks

import (
	context "context"

	github "github.com/google/go-github/github"

	mock "github.com/stretchr/testify/mock"
)

// SearchService is an autogenerated mock type for the SearchService type
type SearchService struct {
	mock.Mock
}

// Issues provides a mock function with given fields: ctx, query, opt
func (_m *SearchService) Issues(ctx context.Context, query string, opt *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error) {
	ret := _m.Called(ctx, query, opt)

	var r0 *github.IssuesSearchResult
	if rf, ok := ret.Get(0).(func(context.Context, string, *github.SearchOptions) *github.IssuesSearchResult); ok {
		r0 = rf(ctx, query, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.IssuesSearchResult)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, *github.SearchOptions) *github.Response); ok {
		r1 = rf(ctx, query, opt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, *github.SearchOptions) error); ok {
		r2 = rf(ctx, query, opt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
