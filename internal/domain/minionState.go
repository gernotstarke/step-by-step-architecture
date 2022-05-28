package domain

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"log"
	"os"
)

var Config MinionState

type MinionState struct {

	// MinionState represents the "state" of the application -
	// which files/directories are  selected and which processing options are configured,
	// plus status messages

	// directories
	// ************************************************
	sourceDirectory os.File

	sourceDirName string

	// status information (in UI: shown in status line at bottom of UI)
	statusInfo string
} // MinionState

// getter functions for better data encapsulation
// ==============================================
func StatusInfo() string    { return Config.statusInfo }
func SourceDirName() string { return Config.sourceDirName }

// setter functions to avoid uncontrolled changes to global data...
// ================
func SetStatusInfo(msg string) {
	Config.statusInfo = msg
}

// SetSourceDirName also sets status messages
func SetSourceDirName(srcDir string) {
	log.Println("SetSourceDirName called")
	if CheckSrcDirectoryStatus(srcDir) {
		Config.sourceDirName = srcDir
	} else {
		SetStatusInfo("illegal source dir")
	}
}

func SetupConfiguration() {

	setupLanguageNeutralConfig()

	switch lang := checkPreferredLanguage(); lang {
	case "German":
		log.Println("Deutsch als Sprache identifiziert.")
	case "English":
		log.Println("English identified as user language.")
	default:
		log.Println("Unknown language. Falling back to EN.")
	}
}

// directories, default Config options
func setupLanguageNeutralConfig() {

	Config.sourceDirName = GetUserHomeDirectory()

}

func checkPreferredLanguage() string {

	var userPrefs = []language.Tag{
		language.Make("de"), // German
		//language.Make("fr"),  // French
	}

	var serverLangs = []language.Tag{
		language.AmericanEnglish, // en-US fallback
		language.German,          // de
	}

	var matcher = language.NewMatcher(serverLangs)

	tag, _, _ := matcher.Match(userPrefs...)

	fmt.Printf("best match: %s (%s)\n",
		display.English.Tags().Name(tag),
		display.Self.Name(tag))

	return display.English.Tags().Name(tag)
}

// a single PDF file
type singleFileToProcess struct {
	srcDirectory     string
	srcFilename      string
	origPageCount    int
	hasBeenEvenified bool
}

// our "to-do" list with all files that need to be processed
type processingState struct {
	totalPageCount        int
	pagesAlreadyProcessed int

	filesToProcess []singleFileToProcess
}
