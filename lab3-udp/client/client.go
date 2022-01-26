package main

import (
	"flag"
	"fmt"
	"github.com/mgutz/logxi/v1"
	//"math/rand"
	"net"
	"os"
	"strconv"
	"github.com/skorobogatov/input"
	"time"

	"proto"
	"encoding/json"
)

func interact(conn *net.UDPConn, command string, data interface{}, id uint)  {
	var rawData json.RawMessage
	rawData, _ = json.Marshal(data)
	ident := strconv.Itoa(int(id))
	rawReq, _ := json.Marshal(&proto.Request{command, &rawData, ident})
	buf := make([]byte, 1000)
	for {
		conn.SetDeadline(time.Now().Add(3*time.Second))
		if _, err := conn.Write(rawReq); err != nil {
			log.Error("sending request to server", "error", err)
			log.Info("sending the request again")
			continue
		}
		conn.SetDeadline(time.Now().Add(1*time.Second))
		if bytesRead, err := conn.Read(buf); err != nil {
			log.Error("receiving answer from server", "error", err)
			continue
		} else {
			conn.SetDeadline(time.Now().Add(2*time.Minute))
			var resp proto.Response
			if err := json.Unmarshal(buf[:bytesRead], &resp); err != nil {
				log.Error("cannot parse answer", "answer", buf, "error", err)
				log.Info("tut")
			} else {
				switch resp.Status {
				case "ok":
					if resp.Ident == ident {
						log.Info("client is off")
						return
					}
				case "added":
					var elem string
					if err := json.Unmarshal(*resp.Data, &elem); err != nil {
						log.Error("cannot parse answer", "answer", resp.Data, "error", err)
					} else {
						if resp.Ident == ident {
							log.Info("successful interaction with server", "added", elem)
							return
						}
					}
				case "failed":
					var reason string
					if err := json.Unmarshal(*resp.Data, &reason); err != nil {
						log.Error("cannot parse answer", "answer", resp.Data, "error", err)
					} else {
						if resp.Ident == ident {
							log.Error("failed", "reason", reason)
							return
						}
					}
				case "peak", "deq", "len":
					var res string
					if err := json.Unmarshal(*resp.Data, &res); err != nil {
						log.Error("cannot parse answer", "answer", resp.Data, "error", err)
					} else {
						if resp.Ident == ident {
							log.Info("successful interaction with server", resp.Status, res)
							fmt.Printf("result: " + res + "\n")
							return
						}
					}
				default:
					log.Error("server reports unknown status %q\n", resp.Status)
				}
			}
		}
	}
}

func main() {
	var (
		serverAddrStr string
		n             uint
		helpFlag      bool
	)
	flag.StringVar(&serverAddrStr, "server", "127.0.0.1:6000", "set server IP address and port")
	flag.UintVar(&n, "n", 10, "set the number of requests")
	flag.BoolVar(&helpFlag, "help", false, "print options list")

	if flag.Parse(); helpFlag {
		fmt.Fprint(os.Stderr, "client [options]\n\nAvailable options:\n")
		flag.PrintDefaults()
	} else if serverAddr, err := net.ResolveUDPAddr("udp", serverAddrStr); err != nil {
		log.Error("resolving server address", "error", err)
	} else if conn, err := net.DialUDP("udp", nil, serverAddr); err != nil {
		log.Error("creating connection to server", "error", err)
	} else {
		defer conn.Close()

		for i := uint(0); i < n; i++ {
			fmt.Printf("command = ")
			cmd := input.Gets()
			switch cmd {
			case "enq":
				var elem string
				fmt.Printf("value = ")
				elem = input.Gets()
				interact(conn, cmd, elem, i)
			case "peak", "deq", "len":
			  interact(conn, cmd, nil, i)
			case "quit":
				interact(conn, cmd, nil, i)
					return
			default:
				log.Error("unknown command")
				i--
				continue
			}
		}
	}
}
