package utils
import {
  "net/http"
  "encoding/json"
  "io/ioutil"
}

func ParseBody(r *http.request,x inteface{}) {
  body,err := ioutil.ReadAll(r.Body); err == nil{
    if err := json.Unmarshal(([]byte Body),x);err != nil{
      return
    }
  }
}
