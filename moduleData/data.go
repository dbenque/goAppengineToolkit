package data

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"

    "appengine"

    "github.com/dbenque/goAppengineToolkit/datastoreEntity"

)

func init() {

    r := mux.NewRouter()
    r.HandleFunc("/{uri}/friend/{name}", handleGetFriend).Methods("GET")
    r.HandleFunc("/{uri}/friend/{name}/{phone}", handleSetFriend).Methods("GET")
    http.Handle("/", r)

}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "data")
}

func handleGetFriend(w http.ResponseWriter, r *http.Request) {

  vars := mux.Vars(r)
  c := appengine.NewContext(r)
  friend := Friend{Name: vars["name"]}

  err:=datastoreEntity.Retrieve(&c, &friend)

  if err!=nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  jdata, _ := json.Marshal(friend)
  w.WriteHeader(http.StatusOK)
  w.Write(jdata)
}

func handleSetFriend(w http.ResponseWriter, r *http.Request) {

  vars := mux.Vars(r)
  c := appengine.NewContext(r)
  friend := Friend{Name: vars["name"],Phone: vars["phone"]}

  err:=datastoreEntity.Store(&c, &friend)

  if err!=nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  jdata, _ := json.Marshal(friend)
  w.WriteHeader(http.StatusOK)
  w.Write(jdata)
}
