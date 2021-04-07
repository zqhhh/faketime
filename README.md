# faketime
> go语言，伪造假时间

**不要在生产环境中使用!**

### 安装
go get -v github.com/zqhhh/faketime

### 示例
```
package main

import (
	"fmt"
	"time"
	
	"github.com/zqhhh/faketime"
)

func main(){
	fmt.Printf("没有hook之前:%s\n",time.Now()) // 2021-03-18 23:45:12.34313 +0800 CST
	ftime := faketime.Add(1*time.Hour)
	ftime.Hook()
	fmt.Printf("hook之后:%s\n",time.Now()) // 2021-03-19 00:45:12.3731311 +0800 CST
	ftime.UnHook()
	fmt.Printf("unhook之后:%s\n",time.Now()) // 2021-03-18 23:45:12.3731311 +0800 CST
}
```

