package gophervideo

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"

	"honnef.co/go/js/dom"
)

// NewPlayer returns a new gopher video player and the contained video
func (p *Player) setupHTML() {

	// div container for the video and controls
	container := document.CreateElement("div").(*dom.HTMLDivElement)
	container.SetClass("GopherVideo")
	container.SetID(fmt.Sprintf("%s", p.ID))

	// the video
	video := document.CreateElement("video").(*dom.HTMLVideoElement)
	video.SetClass("GopherVideo-video")

	// the source for the video
	source := document.CreateElement("source").(*dom.HTMLSourceElement)
	source.SetAttribute("src", p.URL)
	video.AppendChild(source)
	container.AppendChild(video)

	// div for the controls
	controls := document.CreateElement("div").(*dom.HTMLDivElement)
	controls.SetClass("GopherVideo-controls")

	bottomControls := document.CreateElement("div").(*dom.HTMLDivElement)
	bottomControls.SetClass("GopherVideo-bottom-controls")

	// a button to play/pause the video
	object := js.Global.Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "svg")
	playpause := objectToBasicHTMLElement(object)
	playpause.SetAttribute("xmlns", "http://www.w3.org/2000/svg")
	playpause.SetAttribute("class", "GopherVideo-playpause")
	playpause.SetAttribute("width", "20px")
	playpause.SetAttribute("height", "20px")
	playpause.SetAttribute("viewBox", "0 0 8 8")

	object = js.Global.Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "path")
	playpausePath := objectToBasicHTMLElement(object)
	playpausePath.SetAttribute("d", "M0 0v6l6-3-6-3z")
	playpausePath.SetAttribute("transform", "translate(1 1)")
	playpause.AppendChild(playpausePath)
	bottomControls.AppendChild(playpause)

	// the current playtime text
	timeText := document.CreateElement("pre").(*dom.HTMLPreElement)
	timeText.SetClass("GopherVideo-time")
	timeText.SetTextContent("0:00")
	bottomControls.AppendChild(timeText)

	// the progress bar for the video
	progressBar := document.CreateElement("input").(*dom.HTMLInputElement)
	progressBar.SetClass("GopherVideo-progressbar")
	progressBar.SetAttribute("type", "range")
	progressBar.SetAttribute("min", "0")
	progressBar.SetAttribute("max", "1")
	progressBar.Value = "0"
	bottomControls.AppendChild(progressBar)

	// the video duration text
	durationText := document.CreateElement("pre").(*dom.HTMLPreElement)
	durationText.SetClass("GopherVideo-duration")
	durationText.SetTextContent("0:00")
	bottomControls.AppendChild(durationText)

	// a button to enter fullscreen
	object = js.Global.Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "svg")
	fullscreen := objectToBasicHTMLElement(object)
	fullscreen.SetAttribute("xmlns", "http://www.w3.org/2000/svg")
	fullscreen.SetAttribute("class", "GopherVideo-fullscreen")
	fullscreen.SetAttribute("width", "20px")
	fullscreen.SetAttribute("height", "20px")
	fullscreen.SetAttribute("viewBox", "0 0 8 8")

	object = js.Global.Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "path")
	fullscreenPath := objectToBasicHTMLElement(object)
	fullscreenPath.SetAttribute("d", "M0 0v4l1.5-1.5 1.5 1.5 1-1-1.5-1.5 1.5-1.5h-4zm5 4l-1 1 1.5 1.5-1.5 1.5h4v-4l-1.5 1.5-1.5-1.5z")
	fullscreen.AppendChild(fullscreenPath)
	bottomControls.AppendChild(fullscreen)

	controls.AppendChild(bottomControls)
	container.AppendChild(controls)

	p.Container = container
	p.Video = video
	p.Controls = controls
	p.PlayPause = playpause
	p.ProgressBar = progressBar
	p.TimeText = timeText
	p.DurationText = durationText
	p.FullscreenButton = fullscreen
}

func objectToBasicHTMLElement(object *js.Object) *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{object}}}
}
