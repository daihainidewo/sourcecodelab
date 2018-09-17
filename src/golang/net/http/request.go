// Package main main
// file create by daihao, time is 2018/7/30 17:21
package main

import (
	"net/url"
	"io"
	"mime/multipart"
	"crypto/tls"
	"context"
	"net/http"
)

type Request struct {
	// 请求方法
	Method string

	// 请求URL
	URL *url.URL

	// HTTP协议版本
	Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0

	// 请求头部信息
	Header http.Header

	// 请求体
	Body io.ReadCloser

	// 复制请求头，对于服务器请求未使用
	GetBody func() (io.ReadCloser, error)

	// Content-length 数值
	ContentLength int64

	// 记录从外到内的转移编码
	TransferEncoding []string

	// 是否关闭tcp连接
	Close bool

	// 请求host
	Host string

	// URI参数
	Form url.Values

	// body参数
	PostForm url.Values

	// 存放多表单
	MultipartForm *multipart.Form

	Trailer http.Header

	// remote addr
	RemoteAddr string

	// 请求URI
	RequestURI string

	// 传输层安全协议
	TLS *tls.ConnectionState

	Cancel <-chan struct{}

	// 只要在客户端重定向期间填充
	Response *http.Response

	// 服务器上下文， 只能通过withcontext复制整个请求来修改它
	ctx context.Context
}

// main main
func main() {
	// TODO

}
