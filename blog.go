package main

import "time"

type BlogTitle struct {
	Title string
	date  time.Time
}

func getBlogList() ([]BlogTitle, error) {
	return []BlogTitle{}, nil
}
