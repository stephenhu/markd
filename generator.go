package main

import (
	"fmt"
	"html"
	"html/template"
  "io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/eknkc/amber"
	"github.com/russross/blackfriday"
)

const (

	LAYOUT 			= "layout.amber"
	ARTICLE     = "article.amber"
  AMBER       = ".amber"
	MD          = ".md"

)

type Article struct {
	Date 			string
	Summary		template.HTML
	Title     string
	URL				string
}

var compiler = amber.New()

func cwd() string {

  d, err := os.Getwd()

  if err != nil {
		color.Red("[Error] %s", err)
		os.Exit(1)
	}

	return d

} // cwd

func readFiles() {

  d := cwd()
	
	files, err := ioutil.ReadDir(d)

  if err != nil {
		
		color.Red("[Error] %s", err)
		os.Exit(1)

	}

	for _, f := range files {

    filename := f.Name()
		if strings.Contains(filename, MD) && filename[0] != '.' {

		  compile(filename)

		}

	}

} // readFiles


func mdArticle(filename string) *Article {

  buf, readErr := ioutil.ReadFile(filename)

  if readErr != nil {

		color.Red("[Error] %s", readErr.Error())
		return nil

	}

	f, statErr := os.Stat(filename)

  if statErr != nil {

		color.Red("[Error] %s", statErr.Error())
		return nil

	}

	article := Article{}

	md := blackfriday.MarkdownCommon(buf)
  
	name := strings.TrimRight(filename, MD)

  color.Magenta(html.EscapeString(string(md))[0:20])
	color.Magenta(html.UnescapeString(string(md))[0:20])
	color.Magenta(string(md)[0:20])

	article.Summary = template.HTML(string(md))
  article.Date		= f.ModTime().String()
	article.Title   = strings.Title(name)
	article.URL			= fmt.Sprintf("%s.html", name)

	return &article

} // mdArticle


func compile(filename string) {

  color.Cyan("+ Generating html for %s", filename)

	parseErr := compiler.ParseFile(ARTICLE)

	if parseErr != nil {
		
		color.Yellow("[Error] %s", parseErr.Error())
		return

	}

	article := mdArticle(filename)

  _, statErr := os.Stat(article.URL)

	if statErr != nil {

    f, createErr := os.Create(article.URL)

		if createErr != nil {

			color.Red("[Error] %s", createErr.Error())
		  return

		}

		defer f.Close()

    tmpl, tmplErr := compiler.Compile()

		if tmplErr != nil {

			color.Red("[ERROR] %s", tmplErr.Error())
			os.Exit(1)

		}

		color.Yellow(article.Date)
		tmpl.Execute(f, article)
		
	} else {

	}

} // compile

func compileIndex() {

} // compileIndex
