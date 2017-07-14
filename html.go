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
	video.SetAttribute("preload", "none")
	video.SetAttribute("src", p.URL)
	container.AppendChild(video)

	// div for the controls
	controls := document.CreateElement("div").(*dom.HTMLDivElement)
	controls.SetClass("GopherVideo-controls")

	bottomControls := document.CreateElement("div").(*dom.HTMLDivElement)
	bottomControls.SetClass("GopherVideo-bottom-controls")

	// a button to play/pause the video
	playpause := createSVG("M0 0v6l6-3-6-3z", "translate(1 1)", "0 0 8 8")
	playpause.SetAttribute("class", "GopherVideo-playpause")
	playpause.SetAttribute("width", "20px")
	playpause.SetAttribute("height", "20px")
	bottomControls.AppendChild(playpause)

	// the current playtime text
	timeText := document.CreateElement("span").(*dom.HTMLSpanElement)
	timeText.SetClass("GopherVideo-time")
	timeText.SetTextContent("0:00")
	bottomControls.AppendChild(timeText)

	// the background of the progress bar
	progressBarBack := document.CreateElement("div").(*dom.HTMLDivElement)
	progressBarBack.SetClass("GopherVideo-progressbar-back")
	bottomControls.AppendChild(progressBarBack)

	// the foreground of the progress bar
	progressBarFront := document.CreateElement("div").(*dom.HTMLDivElement)
	progressBarFront.SetClass("GopherVideo-progressbar-front")
	progressBarBack.AppendChild(progressBarFront)
	bottomControls.AppendChild(progressBarBack)

	// the video duration text
	durationText := document.CreateElement("span").(*dom.HTMLSpanElement)
	durationText.SetClass("GopherVideo-duration")
	durationText.SetTextContent("0:00")
	bottomControls.AppendChild(durationText)

	// a volume icon
	volumeIcon := createSVG("M3.34 0l-1.34 2h-2v4h2l1.34 2h.66v-8h-.66zm1.66 1v1c.17 0 .34.02.5.06.86.22 1.5 1 1.5 1.94s-.63 1.72-1.5 1.94c-.16.04-.33.06-.5.06v1c.25 0 .48-.04.72-.09h.03c1.3-.33 2.25-1.51 2.25-2.91 0-1.4-.95-2.58-2.25-2.91-.23-.06-.49-.09-.75-.09zm0 2v2c.09 0 .18-.01.25-.03.43-.11.75-.51.75-.97 0-.46-.31-.86-.75-.97-.08-.02-.17-.03-.25-.03z", "", "0 0 8 8")
	volumeIcon.SetAttribute("class", "GopherVideo-volume")
	volumeIcon.SetAttribute("width", "20px")
	volumeIcon.SetAttribute("height", "20px")
	bottomControls.AppendChild(volumeIcon)

	// the progress bar for the video
	volumeBar := document.CreateElement("input").(*dom.HTMLInputElement)
	volumeBar.SetClass("GopherVideo-volumebar")
	volumeBar.SetAttribute("type", "range")
	volumeBar.SetAttribute("min", "0")
	volumeBar.SetAttribute("max", "100")
	volumeBar.SetAttribute("orient", "vertical")
	volumeBar.Value = "70"
	bottomControls.AppendChild(volumeBar)

	// a button to enter fullscreen
	fullscreenButton := createSVG("M0 0v4l1.5-1.5 1.5 1.5 1-1-1.5-1.5 1.5-1.5h-4zm5 4l-1 1 1.5 1.5-1.5 1.5h4v-4l-1.5 1.5-1.5-1.5z", "", "0 0 8 8")
	fullscreenButton.SetAttribute("class", "GopherVideo-fullscreen")
	fullscreenButton.SetAttribute("width", "20px")
	fullscreenButton.SetAttribute("height", "20px")
	bottomControls.AppendChild(fullscreenButton)

	controls.AppendChild(bottomControls)
	container.AppendChild(controls)

	p.Container = container
	p.Video = video
	p.Controls = controls
	p.PlayPause = playpause
	p.ProgressBarBack = progressBarBack
	p.ProgressBarFront = progressBarFront
	p.TimeText = timeText
	p.DurationText = durationText
	p.VolumeIcon = volumeIcon
	p.VolumeBar = volumeBar
	p.FullscreenButton = fullscreenButton
}

func objectToBasicHTMLElement(object *js.Object) *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{BasicElement: &dom.BasicElement{BasicNode: &dom.BasicNode{Object: object}}}
}

func createSVG(d, transform, viewbox string) *dom.BasicHTMLElement {
	object := js.Global.Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "svg")
	svg := objectToBasicHTMLElement(object)
	svg.SetAttribute("xmlns", "http://www.w3.org/2000/svg")
	svg.SetAttribute("viewBox", viewbox)

	object = js.Global.Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "path")
	path := objectToBasicHTMLElement(object)
	path.SetAttribute("d", d)
	path.SetAttribute("transform", transform)
	svg.AppendChild(path)

	return svg
}
