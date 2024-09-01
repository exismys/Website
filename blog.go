package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type BlogTitle struct {
	Title  string
	Date   time.Time
	Status string
}

type BlogPost struct {
	Title   string
	Date    time.Time
	Tags    []string
	Content string
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
			dataBytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			parts := bytes.SplitN(dataBytes, []byte("\n\n"), 2)

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
				blogTitles = append(blogTitles, BlogTitle{Title: title, Date: date, Status: status})
			} else {
				blogTitles = append(blogTitles, BlogTitle{Title: title})
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return blogTitles
}
