package main

import (
	"os"
	"sort"
	"text/template"
)

func genAllPosts() error {
	tmpl, err := template.ParseFiles(pageTemplateName, allPostsTemplateName)
	if err != nil {
		return err
	}

	posts, err := loadAllPosts()
	if err != nil {
		return err
	}

	sort.Slice(posts, func(i, j int) bool {
		t1 := posts[i].ParsedDate
		t2 := posts[j].ParsedDate
		// Sort in descending order (newest to oldest).
		return t2.Before(t1)
	})

	blogConfig["Posts"] = posts
	return tmpl.Execute(os.Stdout, blogConfig)
}
