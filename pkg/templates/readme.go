package templates

import (
	"bytes"
	"context"
	_ "embed"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/brittonhayes/brittonhayes/pkg/spotify"
	"github.com/joho/godotenv"
)

//go:embed readme.tmpl
var readme string

type Document struct {
	Title    string
	Subtitle string
	File     string
	Sections map[string]Section
}

type Section struct {
	Title    string
	Subtitle string
	Items    []string
	Images   []string
}

func (d *Document) Render() string {

	t := template.Must(template.New("template").Funcs(template.FuncMap{
		"myPlaylists": func() string {
			ctx := context.Background()
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}

			user := os.Getenv("SPOTIFY_USER_ID")
			playlists, err := spotify.UserPlaylists(ctx, spotify.New(), user)
			if err != nil {
				log.Fatal(err)
			}

			return strings.Join(playlists, "\n\n")
		},
	}).Parse(readme))

	buf := bytes.NewBuffer([]byte{})

	err := t.Execute(buf, &d)
	if err != nil {
		log.Fatal(err)
	}

	if d.File != "" {
		f, err := os.Create(d.File)
		if err != nil {
			log.Fatal(err)
		}

		_, err = f.Write(buf.Bytes())
		if err != nil {
			log.Fatal(err)
		}
	}

	return buf.String()
}
