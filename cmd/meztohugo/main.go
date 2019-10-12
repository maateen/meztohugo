package main

import (
	"github.com/maateen/meztohugo/internal/exporter"
	"github.com/maateen/meztohugo/internal/importer"
)

func main() {
	posts := importer.ImportBlogPosts()
	exporter.ExportBlogPosts(posts)
}
