# version

## 概述

version包用来给二进制命令添加版本功能：
```
./auth-apiserver --version
GitVersion:     39f91a2-dirty
BuildDate:      2019-03-13T06:47:57Z
GoVersion:      go1.11.4
Compiler:       gc
Platform:       linux/amd64
```

## 包函数列表

- **func AddCommand(cmd *cobra.Command)** 给cobra Command添加version option
- **func String() string** 返回string格式的版本信息
- **Get() *Info** 获取版本详细信息
- **func (info Info) String()** 返回string格式的版本信息
