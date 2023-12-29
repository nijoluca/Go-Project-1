package main
import {
  "net/http"
  "log"
  "github.com/gorilla/mux"
  "github.com/nijoluca/project_2/pkg/routes"
}
func main(){
  r := mux.NewRouter()
  routes.BookStoreRoutes(r)
  http.Handle("/",r)
  log.Fatal(http.ListenAndServe(":8080",nil))
  
}