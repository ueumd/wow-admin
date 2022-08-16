# Wow-Admin
> Developing a RESTful API with Go and Gin

# Usage
```bash
git clone https://github.com/ueumd/wow-admin.git

cd wow-admin

# config.yaml
# db/wow_admin.sql
# redis

go mod tidy

fresh
```


## Api
```bash
get:  http://127.0.0.1:5173/v1/common/getSystemDate

{
    code: 200,
    message: "success",
    data: {
        date: "2022-08-16 16:31:33",
        timestamp: 1660638693000
    }
}

post: http://127.0.0.1:5173/v1/login/phoneLogin
}

{
    "code": 200,
    "message": "success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbiI6ImQ2N2Y4MjI0NmJmZTRkN2JiNTFlMmRkYTViZDJkZGMxIiwidXNlcklkIjoxMDAxMiwidW5pb25pZCI6Im9JVGlhamhJZS1DNTJMRDBhVFYyWFBYaXlJVEEiLCJsb2dpblR5cGUiOiJ3ZWNoYXQiLCJwaG9uZSI6IiIsInRpbWVzdGFtcCI6MTY2MDY0MjcxMH0.qzrfcgC4t2a1XfmN_u0XsHUr4HBACafUEYD8DnGixU0"
    }
}
```


## Using Go Modules

| 指令           | 	功能                               | 
|--------------|-----------------------------------|
| go mod download | 下载依赖包到本地(默认下载目录为 $GOPATH/pkg/mod) |
| go mod edit	 | 编辑 go.mod 文件                      |
| go mod graph	 | 输出 go mod的依赖图                     |
| go mod init	 | 初始化一个mod项目，即在包中示创建一个go.mod文件      |
| go mod tidy	 | 自动补全新增(删除减少)的依赖包                  |
| go mod verify | 校验依赖包                             |
| go mod why	 | 解释为什么要依赖，即显示这个包在哪里被导入了            |