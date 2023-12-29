package models
import {
  "github/jinzhu/gorm"
  "github/nijoluca/project_2/pkg/config"
}
var db *gorm.DB
type Book struct {
  gorm.model
  name string `gorm:""json:"name"`
  Author string 'json:"author"'
  Publications string `json:"publications"`
  
}
func init(){
  config.connect()
    db = config.GetDB()
    ab.AutoMigrate(&Book{})
}

func (b *Book) createBook() *Book {
  db.NewRecord(b)
  db.create(&b)
  return db
  
}

func GetAllBooks() []Book {
  var Books []Book
  db.Find(&Books)
  return Books
}

func GetBookByID(Id int64) (*Book,*gorm.DB){
  var getBook Book
  db := db.where("ID=?",id).Find(&getBook)
  return &getBook,db
}

func Delete(Id int64) Book{
  var book Book
  db.where("ID=?",id).Delete(book)
  return book
}