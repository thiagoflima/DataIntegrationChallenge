/****************************************************
 *              Packages and Imports                *
 ****************************************************/

package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

/****************************************************
 *              Script operations                   *
 ****************************************************/

//  Starts the HTTP server in order to receive requests.
func RunServer() {

    router := mux.NewRouter()

    router.HandleFunc("/companies", GetAllCompanies).Methods("GET")
    router.HandleFunc("/companies/id/{id}", GetCompanyById).Methods("GET")
    router.HandleFunc("/companies/like/{name}/{zipcode}", GetCompaniesLike).Methods("GET")
    router.HandleFunc("/companies/equals/{name}/{zipcode}", GetCompaniesEquals).Methods("GET")
    router.HandleFunc("/companies/bulk", BulkCompaniesWithFile).Methods("POST")

    log.Fatal(http.ListenAndServe(":8000", router))
}

//  Returns all the companies on the database.
func GetAllCompanies(wtr http.ResponseWriter, rdr *http.Request) {
    
    response := Companies{Status: 200, Message: "OK", Companies: SelectAllDB()}

    json.NewEncoder(wtr).Encode(response)
}

//  Returns a company based on its id, from the GET parameters.
func GetCompanyById(wtr http.ResponseWriter, rdr *http.Request) {
    
    params := mux.Vars(rdr)

    response := Companies{Status: 200, Message: "OK", Companies: SelectByIdDB(params["id"])}
    
    json.NewEncoder(wtr).Encode(response)
}

//  Returns a company based on the GET parameters.
func GetCompaniesLike(wtr http.ResponseWriter, rdr *http.Request) {
    
    params := mux.Vars(rdr)

    response := Companies{Status: 200, Message: "OK", Companies: SelectLikeDB(params["name"], params["zipcode"])}
    
    json.NewEncoder(wtr).Encode(response)
}

//  Returns a company with the GET parameters.
func GetCompaniesEquals(wtr http.ResponseWriter, rdr *http.Request) {
    
    params := mux.Vars(rdr)

    response := Companies{Status: 200, Message: "OK", Companies: SelectEqualsDB(params["name"], params["zipcode"])}
    
    json.NewEncoder(wtr).Encode(response)
}

//  Updates the companies using the file from a POST request.
func BulkCompaniesWithFile(wtr http.ResponseWriter, rdr *http.Request) {

    rdr.ParseMultipartForm(32 << 20)

    file, hndlr, err := rdr.FormFile("file")
    CheckError(err)

    SaveFile(hndlr.Filename, file)
    file.Close()

    q2 := LoadFile("Q2")
    UpdateFileToDB(q2)

    response := Response{Status: 200, Message: "OK"}

    json.NewEncoder(wtr).Encode(response)
}

/****************************************************/