package main

import "github.com/mrohadi/snippetbox/pkg/models"

// Define a template data type as the holding structure for
// any dynamic data that we want to pass to out HTMP template.
// At the moment it only contain one field, but we'll add more
// to it as the build process.
type templateStruct struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
