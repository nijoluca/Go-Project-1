package routes
import {
  "github.com/gorilla/mux"
  "github.com/nijoluca/project_2/pkg/controllers"
}
var BookStoreRoutes = func(router *mux.Router){
  router.HandleFunc("/books/",controllers.CreateBook).Methods("POST")
  router.HandleFunc("/books/",controllers.GetBook).Methods("GET")
  router.HandleFunc("/books/{bookId}",controllers.GetBookByID).Methods("GET")
  router.HandleFunc("/books/{bookId}",controllers.UpdateBook).Methods("PUT")
  router.HandleFunc("/books/{bookId}",controllers.DeleteBook).Methods("DELETE")
  
  
  
  
}