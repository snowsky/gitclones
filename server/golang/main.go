package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type ClonedRepos struct {
	Name string
	Url  string
}

var repos = []ClonedRepos{}
var DBFile = "./gitclones.db"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	db, err := sql.Open("sqlite3", DBFile) 
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
        create table IF NOT EXISTS cloned_repos (id integer not null primary key, name text, url text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

//	conn, err := db.Prepare("insert into cloned_repos (id, name, url) values(?, ?, ?)")
//	checkErr(err)
//
//	res, err := conn.Exec(1, "name", "2012-12-09")
//	checkErr(err)
//
//	_, err = res.LastInsertId()
//	checkErr(err)
}

func get_repos() (r []ClonedRepos) {
	db, err := sql.Open("sqlite3", DBFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, name, url from cloned_repos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var url string
		err = rows.Scan(&id, &name, &url)
		if err != nil {
			log.Fatal(err)
		}
		r = append(r, ClonedRepos{name, url})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})

		v1.GET("/repos", func(c *gin.Context) {
			repos = get_repos()
			c.JSON(200, repos)
		})

		v1.POST("/repos", func(c *gin.Context) {
			var repo ClonedRepos
			c.Bind(&repo)
			if repo.Name != "" && repo.Url != "" {
				db, err := sql.Open("sqlite3", DBFile) 
				if err != nil {
					log.Fatal(err)
				}
				defer db.Close()

				sqlStmt := fmt.Sprintf("insert into cloned_repos (name, url) values('%s', '%s')", repo.Name, repo.Url);
				_, err = db.Exec(sqlStmt)
				if err != nil {
					log.Printf("%q: %s\n", err, sqlStmt)
					return
				}
				fmt.Printf("%s\n", repo)
			}
			log.Printf("%s\n", c)
		})
	}

	r.Run(":8080")
}
