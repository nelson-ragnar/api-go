package main

import {
	"api/src/router"
	"fmt"
	"log"
	"http"
}

func main(){
	fmt.Println("Rodando API!")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))

}
