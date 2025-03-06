package main

import "snippetbox.leorafaelmb.net/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
