# No-Sleep
[![license](https://img.shields.io/badge/license-apache-brightgreen.svg?style=flat)](https://github.com/SimpleBiz/no-sleep/blob/main/LICENSE)  [![Go](https://img.shields.io/badge/language-Go-blue.svg)]()

[简体中文](README.md) | [English](README_EN.md)

今夜无眠……

这是一个简单的Go程序，可以用于学习Go的基本并发操作。

更重要的是，运行它，可以无视一些全局管理性的睡眠控制，达到强制阻断电脑休眠的效果。

## 安装

确保你已经安装了Go环境，建议Go1.19及更高版本。

1. 克隆此仓库：
git clone https://github.com/SimpleBiz/no-sleep

2. 进入项目目录并安装依赖：
```
cd no-sleep
go mod tidy
```

3. 运行程序：
```
go run cmd/main.go
```

## 功能

- 基于模拟设备输入阻断睡眠。
- 当5秒内没有任何活动时，启用阻断睡眠机制。

## 疑问

请直接在此仓库提交issue

## 贡献

欢迎对此项目进行贡献！请先fork此仓库，然后提交Pull Request。

## 许可证

Apache License Version 2.0


