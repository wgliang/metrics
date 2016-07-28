/*
	this is s struct for event-flow, and you can define it's size
*/
package poetry

import (
	"container/list"
	"sync"
	"errors"
)

type Flow interface {
	// Clear all flows.
	Clear()

	// Return the size of the flows, which is at most the reservoir size.
	Size() int64

	// Return all the values in the flow.
	Values() (values []string)

	// right pop and push operation
	RPop()	(string, error)
	RPush(string)
	// left pop and push operation
	LPop()	(string, error)
	LPush(string)
}

type flowData struct {
	// golang's list
	values *list.List
	// just from Concurrent operation
	mutex sync.Mutex
	// flow's max length
	Maxsize int64
	// Concurrent flow's length
	Length int64
	// callback function for user
	OnEvicted func()
}

// New a flow
func NewFlow(size int64) Flow {
	f := &flowData{
		values:		list.New(),
		Maxsize: 	size,
		Length:		0,
		OnEvicted: 	func() {},
	}

	return f
}

// Push data from right
func (f *flowData) RPush(value string) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.values.PushFront(value)
	f.Length += 1

	for {
		if f.Maxsize != 0 && f.Length > f.Maxsize {
			ele := f.values.Back()
			f.Length -= 1
			f.values.Remove(ele)
			if f.OnEvicted != nil {
				f.OnEvicted()
			}
		} else {
			break
		}
	}
}

// Push data from left
func (f *flowData) LPush(value string) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.values.PushBack(value)
	f.Length += 1

	for {
		if f.Maxsize != 0 && f.Length > f.Maxsize {
			ele := f.values.Front()
			f.Length -= 1
			f.values.Remove(ele)
			if f.OnEvicted != nil {
				f.OnEvicted()
			}
		} else {
			break
		}
	}
}

// get all values in list
func (f *flowData) Values() (values []string){
	for iter := f.values.Front(); iter != nil ;iter = iter.Next() {
    	values =  append(values,(iter.Value).(string))
	}
	return values
}

// pop data from left
func (f *flowData) LPop() (string, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	ele := f.values.Back()
	if ele == nil {
		return "", errors.New("flow: null list")
	}
	f.values.Remove(ele)
	f.Length -= 1
	return ele.Value.(string), nil
}

//pop data from right
func (f *flowData) RPop() (string, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	ele := f.values.Front()
	if ele == nil {
		return "", errors.New("flow: null list")
	}
	f.values.Remove(ele)
	f.Length -= 1
	return ele.Value.(string), nil
}

// get current size
func (f *flowData) Size() int64{
	return f.Length
}

// Clear list flow
func (f *flowData) Clear() {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.values.Init()
	f.Length = 0

	return 
}