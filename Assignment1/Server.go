/* EchoChanServer
 */
 package main

 import (
	 "fmt"
	 "os"
	 "old/netchan"
	 "strconv"
 )
 
 var count int = 0
 
 func main() {
 
	 exporter := netchan.NewExporter()
	 err := exporter.ListenAndServe("tcp", ":2345")
	 checkError(err)
 
	 echo := make(chan string)
	 exporter.Export("echo", echo, netchan.Send)
	 for {
		 sCount := strconv.Itoa(count)
		 lock := make(chan string)
		 go handleSession(exporter, sCount, lock)
 
		 <-lock
		 echo <- sCount
		 count++
		 exporter.Drain(-1)
	 }
 }
 
 func handleSession(exporter *netchan.Exporter, sCount string, lock chan string) {
	 echoIn := make(chan string)
	 exporter.Export("echoIn"+sCount, echoIn, netchan.Send)
 
	 echoOut := make(chan string)
	 exporter.Export("echoOut"+sCount, echoOut, netchan.Recv)
	 fmt.Println("made " + "echoOut" + sCount)
 
	 lock <- "done"
 
	 for {
		 s := <-echoOut
		 echoIn <- s
	 }
	 // should unexport net channels
 }
 
 func checkError(err error) {
	 if err != nil {
		 fmt.Println("Fatal error ", err.Error())
		 os.Exit(1)
	 }
 }
 and a client is
 
 /* EchoChanClient
  */
 package main
 
 import (
	 "fmt"
	 "old/netchan"
	 "os"
 )
 
 func main() {
	 if len(os.Args) != 2 {
		 fmt.Println("Usage: ", os.Args[0], "host:port")
		 os.Exit(1)
	 }
	 service := os.Args[1]
 
	 importer, err := netchan.Import("tcp", service)
	 checkError(err)
 
	 fmt.Println("Got importer")
	 echo := make(chan string)
	 importer.Import("echo", echo, netchan.Recv, 1)
	 fmt.Println("Imported in")
 
	 count := <-echo
	 fmt.Println(count)
 
	 echoIn := make(chan string)
	 importer.Import("echoIn"+count, echoIn, netchan.Recv, 1)
 
	 echoOut := make(chan string)
	 importer.Import("echoOut"+count, echoOut, netchan.Send, 1)
 
	 for n := 1; n < 10; n++ {
		 echoOut <- "hello "
		 s := <-echoIn
		 fmt.Println(s, n)
	 }
	 close(echoOut)
	 os.Exit(0)
 }
 
 func checkError(err error) {
	 if err != nil {
		 fmt.Println("Fatal error ", err.Error())
		 os.Exit(1)
	 }
 }