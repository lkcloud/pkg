# core

## 概述

core 包实现了所有的核心函数

## 包函数列表

- **type Response** API 返回格式
- **func WriteResponse(c *rf.Context, err error, data interface{})** 根据传入的error和data构造需要的HTTP返回数据个数并返回
