package ui

import (
    "io"
    "log"
    "bytes"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/storage"
    "sync"
)

var resourceCache = make(map[string]*fyne.StaticResource)
var cacheMutex sync.Mutex

func readImage(file fyne.URIReadCloser) []byte {
	defer file.Close()

	var buf bytes.Buffer
	_, err := io.Copy(&buf, file)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func loadImage(imageName string) *fyne.StaticResource {
    cacheMutex.Lock()
    defer cacheMutex.Unlock()

    if res, ok := resourceCache[imageName]; ok {
        return res
    }

	imageURI, _ := storage.ParseURI("file://resources/" + imageName)
	imageFile, _ := storage.OpenFileFromURI(imageURI)
	imageResource := fyne.NewStaticResource(imageFile.URI().Name(), readImage(imageFile))
	resourceCache[imageName] = imageResource

	return imageResource
}

func GetComputerIcon() *fyne.StaticResource {
    return loadImage("computer.png")
}

func GetHumanIcon() *fyne.StaticResource {
    return loadImage("human.png")
}
