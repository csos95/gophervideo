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
	Container        *dom.HTMLDivElement
	Video            *dom.HTMLVideoElement
	Controls         *dom.HTMLDivElement
	PlayPause        *dom.HTMLButtonElement
	ProgressBar      *dom.HTMLInputElement
	Time             *dom.HTMLPreElement
	FullscreenButton *dom.HTMLButtonElement
	Duration         int
	Fullscreen       bool
	FirstPlay        bool
}

// NewPlayer returns a new gopher video player and the contained video
func NewPlayer(url string) *Player {
	id := "1"

	if !cssSet {
		setCss()
	}

	// div container for the video and controls
	container := document.CreateElement("div").(*dom.HTMLDivElement)
	container.SetClass("gopherVideo")
	container.SetID(fmt.Sprintf("%s", id))

	// the video
	video := document.CreateElement("video").(*dom.HTMLVideoElement)
	video.SetClass("gopherVideo-video")

	// the source for the video
	source := document.CreateElement("source").(*dom.HTMLSourceElement)
	source.SetAttribute("src", url)
	video.AppendChild(source)
	container.AppendChild(video)

	// div for the controls
	controls := document.CreateElement("div").(*dom.HTMLDivElement)
	controls.SetClass("gopherVideo-controls")

	bottomControls := document.CreateElement("div").(*dom.HTMLDivElement)
	bottomControls.SetClass("gopherVideo-bottom-controls")

	// a button to play/pause the video
	playpause := document.CreateElement("button").(*dom.HTMLButtonElement)
	playpause.SetClass("gopherVideo-playpause")
	playpause.SetTextContent("playpause")
	bottomControls.AppendChild(playpause)

	// the progress bar for the video
	progressBar := document.CreateElement("input").(*dom.HTMLInputElement)
	progressBar.SetClass("gopherVideo-progressbar")
	progressBar.SetAttribute("type", "range")
	progressBar.SetAttribute("min", "0")
	progressBar.SetAttribute("max", "1")
	progressBar.Value = "0"
	bottomControls.AppendChild(progressBar)

	// the current playtime text
	timeText := document.CreateElement("pre").(*dom.HTMLPreElement)
	timeText.SetClass("gopherVideo-time")
	timeText.SetTextContent("0:00/0:00")
	bottomControls.AppendChild(timeText)

	// a button to enter fullscreen
	fullscreen := document.CreateElement("button").(*dom.HTMLButtonElement)
	fullscreen.SetClass("gopherVideo-fullscreen")
	fullscreen.SetTextContent("fullscreen")
	bottomControls.AppendChild(fullscreen)

	controls.AppendChild(bottomControls)
	container.AppendChild(controls)

	player := &Player{
		ID:               id,
		Container:        container,
		Video:            video,
		Controls:         controls,
		PlayPause:        playpause,
		ProgressBar:      progressBar,
		Time:             timeText,
		FullscreenButton: fullscreen,
		Fullscreen:       false,
		FirstPlay:        true,
	}

	return player
}

// Play the video
func (p *Player) Play() {
	p.Video.Play()
	// if this if the first time the video has been played, get the duration of the video
	if p.FirstPlay {
		go p.Setup()
	}
}

// Pause the video
func (p *Player) Pause() {
	p.Video.Pause()
}

// TogglePlay toggles the play state of the video
func (p *Player) TogglePlay() {
	if p.Video.Paused {
		p.Video.Play()
	} else {
		p.Video.Pause()
	}
}

// Setup the listeners and controls for the video player
func (p *Player) Setup() {
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
		p.Time.SetTextContent(fmt.Sprintf("%s/%s", p.timeFormat(currentTime), p.timeFormat(p.Duration)))
	})

	// seek through the video by dragging the progress bar
	p.ProgressBar.AddEventListener("input", true, func(event dom.Event) {
		event.PreventDefault()
		currentTime := p.Video.Get("currentTime").Int()
		p.Video.Set("currentTime", p.ProgressBar.Value)
		p.Time.SetTextContent(fmt.Sprintf("%s/%s", p.timeFormat(currentTime), p.timeFormat(p.Duration)))
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
		fmt.Printf("|%s|\n", key)
		switch key {
		case " ":
			p.TogglePlay()
		case "f":
			p.ToggleFullscreenState()
		}
	})

	time.Sleep(500 * time.Millisecond)
	p.Duration = p.Video.Get("duration").Int()
	p.ProgressBar.SetAttribute("max", fmt.Sprintf("%d", p.Duration))
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

// set the css for the gopherVideo player
func setCSS() {
	css := `
	.gopherVideo {
		background-color: #000;
		position: relative;
		width: 640px;
	}
	.gopherVideo-video {
		display: flex;
		width: 100%;
		height: 100%;
	}
	.gopherVideo-controls {
		position: absolute;
		display: none;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background-size: contain;
	}
	.gopherVideo:hover > .gopherVideo-controls {
		display: inline-block;
	}
	.gopherVideo-bottom-controls {
		position: absolute;
		height: 40px;
		left: 0;
		right: 0;
		bottom: 0;
		background-color: rgba(0,0,0,0.5);
	}
	.gopherVideo-playpause {
		margin: auto;
		position: absolute;
		top: 10px;
		left: 10px;
		bottom: 10px;
	}
	.gopherVideo-progressbar {
		margin: auto;
		position: absolute;
		top: 10px;
		left: 100px;
		right: 110px;
		bottom: 10px;
		width: 430px;
	}
	.gopherVideo-time {
		margin: auto;
		position: absolute;
		top: 10px;
		right: 10px;
		bottom: 10px;
		color: #fff;
	}
	`
	style := js.Global.Get("document").Call("createElement", "style")
	style.Set("innerHTML", css)
	// insert style of tooltip at the end of head element
	js.Global.Get("document").Call("getElementsByTagName", "head").Call("item", 0).Call("appendChild", style)
}
