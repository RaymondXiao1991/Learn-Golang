package mapbench

import "sync"

type MyMap struct {
	sync.Mutex
	m map[int]int
}

type MyRwMap struct {
	sync.RWMutex
	m map[int]int
}

var myMap *MyMap
var myRwMap *MyRwMap
var syncMap *sync.Map

func init() {
	myMap = &MyMap{
		m: make(map[int]int, 100),
	}
	myRwMap = &MyRwMap{
		m: make(map[int]int, 100),
	}
	syncMap = &sync.Map{}
}

func builtinMapStore(k, v int) {
	myMap.Lock()
	defer myMap.Unlock()

	myMap.m[k] = v
}

func builtinMapLookup(k int) int {
	myMap.Lock()
	defer myMap.Unlock()

	if v, ok := myMap.m[k]; !ok {
		return -1
	} else {
		return v
	}
}

func builtinMapDelete(k int) {
	myMap.Lock()
	defer myMap.Unlock()

	if v, ok := myMap.m[k]; !ok {
		return
	} else {
		delete(myMap.m, v)
	}
}

func builtinRwMapStore(k, v int) {
	myRwMap.Lock()
	defer myRwMap.Unlock()

	myRwMap.m[k] = v
}

func builtinRwMapLookup(k int) int {
	myRwMap.RLock()
	defer myRwMap.RUnlock()

	if v, ok := myRwMap.m[k]; !ok {
		return -1
	} else {
		return v
	}
}

func builtinRwMapDelete(k int) {
	myRwMap.Lock()
	defer myRwMap.Unlock()

	if v, ok := myRwMap.m[k]; !ok {
		return
	} else {
		delete(myRwMap.m, v)
	}
}

func syncMapStore(k, v int) {
	syncMap.Store(k, v)
}

func syncMapLookup(k int) int {
	if v, ok := syncMap.Load(k); !ok {
		return -1
	} else {
		return v.(int)
	}
}

func syncMapDelete(k int) {
	syncMap.Delete(k)
}
