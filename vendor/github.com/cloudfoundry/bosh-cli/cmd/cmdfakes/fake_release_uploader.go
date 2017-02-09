// This file was generated by counterfeiter
package cmdfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/cmd"
)

type FakeReleaseUploader struct {
	UploadReleasesStub        func([]byte) ([]byte, error)
	uploadReleasesMutex       sync.RWMutex
	uploadReleasesArgsForCall []struct {
		arg1 []byte
	}
	uploadReleasesReturns struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReleaseUploader) UploadReleases(arg1 []byte) ([]byte, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.uploadReleasesMutex.Lock()
	fake.uploadReleasesArgsForCall = append(fake.uploadReleasesArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("UploadReleases", []interface{}{arg1Copy})
	fake.uploadReleasesMutex.Unlock()
	if fake.UploadReleasesStub != nil {
		return fake.UploadReleasesStub(arg1)
	}
	return fake.uploadReleasesReturns.result1, fake.uploadReleasesReturns.result2
}

func (fake *FakeReleaseUploader) UploadReleasesCallCount() int {
	fake.uploadReleasesMutex.RLock()
	defer fake.uploadReleasesMutex.RUnlock()
	return len(fake.uploadReleasesArgsForCall)
}

func (fake *FakeReleaseUploader) UploadReleasesArgsForCall(i int) []byte {
	fake.uploadReleasesMutex.RLock()
	defer fake.uploadReleasesMutex.RUnlock()
	return fake.uploadReleasesArgsForCall[i].arg1
}

func (fake *FakeReleaseUploader) UploadReleasesReturns(result1 []byte, result2 error) {
	fake.UploadReleasesStub = nil
	fake.uploadReleasesReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseUploader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.uploadReleasesMutex.RLock()
	defer fake.uploadReleasesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeReleaseUploader) recordInvocation(key string, args []interface{}) {
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

var _ cmd.ReleaseUploader = new(FakeReleaseUploader)
