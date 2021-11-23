# dm
### 介绍
达梦数据库是国产化的数据库，该类库是达梦数据的SDK
看到里面奇怪的文件名和函数名，不要喷，猜测是原版通过别的语言自动生成的吧。

### 特性
Fork自`https://gitee.com/chunanyong/dm`，并做了一些优化
* 去除全局配置
* 优化命名方式
* 增加mac方式编译

``` 
go get github.com/gotomicro/dmgo 
```  

### 使用原生方式
```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/gotomicro/dmgo"
)

func main() {
	obj, err := sql.Open("dm", "dm://username:password@ip:5236")
	if err != nil {
		panic(err)
		return
	}
	rows, err := obj.Query("select  TABLE_NAME,comments TABLE_COMMENT from user_tab_comments")
	for rows.Next() {
		a := TableStruct{}
		rows.Scan(&a.TableName, &a.TableComment)
		fmt.Printf("a--------------->"+"%+v\n", a)
	}
}

type TableStruct struct {
	TableName    string //表名
	TableComment string //表注释
}
```

### 使用gorm方式

### DSN  
dm://userName:password@ip:port  
用户名(userName)就是数据库的名称,达梦用户模式和数据库名称是对应的 
建议达梦使用UTF-8字符编码,不区分大小写    

### 版本号  
golang三段位版本号和达梦四段位版本号不兼容,统一使用1.达梦主版本号.发布的小版本号,具体查看标签的备注  

* v1.8.0 备注是 达梦8.1.1.126  
* v1.8.1 备注是 达梦8.1.1.190  
* v1.8.2 备注是 达梦8.1.2.18  
* v1.8.4 备注是 达梦8.1.2.38        




