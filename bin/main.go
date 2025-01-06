package main

import (
	"context"
	"flag"
	"fmt"
	"os"
)

const (
	pageTemplateName     = "templates/page.template"
	allPostsTemplateName = "templates/all_posts.template"
	postTemplateName     = "templates/post.template"


	// Date format that the author uses in the Markdown posts.
	postParseDateFormat = "2006/01/02"
	// Date format that the final rendered blog uses.
	postDisplayDateFormat = "January 02, 2006"

	outputDirectory      = "out"
	outputDistDirectory  = "out/dist"
	outputPostsDirectory = "out/posts"

	authorEmail     = "jadesmith@email.com"
	authorName      = "Jade Smith"
	authorURL       = "https://github.com/jadesmith"
	blogDescription = "Jade Smith's cool comic"
	blogImage       = "images/cpphs-thumbnail.jpg"
	blogLanguage    = "en"
	blogName        = "Cool comic"
	license         = "&copy;"
)

var blogConfig = map[string]interface{}{
	"Title":                  blogName,
	"BlogDescription":        blogDescription,
	"BlogImage":              blogImage,
	"BlogLanguage":           blogLanguage,
	"BlogName":               blogName,
	"License":                license,
	"AuthorURL":              authorURL,
	"AuthorName":             authorName,
	"AuthorEmail":            authorEmail,
	"PostDisplayDateFormat":  postDisplayDateFormat,
}

func main() {
	ctx := context.Background()

	genAllPostsCmd := flag.NewFlagSet("gen-all-posts", flag.ExitOnError)
	genPostCmd := flag.NewFlagSet("gen-post", flag.ExitOnError)
	postifyCmd := flag.NewFlagSet("postify", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected subcommand, e.g., 'gen-all-posts', 'gen-all-tags', etc")
		os.Exit(1)
	}

	command := os.Args[1]
	var err error
	switch command {
	case "gen-all-posts":
		genAllPostsCmd.Parse(os.Args[2:])
		err = genAllPosts()

	case "gen-post":
		genPostCmd.Parse(os.Args[2:])

		args := genPostCmd.Args()
		if len(args) <= 0 {
			fmt.Fprintf(os.Stderr, "command %q expects 1 argument\n", command)
			os.Exit(1)
		}

		err = genPost(args[0])

	case "postify":
		postifyCmd.Parse(os.Args[2:])

		if len(postifyCmd.Args()) != 1 {
			fmt.Fprintf(os.Stderr, "command %q expects 1 argument\n", command)
			os.Exit(1)
		}

		err = postify(ctx, postifyCmd.Args()[0])

	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n", command)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute command %q: %v\n", command, err)
		os.Exit(1)
	}
}
