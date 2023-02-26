package main


 import (

	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

)


// movie struct (model) 

type  Movie struct{

	ID string `json:"id"` // json tag

	isbn string `json:"isbn"` // json tag

	Title string `json:"title"` // json tag

	Director *Director `json:"director"` // json tag

	
}

var movies []Movie // slice of movies


// director struct (model)  associated with movie struct , ex  -  every movie has a director
type Director struct{

	firstName string `json:"firstName"`
	lastName string `json:"lastName"`
}


// get all movies from the slice of movies
func getMovies(w http.ResponseWriter , r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)

}


// delete a movie from the slice of movies
func deleteMovie(w http.ResponseWriter  , r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index ,  item := range movies{

		if item.ID == params["id"] {
			

			// delete the movie from the slice of movies
			movies = append(movies[:index], movies[index+1:]...)

			break
		}
	}

	json.NewEncoder(w).Encode(movies) 
	
	
}

// get a movie from the slice of movies
func getMovie(w http.ResponseWriter , r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies{

		if item.ID == params["id"] {

			json.NewEncoder(w).Encode(item)
			return
		}
	}


}

// creating a new movie and appending it to the slice of movies

func createMovie( w http.ResponseWriter , r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie) // 

	movie.ID = strconv.Itoa(rand.Intn(1000000))

	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)

}


func updateMovie( w http.ResponseWriter , r *http.Request){

	// define a content type for the response

	w.Header().Set("Content-Type", "application/json")


	params := mux.Vars(r)

	for index , item := range movies{

		if item.ID == params["id"] {

			
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie)

			movie.ID = params["id"]

			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movie)

			return


		}
	}

	
}

func main() {

	mux := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Title: "Movie One", Director: &Director{firstName: "John", lastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Title: "Movie Two", Director: &Director{firstName: "Steve", lastName: "Smith"}})

	mux.HandleFunc("/movies", getMovies).Methods("GET")
	mux.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	mux.HandleFunc("/movies", createMovie).Methods("POST")
	mux.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	mux.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")


	fmt.Printf("Starting server at port 8000\n")

	log.Fatal(http.ListenAndServe(":8000", mux))




}





