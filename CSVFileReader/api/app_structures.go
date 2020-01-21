/****************************************************
 *				Packages and Imports				*
 ****************************************************/

package main

/****************************************************
 *					Structures						*
 ****************************************************/

//	The company type.
type Company struct {
	Id int
    Name string 
    ZipCode string 
    Website string
}

//	A single HTTP response.
type Response struct {
	Status int
	Message string
}

//	The HTTP response with a set of companies.
type Companies struct {
	Status int
	Message string
	Companies []Company
}

/****************************************************/