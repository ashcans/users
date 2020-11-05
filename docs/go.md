#go环境配置

##下载golang
官网地址: https://golang.org/dl/
##安装
安装时会自动配置环境变量
##Go Modules配置
中国大陆不使用科学上网的方式的话就要配置这个了
参考 https://juejin.im/post/6844903954879348750
```bash
#修改 GOBIN 路径（可选）：
go env -w GOBIN=$HOME/bin
#打开 Go modules：
go env -w GO111MODULE=on
#设置 GOPROXY：
go env -w GOPROXY=https://goproxy.cn,direct
# 在中国是必须的，因为它的默认值被墙了。

#第三步(可选): 按照你喜欢的目录结构重新组织你的所有项目。
#第四步: 在你项目的根目录下执行 go mod init <OPTIONAL_MODULE_PATH> 以生成 go.mod 文件。
go mod init #默认路径github.com/ashcans/authorization
#第五步: 想办法说服你身边所有的人都去走一下前四步。


```