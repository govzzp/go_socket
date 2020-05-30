package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn)  {
	defer conn.Close()
	for  {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n , err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("Read From Coon Failed,err:%v\n",err)
			break
		}
		recv := string(buf[:n])
		fmt.Printf("Data has receviced:",recv)
		conn.Write(buf[:n])
		conn.Write([]byte("ok"))



	}
}
func main()  {
	listen,err := net.Listen("tcp","127.0.0.1:12345")
	if err != nil{
		fmt.Printf("Listen Error err:%v\n",err)
		return
	}
	for  {
		conn ,err := listen.Accept()
		if err !=nil  {
			fmt.Printf("Accept Failed,err:%v\n",err)
			continue
		}
		go process(conn)
	}
}
