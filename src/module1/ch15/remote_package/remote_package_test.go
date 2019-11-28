package remote_package

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestRemotePackage(t *testing.T) {
	concurrentMap := cm.CreateConcurrentMap(10)
	concurrentMap.Set(cm.StrKey("key"), 1)
	t.Log(concurrentMap.Get(cm.StrKey("key")))
}

//glide init
//glide install
