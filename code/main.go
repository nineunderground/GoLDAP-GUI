package main

import (
	"fmt"
	"os"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/fyne-io/examples/img/icon"
)

// MyApp ...
var MyApp fyne.App = app.New()

// MainWindow ...
var MainWindow fyne.Window

// Documentation: https://godoc.org/fyne.io/fyne

var Properties map[string]string

func maine() {
	Properties = GetSessionProfiles()
	fmt.Println("prop: ", Properties["SESSION_1_BASE_DN"])
	// for idx, prop := range Properties {
	// 	fmt.Println("idx: ", idx)
	// 	fmt.Println("prop: ", prop)
	// }
}

func main() {
	Properties = GetSessionProfiles()

	MyApp.SetIcon(icon.TextEditorBitmap)

	// Basic setup window
	MainWindow = MyApp.NewWindow(MainWindowTitle)
	MainWindow.Resize(fyne.Size{Height: 500, Width: 750})
	MainWindow.CenterOnScreen()

	// Window content
	MainWindow.SetContent(CreateMainContent())

	// Menu bar
	MainWindow.SetMainMenu(CreateMenuBar())

	// Show window
	MainWindow.ShowAndRun()
}

// ******************************************************

// do Test the LDAP actions
func do() {
	run(LdapUser, LdapPass)
}

// run test login and search actions
func run(user string, pass string) {
	client := Connect(ModeNonTLS)
	fmt.Println("user: " + user + " pass: " + pass)
	isLogged := Bind(client, user, pass)
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

	result := Search(client, filterToMatch, AllAttr)
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
	Close(client)
	os.Exit(1)
}
