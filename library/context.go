package context

import (
	"sync"
	"time"
	"github.com/ant0ine/go-json-rest/rest"
)

var (
	mutex sync.RWMutex
	data  = make(map[*rest.Request]map[interface{}]interface{})
	datat = make(map[*rest.Request]int64)
)

// Set stores a value for a given key in a given request.
func Set(r *rest.Request, key, val interface{}) {
	mutex.Lock()
	if data[r] == nil {
		data[r] = make(map[interface{}]interface{})
		datat[r] = time.Now().Unix()
	}
	data[r][key] = val
	mutex.Unlock()
}

// Get returns a value stored for a given key in a given request.
func Get(r *rest.Request, key interface{}) interface{} {
	mutex.RLock()
	if ctx := data[r]; ctx != nil {
		value := ctx[key]
		mutex.RUnlock()
		return value
	}
	mutex.RUnlock()
	return nil
}

// GetOk returns stored value and presence state like multi-value return of map access.
func GetOk(r *rest.Request, key interface{}) (interface{}, bool) {
	mutex.RLock()
	if _, ok := data[r]; ok {
		value, ok := data[r][key]
		mutex.RUnlock()
		return value, ok
	}
	mutex.RUnlock()
	return nil, false
}

// GetAll returns all stored values for the request as a map. Nil is returned for invalid requests.
func GetAll(r *rest.Request) map[interface{}]interface{} {
	mutex.RLock()
	if context, ok := data[r]; ok {
		result := make(map[interface{}]interface{}, len(context))
		for k, v := range context {
			result[k] = v
		}
		mutex.RUnlock()
		return result
	}
	mutex.RUnlock()
	return nil
}

// GetAllOk returns all stored values for the request as a map and a boolean value that indicates if
// the request was registered.
func GetAllOk(r *rest.Request) (map[interface{}]interface{}, bool) {
	mutex.RLock()
	context, ok := data[r]
	result := make(map[interface{}]interface{}, len(context))
	for k, v := range context {
		result[k] = v
	}
	mutex.RUnlock()
	return result, ok
}

// Delete removes a value stored for a given key in a given request.
func Delete(r *rest.Request, key interface{}) {
	mutex.Lock()
	if data[r] != nil {
		delete(data[r], key)
	}
	mutex.Unlock()
}

// Clear removes all values stored for a given request.
//
// This is usually called by a handler wrapper to clean up request
// variables at the end of a request lifetime. See ClearHandler().
func Clear(r *rest.Request) {
	mutex.Lock()
	clear(r)
	mutex.Unlock()
}

// clear is Clear without the lock.
func clear(r *rest.Request) {
	delete(data, r)
	delete(datat, r)
}

// Purge removes request data stored for longer than maxAge, in seconds.
// It returns the amount of requests removed.
//
// If maxAge <= 0, all request data is removed.
//
// This is only used for sanity check: in case context cleaning was not
// properly set some request data can be kept forever, consuming an increasing
// amount of memory. In case this is detected, Purge() must be called
// periodically until the problem is fixed.
func Purge(maxAge int) int {
	mutex.Lock()
	count := 0
	if maxAge <= 0 {
		count = len(data)
		data = make(map[*rest.Request]map[interface{}]interface{})
		datat = make(map[*rest.Request]int64)
	} else {
		min := time.Now().Unix() - int64(maxAge)
		for r := range data {
			if datat[r] < min {
				clear(r)
				count++
			}
		}
	}
	mutex.Unlock()
	return count
}
