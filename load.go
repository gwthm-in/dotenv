package dotenv

import "github.com/fsnotify/fsnotify"

var d = new()

func Load() error {
	return d.Load()
}

func Reset() {
	d = new()
}

func Overload() error {
	return d.Overload()
}

func OverloadWatchFiles() error {
	return d.OverloadWatchFiles()
}

func OptLookupGit() {
	d.mu.Lock()
	d.opts.lookupGit = true
	d.mu.Unlock()
}

func OptLookupMod() {
	d.mu.Lock()
	d.opts.lookupMod = true
	d.mu.Unlock()
}

func OptDisableFileExpand() {
	d.mu.Lock()
	d.opts.disableFileExpand = true
	d.mu.Unlock()
}

func OptDisablePathExpand() {
	d.mu.Lock()
	d.opts.disablePathExpand = true
	d.mu.Unlock()
}

func OptDebug() {
	d.mu.Lock()
	d.opts.debug = true
	d.mu.Unlock()
}

func OptLookupFile(file string) {
	d.mu.Lock()
	d.opts.lookupFile = append(d.opts.lookupFile, file)
	d.mu.Unlock()
}

func OptLookupWatchFile(file string) {
	d.opts.lookupWatchFile = append(d.opts.lookupWatchFile, file)
}

func WatchConfig() {
	d.opts.watchConfig = true
	files := d.opts.ParseWatchFilePaths()
	for _, file := range files {
		go d.WatchConfig(file)
	}
}

func OnConfigChange(fn func(fsnotify.Event)) {
	d.opts.onConfigChange = fn
}
