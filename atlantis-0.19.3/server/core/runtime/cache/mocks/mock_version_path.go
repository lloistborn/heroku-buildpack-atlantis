// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/runtime/cache (interfaces: ExecutionVersionCache)

package mocks

import (
	"reflect"
	"time"

	go_version "github.com/hashicorp/go-version"
	pegomock "github.com/petergtz/pegomock"
)

type MockExecutionVersionCache struct {
	fail func(message string, callerSkip ...int)
}

func NewMockExecutionVersionCache(options ...pegomock.Option) *MockExecutionVersionCache {
	mock := &MockExecutionVersionCache{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockExecutionVersionCache) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockExecutionVersionCache) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockExecutionVersionCache) Get(key *go_version.Version) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockExecutionVersionCache().")
	}
	params := []pegomock.Param{key}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Get", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockExecutionVersionCache) VerifyWasCalledOnce() *VerifierMockExecutionVersionCache {
	return &VerifierMockExecutionVersionCache{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockExecutionVersionCache) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockExecutionVersionCache {
	return &VerifierMockExecutionVersionCache{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockExecutionVersionCache) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockExecutionVersionCache {
	return &VerifierMockExecutionVersionCache{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockExecutionVersionCache) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockExecutionVersionCache {
	return &VerifierMockExecutionVersionCache{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockExecutionVersionCache struct {
	mock                   *MockExecutionVersionCache
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockExecutionVersionCache) Get(key *go_version.Version) *MockExecutionVersionCache_Get_OngoingVerification {
	params := []pegomock.Param{key}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Get", params, verifier.timeout)
	return &MockExecutionVersionCache_Get_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockExecutionVersionCache_Get_OngoingVerification struct {
	mock              *MockExecutionVersionCache
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockExecutionVersionCache_Get_OngoingVerification) GetCapturedArguments() *go_version.Version {
	key := c.GetAllCapturedArguments()
	return key[len(key)-1]
}

func (c *MockExecutionVersionCache_Get_OngoingVerification) GetAllCapturedArguments() (_param0 []*go_version.Version) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*go_version.Version, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(*go_version.Version)
		}
	}
	return
}
