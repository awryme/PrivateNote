package main

import (
	"fmt"

	"andreyko0/PrivateNote/packages/notes"
	"andreyko0/PrivateNote/packages/notestorage"
	"strconv"

	"io/ioutil"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	storage := notestorage.NewStorage()
	bs, err := ioutil.ReadFile("id.html")
	if err != nil {
		log.Fatal(err)
	}
	idtmpl := string(bs)
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.File("index.html")
	})
	e.GET("/:id", func(context echo.Context) error {
		id := context.Param("id")
		return context.HTML(200, fmt.Sprintf(idtmpl, id))
	})
	e.GET("/files/:name", func(context echo.Context) error {
		return context.File(context.Param("name"))
	})
	gr := e.Group("/api")
	gr.POST("/new/:ttl", func(context echo.Context) error {
		ttl, err := strconv.Atoi(context.Param("ttl"))
		if err != nil {
			context.Error(err)
		}
		var cont struct {
			Cont string `json:"cont"`
		}
		err = context.Bind(&cont)
		if err != nil {
			context.Error(err)
		}
		id := RndString()
		for storage.NoteExists(id) {
			id = RndString()
		}
		n := notes.NewNote(id, cont.Cont, ttl)
		storage.AddNote(n)
		return context.String(200, id)
	})
	gr.GET("/:id", func(context echo.Context) error {
		id := context.Param("id")
		n := storage.GetNote(id)
		if n == nil {
			return context.String(404, "Nothing")
		}
		return context.JSON(200, n)
	})
	host := "localhost:1234"
	fmt.Println("Running on " + host)
	log.Fatal(e.Start(host))
}
