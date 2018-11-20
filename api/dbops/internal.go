package dbops

import (
	"git.umlife.net/gateway/openstat/log"
	"strconv"
	"sync"
	"w3u72y/video_server_1_5/api/defs"
)

func InsertSession(session_id string, ttl string, login_name string) error {
	stmt, err := dbConn.Prepare("INSERT into sessions(session_id,TTL,login_name) " +
		"VALUES (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(session_id, ttl, login_name)
	if err != nil {
		return err
	}

	return nil
}

func RetrieveSession(session_id string) (*defs.SimpleSession, error) {
	session := &defs.SimpleSession{}
	stmt, err := dbConn.Prepare("SELECT TTL,login_name FROM sessions where session_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var login_name string
	var ttl string
	err = stmt.QueryRow(session_id).Scan(&ttl, &login_name)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		session.TTL = res
		session.Username = login_name
		return session, nil
	}
	return nil, err
}

func RetrieveAllSession() (*sync.Map, error) {
	m := &sync.Map{}
	stmt, err := dbConn.Prepare("SELECT session_id,TTL,login_name FROM")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var sid string
	var ttl string
	var login_name string
	for rows.Next() {
		if err := rows.Scan(&sid, &ttl, &login_name); err != nil {
			log.Error(err)
			break
		}

		if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
			session := &defs.SimpleSession{Username: login_name, TTL: res}
			m.Store(sid, session)
			log.Printf("session id:%d, ttl: %d", sid, session.TTL)
		}
	}
	return m, nil
}

func DeleteSession(session_id string) error {
	stmt, err := dbConn.Prepare("DELETE FROM sessions where session_id = ?")
	if err != nil {
		log.Error(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(session_id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
