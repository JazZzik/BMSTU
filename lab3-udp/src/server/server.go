package main

import (
	"flag"
	"fmt"
	"github.com/mgutz/logxi/v1"
	"net"
	"os"
	"strconv"


	"proto"
	"encoding/json"
)

type Queue struct {
	Elements []int
}

func (q *Queue) Enqueue(elem int) {
	q.Elements = append(q.Elements, elem)
}

func (q *Queue) Dequeue() int {
	var f int
	f, q.Elements = q.Elements[0], q.Elements[1:]
	return f
}

func (q *Queue) Peek() int {
	return q.Elements[0]
}

func (q *Queue) GetLength() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

var queue = Queue{}



type Client struct {
	resp   map[int]proto.Response
	count  int
}

func NewClient() *Client {
	return &Client{
		resp: 	make(map[int]proto.Response),
		count:  0,
	}
}


func serveToClients(conn *net.UDPConn) {
	clientMap := make(map[string]*Client)
	buf := make([]byte, 1000)

	for {
		
		if bytesRead, clientAddr, err := conn.ReadFromUDP(buf); err != nil {
			log.Error("receiving message from client", "error", err)
		} else {
				clientAddrStr := clientAddr.String()
				_, found := clientMap[clientAddrStr]
				if !found {
					log.Info("client is on",  "client", clientAddrStr)
					clientMap[clientAddrStr] = NewClient()
				}
				var req proto.Request
				if err := json.Unmarshal(buf[:bytesRead], &req); err != nil {
					log.Error("cannot parse request", "request", buf[:bytesRead], "error", err)
					respond("failed", err, "-1", clientAddr, conn)
				} else {
						id, _ := strconv.Atoi(req.Ident)
						resp, found := clientMap[clientAddrStr].resp[id]
						if found {
							log.Info("Sending the response again")
							respond(resp.Status, resp.Data, resp.Ident, clientAddr, conn)
						} else {
								switch req.Command {
								case "enq":
									errorMsg := ""
									var elem string
									if req.Data == nil {
										errorMsg = "data field is absent"
									} else {
										if err := json.Unmarshal(*req.Data, &elem); err != nil {
											errorMsg = "cannot parse request"
											log.Info("hello")
										} else {
											e, _ := strconv.Atoi(elem)
											queue.Enqueue(e)
										}
									}
									if errorMsg == "" {
										var rawData json.RawMessage
										rawData, _ = json.Marshal(elem)
										clientMap[clientAddrStr].resp[id] = proto.Response{"added", &rawData, req.Ident}
										if respond("added", elem, req.Ident, clientAddr, conn) {
											log.Info("successful interaction with client", "added", elem, "client", clientAddrStr)

										}
									} else {
										log.Error("addition failed", "reason", errorMsg)
										respond("failed", errorMsg, req.Ident, clientAddr, conn)
									}
								case "peak":
									c := strconv.Itoa(queue.Peek())
									var rawData json.RawMessage
									rawData, _ = json.Marshal(c)
									clientMap[clientAddrStr].resp[id] = proto.Response{req.Command, &rawData, req.Ident}
									if respond("peak", c, req.Ident, clientAddr, conn) {
										log.Info("successful interaction with client", req.Command, c, "client", clientAddrStr)
									}
								case "deq":
									c := strconv.Itoa(queue.Dequeue())
									var rawData json.RawMessage
									rawData, _ = json.Marshal(c)
									clientMap[clientAddrStr].resp[id] = proto.Response{req.Command, &rawData, req.Ident}
									if respond("deq", c, req.Ident, clientAddr, conn) {
										log.Info("successful interaction with client", req.Command, c, "client", clientAddrStr)
									}
								case "len":
									c := strconv.Itoa(queue.GetLength())
									var rawData json.RawMessage
									rawData, _ = json.Marshal(c)
									clientMap[clientAddrStr].resp[id] = proto.Response{req.Command, &rawData, req.Ident}
									if respond("len", c, req.Ident, clientAddr, conn) {
										log.Info("successful interaction with client", req.Command, c, "client", clientAddrStr)
									}
								case "quit":
									clientMap[clientAddrStr].resp[id] = proto.Response{"ok", nil, req.Ident}
									if respond("ok", nil, req.Ident, clientAddr, conn) {
										log.Info("client is off",  "client", clientAddrStr)
									}
								}
							}
			}
		}
	}
}

// respond - вспомогательный метод для передачи ответа с указанным статусом и данными
func respond(status string, data interface{}, ident string, addr *net.UDPAddr, conn *net.UDPConn) bool {
	var rawData json.RawMessage
	rawData, _ = json.Marshal(data)
	rawResp, _ := json.Marshal(&proto.Response{status, &rawData, ident})
	if _, err := conn.WriteToUDP(rawResp, addr); err != nil {
		log.Error("sending response to client", "error", err)
		return false
	}
	return true
}

func main() {
	var (
		serverAddrStr string
		helpFlag      bool
	)
	flag.StringVar(&serverAddrStr, "addr", "127.0.0.1:6000", "set server IP address and port")
	flag.BoolVar(&helpFlag, "help", false, "print options list")

	if flag.Parse(); helpFlag {
		fmt.Fprint(os.Stderr, "server [options]\n\nAvailable options:\n")
		flag.PrintDefaults()
	} else if serverAddr, err := net.ResolveUDPAddr("udp", serverAddrStr); err != nil {
		log.Error("resolving server address", "error", err)
	} else if conn, err := net.ListenUDP("udp", serverAddr); err != nil {
		log.Error("creating listening connection", "error", err)
	} else {
		log.Info("server listens incoming messages from clients")
		serveToClients(conn)
	}
}
