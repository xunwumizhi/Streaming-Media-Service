package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"Streaming-Media-Service/api/defs"
	"Streaming-Media-Service/api/utils"
	"time"
)



func AddUser(loginName string, pwd string) error {

	stmtIn, err := dbConn.Prepare("INSERT INTO users(login_name, pwd) values(?, ?)")  //本包中的变量dbConn
	if err != nil {
		return err
	}
	_, err = stmtIn.Exec(loginName, pwd) //注意err类型已由前面确定了，这里不需要使用`:=`，否则出现编译错误`no new variables on left side of :=`
	if err != nil {
		return err
	}

	defer stmtIn.Close()
	return nil
}

//获取信息
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("SELECT error: %s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmtOut.Close()
	return pwd, nil
}

//返回实体
func GetUser(loginName string) (*defs.User, error) {
	stmtOut, err := dbConn.Prepare("SELECT id, pwd FROM users WHERE login_name = ?")
	defer stmtOut.Close()
	if err != nil {
		log.Printf("SELECT error: %s", err)
	}
	var (
		id int
		pwd string
	)
	err = stmtOut.QueryRow(loginName).Scan(&id, &pwd)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	res := &defs.User{Id:id, LoginName:loginName, Pwd:pwd}
	return res, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Printf("DELETE error: %s", err)
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

//video操作
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	ctime := time.Now().Format("Jan 02 2006, 15:04:05") 

	stmtIn, errDB := dbConn.Prepare("INSERT INTO video_info(id,author_id,name,display_ctime,create_time) VALUES(?,?,?,?,now())")
	defer stmtIn.Close()
	if errDB != nil {
		return nil, errDB
	}
	_, err = stmtIn.Exec(vid, aid, name, ctime)
	if err!= nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id:vid, AuthorId:aid, Name:name, DisplayCtime:ctime}
	return res, nil

}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id = ?")
	defer stmtOut.Close()
	if err != nil {
		return nil, err
	}
	var (
		aid int
		name string
		ctime string
	)
	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &ctime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	res := &defs.VideoInfo{Id:vid, AuthorId:aid, Name:name, DisplayCtime:ctime}
	return res, nil
}

// func ListVideoInfo(uname string, from, to int) ([]*defs.VideoInfo, error) {
// 	stmtOut, err := dbConn.Prepare(`SELECT video_info.id, video_info.author_id, video_info.name, video_info.display_ctime FROM video_info
// 		INNER JOIN users ON video_info.author_id = users.id
// 		WHERE users.login_name=? AND video_info.create_time > FROM_UNIXTIME(?) AND video_info.create_time<=FROM_UNIXTIME(?)
// 		ORDER BY video_info.create_time DESC`)
func ListVideoInfo(uname string) ([]*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare(`SELECT video_info.id, video_info.author_id, video_info.name, video_info.display_ctime FROM video_info
		INNER JOIN users ON video_info.author_id = users.id
		WHERE users.login_name=?
		ORDER BY video_info.create_time DESC`)
	defer stmtOut.Close()

	var res []*defs.VideoInfo
	if err != nil {
		return res, err
	}
	// rows, errDB := stmtOut.Query(uname, from, to)
	rows, errDB := stmtOut.Query(uname)
	if errDB != nil {
		log.Printf("SELECT error: %s", errDB)
		return res, errDB
	}
	for rows.Next() {
		var (
			vid, name, ctime string
			aid int
		)

		if err := rows.Scan(&vid, &aid, &name, &ctime); err != nil {
			return res, err
		}
		one := &defs.VideoInfo{Id:vid, AuthorId:aid, Name:name, DisplayCtime:ctime}
		res = append(res, one)

	}
	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	defer stmtDel.Close()
	if err!=nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err!=nil {
		return err
	}
	return nil
}


//comments操作
func AddNewComment(vid string, aid int, content string) error {
	stmtIn, err := dbConn.Prepare("INSERT INTO comments(id, video_id, author_id, content, time) values(?,?,?,?,now())")
	defer stmtIn.Close()
	if err!=nil {
		return err
	}

	id, errID := utils.NewUUID()
	if errID!= nil {
		return errID
	}

	_, err = stmtIn.Exec(id, vid, aid, content)
	if err!=nil {
		return err
	}

	return nil
}

// func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
// 	stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name, comments.content FROM comments 
// 		INNER JOIN users ON comments.author_id = users.id
// 		WHERE comments.video_id = ? AND comments.time >  FROM_UNIXTIME(?) AND comments.time<=FROM_UNIXTIME(?)
// 		OREDER BY comments.time DESC`)
func ListComments(vid string) ([]*defs.Comment, error) {
	// stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name, comments.content FROM comments 
	// 	INNER JOIN users ON comments.author_id = users.id
	// 	WHERE comments.video_id = ?
	// 	ORDER BY comments.time DESC`)
	stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name, comments.content FROM comments 
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ?
		ORDER BY comments.time DESC`)
	defer stmtOut.Close()
	
	var res []*defs.Comment
	if err!=nil {
		return res, err
	}
	// rows, errDB := stmtOut.Query(vid, from, to)
	rows, errDB := stmtOut.Query(vid)
	if errDB!= nil {
		return res, errDB
	}

	for rows.Next() {
		var id, name, content string
		if err:=rows.Scan(&id, &name, &content); err!=nil {
			return res, err
		}
		one := &defs.Comment{Id:id, VideoId:vid, Author:name, Content:content}
		res = append(res, one)
	}
	return res, nil
}