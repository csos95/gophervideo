package gopherVideo

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()
var documentElement = js.Global.Get("document")
var body = document.DocumentElement().GetElementsByTagName("body")[0].(*dom.HTMLBodyElement)
var cssSet = false

// Player represents a gopher video player
type Player struct {
	ID               string
	URL              string
	Parent           dom.HTMLElement
	Container        *dom.HTMLDivElement
	Video            *dom.HTMLVideoElement
	Controls         *dom.HTMLDivElement
	PlayPause        *dom.HTMLButtonElement
	ProgressBar      *dom.HTMLInputElement
	TimeText         *dom.HTMLPreElement
	DurationText     *dom.HTMLPreElement
	FullscreenButton *dom.HTMLButtonElement
	Duration         int
	Fullscreen       bool
	FirstPlay        bool
	Removed          bool
}

// NewPlayer returns a new gopher video player and the contained video
func NewPlayer(parent dom.HTMLElement, url string) *Player {
	id := "1"

	player := &Player{
		ID:         id,
		URL:        url,
		Parent:     parent,
		Fullscreen: false,
		FirstPlay:  true,
		Removed:    false,
	}

	if !cssSet {
		player.setupCSS()
	}
	player.setupHTML()
	player.setupListeners()

	player.Parent.AppendChild(player.Container)

	return player
}

// Remove the player from the document
func (p *Player) Remove() {
	p.Parent.RemoveChild(p.Container)
	p.Removed = true
}

// formats the time in days:hours:minutes:seconds leaving off empty fields to the left
func (p *Player) timeFormat(seconds int) string {
	if p.Duration < 60 {
		return fmt.Sprintf("%d", seconds)
	} else if p.Duration < 3600 {
		return fmt.Sprintf("%d:%02d", seconds/60, seconds%60)
	} else if p.Duration < 86400 {
		return fmt.Sprintf("%d:%02d:%02d", seconds/3600, seconds/60%60, seconds%60)
	}
	return fmt.Sprintf("%d:%02d:%02d:%02d", seconds/86400, seconds/3600%24, seconds/60%60, seconds%60)
}
