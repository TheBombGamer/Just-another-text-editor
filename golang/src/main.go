package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gotk3/gotk3/gtk"
	"github.com/rfd/rfd"
)

func main() {
	gtk.Init(nil)

	application, _ := gtk.ApplicationNew("com.example.advanced_text_editor", 0)
	application.Connect("activate", func() {
		window, _ := gtk.ApplicationWindowNew(application)
		window.SetTitle("Advanced Text Editor")
		window.SetDefaultSize(800, 600)

		vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
		menuBar, _ := gtk.MenuBarNew()
		fileMenu, _ := gtk.MenuNew()
		fileMenuItem, _ := gtk.MenuItemNewWithLabel("File")
		openItem, _ := gtk.MenuItemNewWithLabel("Open")
		saveItem, _ := gtk.MenuItemNewWithLabel("Save")
		exitItem, _ := gtk.MenuItemNewWithLabel("Exit")

		openItem.Connect("activate", func() {
			dialog := rfd.FileDialog{
				Title: "Open File",
			}
			if path, err := dialog.Load(); err == nil {
				content, err := ioutil.ReadFile(path)
				if err == nil {
					fmt.Println(string(content)) 
				}
			}
		})

		saveItem.Connect("activate", func() {
			dialog := rfd.FileDialog{
				Title: "Save File",
			}
			if path, err := dialog.Save(); err == nil {
				err := ioutil.WriteFile(path, []byte("Your content here"), 0644)
				if err != nil {
					fmt.Println("Unable to write file:", err)
				}
			}
		})
		exitItem.Connect("activate", func() {
			gtk.MainQuit()
		})
		fileMenu.Append(openItem)
		fileMenu.Append(saveItem)
		fileMenu.Append(exitItem)
		fileMenuItem.SetSubmenu(fileMenu)
		menuBar.Append(fileMenuItem)
		textView, _ := gtk.TextViewNew()
		scrolledWindow, _ := gtk.ScrolledWindowNew(nil, nil)
		scrolledWindow.Add(textView)
		vbox.PackStart(menuBar, false, false, 0)
		vbox.PackStart(scrolledWindow, true, true, 0)
		window.Add(vbox)
		window.ShowAll()
	})
	application.Run(nil)
}
