package main

import (
	"fmt"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "server/redis"
	_ "github.com/davyxu/cellnet/peer/http"
	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/peer/gorillaws"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/http"
	_ "github.com/davyxu/cellnet/proc/tcp"
	_ "github.com/davyxu/cellnet/proc/gorillaws"
	"github.com/davyxu/golog"
	"server/login"
	_ "server/mysql"
	"server/proto"
	"time"
)

var (
	log = golog.New("server")
	TCP_ADDRESS = "0.0.0.0:3600"
	WS_ADDRESS = "0.0.0.0:3601"
	HTTP_ADDRESS = "0.0.0.0:3602"
)

func HttpPeer(ch chan<- int) {
	queue := cellnet.NewEventQueue()
	p := peer.NewGenericPeer("http.Acceptor", "WebServer", HTTP_ADDRESS, queue)
	proc.BindProcessorHandler(p, "http", func(ev cellnet.Event) {
		switch ev.Message().(type) {
		case *cellnet.SessionAccepted:
			log.Debugln("server accepted")
		}
	})
	// 开始侦听
	p.Start()
	// 事件队列开始循环
	queue.StartLoop()
	// 阻塞等待事件队列结束退出( 在另外的goroutine调用queue.StopLoop() )
	queue.Wait()
	ch <- 1
}

func WsPeer(ch chan <- int) {
	queue := cellnet.NewEventQueue()
	p := peer.NewGenericPeer("gorillaws.Acceptor", "WsServer", WS_ADDRESS, queue)
	proc.BindProcessorHandler(p, "tcp.ltv", func(ev cellnet.Event) {
		switch ev.Message().(type) {
		case *cellnet.SessionAccepted:
			log.Debugln("Server accepted")
		}
	})

	// 开始侦听
	p.Start()
	// 事件队列开始循环
	queue.StartLoop()
	// 阻塞等待事件队列结束退出( 在另外的goroutine调用queue.StopLoop() )
	queue.Wait()
	ch <- 3
}

func callback(ev cellnet.Event) {
	switch ev.Message().(type) {
	// 有新的连接
	case *cellnet.SessionAccepted:
		log.Debugln("server accepted")
		// 有连接断开
	case *cellnet.SessionClosed:
		log.Debugln("session closed: ", ev.Session().ID())
		// 收到某个连接的ChatREQ消息
	case *proto.LoginReq:
		login.LoginRequestCallback(ev)
	}
}

func TcpPeer(ch chan<- int) {
	// 创建一个事件处理队列，整个服务器只有这一个队列处理事件，服务器属于单线程服务器
	queue := cellnet.NewEventQueue()

	// 创建一个tcp的侦听器，名称为server，连接地址为127.0.0.1:8801，所有连接将事件投递到queue队列,单线程的处理（收发封包过程是多线程）
	p := peer.NewGenericPeer("tcp.Acceptor", "server", TCP_ADDRESS, queue)
	// 设置30秒读超时和5秒写超时
	p.(cellnet.TCPSocketOption).SetSocketDeadline(time.Second * 30, time.Second * 5)
	// 设定封包收发处理的模式为tcp的ltv(Length-Type-Value), Length为封包大小，Type为消息ID，Value为消息内容
	// 每一个连接收到的所有消息事件(cellnet.Event)都被派发到用户回调, 用户使用switch判断消息类型，并做出不同的处理
	proc.BindProcessorHandler(p, "tcp.ltv", callback)

	// 开始侦听
	p.Start()

	// 事件队列开始循环
	queue.StartLoop()

	// 阻塞等待事件队列结束退出( 在另外的goroutine调用queue.StopLoop() )
	queue.Wait()
	ch <- 2
}

func main() {
	ch := make(chan int, 3)
	go HttpPeer(ch)
	go TcpPeer(ch)
	go WsPeer(ch)
	select {
	case n := <-ch:
		if n == 1 {
			fmt.Println("Http server start.")
		} else if n == 2 {
			fmt.Println("Tcp server start.")
		}
	}
	fmt.Println(WS_ADDRESS)
}
