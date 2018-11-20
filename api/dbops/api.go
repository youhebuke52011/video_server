package dbops

import (
	"database/sql"
	"git.umlife.net/gateway/openstat/log"
	"time"
	"video_server/api/utils"
	"w3u72y/video_server_1_5/api/defs"
)

func AddUserCredential(loginName string, password string) error {
	stmt, err := dbConn.Prepare("INSERT into users (login_name, pwd) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	stmt.Exec(loginName, password)
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmt, err := dbConn.Prepare("SELECT pwd FROM users where login_name = ?")
	if err != nil {
		log.Error(err)
		return "", err
	}
	defer stmt.Close()

	var password string
	err = stmt.QueryRow(loginName).Scan(&password)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	return password, nil

}

func DeleteUser(loginName string, pwd string) error {
	stmt, err := dbConn.Prepare("DELETE from user where login_name = ? AND pwd = ?")
	if err != nil {
		log.Error(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(loginName, pwd)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func AddNewVedio(aid int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Mon Jan 02 2006 15:04:05")

	stmt, err := dbConn.Prepare("INSERT INTO video_info(id,author_id,name,display_ctime)" +
		" VALUES (?,?,?,?)")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	videoInfo := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	return videoInfo, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmt, err := dbConn.Prepare("SELECT author_id,name,display_ctime FROM video_info where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var aid int
	var name string
	var dctime string
	err = stmt.QueryRow(vid).Scan(&aid, &name, &dctime)
	if err != nil {
		return nil, err
	}

	videoInfo := &defs.VideoInfo{AuthorId: aid, Name: name, DisplayCtime: dctime}
	return videoInfo, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmt, err := dbConn.Prepare("SELECT comments.id,users.name,comments.content from comments " +
		"INNER JOIN users ON users.id = comments.author_id where vedio_id = ? and time > " +
		"FROM_UNIXTIME(?) and time <= FROM_UNIXTIME(?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(vid, from, to)
	if err != nil {
		return nil, err
	}

	var cid, uname, content string
	var res []*defs.Comment

	for rows.Next() {
		err = rows.Scan(&cid, &uname, &content)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		comment := &defs.Comment{Id: cid, Author: uname, Content: content}
		res = append(res, comment)
	}
	return res, nil
}
