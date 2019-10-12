package exporter

import (
	"os"
	"strings"

	"github.com/lunny/html2md"
	"github.com/maateen/meztohugo/config"
	"github.com/maateen/meztohugo/internal/importer"
)

// ExportBlogPosts ...
func ExportBlogPosts(posts []importer.Post) {
	conf := config.GetConfig()
	if _, err := os.Stat(conf.Hugo.ContentDir); os.IsNotExist(err) {
		os.Mkdir(conf.Hugo.ContentDir, os.ModePerm)
	}

	for _, post := range posts {
		f, err := os.Create(conf.Hugo.ContentDir + "/" + post.Slug + ".md")
		defer f.Close()

		if err != nil {
			panic(err)
		}

		f.WriteString("---\n")
		f.WriteString("title: \"" + post.Title + "\"\n")
		f.WriteString("slug: " + post.Slug + "\n")

		if post.Status == 1 {
			f.WriteString("draft: true\n")
		}

		f.WriteString("date: " + post.Created.Format("2006-01-02") + "\n")
		f.WriteString("lastmod: " + post.Updated.Format("2006-01-02") + "\n")
		f.WriteString("publishDate: " + post.PublishDate.Format("2006-01-02") + "\n")

		if post.ExpiryDate.Format("2006-01-02") != "0001-01-01" {
			f.WriteString("expiryDate: " + post.ExpiryDate.Format("2006-01-02") + "\n")
		}

		f.WriteString("type: post\n")

		if post.FeaturedImage != "" {
			f.WriteString("coverImage: /" + post.FeaturedImage + "\n")
		} else {
			f.WriteString("coverImage: /" + conf.Hugo.DefaultCoverImage + "\n")
		}

		f.WriteString("categories:\n")

		for _, category := range post.Category {
			f.WriteString("- " + category + "\n")
		}

		f.WriteString("description: \"" + post.Description + "\"\n")
		f.WriteString("---\n")
		f.WriteString(getMarkdown(post.Content))
	}
}

func getMarkdown(content string) string {
	md := html2md.Convert(content)

	md = strings.ReplaceAll(md, "&nbsp;", " ")
	md = strings.ReplaceAll(md, "&lt;", "<")
	md = strings.ReplaceAll(md, "&gt;", ">")
	md = strings.ReplaceAll(md, "&amp;", "&")
	md = strings.ReplaceAll(md, "&rsquo;", "'")
	md = strings.ReplaceAll(md, "&lsquo;", "'")
	md = strings.ReplaceAll(md, "&ldquo;", "\"")
	md = strings.ReplaceAll(md, "&rdquo;", "\"")
	md = strings.ReplaceAll(md, "static/media/uploads", "img")

	return md
}
