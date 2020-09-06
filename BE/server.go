package main

import (
  "os"
  "os/exec"
  "log"
  "net/http"
  "strings"
  "strconv"
  "fmt"
)

func enalbeCors(w *http.ResponseWriter) {
  (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {

  for name, headers := range r.Header {
    for _, h := range headers {
      log.Printf("%v: %v\n", name, h)
    }
  }

  enalbeCors(&w)

  r.ParseForm()

  for k, v := range r.Form {
    log.Println("key: ", k)
    log.Println("val: ", strings.Join(v, ""))
  }

  log.Println(strings.Split(r.Referer(), "/"))
  tokens := strings.Split(r.Referer(), "/")
  botName := tokens[len(tokens) - 1]


  cmd := exec.Command("scbw.play", "--bots", botName, "--human")
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  must(cmd.Run())




  fmt.Fprintf(w, "Golang wevserver is working!\n")
}

func must (err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  port := 9090

  http.HandleFunc("/", defaultHandler)

  log.Println(":" + strconv.Itoa(port))

  err := http.ListenAndServe(":" + strconv.Itoa(port), nil)

  if err != nil {
    log.Fatalln("Hey, the server is down.")
  } else {
    log.Printf("The server is running -> port %d", port)
  }
}
