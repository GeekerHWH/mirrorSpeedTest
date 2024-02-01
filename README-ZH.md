<div align="center">
  <a href="README.md">Click me for English version</a>
</div>
# 测速 Debian/Ubuntu 镜像
该工具主要帮助Debian系用户筛选最佳的apt仓库镜像

如果你有任何的建议或者发现了任何的bug欢迎随时提出issue：）

# To Do List
- [ ] 在多选模式中支持一键多选
- [ ] 支持英文注释
- [ ] 支持英文版本
- [ ] 检查主机系统
- [ ] 支持更多镜像源
- [ ] 支持多线程测试
- [ ] 支持基于地域的镜像测试
- [ ] 支持网络延时测试
- [ ] 支持内嵌的换源操作

# 如何使用
1. 确保你的电脑安装有Go的运行环境
```bash
go env
```
2. 下载解压该项目文件后，进入项目文件夹运行下命令
```bash
go run main/main.go
```
3. 跟随app指令交互即可