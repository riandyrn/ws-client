package main

import (
  "github.com/gorilla/websocket"
  "net/http"
  "log"
)

var upgrader = websocket.Upgrader{
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
}

func main(){

  // handle main page
  http.Handle("/", http.FileServer(http.Dir("./public")))

  // handle websocket connections
  http.HandleFunc("/websocket", handleWebsocketConn)

  // start server
  log.Println("http server started on :3000")
  log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleWebsocketConn(w http.ResponseWriter, r *http.Request) {

  // attempt to upgrade http connection to websocket
  ws, err := upgrader.Upgrade(w, r, nil)
  if checkError(err) {
    return
  }
  defer ws.Close()

  for {
    msgType, msg, err := ws.ReadMessage()
    if checkError(err) {
      break
    }
    err = ws.WriteMessage(msgType, msg)
    if checkError(err) {
      break
    }
  }
}

// define short function for checking error
func checkError(err error) bool{
  var isError bool = (err != nil)
    if isError {
      log.Println(err)
  }
  return isError
}
