# Modlib

实现 Go 模块结构，包含软件包（lib）和命令行工具（cmd）。

只需在项目中引入 Go 模块，就可以轻松使用：

```
import "github.com/zhong-my/modlib"
import "github.com/zhong-my/modlib/clientlib"
import "github.com/zhong-my/modlib/serverlib"

config := modlib.Config()
helloc := clientlib.Hello()
hellos := serverlib.Hello()
```

## 阅读更多

- [Simple Go Project Layout With Modules](https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/)
- [modlib](https://github.com/eliben/modlib)

## 版权

MIT
