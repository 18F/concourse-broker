// This file was generated by counterfeiter
package volumeserverfakes

import (
	"sync"

	"github.com/concourse/atc/api/volumeserver"
	"github.com/concourse/atc/db"
)

type FakeVolumesDB struct {
	GetVolumesStub        func() ([]db.SavedVolume, error)
	getVolumesMutex       sync.RWMutex
	getVolumesArgsForCall []struct{}
	getVolumesReturns     struct {
		result1 []db.SavedVolume
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolumesDB) GetVolumes() ([]db.SavedVolume, error) {
	fake.getVolumesMutex.Lock()
	fake.getVolumesArgsForCall = append(fake.getVolumesArgsForCall, struct{}{})
	fake.recordInvocation("GetVolumes", []interface{}{})
	fake.getVolumesMutex.Unlock()
	if fake.GetVolumesStub != nil {
		return fake.GetVolumesStub()
	} else {
		return fake.getVolumesReturns.result1, fake.getVolumesReturns.result2
	}
}

func (fake *FakeVolumesDB) GetVolumesCallCount() int {
	fake.getVolumesMutex.RLock()
	defer fake.getVolumesMutex.RUnlock()
	return len(fake.getVolumesArgsForCall)
}

func (fake *FakeVolumesDB) GetVolumesReturns(result1 []db.SavedVolume, result2 error) {
	fake.GetVolumesStub = nil
	fake.getVolumesReturns = struct {
		result1 []db.SavedVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumesDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getVolumesMutex.RLock()
	defer fake.getVolumesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeVolumesDB) recordInvocation(key string, args []interface{}) {
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

var _ volumeserver.VolumesDB = new(FakeVolumesDB)
