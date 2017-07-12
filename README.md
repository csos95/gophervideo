# gopherVideo
A video player written with gopherjs.

SVG icons provided by [Open Iconic](www.useiconic.com/open)

## Installation and usage

1. Install or update  
	`go get -u github.com/csos95/gopherVideo`

2. Write a GopherJS project that uses gopherVideo  
	Simple example that adds a video to the page:
	```Go
	package main

	import (
		"github.com/csos95/gopherVideo"
		"honnef.co/go/js/dom"
	)

	func main() {
		// get the document and body elements
		document := dom.GetWindow().Document()
		body := document.DocumentElement().GetElementsByTagName("body")[0].(*dom.HTMLBodyElement)

		// append a new gopherVideo player to the body
		player := gopherVideo.NewPlayer(body, "http://example.com/video.mp4")
	}
	```
3. Run `gopherjs build -m -o myscript.js` to compile a minified version
4. Use the script in a html file  
	`<script type="text/javascript" src="myscript.js"></script>`
  

## Features/Todo List

| Feature | Status | Notes |
|-|-|-|
| play/pause | done |
| fullscreen | done |
| time/duration text | working, not decorated |
| progress bar | working, not decorated |
| volume bar | not started |
| show controls on hover | done |
| keybinds | done | see [Controls](#controls) |
| title bar | not started |
| close button | not started | will be optional |
| settings cog | not started | for overflow settings |
| playback speed | not started |
| buffering animation | not started |
| show buffered data on progress bar | not started | will do when decorating the progress bar |
| resize elements as needed | not started | right now, this means resize the progress bar to fill the space in fullscreen |
| show time on hover over progress bar | not started |
| click in window to play/pause | not started |
| double click in window to enter/exit fullscreen | not started |

## Controls
| Key | Action |
|-|-|
| space | play/pause |
| f | enter/exit fullscreen |
| k | play/pause |
| j | go backward 10 seconds |
| l | go forward 10 seconds |

## Frequently Asked Questions<sup><sup><sup>that no one has actually asked yet</sup></sup></sup>

### Why should I use this?

You probably shouldn't. This project is not very mature and was started so I could work with GopherJS in a project

### The Javascript file is massive!

GopherJS compiles the Go runtime and all dependencies into pure Javascript.  
Because of this, the output files can get pretty big.  
If you use the `-m` flag on the GopherJS compiler and gzip the output, it helps a lot.
