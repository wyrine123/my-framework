package main

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log/slog"
	"os"
	"sync"
)

var (
	configsFolderPath = "./configs"
	initConfigOnce          sync.Once
	configFolderWatcherOnce sync.Once
)

// SetConfigsFolderPath set configs folder path according to input string
// panic if folderPath is empty
func SetConfigsFolderPath(folderPath string) {
	if folderPath == "" {
		panic("folder path empty")
	}
	configsFolderPath = folderPath
}


func InitConfig(ctx context.Context) error {
	var err error
	initConfigOnce.Do(func() {
		var f *os.File
		f, err = os.OpenFile("./configs/server.yaml", os.O_RDONLY, 0666)
		if err != nil {
			return
		}
		fmt.Println(f)
		// decoder := yaml.NewDecoder(f)
		// decoder.Decode()
		//
		// go func() {
		//
		// }()
	})

	return err
}

func configFolderWatcher(ctx context.Context) error {
	var err error
	configFolderWatcherOnce.Do(func() {
		var watcher *fsnotify.Watcher
		watcher, err = fsnotify.NewWatcher()
		if err != nil {
			return
		}
		defer watcher.Close()

		go func() {
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					slog.Debug("watch folder event", "event", event)
					if event.Has(fsnotify.Write) {
						// TODO 更新config配置
						slog.Debug("modify file", "file_name", event.Name)
					}
				case err1, ok := <- watcher.Errors:
					if !ok {
						return
					}
					slog.Error("watch folder error", "error", err1)
				case <-ctx.Done():
					slog.Debug("watch folder receive exit signal")
					return
				}
			}
		}()

		// Add a path
		err = watcher.Add(configsFolderPath)
		if err != nil {
			return
		}

		// Block main goroutine forever.
		<-make(chan struct{})
	})

	return err
}
