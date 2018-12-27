## 下载安装包

<a href="https://golang.google.cn/dl/" target="_blank">golang.google.cn/dl/</a>

---
## 设置GOPATH环境变量

GOPATH 路径根据个人情况设置为自已的实际目录路径

注意：可能还需要同时添加 `GOPATH/bin` 路径到环境变量

#### MacOS设置

```
nano ~/.bash_profile
```

添加以下内容，其中 alias 为代理指定，如无代理此行不加

```
export GOPATH=/Users/yourname/gocode
alias go="https_proxy=socks5://127.0.0.1:1086 go"
```

然后执行

```
source ~/.bash_profile
```

#### Windows设置（未测试）

添加系统环境变量 `GOPATH=D:\gocode`

创建一个 `go-alias.bat` 打开编辑，添加以下内容：

```
@doskey go=set https_proxy=socks5://127.0.0.1:1086 && go $*
```

然后把 `go-alias.bat` 加入到系统自启动中，win+r，输入regedit，修改注册表

HKEY_CURRENT_USER -> Software -> Microsoft -> Command Processor添加一个AUTORUN的字符串值，值为bat的全路径

#### Git代理

因为 `Go` 的包管理主要是依赖于 github 的，所以Git的代理几乎也是必须的，以下为Git代理设置（MacOS示例）：

```
touch ~/.gitconfig
open ~/.gitconfig
```

设置以下内容

```
[http]
  proxy = socks5://127.0.0.1:1086
[https]
  proxy = socks5://127.0.0.1:1086
```