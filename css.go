package gophervideo

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

// set the css for the GopherVideo player
func (p *Player) setupCSS() {
	css := `
	.GopherVideo {
		background-color: #000;
		position: relative;
		width: 640px;
	}
	.GopherVideo-video {
		display: flex;
		width: 100%;
		height: 100%;
	}
	.GopherVideo-controls {
		position: absolute;
		display: none;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background-size: contain;
	}
	.GopherVideo:hover > .GopherVideo-controls {
		display: inline-block;
	}
	.GopherVideo-bottom-controls {
		position: absolute;
		height: 40px;
		left: 0;
		right: 0;
		bottom: 0;
		background-color: rgba(0,0,0,0.5);
	}
	.GopherVideo-playpause {
		margin: auto;
		position: absolute;
		left: 0;
		bottom: 0;
		fill: #fff;
		padding: 10px;
	}
	.GopherVideo-time {
		margin: auto;
		position: absolute;
		font-size: 20px;
		left: 40px;
		bottom: 8px;
		color: #fff;
	}
	.GopherVideo-progressbar {
		margin: auto;
		position: absolute;
		left: 90px;
		bottom: 10px;
		width: 420px;
	}
	.GopherVideo-progressbar-back {
		margin: auto;
		position: absolute;
		left: 90px;
		bottom: 15px;
		width: 420px;
		height: 10px;
		background-color: #666;
	}
	.GopherVideo-progressbar-front {
		margin: auto;
		position: absolute;
		left: 0px;
		bottom: 0px;
		width: 420px;
		height: 10px;
		background-color: #ccc;
	}
	.GopherVideo-duration {
		margin: auto;
		position: absolute;
		font-size: 20px;
		right: 70px;
		bottom: 8px;
		color: #fff;
	}
	.GopherVideo-volume {
		margin: auto;
		position: absolute;
		right: 40px;
		bottom: 0;
		fill: #fff;
		padding: 10px 0px; 10px; 10px;
	}
	.GopherVideo-volumebar {
		margin: auto;
		position: absolute;
		display: none;
		right: 30px;
		bottom: 40px;
		width: 20px;
		height: 40px;
		padding: 10px;
		-webkit-appearance: slider-vertical; /* WebKit */
	}
	.GopherVideo-volume:hover~.GopherVideo-volumebar {
		display: inline-block;
	}
	.GopherVideo-volumebar:hover {
		display: inline-block;
	}
	.GopherVideo-fullscreen {
		margin: auto;
		position: absolute;
		top: 0;
		right: 0;
		bottom: 0;
		fill: #fff;
		padding: 10px;
	}
	`
	style := js.Global.Get("document").Call("createElement", "style")
	style.Set("innerHTML", css)
	// insert style of tooltip at the end of head element
	js.Global.Get("document").Call("getElementsByTagName", "head").Call("item", 0).Call("appendChild", style)
}

func (p *Player) styleProgressBar() {
	p.Controls.SetAttribute("style", "display:inline-block;")
	p.TimeTextWidth = int(p.DurationText.OffsetWidth())

	// left is the distance from the left side, right is from the right side
	// the first values are the sizes hardcoded in css, second is the width of the time text elements,
	// and third is the space to put between the progress bar and the other elements
	left := 40 + p.TimeTextWidth + 10
	right := 70 + p.TimeTextWidth + 10
	// width is how wide the progress bar needs to be to fill the space
	p.ProgressBarWidth = int(p.Video.OffsetWidth()) - left - right

	currentTime := p.Video.Get("currentTime").Int()
	x := currentTime * p.ProgressBarWidth / p.Duration

	p.ProgressBarBack.SetAttribute("style", fmt.Sprintf("left:%dpx;width:%dpx;", left, p.ProgressBarWidth))
	p.ProgressBarFront.SetAttribute("style", fmt.Sprintf("width:%dpx;", x))

	p.Controls.SetAttribute("style", "")
}

func (p *Player) progressBarUpdate() {
	currentTime := p.Video.Get("currentTime").Int()

	if p.Duration != 0 {
		x := currentTime * p.ProgressBarWidth / p.Duration
		p.ProgressBarFront.SetAttribute("style", fmt.Sprintf("width:%dpx;", x))
	}
}
