# go-sample

## 问题

1. supervisor 脚本如何优雅的管理
    - 如果supervisor任务不打二进制包,直接运行代码,留存superviosr配置文件即可, 更新supervisor只需要拉代码,重启脚本即可
    - 如果supervisor任务打成二进制包,就麻烦了,如何判断二进制包基于哪个git版本呢

## 参考文献

1. 项目文件夹命名 <https://github.com/golang-standards/project-layout/blob/master/README_zh-CN.md>

## 使用的第三方库

1. gin
