/****************************************************
 *              Packages and Imports                *
 ****************************************************/

package main

import (
    "database/sql"
    "strings"
    _ "github.com/mattn/go-sqlite3"
)

/****************************************************
 *              Basic database funcions             *
 ****************************************************/

//  Initializes the database.
func InitializeDB () bool {

    var query string = ""

    query = query + "CREATE TABLE IF NOT EXISTS `tb_companies` (    "
    query = query + "   `cmp_id` INTEGER PRIMARY KEY AUTOINCREMENT, "
    query = query + "   `cmp_name` VARCHAR(80) NULL,                "
    query = query + "   `cmp_zip` VARCHAR(5) NULL,                  "
    query = query + "   `cmp_website` VARCHAR(100) NULL             "
    query = query + ");                                             "

    db, err := sql.Open("sqlite3", "database/db_companies.db")
    CheckError(err)

    statement, err := db.Prepare(query);
    CheckError(err)

    _, err = statement.Exec()
    CheckError(err)

    db.Close()

    return true
}

//  Inserts a new entry inside the database.
func InsertIntoDB(company Company) bool {

    var query string = ""
    
    query = query + "INSERT INTO `tb_companies`                 "
    query = query + "   (`cmp_name`, `cmp_zip`, `cmp_website`)  "
    query = query + "VALUES (?, ?, ?)                           "

    company.Name = strings.ToUpper(company.Name)
    company.Website = strings.ToLower(company.Website)

    db, err := sql.Open("sqlite3", "database/db_companies.db")
    CheckError(err)

    statement, err := db.Prepare(query);
    CheckError(err)

    _, err = statement.Exec(company.Name, company.ZipCode, company.Website)
    CheckError(err)

    db.Close()

    return true
}

//  Updates a new entry inside the database.
func UpdateIntoDB(company Company) bool {

    var query string = ""
    
    query = query + "UPDATE `tb_companies`                                      "
    query = query + "   SET `cmp_name` = ?, `cmp_zip` = ?, `cmp_website` = ?    "
    query = query + "   WHERE `cmp_id` = ?                                      "

    company.Name = strings.ToUpper(company.Name)
    company.Website = strings.ToLower(company.Website)

    db, err := sql.Open("sqlite3", "database/db_companies.db")
    CheckError(err)

    statement, err := db.Prepare(query);
    CheckError(err)

    _, err = statement.Exec(company.Name, company.ZipCode, company.Website, company.Id)
    CheckError(err)

    db.Close()

    return true
}

//  Selects values from the database based on its id.
func SelectByIdDB(id string) []Company {

    var companies []Company = []Company{}
    var query string = ""
    
    query = query + "SELECT * FROM `tb_companies`   "
    query = query + "   WHERE `cmp_id` = ?          "

    db, err := sql.Open("sqlite3", "database/db_companies.db")
    CheckError(err)

    result, err := db.Query(query, id);
    CheckError(err)

    for result.Next() {         
        
        var id int
        var name string
        var zipcode string
        var website string

        result.Scan(&id, &name, &zipcode, &website)

        companies = append(companies, Company{id, name, zipcode, website})
    }

    db.Close()

    return companies
}

//  Selects values from the database that matches the inputs.
func SelectEqualsDB(name string, zipcode string) []Company {

    var companies []Company = []Company{}
    var query string = ""
    
    query = query + "SELECT * FROM `tb_companies`               "
    query = query + "   WHERE `cmp_name` = ? AND `cmp_zip` = ?  "
    query = query + "   ORDER BY `cmp_id`                       "

    name = strings.ToUpper(name)

    db, err := sql.Open("sqlite3", "database/db_companies.db")
    CheckError(err)

    result, err := db.Query(query, name, zipcode);
    CheckError(err)

    for result.Next() {         
        
        var id int
        var name string
        var zipcode string
        var website string

        result.Scan(&id, &name, &zipcode, &website)

        companies = append(companies, Company{id, name, zipcode, website})
    }

    db.Close()

    return companies
}

//  Selects a value from the database that includes the inputs.
func SelectLikeDB(name string, zipcode string) []Company {

    var companies []Company = []Company{}
    var query string = ""
    
    query = query + "SELECT * FROM `tb_companies`                       "
    query = query + "   WHERE `cmp_name` LIKE ? AND `cmp_zip` LIKE ?    "
    query = query + "   ORDER BY `cmp_id`                               "

    name = "%" + strings.ToUpper(name) + "%"
    zipcode = "%" + zipcode + "%"

    db, err := sql.Open("sqlite3", "database/db_companies.db")
    CheckError(err)

    result, err := db.Query(query, name, zipcode);
    CheckError(err)

    for result.Next() {         
        
        var id int
        var name string
        var zipcode string
        var website string

        result.Scan(&id, &name, &zipcode, &website)

        companies = append(companies, Company{id, name, zipcode, website})
    }

    db.Close()

    return companies
}

//  Selects all the values from the database.
func SelectAllDB() []Company {

    var companies []Company = []Company{}
    var query string = ""
    
    query = query + "SELECT * FROM `tb_companies`   "
    query = query + "   ORDER BY `cmp_id`           "

    db, err := sql.Open("sqlite3", "database/db_companies.db")
    CheckError(err)

    result, err := db.Query(query);
    CheckError(err)

    for result.Next() {         
        
        var id int
        var name string
        var zipcode string
        var website string

        result.Scan(&id, &name, &zipcode, &website)

        companies = append(companies, Company{id, name, zipcode, website})
    }

    db.Close()

    return companies
}

//  Clears the database, removing all of its contents.
func ClearAllDB() bool {

    var query string = ""
    
    query = query + "DELETE FROM `tb_companies` "
    //query = query + "DELETE FROM `sqlite_sequence` WHERE name = `tb_companies`"

    db, err := sql.Open("sqlite3", "database/db_companies.db")
    CheckError(err)

    statement, err := db.Prepare(query);
    CheckError(err)

    _, err = statement.Exec()
    CheckError(err)

    db.Close()

    return true
}

/****************************************************/