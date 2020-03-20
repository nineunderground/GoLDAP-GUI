package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

var tabCont *widget.TabContainer

var nameEntryArray = make([]widget.Entry, 0)
var hostEntryArray = make([]widget.Entry, 0)
var portEntryArray = make([]widget.Entry, 0)
var baseDnEntryArray = make([]widget.Entry, 0)
var userEntryArray = make([]widget.Entry, 0)
var passEntryArray = make([]widget.Entry, 0)
var sslEntryArray = make([]widget.Check, 0)

var nameEntryValues = [5]string{"", "", "", "", ""}
var hostEntryValues = [5]string{"", "", "", "", ""}
var portEntryValues = [5]string{"", "", "", "", ""}
var baseDnEntryValues = [5]string{"", "", "", "", ""}
var userEntryValues = [5]string{"", "", "", "", ""}
var passEntryValues = [5]string{"", "", "", "", ""}
var sslEntryValues = [5]string{"", "", "", "", ""}

// TODO Make rest of ...Values string array

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
				MyApp.Quit()
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
		//widget.NewButton("Quit", func() { MyApp.Quit() }),
		// widget.NewButton("Quit", func() {
		// 	MyApp.Quit()
		// }),
	)
}

// ShowConnectContent It creates and shows the modal window to display LDAP connection profiles
func ShowConnectContent() {
	fmt.Println("ShowConnectContent")
	// TODO Load the saved sessions from local settings...
	var savedSessions = []string{"SESSION_1", "SESSION_2", "SESSION_3", "SESSION_4", "SESSION_5"}
	var tabItems [5]*widget.TabItem

	if Properties == nil {
		panic("Error Properties == nil!")
	}

	for idx, t := range savedSessions {
		sessionPrefix := "SESSION_" + strconv.Itoa(idx+1) + "_"

		nameEntry := widget.NewEntry()
		nameEntry.SetText(Properties[sessionPrefix+"NAME"])
		nameEntryValues[idx] = Properties[sessionPrefix+"NAME"]
		nameEntry.OnChanged = func(newValue string) {
			nameEntryValues[tabCont.CurrentTabIndex()] = newValue
		}
		nameEntryArray = append(nameEntryArray, *nameEntry)

		hostEntry := widget.NewEntry()
		hostEntry.SetText(Properties[sessionPrefix+"HOSTNAME"])
		hostEntryValues[idx] = Properties[sessionPrefix+"HOSTNAME"]
		hostEntry.OnChanged = func(newValue string) {
			selectedIdx := tabCont.CurrentTabIndex()
			hostEntryValues[selectedIdx] = newValue
		}
		hostEntryArray = append(hostEntryArray, *hostEntry)

		portEntry := widget.NewEntry()
		portEntry.SetText(Properties[sessionPrefix+"PORT"])
		portEntryValues[idx] = Properties[sessionPrefix+"PORT"]
		portEntry.OnChanged = func(newValue string) {
			selectedIdx := tabCont.CurrentTabIndex()
			portEntryValues[selectedIdx] = newValue
		}
		portEntryArray = append(portEntryArray, *portEntry)

		baseDNEntry := widget.NewEntry()
		baseDNEntry.SetText(Properties[sessionPrefix+"BASE_DN"])
		baseDnEntryValues[idx] = Properties[sessionPrefix+"BASE_DN"]
		baseDNEntry.OnChanged = func(newValue string) {
			selectedIdx := tabCont.CurrentTabIndex()
			baseDnEntryValues[selectedIdx] = newValue
		}
		baseDnEntryArray = append(baseDnEntryArray, *baseDNEntry)

		userEntry := widget.NewEntry()
		userEntry.SetText(Properties[sessionPrefix+"USER_DN"])
		userEntryValues[idx] = Properties[sessionPrefix+"USER_DN"]
		userEntry.OnChanged = func(newValue string) {
			selectedIdx := tabCont.CurrentTabIndex()
			userEntryValues[selectedIdx] = newValue
		}
		userEntryArray = append(userEntryArray, *userEntry)

		passEntry := widget.NewPasswordEntry()
		passEntry.SetText(Properties[sessionPrefix+"PASSWORD"])
		passEntryValues[idx] = Properties[sessionPrefix+"PASSWORD"]
		passEntry.OnChanged = func(newValue string) {
			selectedIdx := tabCont.CurrentTabIndex()
			passEntryValues[selectedIdx] = newValue
		}
		passEntryArray = append(passEntryArray, *passEntry)

		sslEntry := widget.NewCheck("", nil)
		if Properties[sessionPrefix+"SSL"] == "YES" {
			sslEntry.SetChecked(true)
			sslEntryValues[idx] = Properties[sessionPrefix+"SSL"]
		} else {
			sslEntry.SetChecked(false)
			sslEntryValues[idx] = Properties[sessionPrefix+"SSL"]
		}
		sslEntry.OnChanged = func(isChecked bool) {
			selectedIdx := tabCont.CurrentTabIndex()
			if isChecked {
				sslEntryValues[selectedIdx] = "YES"
			} else {
				sslEntryValues[selectedIdx] = "NO"
			}
		}
		sslEntryArray = append(sslEntryArray, *sslEntry)

		saveButton := widget.NewButton("Save", func() { saveAction() })
		connectButton := widget.NewButton("Conn", func() { conectAction() })
		buttonsLayout := widget.NewHBox(saveButton, connectButton)

		sessionNameForm := &widget.FormItem{Text: "Session name:", Widget: nameEntry}
		hostNameForm := &widget.FormItem{Text: "Hostname:", Widget: hostEntry}
		hostPortForm := &widget.FormItem{Text: "Port:", Widget: portEntry}
		baseDnForm := &widget.FormItem{Text: "Base DN:", Widget: baseDNEntry}
		userDnForm := &widget.FormItem{Text: "User DN:", Widget: userEntry}
		passForm := &widget.FormItem{Text: "Password:", Widget: passEntry}
		sslForm := &widget.FormItem{Text: "SSL:", Widget: sslEntry}
		saveForm := &widget.FormItem{Text: "", Widget: buttonsLayout}
		separatorForm := &widget.FormItem{Text: "", Widget: widget.NewLabel("")}

		form := widget.NewForm(separatorForm, sessionNameForm, hostNameForm, hostPortForm, baseDnForm, userDnForm, passForm, sslForm, saveForm, separatorForm)
		tabItem := widget.NewTabItem(t, form)

		tabItems[idx] = tabItem
	}
	tabCont = widget.NewTabContainer(tabItems[0], tabItems[1], tabItems[2], tabItems[3], tabItems[4])
	tabCont.SelectTabIndex(0)
	connectWinContent := widget.NewVBox(tabCont)
	dialog.ShowCustom("Connect", "Close", connectWinContent, MainWindow)
}

// Button actions
func saveAction() {
	selectedIdx := tabCont.CurrentTabIndex()
	sessionPrefix := "SESSION_" + strconv.Itoa(tabCont.CurrentTabIndex()+1) + "_"

	// Saving into map new values
	properties := Properties
	if Properties == nil {
		panic("Error Properties == nil!")
	}

	properties[sessionPrefix+"NAME"] = nameEntryValues[selectedIdx]
	properties[sessionPrefix+"HOSTNAME"] = hostEntryValues[selectedIdx]
	properties[sessionPrefix+"PORT"] = portEntryValues[selectedIdx]
	properties[sessionPrefix+"BASE_DN"] = baseDnEntryValues[selectedIdx]
	properties[sessionPrefix+"USER_DN"] = userEntryValues[selectedIdx]
	properties[sessionPrefix+"PASSWORD"] = passEntryValues[selectedIdx]
	properties[sessionPrefix+"SSL"] = sslEntryValues[selectedIdx]

	//fmt.Println("Saving, Properties map assigned: ")
	dumpMapToFile(properties)
	//fmt.Println("Saving, dumpMapToFile Done ")
}

func dumpMapToFile(properties map[string]string) {
	// Delete
	var err = os.Remove("./session.profiles")
	if err != nil {
		return
	}
	f, err := os.Create("./session.profiles")
	if err != nil {
		return
	}
	f.Close()
	f, err = os.OpenFile("./session.profiles", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println("Saving new content")
	keys := make([]string, 0, len(properties))
	for k := range properties {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, mapKey := range keys {
		lineToBeWritten := mapKey + "=" + properties[mapKey] + "\n"
		//fmt.Println("dumpMapToFile: saving line:", lineToBeWritten)
		if strings.HasPrefix(lineToBeWritten, "REF:") {
			indexFound := strings.Index(lineToBeWritten, "-")
			if indexFound > -1 {
				lineToBeWritten = lineToBeWritten[indexFound+1:]
			}
		}
		_, err = f.WriteString(lineToBeWritten)
		if err != nil {
			panic(err)
		}
	}

	// Save file changes.
	err = f.Sync()
	if err != nil {
		return
	}
}

func dumpFileToMap() map[string]string {
	return GetSessionProfiles()
}

func conectAction() {
	// TODO
	client := Connect(ModeNonTLS)
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

// GetSessionProfiles ...
func GetSessionProfiles() map[string]string {
	// Create Map
	var properties map[string]string
	properties = make(map[string]string)

	// Scan all properties
	lines, err := scanLines("./session.profiles")
	if err != nil {
		panic(err)
	}

	// Assign entries to map
	cont := 1
	ycont := 1
	for _, line := range lines {
		if strings.HasPrefix(line, "# ") {
			properties["REF:"+strconv.Itoa(cont)+"-"+line] = "Comment"
			cont++
		} else {
			indexFound := strings.Index(line, "=")
			if indexFound == -1 {
				panic("Wrong line")
			}
			//fmt.Println("KEY->" + line[:indexFound])
			//fmt.Println("VALUE->" + line[indexFound+1:])
			properties[line[:indexFound]] = line[indexFound+1:]

			//lineWords := strings.Split(line, "=")
			//properties[lineWords[0]] = lineWords[1]
			ycont++
		}
	}

	//fmt.Println("Total comments detected: ", cont-1)
	//fmt.Println("Total non comment detected: ", ycont-1)
	fmt.Println("Total map: ", len(properties))
	return properties
}

func createDefaultFile() error {
	f, err := os.Create("./session.profiles")
	defer f.Close()
	if err != nil {
		return err
	}
	for _, newLine := range DefaultFileContent {
		lineToBeWritten := newLine + "\n"
		_, err = f.WriteString(lineToBeWritten)
		//err = ioutil.WriteFile("./session.profiles", []byte(lineToBeWritten), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func scanLines(path string) ([]string, error) {
	if _, err := os.Stat("./session.profiles"); os.IsNotExist(err) {
		err = createDefaultFile()
		if err != nil {
			return nil, err
		}
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
