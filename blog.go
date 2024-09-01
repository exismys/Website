package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type BlogTitle struct {
	Title  string
	Date   time.Time
	Status string
	Slug   string
}

type BlogPost struct {
	Title   string
	Date    time.Time
	Tags    []string
	Content template.HTML
}

func getBlogList() []BlogTitle {
	blogTitles := []BlogTitle{}

	var title string
	var date time.Time
	var status string

	err := filepath.Walk("blogs", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".html" {
			filenameWithoutExt := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))

			dataBytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			parts := bytes.SplitN(dataBytes, []byte("\n\n"), 2)
			if len(parts) < 2 {
				parts = bytes.SplitN(dataBytes, []byte("\r\n\r\n"), 2)
			}
			metadata := parts[0]
			// blogContent := parts[1]

			lines := strings.Split(string(metadata), "\n")
			if len(lines) < 3 {
				return fmt.Errorf("invalid metadata in file %s", path)
			}

			for _, line := range lines {
				if strings.HasPrefix(line, "Title: ") {
					title = strings.TrimSpace(strings.TrimPrefix(line, "Title: "))
				} else if strings.HasPrefix(line, "Date: ") {
					d := strings.TrimSpace(strings.TrimPrefix(line, "Date: "))
					if d != "NA" && d != "" {
						date, err = time.Parse("2006-01-02", d)
						if err != nil {
							return err
						}
						status = "Published"
					}
				}
			}

			if status == "Published" {
				blogTitles = append(blogTitles, BlogTitle{Title: title, Date: date, Status: status, Slug: filenameWithoutExt})
			} else {
				blogTitles = append(blogTitles, BlogTitle{Title: title, Slug: filenameWithoutExt})
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return blogTitles
}

func getBlogPost(blogFile string) BlogPost {
	var title string
	var date time.Time
	var tags []string
	var content string

	dataBytes, err := os.ReadFile(blogFile)
	if err != nil {
		panic(err)
	}

	parts := bytes.SplitN(dataBytes, []byte("\n\n"), 2)
	if len(parts) < 2 {
		parts = bytes.SplitN(dataBytes, []byte("\r\n\r\n"), 2)
	}

	metadata := parts[0]
	blogContent := parts[1]

	lines := strings.Split(string(metadata), "\n")
	if len(lines) < 3 {
		panic(fmt.Errorf("invalid metadata in file %s", blogFile))
	}

	for _, line := range lines {
		if strings.HasPrefix(line, "Title: ") {
			title = strings.TrimSpace(strings.TrimPrefix(line, "Title: "))
		} else if strings.HasPrefix(line, "Date: ") {
			d := strings.TrimSpace(strings.TrimPrefix(line, "Date: "))
			if d != "NA" && d != "" {
				date, err = time.Parse("2006-01-02", d)
				if err != nil {
					panic(err)
				}
			}
		} else if strings.HasPrefix(line, "Tags: ") {
			tags = strings.Split(strings.TrimSpace(strings.TrimPrefix(line, "Tags: ")), ",")
		}
	}

	content = string(blogContent)

	return BlogPost{Title: title, Date: date, Tags: tags, Content: template.HTML(content)}
}
