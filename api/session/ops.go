package session

import (
	"sync"
	"time"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/utils"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionFromDB() {
	allSession, err := dbops.RetrieveAllSession()
	if err != nil {
		return
	}

	allSession.Range(func(key, value interface{}) bool {
		session := value.(*defs.SimpleSession)
		sessionMap.Store(key, session)
		return true
	})
}

func GenerateNewSessionId(uname string) string {
	id, _ := utils.NewUUID()
	ctime := nowInMilli()
	ttl := ctime + 30*60*1000

	session := &defs.SimpleSession{UserName: uname, TTL: ttl}
	sessionMap.Store(id, session)
	dbops.InsertSession(id, ttl, uname)
	return id
}

func IsSessionExpired(session_id string) (string, bool) {
	value, ok := sessionMap.Load(session_id)
	if ok {
		if value.(*defs.SimpleSession).TTL < nowInMilli() {
			deleteExpiredSession(session_id)
			return "", true
		}
		return value.(*defs.SimpleSession).UserName, false
	}
	return "", true
}
