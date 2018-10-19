// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/softlayer/softlayer-go/datatypes"
	"github.ibm.com/alchemy-containers/armada-cluster/iaas/softlayer/backend"
)

type NetworkVlanService struct {
	FilterStub        func(string) backend.NetworkVlanService
	filterMutex       sync.RWMutex
	filterArgsForCall []struct {
		arg1 string
	}
	filterReturns struct {
		result1 backend.NetworkVlanService
	}
	filterReturnsOnCall map[int]struct {
		result1 backend.NetworkVlanService
	}
	MaskStub        func(string) backend.NetworkVlanService
	maskMutex       sync.RWMutex
	maskArgsForCall []struct {
		arg1 string
	}
	maskReturns struct {
		result1 backend.NetworkVlanService
	}
	maskReturnsOnCall map[int]struct {
		result1 backend.NetworkVlanService
	}
	IDStub        func(int) backend.NetworkVlanService
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
		arg1 int
	}
	iDReturns struct {
		result1 backend.NetworkVlanService
	}
	iDReturnsOnCall map[int]struct {
		result1 backend.NetworkVlanService
	}
	GetNetworkSpaceStub        func() (string, error)
	getNetworkSpaceMutex       sync.RWMutex
	getNetworkSpaceArgsForCall []struct{}
	getNetworkSpaceReturns     struct {
		result1 string
		result2 error
	}
	getNetworkSpaceReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetVlanStub        func() (datatypes.Network_Vlan, error)
	getVlanMutex       sync.RWMutex
	getVlanArgsForCall []struct{}
	getVlanReturns     struct {
		result1 datatypes.Network_Vlan
		result2 error
	}
	getVlanReturnsOnCall map[int]struct {
		result1 datatypes.Network_Vlan
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NetworkVlanService) Filter(arg1 string) backend.NetworkVlanService {
	fake.filterMutex.Lock()
	ret, specificReturn := fake.filterReturnsOnCall[len(fake.filterArgsForCall)]
	fake.filterArgsForCall = append(fake.filterArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Filter", []interface{}{arg1})
	fake.filterMutex.Unlock()
	if fake.FilterStub != nil {
		return fake.FilterStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.filterReturns.result1
}

func (fake *NetworkVlanService) FilterCallCount() int {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	return len(fake.filterArgsForCall)
}

func (fake *NetworkVlanService) FilterArgsForCall(i int) string {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	return fake.filterArgsForCall[i].arg1
}

func (fake *NetworkVlanService) FilterReturns(result1 backend.NetworkVlanService) {
	fake.FilterStub = nil
	fake.filterReturns = struct {
		result1 backend.NetworkVlanService
	}{result1}
}

func (fake *NetworkVlanService) FilterReturnsOnCall(i int, result1 backend.NetworkVlanService) {
	fake.FilterStub = nil
	if fake.filterReturnsOnCall == nil {
		fake.filterReturnsOnCall = make(map[int]struct {
			result1 backend.NetworkVlanService
		})
	}
	fake.filterReturnsOnCall[i] = struct {
		result1 backend.NetworkVlanService
	}{result1}
}

func (fake *NetworkVlanService) Mask(arg1 string) backend.NetworkVlanService {
	fake.maskMutex.Lock()
	ret, specificReturn := fake.maskReturnsOnCall[len(fake.maskArgsForCall)]
	fake.maskArgsForCall = append(fake.maskArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Mask", []interface{}{arg1})
	fake.maskMutex.Unlock()
	if fake.MaskStub != nil {
		return fake.MaskStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.maskReturns.result1
}

func (fake *NetworkVlanService) MaskCallCount() int {
	fake.maskMutex.RLock()
	defer fake.maskMutex.RUnlock()
	return len(fake.maskArgsForCall)
}

func (fake *NetworkVlanService) MaskArgsForCall(i int) string {
	fake.maskMutex.RLock()
	defer fake.maskMutex.RUnlock()
	return fake.maskArgsForCall[i].arg1
}

func (fake *NetworkVlanService) MaskReturns(result1 backend.NetworkVlanService) {
	fake.MaskStub = nil
	fake.maskReturns = struct {
		result1 backend.NetworkVlanService
	}{result1}
}

func (fake *NetworkVlanService) MaskReturnsOnCall(i int, result1 backend.NetworkVlanService) {
	fake.MaskStub = nil
	if fake.maskReturnsOnCall == nil {
		fake.maskReturnsOnCall = make(map[int]struct {
			result1 backend.NetworkVlanService
		})
	}
	fake.maskReturnsOnCall[i] = struct {
		result1 backend.NetworkVlanService
	}{result1}
}

func (fake *NetworkVlanService) ID(arg1 int) backend.NetworkVlanService {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("ID", []interface{}{arg1})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.iDReturns.result1
}

func (fake *NetworkVlanService) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *NetworkVlanService) IDArgsForCall(i int) int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return fake.iDArgsForCall[i].arg1
}

func (fake *NetworkVlanService) IDReturns(result1 backend.NetworkVlanService) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 backend.NetworkVlanService
	}{result1}
}

func (fake *NetworkVlanService) IDReturnsOnCall(i int, result1 backend.NetworkVlanService) {
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 backend.NetworkVlanService
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 backend.NetworkVlanService
	}{result1}
}

func (fake *NetworkVlanService) GetNetworkSpace() (string, error) {
	fake.getNetworkSpaceMutex.Lock()
	ret, specificReturn := fake.getNetworkSpaceReturnsOnCall[len(fake.getNetworkSpaceArgsForCall)]
	fake.getNetworkSpaceArgsForCall = append(fake.getNetworkSpaceArgsForCall, struct{}{})
	fake.recordInvocation("GetNetworkSpace", []interface{}{})
	fake.getNetworkSpaceMutex.Unlock()
	if fake.GetNetworkSpaceStub != nil {
		return fake.GetNetworkSpaceStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getNetworkSpaceReturns.result1, fake.getNetworkSpaceReturns.result2
}

func (fake *NetworkVlanService) GetNetworkSpaceCallCount() int {
	fake.getNetworkSpaceMutex.RLock()
	defer fake.getNetworkSpaceMutex.RUnlock()
	return len(fake.getNetworkSpaceArgsForCall)
}

func (fake *NetworkVlanService) GetNetworkSpaceReturns(result1 string, result2 error) {
	fake.GetNetworkSpaceStub = nil
	fake.getNetworkSpaceReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *NetworkVlanService) GetNetworkSpaceReturnsOnCall(i int, result1 string, result2 error) {
	fake.GetNetworkSpaceStub = nil
	if fake.getNetworkSpaceReturnsOnCall == nil {
		fake.getNetworkSpaceReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getNetworkSpaceReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *NetworkVlanService) GetVlan() (datatypes.Network_Vlan, error) {
	fake.getVlanMutex.Lock()
	ret, specificReturn := fake.getVlanReturnsOnCall[len(fake.getVlanArgsForCall)]
	fake.getVlanArgsForCall = append(fake.getVlanArgsForCall, struct{}{})
	fake.recordInvocation("GetVlan", []interface{}{})
	fake.getVlanMutex.Unlock()
	if fake.GetVlanStub != nil {
		return fake.GetVlanStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getVlanReturns.result1, fake.getVlanReturns.result2
}

func (fake *NetworkVlanService) GetVlanCallCount() int {
	fake.getVlanMutex.RLock()
	defer fake.getVlanMutex.RUnlock()
	return len(fake.getVlanArgsForCall)
}

func (fake *NetworkVlanService) GetVlanReturns(result1 datatypes.Network_Vlan, result2 error) {
	fake.GetVlanStub = nil
	fake.getVlanReturns = struct {
		result1 datatypes.Network_Vlan
		result2 error
	}{result1, result2}
}

func (fake *NetworkVlanService) GetVlanReturnsOnCall(i int, result1 datatypes.Network_Vlan, result2 error) {
	fake.GetVlanStub = nil
	if fake.getVlanReturnsOnCall == nil {
		fake.getVlanReturnsOnCall = make(map[int]struct {
			result1 datatypes.Network_Vlan
			result2 error
		})
	}
	fake.getVlanReturnsOnCall[i] = struct {
		result1 datatypes.Network_Vlan
		result2 error
	}{result1, result2}
}

func (fake *NetworkVlanService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	fake.maskMutex.RLock()
	defer fake.maskMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.getNetworkSpaceMutex.RLock()
	defer fake.getNetworkSpaceMutex.RUnlock()
	fake.getVlanMutex.RLock()
	defer fake.getVlanMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *NetworkVlanService) recordInvocation(key string, args []interface{}) {
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

var _ backend.NetworkVlanService = new(NetworkVlanService)
