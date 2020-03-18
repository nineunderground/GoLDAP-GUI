package main

import (
	"fmt"
	"os"
	"strconv"

	adapi "./modules"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	//fmt.Println("Test 1,2,3!")
	app := app.New()
	// app.SetIcon() TODO

	// Basic setup window
	win := app.NewWindow(adapi.MainWindowTitle)
	win.Resize(fyne.Size{Height: 500, Width: 750})
	win.CenterOnScreen()

	// Window content
	win.SetContent(CreateMainContent(app))

	// Menu bar
	win.SetMainMenu(CreateMenuBar())

	// Show window
	win.ShowAndRun()
}

func CreateMenuBar() *fyne.MainMenu {
	return fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New", func() {
				fmt.Println("File / New")
			}),
		),
		// fyne.NewMenu("Edit",
		// 	fyne.NewMenuItem("Cut", editor.cut),
		// 	fyne.NewMenuItem("Copy", editor.copy),
		// 	fyne.NewMenuItem("Paste", editor.paste),
		// ),
	)
}

// CreateMainContent It creates the main content to display after initial startup
func CreateMainContent(app fyne.App) *widget.Box {
	return widget.NewVBox(
		widget.NewLabel("Hello World!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	)
}

// do Test the LDAP actions
func do() {
	run(adapi.LdapUser, adapi.LdapPass)
}

// run test login and search actions
func run(user string, pass string) {
	client := adapi.Connect(adapi.ModeNonTLS)
	fmt.Println("user: " + user + " pass: " + pass)
	isLogged := adapi.Bind(client, user, pass)
	if isLogged {
		fmt.Println("user: " + user + " logged succesfully.")
	} else {
		fmt.Println("user: " + user + " not logged!")
	}

	var filterToMatch = []string{
		"(cn=uid)",
		"(&(owner=*)(cn=cis-fac))",
		"(&(objectclass=rfc822mailgroup)(cn=*Computer*))",
		"(&(objectclass=rfc822mailgroup)(cn=*Mathematics*))"}

	result := adapi.Search(client, filterToMatch, adapi.AllAttr)
	fmt.Println("Results: ")
	if len(result) > 0 {
		fmt.Println()
		for _, entry := range result {
			fmt.Println("DN: ")
			fmt.Println(entry.DN)
			fmt.Println()
			fmt.Println("Total attributes: " + strconv.Itoa(len(entry.Attributes)))
			for _, attr := range entry.Attributes {
				fmt.Println("Attribute: " + attr.Name)
				for _, val := range attr.Values {
					fmt.Println("Value: " + val)
				}
			}
		}
		fmt.Println()
	}
	adapi.Close(client)
	os.Exit(1)
}
