package main

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

func main() {

	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile("DigiCertGlobalRootCA.crt.pem")
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}
	mysql.RegisterTLSConfig("custom", &tls.Config{RootCAs: rootCertPool})
	var connectionString string
	connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true&tls=custom",
		"obspixieuser", "Test@1995", "obs-pixie-mysql.mysql.database.azure.com", "pixie")
	db, _ := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO Persons (PersonID, LastName, FirstName, Address, City) VALUES (1, 'Tom B. Erichsen', 'Skagen 21', 'Stavanger', '4006');")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Yay, values added!")
}
