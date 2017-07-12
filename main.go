package gopherVideo

import (
	"fmt"
	"time"

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
}

// NewPlayer returns a new gopher video player and the contained video
func NewPlayer(url string) *Player {
	id := "1"

	player := &Player{
		ID:         id,
		URL:        url,
		Fullscreen: false,
		FirstPlay:  true,
	}

	if !cssSet {
		player.setupCSS()
	}
	player.setupHTML()
	player.setupListeners()

	return player
}

// Play the video
func (p *Player) Play() {
	fmt.Println("in play")
	p.Video.Play()
	// if this if the first time the video has been played, duration text and progressbar size
	if p.FirstPlay {
		go func() {
			// sleep to give to video a chance to start loading
			time.Sleep(500 * time.Millisecond)
			p.Duration = p.Video.Get("duration").Int()
			p.DurationText.SetTextContent(p.timeFormat(p.Duration))
			p.ProgressBar.SetAttribute("max", fmt.Sprintf("%d", p.Duration))
			p.Controls.SetAttribute("style", "display:inline-block;")
			fmt.Printf("%f\n", p.DurationText.OffsetWidth())
			p.Controls.SetAttribute("style", "")
		}()
	}
	p.FirstPlay = false
}

// Pause the video
func (p *Player) Pause() {
	p.Video.Pause()
}

// TogglePlay toggles the play state of the video
func (p *Player) TogglePlay() {
	if p.Video.Paused {
		p.Play()
	} else {
		p.Pause()
	}
}

// ToggleFullscreenState toggles the fullscreen state of the container
func (p *Player) ToggleFullscreenState() {
	if p.Fullscreen {
		if documentElement.Get("exitFullscreen") != js.Undefined {
			documentElement.Call("exitFullscreen")
		} else if documentElement.Get("webkitExitFullscreen") != js.Undefined {
			documentElement.Call("webkitExitFullscreen")
		} else if documentElement.Get("mozCancelFullScreen") != js.Undefined {
			documentElement.Call("mozCancelFullScreen")
		} else if documentElement.Get("msExitFullscreen") != js.Undefined {
			documentElement.Call("msExitFullscreen")
		} else {
			fmt.Println("can't exit fullscreen")
		}
	} else {
		if p.Container.Get("requestFullscreen") != js.Undefined {
			p.Container.Call("requestFullscreen")
		} else if p.Container.Get("webkitRequestFullScreen") != js.Undefined {
			p.Container.Call("webkitRequestFullScreen")
		} else if p.Container.Get("mozRequestFullScreen") != js.Undefined {
			p.Container.Call("mozRequestFullScreen")
		} else if p.Container.Get("msRequestFullscreen") != js.Undefined {
			p.Container.Call("msRequestFullscreen")
		} else {
			fmt.Println("can't enter fullscreen")
		}
	}
}

// ToggleFullscreenStyle toggles the fullscreen style of the container and fullscreen variables
func (p *Player) ToggleFullscreenStyle() {
	if p.Fullscreen {
		p.Container.SetAttribute("style", "width: 640px;")
	} else {
		p.Container.SetAttribute("style", "width:100%;height:100%;top:0;left:0;")
	}
	p.Fullscreen = !p.Fullscreen
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
