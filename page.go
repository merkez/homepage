package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"path"
	"strings"
)

type Page struct {
	Title        string
	Body         template.HTML
	RepoURL      string
	IsDir        bool
	Files        []string
	UrlPath      string
	LastModified string
}

var (
	ErrorNotFound = errors.New(`<a href="/search">search</a> and you will find`)
)

// RenderPage renders a file or a directory as an HTML page.
func RenderPage(repoURL string, repoPath string, urlPath string) (*Page, error) {
	dataPath := repoPath + "/data"

	// map URL path to filesystem path (without .md)
	filePath := dataPath + "/" + urlPath

	title := path.Base(filePath)

	if IsDir(filePath) { // if file is a dir return list of files within dir
		files := ListFiles(filePath, false)
		if strings.HasSuffix(filePath, "blog") {
			files = ListFiles(filePath, true)
		}
		return &Page{Title: title, Files: files, RepoURL: repoURL, UrlPath: urlPath, IsDir: true}, nil
	} else if IsFile(filePath + ".md") { // if file is a file return file contents
		lastModified, err := LastModified(repoPath, filePath+".md")
		if err != nil {
			return nil, err
		}

		data, err := ioutil.ReadFile(filePath + ".md")
		if err != nil {
			return nil, err
		}

		body := MdToHtml(data)

		return &Page{Title: title, Body: body, RepoURL: repoURL, UrlPath: urlPath, LastModified: lastModified}, nil
	}

	err := ErrorNotFound
	return nil, err
}
