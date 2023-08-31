package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}
// caracterista do GO sempre tera um if para verificar se há erros
//"movies" é uma variavel
var movies = []Movie{}

func CreateMovie(c *gin.Context){
	var newMovie Movie
	if err := c.ShouldBindJSON(&newMovie); err != nil { //nil é nulo
	c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		}

	newMovie.ID = len(movies) + 1   // variavel ID pega o tamanho de array e adiciona mais
	movies = append(movies, newMovie) // apeend Movie é adiciona e concatenar no array
	c.JSON(http.StatusOK, newMovie) //sucesso na porta 200
}

func GetMovies(c *gin.Context){
	c.JSON(http.StatusOK, movies)
}



func main (){
	r := gin.Default()	
	r.POST("/movie" , CreateMovie)
	r.GET("/movie", GetMovies)
	r.Run(":8080")
}