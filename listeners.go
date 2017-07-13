package gophervideo

import (
	"fmt"
	"strconv"

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
		p.TimeText.SetTextContent(p.timeFormat(currentTime))

		left := 40 + p.TimeTextWidth + 10
		x := currentTime * p.ProgressBarWidth / p.Duration
		p.ProgressBarFront.SetAttribute("style", fmt.Sprintf("left:%dpx;width:%dpx;", left, x))
		fmt.Println(fmt.Sprintf("width:%dpx;", x))
	})

	// seek through the video by dragging the progress bar
	// p.ProgressBarBack.AddEventListener("input", true, func(event dom.Event) {
	// 	event.PreventDefault()
	// 	seekTime, _ := strconv.Atoi(p.ProgressBar.Value)
	// 	p.Seek(seekTime)
	// })

	// back.addEventListener('click', function(e) {
	// 		let x = e.pageX - back.offsetLeft;
	// 	let y = e.pageY - back.offsetTop;
	// 	front.setAttribute('style', 'width:' + x + 'px;');
	// 	currentTime = x * maxTime / maxWidth;
	// 	console.log(currentTime);
	// })

	// back.addEventListener('mousemove', function(e) {
	// 	if (seeking) {
	// 		let x = e.pageX - back.offsetLeft;
	// 		let y = e.pageY - back.offsetTop;
	// 		front.setAttribute('style', 'width:' + x + 'px;');
	// 		currentTime = x * maxTime / maxWidth;
	// 		console.log(currentTime);
	// 	}
	// })

	// change the volume dragging the volume bar
	p.VolumeBar.AddEventListener("input", true, func(event dom.Event) {
		event.PreventDefault()
		volume, _ := strconv.Atoi(p.VolumeBar.Value)
		fmt.Println(volume)
		p.ChangeVolume(volume)
	})

	// click the fullscreen button to enter/exit fullscreen
	p.FullscreenButton.AddEventListener("click", true, func(event dom.Event) {
		event.PreventDefault()
		p.ToggleFullscreenState()
	})

	// fullscreenchange events to toggle the container style
	document.AddEventListener("fullscreenchange", false, func(event dom.Event) {
		p.toggleFullscreenStyle()
	})
	document.AddEventListener("webkitfullscreenchange", false, func(event dom.Event) {
		p.toggleFullscreenStyle()
	})
	document.AddEventListener("mozfullscreenchange", false, func(event dom.Event) {
		p.toggleFullscreenStyle()
	})
	document.AddEventListener("MSFullscreenChange", false, func(event dom.Event) {
		p.toggleFullscreenStyle()
	})

	// keypress event listener for keybinds
	document.AddEventListener("keypress", false, func(event dom.Event) {
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
