package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// HandleSearch handles requests for /search. It searches paths and contents of
// files in data folder.
func HandleSearch(w http.ResponseWriter, r *http.Request) {
	pattern := r.URL.Query().Get("regexp")

	var searchOnlyPath bool

	if strings.HasPrefix(pattern, "path:") {
		searchOnlyPath = true
		pattern = strings.TrimPrefix(pattern, "path:")
	}

	rx, err := regexp.Compile("(?i)" + pattern) // make pattern case insensitive
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var foundFiles []string

	err = filepath.Walk(repoPath+"/data", func(path string, info os.FileInfo, err error) error {
		matchedFilePath := GrepFilePath(path, rx)
		matchedFileContent := false

		if !searchOnlyPath && !info.IsDir() {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			html := MdToHtml(data)
			matchedFileContent = GrepFileContent(string(html), rx)
		}

		if matchedFilePath || matchedFileContent {
			urlPath, err := FilesystemToURL(path)
			if err != nil {
				return err
			}
			foundFiles = append(foundFiles, urlPath)
		}

		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p := &Page{
		Title: "search",
		IsDir: true,
		Files: foundFiles,
	}

	t, err := template.New("page.html").
		Funcs(template.FuncMap{"removeTrailingSlash": RemoveTralingSlash}).
		ParseFiles("template/page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := t.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandleRest handles all requests except for /search.
func HandleRest(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path[1:] // remove leading "/"

	if urlPath == "" {
		http.Redirect(w, r, "/about", http.StatusFound)
	}

	p, err := RenderPage(repoURL, repoPath, urlPath)
	if err != nil {
		if err == ErrorNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, err.Error())
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	t, err := template.New("page.html").
		Funcs(template.FuncMap{"removeTrailingSlash": RemoveTralingSlash}).
		ParseFiles("template/page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
