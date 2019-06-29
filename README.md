# iris 框架

## 安装 iris

```bash
#  中文文档 https://studyiris.com/doc/irisDoc/Installation.html
cd $GOPATH/src

go get -u github.com/kataras/iris
```

## 热重启

```bash
# 安装依赖
cd $GOPATH/src

# 这个不太好用，如果代码报错后，就没办法重启了。 因此弃用。 改成 gowatch
go get -u github.com/kataras/rizla

# 进度项目目录
rizla main.go

# gowatch 目前用起来，感觉良好
go get -u github.com/silenceper/gowatch

gowatch main.go
```

## 构建二进制文件

```bash
sh build.sh
# 或者直接执行
go build -o dist/iris-admin-api main.go
```
