package main

import (
	"database/sql"
	"encoding/json"
	"enconding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}



func main(){
	
	//connect to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()


	//create router
	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", getUsers(db)).Methods("GET")
	router.HandleFunc("/users", createUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser(db)).Methods("DELETE")

//start server
	log.Fatal(http.ListenAndServe(":8000", jsonContendTypeMiddleware(router)))
		return http.HandleFunc(func(w http.ResponseWriter, r *http.Request){
			w.Header().Set("Contend-Type", "application/json")
			next.ServeHTTP(w, r)
		})


	}
	func getUsers(db *sql.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request){
			rows, err := db.Query("SELECT * FROM users")
			if err != nil{
				log.Fatal(err)
			}
			defer rows.Close()

			users := []User{}
			for rows.Next(){
				var u User
				if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil{
					log.Fatal(err)
				}
				users = append(users, u)
			}
			if err := rows.Err(); err != nil {
				log.Fatal(err)
			}
			json.NewEncoder(w).Encode(users)
		}
	}



		

















