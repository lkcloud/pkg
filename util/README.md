## 概述

util 包定义了一些常用的工具类函数，该包也包含具有特定功能的子包

### 包列表

- [cache](cache/README.md)
- [clock](clock/README.md)
- [cmd](cmd/README.md)
- [diff](diff/README.md)
- [exec](exec/README.md)
- [flag](flag/README.md)
- [idutil](idutil/README.md)
- [id_util](id_util/README.md)
- [net](net/README.md)
- [rand](rand/README.md)
- [sets](sets/README.md)
- [urlutil](urlutil/README.md)
- [uuid](uuid/README.md)
- [validation](validation/README.md)
- [yaml](yaml/README.md)

### 包函数列表

- **func EnsureDirExists(dir string) (err error)** 如果dir不存在则创建，存在则返回成功
- **CombineRequestErr(resp gorequest.Response, body string, errs []error)** `github.com/parnurzeal/gorequest`包错误整合
- **func DelFromSlice(slice []string, elems ...string) []string** 从slice中删除某个元素
- **func GetLocalAddress() string** 获取本地IP地址
- [func GetLocalAddress() string](GetLocalAddress.md)

