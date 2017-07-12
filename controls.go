package gopherVideo

import (
	"fmt"
	"time"

	"honnef.co/go/js/dom"

	"github.com/gopherjs/gopherjs/js"
)

// Play the video
func (p *Player) Play() {
	p.Video.Play()
	// if this if the first time the video has been played, duration text and progressbar size
	if p.FirstPlay {
		go func() {
			try := 0
			// sleep to give to video a chance to start loading
			for p.Duration == 0 && p.FirstPlay && !p.Removed {
				try++
				time.Sleep(500 * time.Millisecond)
				p.Duration = p.Video.Get("duration").Int()
				if try == 20 {
					fmt.Println("trying to reload")
					p.Video.RemoveChild(p.Video.FirstChild())
					source := document.CreateElement("source").(*dom.HTMLSourceElement)
					source.SetAttribute("src", p.URL)
					p.Video.AppendChild(source)
					p.Video.Play()
					try = 0
				}
			}
			p.DurationText.SetTextContent(p.timeFormat(p.Duration))
			p.ProgressBar.SetAttribute("max", fmt.Sprintf("%d", p.Duration))
			p.Controls.SetAttribute("style", "display:inline-block;")
			fmt.Printf("%f\n", p.DurationText.OffsetWidth())
			p.Controls.SetAttribute("style", "")
			p.FirstPlay = false
		}()
	}
}

// Pause the video
func (p *Player) Pause() {
	p.Video.Pause()
}

// TogglePlay toggles the play state of the video
func (p *Player) TogglePlay() {
	if p.Video.Paused {
		p.Play()
	} else {
		p.Pause()
	}
}

// Seek the video to the specified time
func (p *Player) Seek(seekTime int) {
	p.Video.Set("currentTime", seekTime)
	p.TimeText.SetTextContent(p.timeFormat(seekTime))
}

// SeekOffset seeks by an offset. a positive offset seeks forward, a negative offset seeks backward
func (p *Player) SeekOffset(seekOffset int) {
	currentTime := p.Video.Get("currentTime").Int()
	seekTime := currentTime + seekOffset
	if seekTime < 0 {
		seekTime = 0
	} else if seekTime > p.Duration {
		seekTime = p.Duration
	}
	p.Seek(seekTime)
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
func (p *Player) toggleFullscreenStyle() {
	if p.Fullscreen {
		p.Container.SetAttribute("style", "width: 640px;")
	} else {
		p.Container.SetAttribute("style", "width:100%;height:100%;top:0;left:0;")
	}
	p.Fullscreen = !p.Fullscreen
}
