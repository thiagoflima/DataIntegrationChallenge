/****************************************************
 *				Packages and Imports				*
 ****************************************************/

package main

import (
    "encoding/csv"
    "io"
    "os"
)

/****************************************************
 *					File functions					*
 ****************************************************/

//	Loads a file from the app folder, returning a reader for it.
func LoadFile(name string) *csv.Reader {

	csvFile, err := os.Open("files/" + name + ".csv")
    CheckError(err)

	reader := csv.NewReader(csvFile)
	reader.Comma = ';'

	return reader
}

//  Saves a file inside the app folder.
func SaveFile(name string, rdr io.Reader) {

    csvFile, err := os.OpenFile("files/" + name, os.O_WRONLY | os.O_CREATE, 0666)
    CheckError(err)

    io.Copy(csvFile, rdr)

    return
}

/****************************************************/