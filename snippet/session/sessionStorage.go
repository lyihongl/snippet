package session

import (
	//"container/list"
	"time"
)

//var pder = &Provider{list: list.New()}

type SessionStore struct {
	sid          string
	timeAccessed time.Time
	value        map[interface{}]interface{}
}

//func(st *SessionStore) Set(key, value, interface{}) error {
//	st.value[key] = value
//	pder.SessionUpdate(st.sid)
//}
