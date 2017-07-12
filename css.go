package gopherVideo

import "github.com/gopherjs/gopherjs/js"

// set the css for the gopherVideo player
func (p *Player) setupCSS() {
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
		fill: #fff;
	}
	.gopherVideo-time {
		margin: auto;
		position: absolute;
		top: 10px;
		left: 40px;
		bottom: 10px;
		color: #fff;
	}
	.gopherVideo-progressbar {
		margin: auto;
		position: absolute;
		top: 10px;
		left: 90px;
		bottom: 10px;
		width: 370px;
	}
	.gopherVideo-duration {
		margin: auto;
		position: absolute;
		top: 10px;
		right: 100px;
		bottom: 10px;
		color: #fff;
	}
	.gopherVideo-fullscreen {
		margin: auto;
		position: absolute;
		top: 10px;
		right: 10px;
		bottom: 10px;
		fill: #fff;
	}
	`
	style := js.Global.Get("document").Call("createElement", "style")
	style.Set("innerHTML", css)
	// insert style of tooltip at the end of head element
	js.Global.Get("document").Call("getElementsByTagName", "head").Call("item", 0).Call("appendChild", style)
}
