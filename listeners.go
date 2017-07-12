package gopherVideo

import (
	"fmt"

	"honnef.co/go/js/dom"
)

// Setup the listeners and controls for the video player
func (p *Player) setupListeners() {
	// listener to play/pause the video
	p.PlayPause.AddEventListener("click", true, func(event dom.Event) {
		event.PreventDefault()
		p.TogglePlay()
	})

	// listener to update the progress bar and time text
	p.Video.AddEventListener("timeupdate", false, func(event dom.Event) {
		event.PreventDefault()
		currentTime := p.Video.Get("currentTime").Int()
		p.ProgressBar.Value = fmt.Sprintf("%d", currentTime)
		p.TimeText.SetTextContent(p.timeFormat(currentTime))
	})

	// seek through the video by dragging the progress bar
	p.ProgressBar.AddEventListener("input", true, func(event dom.Event) {
		event.PreventDefault()
		currentTime := p.Video.Get("currentTime").Int()
		p.Video.Set("currentTime", p.ProgressBar.Value)
		p.TimeText.SetTextContent(p.timeFormat(currentTime))
	})

	// click the fullscreen button to enter/exit fullscreen
	p.FullscreenButton.AddEventListener("click", true, func(event dom.Event) {
		event.PreventDefault()
		p.ToggleFullscreenState()
	})

	// fullscreenchange events to toggle the container style
	document.AddEventListener("fullscreenchange", false, func(event dom.Event) {
		fmt.Println("fullscreen change")
		p.ToggleFullscreenStyle()
	})
	document.AddEventListener("webkitfullscreenchange", false, func(event dom.Event) {
		fmt.Println("fullscreen change")
		p.ToggleFullscreenStyle()
	})
	document.AddEventListener("mozfullscreenchange", false, func(event dom.Event) {
		fmt.Println("fullscreen change")
		p.ToggleFullscreenStyle()
	})
	document.AddEventListener("MSFullscreenChange", false, func(event dom.Event) {
		fmt.Println("fullscreen change")
		p.ToggleFullscreenStyle()
	})

	// keypress event listener for keybinds
	document.AddEventListener("keypress", false, func(event dom.Event) {
		key := event.(*dom.KeyboardEvent).Key
		if _, ok := event.Target().(*dom.HTMLInputElement); ok {
			fmt.Println("target input")
			return
		}
		if _, ok := event.Target().(*dom.HTMLTextAreaElement); ok {
			fmt.Println("target textarea")
			return
		}
		fmt.Printf("|%s|\n", key)
		switch key {
		case " ":
			p.TogglePlay()
		case "f":
			p.ToggleFullscreenState()
		}
	})
}
