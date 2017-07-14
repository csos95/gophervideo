package gophervideo

import (
	"fmt"
	"strconv"

	"honnef.co/go/js/dom"
)

// Setup the listeners and controls for the video player
func (p *Player) setupListeners() {
	// listener to play/pause the video
	p.playpauseListener = p.PlayPause.AddEventListener("click", false, func(event dom.Event) {
		p.TogglePlay()
	})

	// listener to update the progress bar and time text
	p.videoTimeUpdateListener = p.Video.AddEventListener("timeupdate", false, func(event dom.Event) {
		currentTime := p.Video.Get("currentTime").Int()
		p.TimeText.SetTextContent(p.timeFormat(currentTime))

		if p.Duration != 0 {
			x := currentTime * p.ProgressBarWidth / p.Duration
			p.ProgressBarFront.SetAttribute("style", fmt.Sprintf("width:%dpx;", x))
		}
	})

	// seek to the position clicked on (when seeking forward, the click event goes to the back div)
	p.ProgressBarClickListener = p.ProgressBarBack.AddEventListener("click", false, func(event dom.Event) {
		pageX := event.Underlying().Get("pageX").Int()
		offsetX := p.ProgressBarBack.Get("offsetLeft").Int()
		var containerX int
		if !p.Fullscreen {
			containerX = p.Container.Get("offsetLeft").Int()
		}
		x := pageX - containerX - offsetX

		newTime := x * p.Duration / p.ProgressBarWidth
		p.Seek(newTime)
	})

	// seek through the video by dragging the progress bar
	p.ProgressBarDragListener = p.ProgressBarBack.AddEventListener("mousemove", false, func(event dom.Event) {
		if p.Seeking {
			pageX := event.Underlying().Get("pageX").Int()
			offsetX := p.ProgressBarBack.Get("offsetLeft").Int()
			var containerX int
			if !p.Fullscreen {
				containerX = p.Container.Get("offsetLeft").Int()
			}
			x := pageX - containerX - offsetX

			newTime := x * p.Duration / p.ProgressBarWidth
			p.Seek(newTime)
		}
	})

	// start seeking on mouse down on the back
	p.ProgressBarDownListener = p.ProgressBarBack.AddEventListener("mousedown", false, func(event dom.Event) {
		p.Seeking = true
	})

	// stop seeking on mouse up on the document
	p.ProgressBarUpListener = document.AddEventListener("mouseup", false, func(event dom.Event) {
		p.Seeking = false
	})

	// change the volume dragging the volume bar
	p.volumeBarListener = p.VolumeBar.AddEventListener("input", false, func(event dom.Event) {
		volume, _ := strconv.Atoi(p.VolumeBar.Value)
		fmt.Println(volume)
		p.ChangeVolume(volume)
	})

	// click the fullscreen button to enter/exit fullscreen
	p.fullscreenButtonListener = p.FullscreenButton.AddEventListener("click", false, func(event dom.Event) {
		p.ToggleFullscreenState()
	})

	// fullscreenchange events to toggle the container style
	p.fullscreenListener = document.AddEventListener("fullscreenchange", false, func(event dom.Event) {
		p.toggleFullscreenStyle()
	})
	p.webkitFullscreenListener = document.AddEventListener("webkitfullscreenchange", false, func(event dom.Event) {
		p.toggleFullscreenStyle()
	})
	p.mozillaFullscreenListener = document.AddEventListener("mozfullscreenchange", false, func(event dom.Event) {
		p.toggleFullscreenStyle()
	})
	p.microsoftFullscreenListener = document.AddEventListener("MSFullscreenChange", false, func(event dom.Event) {
		p.toggleFullscreenStyle()
	})

	// keypress event listener for keybinds
	p.keyPressListener = document.AddEventListener("keypress", false, func(event dom.Event) {
		key := event.(*dom.KeyboardEvent).Key
		if input, ok := event.Target().(*dom.HTMLInputElement); ok &&
			input.Attributes()["type"] == "text" {
			fmt.Println("target input")
			return
		}
		if _, ok := event.Target().(*dom.HTMLTextAreaElement); ok {
			fmt.Println("target textarea")
			return
		}
		fmt.Printf("|%s|\n", key)
		switch key {
		case " ", "k":
			p.TogglePlay()
		case "j":
			p.SeekOffset(-10)
		case "l":
			p.SeekOffset(10)
		case "f":
			p.ToggleFullscreenState()
		}
	})
}
