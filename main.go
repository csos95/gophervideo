package gopherVideo

import (
	"fmt"
	"time"

	"github.com/gopherjs/gopherjs/js"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()
var cssSet = false

// Player represents a gopher video player
type Player struct {
	ID          string
	Container   *dom.HTMLDivElement
	Video       *dom.HTMLVideoElement
	Controls    *dom.HTMLDivElement
	PlayPause   *dom.HTMLButtonElement
	ProgressBar *dom.HTMLInputElement
	Time        *dom.HTMLPreElement
	Duration    int
	FirstPlay   bool
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

	controls.AppendChild(bottomControls)
	container.AppendChild(controls)

	player := &Player{
		ID:          id,
		Container:   container,
		Video:       video,
		Controls:    controls,
		PlayPause:   playpause,
		ProgressBar: progressBar,
		Time:        timeText,
		FirstPlay:   true,
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

	p.ProgressBar.AddEventListener("input", true, func(event dom.Event) {
		event.PreventDefault()
		currentTime := p.Video.Get("currentTime").Int()
		p.Video.Set("currentTime", p.ProgressBar.Value)
		p.Time.SetTextContent(fmt.Sprintf("%s/%s", p.timeFormat(currentTime), p.timeFormat(p.Duration)))
	})

	time.Sleep(500 * time.Millisecond)
	p.Duration = p.Video.Get("duration").Int()
	p.ProgressBar.SetAttribute("max", fmt.Sprintf("%d", p.Duration))
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
func setCss() {
	css := `
	.gopherVideo {
		position: relative;
		display: inline-block;
	}
	.gopherVideo-video {
		width: 640px;
		display: flex;
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
