package main

import (
   "fmt"
   "net"
)

var name string
func main() {
   conn,err := net.Dial("tcp", "localhost:8080")
   if err != nil {
      fmt.Println("There was an error: ")
      fmt.Println(err)
   }
   fmt.Println("Enter your name:")
   fmt.Scanln(&name)
   c := make(chan string)
   go sendToServer(conn, c)
   go listenToServer(conn)
   readKeyboard(c)
   fmt.Println("Done")  
}

func sendToServer(conn net.Conn, c chan string) {
   for {
      userLine := name + ">>> "
      userLine += <- c
      bArr := []byte(userLine)
      conn.Write(bArr)
   }
}

func readKeyboard(c chan string) {
   fmt.Println("Ready to chat")
   var err error
   var n int
   for {
      ln := ""
      n, err = fmt.Scanln(&ln)
      if err != nil {
         fmt.Println("There was an error, quitting")
         fmt.Println(err)
         break
      }
      if n > 0 {
         c <- ln
      }
   }
}

func listenToServer(c net.Conn){
   for {
      tbuf := make([]byte, 81920)
      _, err := c.Read(tbuf)
      if err != nil {
         fmt.Println(err)
      }
      message := string(tbuf)
      fmt.Println(message)
   }
}