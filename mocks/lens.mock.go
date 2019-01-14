// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	service "github.com/RTradeLtd/grpc/lens"
	"github.com/RTradeLtd/grpc/lens/request"
	"github.com/RTradeLtd/grpc/lens/response"
	"google.golang.org/grpc"
)

type FakeIndexerAPIClient struct {
	IndexStub        func(context.Context, *request.Index, ...grpc.CallOption) (*response.Index, error)
	indexMutex       sync.RWMutex
	indexArgsForCall []struct {
		arg1 context.Context
		arg2 *request.Index
		arg3 []grpc.CallOption
	}
	indexReturns struct {
		result1 *response.Index
		result2 error
	}
	indexReturnsOnCall map[int]struct {
		result1 *response.Index
		result2 error
	}
	SearchStub        func(context.Context, *request.Search, ...grpc.CallOption) (*response.Results, error)
	searchMutex       sync.RWMutex
	searchArgsForCall []struct {
		arg1 context.Context
		arg2 *request.Search
		arg3 []grpc.CallOption
	}
	searchReturns struct {
		result1 *response.Results
		result2 error
	}
	searchReturnsOnCall map[int]struct {
		result1 *response.Results
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIndexerAPIClient) Index(arg1 context.Context, arg2 *request.Index, arg3 ...grpc.CallOption) (*response.Index, error) {
	fake.indexMutex.Lock()
	ret, specificReturn := fake.indexReturnsOnCall[len(fake.indexArgsForCall)]
	fake.indexArgsForCall = append(fake.indexArgsForCall, struct {
		arg1 context.Context
		arg2 *request.Index
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("Index", []interface{}{arg1, arg2, arg3})
	fake.indexMutex.Unlock()
	if fake.IndexStub != nil {
		return fake.IndexStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.indexReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIndexerAPIClient) IndexCallCount() int {
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	return len(fake.indexArgsForCall)
}

func (fake *FakeIndexerAPIClient) IndexCalls(stub func(context.Context, *request.Index, ...grpc.CallOption) (*response.Index, error)) {
	fake.indexMutex.Lock()
	defer fake.indexMutex.Unlock()
	fake.IndexStub = stub
}

func (fake *FakeIndexerAPIClient) IndexArgsForCall(i int) (context.Context, *request.Index, []grpc.CallOption) {
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	argsForCall := fake.indexArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIndexerAPIClient) IndexReturns(result1 *response.Index, result2 error) {
	fake.indexMutex.Lock()
	defer fake.indexMutex.Unlock()
	fake.IndexStub = nil
	fake.indexReturns = struct {
		result1 *response.Index
		result2 error
	}{result1, result2}
}

func (fake *FakeIndexerAPIClient) IndexReturnsOnCall(i int, result1 *response.Index, result2 error) {
	fake.indexMutex.Lock()
	defer fake.indexMutex.Unlock()
	fake.IndexStub = nil
	if fake.indexReturnsOnCall == nil {
		fake.indexReturnsOnCall = make(map[int]struct {
			result1 *response.Index
			result2 error
		})
	}
	fake.indexReturnsOnCall[i] = struct {
		result1 *response.Index
		result2 error
	}{result1, result2}
}

func (fake *FakeIndexerAPIClient) Search(arg1 context.Context, arg2 *request.Search, arg3 ...grpc.CallOption) (*response.Results, error) {
	fake.searchMutex.Lock()
	ret, specificReturn := fake.searchReturnsOnCall[len(fake.searchArgsForCall)]
	fake.searchArgsForCall = append(fake.searchArgsForCall, struct {
		arg1 context.Context
		arg2 *request.Search
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("Search", []interface{}{arg1, arg2, arg3})
	fake.searchMutex.Unlock()
	if fake.SearchStub != nil {
		return fake.SearchStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.searchReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIndexerAPIClient) SearchCallCount() int {
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	return len(fake.searchArgsForCall)
}

func (fake *FakeIndexerAPIClient) SearchCalls(stub func(context.Context, *request.Search, ...grpc.CallOption) (*response.Results, error)) {
	fake.searchMutex.Lock()
	defer fake.searchMutex.Unlock()
	fake.SearchStub = stub
}

func (fake *FakeIndexerAPIClient) SearchArgsForCall(i int) (context.Context, *request.Search, []grpc.CallOption) {
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	argsForCall := fake.searchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIndexerAPIClient) SearchReturns(result1 *response.Results, result2 error) {
	fake.searchMutex.Lock()
	defer fake.searchMutex.Unlock()
	fake.SearchStub = nil
	fake.searchReturns = struct {
		result1 *response.Results
		result2 error
	}{result1, result2}
}

func (fake *FakeIndexerAPIClient) SearchReturnsOnCall(i int, result1 *response.Results, result2 error) {
	fake.searchMutex.Lock()
	defer fake.searchMutex.Unlock()
	fake.SearchStub = nil
	if fake.searchReturnsOnCall == nil {
		fake.searchReturnsOnCall = make(map[int]struct {
			result1 *response.Results
			result2 error
		})
	}
	fake.searchReturnsOnCall[i] = struct {
		result1 *response.Results
		result2 error
	}{result1, result2}
}

func (fake *FakeIndexerAPIClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIndexerAPIClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ service.IndexerAPIClient = new(FakeIndexerAPIClient)
