package dbops
import (
	"testing"
	"fmt"
	"Streaming-Media-Service/api/utils"
	// "strconv"
	// "time"
)

//清理数据库
func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

//测试主函数
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

//User--测试流安排
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("GetCredential", testGetUserCredential)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
}

//使用了TestMain，测试函数不能大写被外界访问，这里testAddUser小写
func testAddUser(t *testing.T) {
	err := AddUser("gxyu", "123")
	if err != nil {
		t.Errorf("error of AddUser: %v", err)
	}
}

func testGetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("gxyu")
	if err != nil || pwd != "123" {
		t.Errorf("error of GetUserCredential: %v", err)
	}
}

func testGetUser(t *testing.T) {
	_, err := GetUser("gxyu")
	if err!=nil {
		t.Errorf("error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("gxyu", "123")
	if err!=nil {
		t.Errorf("error of DeleteUser: %v", err)
	}

	//再次确认
	pwd, errDB := GetUserCredential("gxyu")
	if errDB != nil || pwd != "" {
		t.Errorf("error of DeleteUser: %v", err)
	}

}

//========================================video
func TestVideoWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("AddVideo", testAddNewVideo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("GetVideoList", testListVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
}

var tempvid string

func testAddNewVideo(t *testing.T) {
	VI, err := AddNewVideo(1, "my-video")
	if err!=nil {
		t.Errorf("error of AddNewVideo: %v", err)
	}
	tempvid = VI.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err!=nil {
		t.Errorf("error of GetVideoInfo: %v", err)
	}
}

func testListVideoInfo(t *testing.T) {
	
	_, err := AddNewVideo(1, "my-video2")
	if err!=nil {
		t.Errorf("error of AddNewVideo: %v", err)
	}

	// from:=1514764800
	// to, _:=strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	// res, errList := ListVideoInfo("gxyu",from, to)
	res, errList := ListVideoInfo("gxyu")
	if errList!=nil {
		t.Errorf("error of ListVideoInfo: %v", err)
	}
	for i, one := range res {
		fmt.Printf("video%d : %v\n", i, one)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err!=nil {
		t.Errorf("error of DeleteVideoInfo: %v", err)		
	}

	VI, err := GetVideoInfo(tempvid)
	if err!=nil || VI!=nil {
		t.Errorf("error of DeleteVideoInfo: %v", err)
	}

}

//===================================comments操作
func TestCommentWorkFlow(t *testing.T) {
	t.Run("AddComment", testAddNewComment)
	t.Run("GetCommentList", testListComments)

}

func testAddNewComment(t *testing.T) {
	vid := "video1"
	aid := 1
	content := "I like this video"
	err := AddNewComment(vid, aid, content)
	if err!=nil {
		t.Errorf("error of AddNewComment: %v", err)		
	}	
}

func testListComments(t *testing.T) {

	// from:=1514764800
	// to, _:=strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	// res, errList := ListComments("video1",from, to)
	res, errList := ListComments("video1")
	if errList!=nil {
		t.Errorf("error of ListComments: %v", err)
	}
	for i, one := range res {
		fmt.Printf("comment%d : %v\n", i, one)
	}	
}

//===================================sessions操作
var tempsid string

func TestSessions(t *testing.T) {

	t.Run("AddSession", testInsertSession)
	t.Run("RetriveOneSession", testRetriveSession)

}

func testInsertSession(t * testing.T) {
	sid, err:=utils.NewUUID()
	if err!=nil{
		t.Errorf("error of UUID, %v", err)
	}
	tempsid = sid
	ttl:=int64(129183174987124)
	err= InsertSession(sid, ttl, "skyone")
	if err!=nil{
		t.Errorf("error of InsertSession: %v", err)
	}
}

func testRetriveSession(t *testing.T) {
	res, err:=RetrieveSession(tempsid)
	if err!=nil{
		t.Errorf("Error of RetriveSession: %v", err)
	}
	fmt.Printf("session: %+v", res)
}
