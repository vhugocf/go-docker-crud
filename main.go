package main

import(
	"fmt"
	"database/sql"
	"enconding/json"
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