package main

import (
	"fmt"

	"github.com/csos95/gopherVideo"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

// get the various elements needed
var (
	document    = dom.GetWindow().Document()
	body        = document.DocumentElement().GetElementsByTagName("body")[0].(*dom.HTMLBodyElement)
	gatewayURL  = document.GetElementByID("gateway-url").(*dom.HTMLInputElement)
	gatewayName = document.GetElementByID("gateway-name").(*dom.HTMLInputElement)
	videoHash   = document.GetElementByID("video-hash").(*dom.HTMLInputElement)
	gateways    = document.GetElementByID("gateway").(*dom.HTMLSelectElement)
	player      *gopherVideo.Player
)

// add a gopher video player
func addVideo(url string) {
	if player != nil {
		player.Remove()
	}
	player = gopherVideo.NewPlayer(body, url)
	player.Play()
}

// add a ipfs gateway
func addGateway(url, name string) {
	gateway := document.CreateElement("option").(*dom.HTMLOptionElement)
	gateway.Value = url
	gateway.SetTextContent(name)
	gateways.AppendChild(gateway)
}

func main() {
	// create two new location options and add them
	// instead of being hardcoded, an api could be called for available locations and those used
	addGateway("http://127.0.0.1:8080/ipfs/", "localhost")
	addGateway("https://ipfs.io/ipfs/", "ipfs.io")
	addGateway("http://45.77.75.143:8080/ipfs/", "vultr")

	// setup the funcmap
	// accessible in js with funcmap.function()
	js.Global.Set("funcmap", map[string]interface{}{
		"addVideo": func() {
			addVideo(fmt.Sprintf("%s%s", gateways.Options()[gateways.SelectedIndex].Value, videoHash.Value))
		},
		"addGateway": func() {
			addGateway(gatewayURL.Value, gatewayName.Value)
		},
	})
}
