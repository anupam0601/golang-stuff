package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// Creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Println("EVENT! %#v\n", event)

				//Watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add("file_to_watch.log"); err != nil {
		fmt.Println("ERROR", err)
	}

	<-done

}
