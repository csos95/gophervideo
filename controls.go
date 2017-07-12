package gophervideo

import (
	"fmt"
	"time"

	"honnef.co/go/js/dom"

	"github.com/gopherjs/gopherjs/js"
)

// Play the video
func (p *Player) Play() {
	p.Video.Play()
	playpausePath := p.PlayPause.FirstChild().(dom.Element)
	playpausePath.SetAttribute("d", "M0 0v6h2v-6h-2zm4 0v6h2v-6h-2z")

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
	playpausePath := p.PlayPause.FirstChild().(dom.Element)
	playpausePath.SetAttribute("d", "M0 0v6l6-3-6-3z")
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

// ChangeVolume sets the volume 0-100
func (p *Player) ChangeVolume(volume int) {
	p.Video.Set("volume", float64(volume)*0.01)

	volumePath := p.VolumeIcon.FirstChild().(dom.Element)
	if volume == 0 {
		volumePath.SetAttribute("d", "M3.34 0l-1.34 2h-2v4h2l1.34 2h.66v-8h-.66z")
		volumePath.SetAttribute("transform", "translate(2)")
	} else if volume < 50 {
		volumePath.SetAttribute("d", "M3.34 0l-1.34 2h-2v4h2l1.34 2h.66v-8h-.66zm1.66 3v2c.09 0 .18-.01.25-.03.43-.11.75-.51.75-.97 0-.46-.31-.86-.75-.97-.08-.02-.17-.03-.25-.03z")
		volumePath.SetAttribute("transform", "translate(1)")
	} else {
		volumePath.SetAttribute("d", "M3.34 0l-1.34 2h-2v4h2l1.34 2h.66v-8h-.66zm1.66 1v1c.17 0 .34.02.5.06.86.22 1.5 1 1.5 1.94s-.63 1.72-1.5 1.94c-.16.04-.33.06-.5.06v1c.25 0 .48-.04.72-.09h.03c1.3-.33 2.25-1.51 2.25-2.91 0-1.4-.95-2.58-2.25-2.91-.23-.06-.49-.09-.75-.09zm0 2v2c.09 0 .18-.01.25-.03.43-.11.75-.51.75-.97 0-.46-.31-.86-.75-.97-.08-.02-.17-.03-.25-.03z")
		volumePath.SetAttribute("transform", "")
	}
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
		fullscreenPath := p.FullscreenButton.FirstChild().(dom.Element)
		fullscreenPath.SetAttribute("d", "M0 0v4l1.5-1.5 1.5 1.5 1-1-1.5-1.5 1.5-1.5h-4zm5 4l-1 1 1.5 1.5-1.5 1.5h4v-4l-1.5 1.5-1.5-1.5z")
	} else {
		p.Container.SetAttribute("style", "width:100%;height:100%;top:0;left:0;")
		fullscreenPath := p.FullscreenButton.FirstChild().(dom.Element)
		fullscreenPath.SetAttribute("d", "M1 0l-1 1 1.5 1.5-1.5 1.5h4v-4l-1.5 1.5-1.5-1.5zm3 4v4l1.5-1.5 1.5 1.5 1-1-1.5-1.5 1.5-1.5h-4z")
	}
	p.Fullscreen = !p.Fullscreen
}
