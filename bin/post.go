package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Post struct {
	// './mypage.html'
	PostURL string
	// 'Post title'
	PostTitle string
	// 'November 4, 2022'.
	PostDate string
	// Parsed 'PostDate'.
	ParsedDate time.Time
	// 'path/to/image.png'
	PostImage string
	// Tags, e.g., 'poetry', 'prose'.
	Tags []string
	// e.g., 'mypost.md'
	MarkdownFilename string
	// Post content without the title, date, tags, etc.
	PostContent string
	// HTML content after it was rendered by the markdown program,
	// without the title, date, tags, etc.
	HTMLContent string
}

func postifiedFilename(filename string) string {
	filename = path.Base(filename)
	filename = strings.TrimSuffix(filename, path.Ext(filename))
	filename = strings.TrimSuffix(filename, ".")
	filename = filename + ".post"
	return path.Join(outputPostsDirectory, filename)
}

func loadPost(filename string) (Post, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Post{}, err
	}

	var post Post
	if err := json.Unmarshal(data, &post); err != nil {
		return Post{}, err
	}

	return post, nil
}

func loadAllPosts() ([]Post, error) {
	filenames, err := filepath.Glob(fmt.Sprintf("%s/*.post", outputPostsDirectory))
	if err != nil {
		return nil, err
	}

	posts := make([]Post, 0, len(filenames))
	for _, filename := range filenames {
		post, err := loadPost(filename)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func storePost(filename string, post Post) error {
	postJson, err := json.Marshal(post)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	if _, err := outputFile.Write(postJson); err != nil {
		return err
	}

	return outputFile.Close()
}
