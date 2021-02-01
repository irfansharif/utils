package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var api = "https://www.googleapis.com/books/v1/volumes?q=intitle:%s"

type Result struct {
	TotalItems int
	Items      []*Item
}

type Item struct {
	VolumeInfo `json:"volumeInfo"`
}

type VolumeInfo struct {
	Title         string
	Subtitle      string
	Authors       []string
	Publisher     string
	PublishedDate string
	Description   string
}

func fetch(title string) (*book, error) {
	title = url.QueryEscape(title)
	url := fmt.Sprintf(api, title)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code")
	}

	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if result.TotalItems == 0 {
		return nil, errors.New("no results found")
	}

	b := result.Items[0]
	book := book{
		Title:   b.Title,
		Details: b.Subtitle,
	}
	if len(b.Authors) != 0 {
		book.Author = b.Authors[0]
	}
	return &book, nil
}
