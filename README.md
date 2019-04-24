# 备忘录

在代码一些细节点注意加上注释



# Streaming-Media-Service

- 视频网站包含Go在实战项目中大部分技能点

![1556075557151](README.assets/1556075557151.png)



# 环境及有关说明

- OS：Windows



- 仓库版本目录为`/Streaming-Media-Service/.git/`，整个文件目录为`$GOPATH/src/Streaming-Media-Service/`

- 编译某一个包时`go get`自动获取依赖的，依赖包存放目录`$GOPATH/src/`中
- 编译测试某一个包时，使用`go build`在当前包目录生成可执行文件，点击直接运行。而使用`go install`可执行文件输出到`$GOPATH/bin/`中

- 数据库、表

```sql
CREATE TABLE video_info (
  id VARCHAR(255) PRIMARY KEY NOT NULL,
  author_id INT,  --同上传用户ID，不使用外键约束而在代码中进行检查，减去数据库负荷
  name TEXT,
  display_ctime TEXT, --展示给用户
  create_time DATETIME --入库时间，类型为DATETIME
);

CREATE TABLE sessions (
  session_id VARCHAR(255) PRIMARY KEY NOT NULL,
  TTL TINYTEXT, --过期信息
  login_name VARCHAR(255)
);
```

![1556092916184](README.assets/1556092916184.png)

- - 数据库驱动`"github.com/go-sql-driver/mysql"`



# 原生、扩展API说明

```go
"io"
"net/http"
"github.com/julienschmidt/httprouter"
"github.com/go-sql-driver/mysql"

```




# 各模块

## 后端API



|API几个特点

- REST设计风格；RESTful使用HTTP通信协议，JSON作为数据格式

- 无状态

![1556075968522](README.assets/1556075968522.png)

|用户

注册：URL:/user Method: POST, SC: 201, 400, 500

登录：URL:/user/username Method: POST, SC: 200, 400, 500

获取用户信息：URL:/user/username Method: GET, SC: 200, 400, 401, 403,  500

注销：URL:/user/username Method: DELETE, SC: 204, 400, 401, 403, 500

|视频资源

.../user/username/videos/vid-id

|评论

.../videos/vid-id/commets/commets-id

|`api/handler.go`