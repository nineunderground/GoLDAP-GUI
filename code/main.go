package main

import (
	"fmt"
	"os"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/fyne-io/examples/img/icon"

	adapi "./modules" // TODO Refactor with no folders!
)

//var textMessage string = "Hola"

// Global variables
var myApp fyne.App = app.New()
var mainWindow fyne.Window

// Documentation: https://godoc.org/fyne.io/fyne
func main() {
	//fmt.Println("Test 1,2,3!")
	myApp.SetIcon(icon.TextEditorBitmap)

	// Basic setup window
	mainWindow = myApp.NewWindow(adapi.MainWindowTitle)
	mainWindow.Resize(fyne.Size{Height: 500, Width: 750})
	mainWindow.CenterOnScreen()

	// Window content
	mainWindow.SetContent(CreateMainContent())

	// Menu bar
	mainWindow.SetMainMenu(CreateMenuBar())

	// Show window
	mainWindow.ShowAndRun()

	// mainSettings := myApp.Settings() TODO
}

// CreateMenuBar It creates a menu toolbar
func CreateMenuBar() *fyne.MainMenu {
	return fyne.NewMainMenu(
		fyne.NewMenu("Archivo",
			fyne.NewMenuItem("Connect", func() {
				ShowConnectContent()
			}),
			fyne.NewMenuItem("Disconnect", func() {
				fmt.Println("File / Disconnect")
			}),
			fyne.NewMenuItem("Close", func() {
				fmt.Println("File / Close")
			}),
			fyne.NewMenuItem("Exit", func() {
				fmt.Println("File / Exit")
				myApp.Quit()
			}),
		),
		fyne.NewMenu("Edit"),
		fyne.NewMenu("View"),
		fyne.NewMenu("LDIF"),
		fyne.NewMenu("Help"),
	)
}

// CreateMainContent It creates the main content to display after initial startup
func CreateMainContent() *widget.Box {
	//var ti widget.ToolbarItem
	return widget.NewVBox(
		CreateToolbar(),
		//widget.NewLabel(textMessage),
		//widget.NewButton("Quit", func() { myApp.Quit() }),
		// widget.NewButton("Quit", func() {
		// 	myApp.Quit()
		// }),
	)
}

func radioSelected(sel string) {
	// TODO
}

// ShowConnectContent It creates and shows the modal window to display LDAP connection profiles
func ShowConnectContent() {

	// TODO Load the saved sessions from local settings...
	var savedSessions = []string{"SESSION_1", "SESSION_2", "SESSION_3", "SESSION_4", "SESSION_5"}
	var tabItems [5]*widget.TabItem

	for idx, t := range savedSessions {
		nameEntry := widget.NewEntry()
		hostEntry := widget.NewEntry()
		portEntry := widget.NewEntry()
		baseDNEntry := widget.NewEntry()
		userEntry := widget.NewEntry()
		passEntry := widget.NewPasswordEntry()
		isSSl := widget.NewCheck("", nil)
		saveButton := widget.NewButton("Save", func() { saveAction() })
		connectButton := widget.NewButton("Conn", func() { conectAction() })
		buttonsLayout := widget.NewHBox(saveButton, connectButton)

		sessionNameForm := &widget.FormItem{Text: "Session name:", Widget: nameEntry}
		hostNameForm := &widget.FormItem{Text: "Hostname:", Widget: hostEntry}
		hostPortForm := &widget.FormItem{Text: "Port:", Widget: portEntry}
		baseDnForm := &widget.FormItem{Text: "Base DN:", Widget: baseDNEntry}
		userDnForm := &widget.FormItem{Text: "User DN:", Widget: userEntry}
		passForm := &widget.FormItem{Text: "Password:", Widget: passEntry}
		sslForm := &widget.FormItem{Text: "SSL:", Widget: isSSl}
		saveForm := &widget.FormItem{Text: "", Widget: buttonsLayout}
		separatorForm := &widget.FormItem{Text: "", Widget: widget.NewLabel("")}

		form := widget.NewForm(separatorForm, sessionNameForm, hostNameForm, hostPortForm, baseDnForm, userDnForm, passForm, sslForm, saveForm, separatorForm)
		tabItem := widget.NewTabItem(t, form)

		tabItems[idx] = tabItem
	}
	tabCont := widget.NewTabContainer(tabItems[0], tabItems[1], tabItems[2], tabItems[3], tabItems[4])
	tabCont.SelectTabIndex(0)
	connectWinContent := widget.NewVBox(tabCont)
	dialog.ShowCustom("Connect", "Close", connectWinContent, mainWindow)
}

// Button actions
func saveAction() {
	// TODO Save into properties file
}

func conectAction() {
	client := adapi.Connect(adapi.ModeNonTLS)
	if client != nil {
		client.Close()
	}
}

func disconectAction() {
	// TODO Disconnect & refresh the view
}

// CreateToolbar Creates a toolbar
func CreateToolbar() *widget.Toolbar {
	return widget.NewToolbar(widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
		//e.entry.SetText("")
		fmt.Println("Action 1")
	}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {
			//e.cut()
			fmt.Println("Action 2")
		}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			//e.copy()
			fmt.Println("Action 3")
		}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
			// e.paste()
			fmt.Println("Action 4")
		}))
}

// ******************************************************

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

// ******************************************************
