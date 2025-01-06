package main

import (
	"os"
	"text/template"
)

func genAllPosts() error {
	tmpl, err := template.ParseFiles(pageTemplateName, allPostsTemplateName)
	if err != nil {
		return err
	}

	posts, err := loadAllPostsSortedDescending()
	if err != nil {
		return err
	}

	blogConfig["Posts"] = posts
	return tmpl.Execute(os.Stdout, blogConfig)
}
