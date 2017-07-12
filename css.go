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
		left: 0;
		bottom: 0;
		fill: #fff;
		padding: 10px;
	}
	.GopherVideo-time {
		margin: auto;
		position: absolute;
		left: 40px;
		bottom: 10px;
		color: #fff;
	}
	.GopherVideo-progressbar {
		margin: auto;
		position: absolute;
		left: 90px;
		bottom: 10px;
		width: 420px;
	}
	.GopherVideo-duration {
		margin: auto;
		position: absolute;
		right: 70px;
		bottom: 10px;
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
