package ui

import (
    "fyne.io/fyne/v2"
    _ "embed"
    "sync"
)

//go:embed resources/computer.png
var computerIcon []byte

//go:embed resources/human.png
var humanIcon []byte

var resourceCache = make(map[string]*fyne.StaticResource)
var cacheMutex sync.Mutex

func getImage(imageName string, bytes *[]byte) *fyne.StaticResource {
    cacheMutex.Lock()
    defer cacheMutex.Unlock()

    if res, ok := resourceCache[imageName]; ok {
        return res
    }

	imageResource := fyne.NewStaticResource(imageName, *bytes)
	resourceCache[imageName] = imageResource

	return imageResource
}

func GetComputerIcon() *fyne.StaticResource {
    return getImage("computer.png", &computerIcon)
}

func GetHumanIcon() *fyne.StaticResource {
    return getImage("human.png", &humanIcon)
}
