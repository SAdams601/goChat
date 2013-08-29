package main

import (
   "fmt"
   "net"
   "io"
   )

var knownClients []net.Conn

func main() {
   ln, err := net.Listen("tcp", ":8080")
   if err != nil {
      fmt.Println("Listen failure, shutting down")
      return
   }
   for {
      conn, err := ln.Accept()
      knownClients = append(knownClients, conn)
      pos := len(knownClients) - 1
      if err != nil {
         fmt.Println("There was an error: ")
         fmt.Println(err)
         continue
      }
      go handleConnection(conn, pos)
   }
}

func handleConnection(c net.Conn, pos int) {
   tbuf := make([]byte, 81920)
   totalBytes := 0
   for {
      n, err := c.Read(tbuf)
      totalBytes += n
      if err != nil {
         if err != io.EOF {
            fmt.Println("Read error")
         }
         break
      }
      fmt.Println(n)
      s := string(tbuf)
      go sendChats(s, pos)
      fmt.Println(s)
   }
   fmt.Printf("%d bytes read\n", totalBytes)
   c.Close()
}

func sendChats(message string, pos int) {
   for index, value := range knownClients {
      if index != pos {
         go sendStr(message, value)
      }
   }
}

func sendStr(message string, c net.Conn){
   bArr := []byte(message)
   c.Write(bArr)
}