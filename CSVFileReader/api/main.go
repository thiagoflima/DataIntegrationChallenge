/****************************************************
 *				Packages and Imports				*
 ****************************************************/

package main

import (
    "fmt"   
    "os"
    _ "github.com/mattn/go-sqlite3"
)

/****************************************************
 *					Main Method						*
 ****************************************************/

//	The mmain method.
func main() {

	if len(os.Args) == 2 {

		//	Initializes the database.	
		InitializeDB()

		// Does an operation based on the program argument.
		switch os.Args[1] {

			//	Bulks Q1 inside the database.
			case "-bulk":
				q1 := LoadFile("Q1")
				BulkFileToDB(q1)
			
			// Updates Q2 to the database.
			case "-update":
				q2 := LoadFile("Q2")
				UpdateFileToDB(q2)

			// Runs the API.
			case "-run":
				StartAPI()

			// Clears the database.
			case "-clear":
				ClearInfoInDB()

			//	Prints the database.
			case "-print":
				PrintDB()	
		}

	} else {

		fmt.Printf("\nMissing argument. They must be one of the following:")
		fmt.Printf("\n -bulk  : bulks the Q1 file to the database.")
		fmt.Printf("\n -update: updates the Q2 file into the database.")
		fmt.Printf("\n -run   : runs the REST API in order to do requests.")
		fmt.Printf("\n -clear : delete all the data inside the database.")
		fmt.Printf("\n -print : prints all the data inside the database.\n")
	}	
}

/****************************************************/