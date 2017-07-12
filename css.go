package gophervideo

import "github.com/gopherjs/gopherjs/js"

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
		top: 0;
		left: 0;
		bottom: 0;
		fill: #fff;
		padding: 10px;
	}
	.GopherVideo-time {
		margin: auto;
		position: absolute;
		top: 10px;
		left: 40px;
		bottom: 10px;
		color: #fff;
	}
	.GopherVideo-progressbar {
		margin: auto;
		position: absolute;
		top: 10px;
		left: 90px;
		bottom: 10px;
		width: 450px;
	}
	.GopherVideo-duration {
		margin: auto;
		position: absolute;
		top: 10px;
		right: 40px;
		bottom: 10px;
		color: #fff;
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
