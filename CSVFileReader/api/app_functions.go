/****************************************************
 *              Packages and Imports                *
 ****************************************************/

package main

import (
    "encoding/csv"
    "fmt"   
    "io"
)

/****************************************************
 *              Script operations                   *
 ****************************************************/

//  Bulks data from the reader inside the database.
func BulkFileToDB(fileReader *csv.Reader) {

    var lines int = 0
    var header bool = true
    var company Company = Company{}
    
    for {
        
        line, err := fileReader.Read()
        
        if err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }

        if !header {
                    
            company.Id = 0
            company.Name = line[0]
            company.ZipCode = line[1]
            company.Website = ""

            InsertIntoDB(company)

            lines++ 

        } else {
            header = false
        }   
    }

    fmt.Printf("\n|--------------------------------------------|")
    fmt.Printf("\n| %d lines affected.", lines)
    fmt.Printf("\n|--------------------------------------------|\n")
}

//  Updates data from the reader inside the database.
func UpdateFileToDB(fileReader *csv.Reader) {

    var lines int = 0
    var header bool = true
    
    for {
        
        line, err := fileReader.Read()
        
        if err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }

        if !header {

            companies := SelectEqualsDB(line[0], line[1])
            
            if len(companies) == 1 {

                companies[0].Website = line[2]

                UpdateIntoDB(companies[0])

                lines++ 
            }

        } else {
            header = false
        }   
    }

    fmt.Printf("\n|--------------------------------------------|")
    fmt.Printf("\n| %d lines affected.", lines)
    fmt.Printf("\n|--------------------------------------------|\n")
}

//  Starts the webservice from the REST API.
func StartAPI() {

    fmt.Printf("\n|--------------------------------------------|")
    fmt.Printf("\n| The API is now running...")
    fmt.Printf("\n|--------------------------------------------|\n")

    RunServer()
}

//  Clears the database, removing all of its contents.
func ClearInfoInDB() {

    ClearAllDB()

    fmt.Printf("\n|--------------------------------------------|")
    fmt.Printf("\n| Database cleared.")
    fmt.Printf("\n|--------------------------------------------|\n")
}

// Shows the content of "tb_companies" on screen.
func PrintDB() {

    var companies []Company = []Company{}
    var lines int = 0

    companies = SelectAllDB()

    fmt.Printf("\n|--------------------------------------------|")
    fmt.Printf("\n|                TB_COMPANIES                |")
    fmt.Printf("\n|--------------------------------------------|\n")

    for i := 0; i < len(companies); i++ {

        if len(companies[i].Website) > 29 {

            runes := []rune(companies[i].Website)
            companies[i].Website = string(runes[0:29]) + "[...]"
        }

        fmt.Printf("|--------------------------------------------|\n")
        fmt.Printf("| id     : %d\n", companies[i].Id)
        fmt.Printf("| Name   : %s\n", companies[i].Name)
        fmt.Printf("| ZipCode: %s\n", companies[i].ZipCode)     
        fmt.Printf("| Website: %s\n", companies[i].Website)

        lines++
    }

    fmt.Printf("|--------------------------------------------|\n")
    fmt.Printf("| %d lines found.\n", len(companies))
    fmt.Printf("|--------------------------------------------|\n")
}

//  Checks if there was an error during an operation.
func CheckError(err error) {
    
    if err != nil {
        panic(err)
    }
}

/****************************************************/