package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	str "strings"
	"time"

	colour "github.com/fatih/color"
)

// Trims whitespace and lowercases the input string.
func inputCleanse(input string) string {
	return str.TrimSpace(str.ToLower(input))
}

// Checks whether the provided path leads to a directory or a file.
func isDirectory(path string) bool {
	var fileStat, err = os.Stat(path)
	if err != nil {
		fmt.Println("There was an error while checking the entered path. It may not exist, or Year in Review may not have permission to access it.")
		fmt.Print(err)
		os.Exit(1)
	}
	return fileStat.IsDir()
}

/* func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
} */ // For potential later use.

// Checks whether a particular string is present in the Names of the provided []fs.DirEntry, and returns the index (if any).
func containsIdx(slice []fs.DirEntry, name string) (bool, int) {
	for idx, v := range slice {
		if v.Name() == name {
			return true, idx
		}
	}
	return false, -1
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Sorry, but this program doesn't work with Windows. Please use macOS or another UNIX or Unix-like system (eg: Linux, BSD, etc...) or run this program in an online Unix-like IDE (eg: Replit).")
		os.Exit(1)
	}

	fmt.Print("\033c\033[H")

	var boldHiRed = colour.New(colour.Bold, colour.FgHiRed).SprintFunc()
	fmt.Println("Welcome to the unofficial", boldHiRed("Apple Music Year in Review")+".")
	fmt.Println("\n " + boldHiRed("1.") + " First of all, make sure you have obtained a copy of your Apple Media Services information from " + boldHiRed("https://privacy.apple.com") + ". (note this process can take up to 7 days)")
	fmt.Println(" " + boldHiRed("2.") + " Once you have your data, open the '" + boldHiRed("Apple Media Services information") + "' folder you downloaded (this may come as a zip you need to extract).")
	fmt.Println(" " + boldHiRed("3.") + " Then " + boldHiRed("UNZIP") + " the '" + boldHiRed("Apple_Media_Services.zip") + "' file within and open the '" + boldHiRed("Apple_Media_Services") + "' folder.")
	fmt.Println(" " + boldHiRed("4.") + " Within that folder, you will find the folder '" + boldHiRed("Apple Music Activity") + "'. This is the magic folder that you want.")
	fmt.Println("  - It contains all of your Apple Music activity, from your library to your listen history, and so much more.")

	var userHome, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	var input string = ""
	var pathToSearch string = ""
	fmt.Println("\nIt is likely that you downloaded your Apple Music data to '" + boldHiRed(userHome+"/Downloads") + "'.")
	fmt.Print("Would you like for " + boldHiRed("Year in Review") + " to look for your data there? (y/n) ")
	_, err = fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Failed to read your y/n answer.")
		fmt.Println(err)
		os.Exit(1)
	}
	if inputCleanse(input) == "y" {
		pathToSearch = fmt.Sprintf("%s%sDownloads", userHome, string(os.PathSeparator))
	} else if inputCleanse(input) == "n" {
		fmt.Printf("Enter an absolute path (eg: '/Users/me/Documents') to a folder for searching (your home folder is '%s'): ", userHome)
		var reader = bufio.NewReader(os.Stdin)
		pathToSearch, err = reader.ReadString('\n')
		pathToSearch = str.Trim(pathToSearch, "\n")
		if err != nil {
			fmt.Println("Failed to read your entered path.")
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Answer must be 'y' or 'n'.")
		os.Exit(1)
	}
	//fmt.Println(pathToSearch)
	var regex = regexp.MustCompile(`^\/(?:([A-z0-9\/\-\_ ]+)\/?)?$`)
	if !regex.MatchString(pathToSearch) {
		fmt.Println("Not a valid Unix-like path. Remember to use forward slashes '/'!")
		os.Exit(1)
	}
	if !isDirectory(pathToSearch) {
		fmt.Println("The provided path links to a " + boldHiRed("file") + ", not a " + boldHiRed("folder") + ". Make sure you typed the path correctly.")
		os.Exit(1)
	}

	fmt.Printf("Searching '%s'...\n", pathToSearch)
	var parentsPresent = map[string][]int{
		"Apple Media Services information": {0},
		"Apple_Media_Services":             {0},
		"Apple_Media_Services.zip":         {0},
		"Apple Music Activity":             {0},
	}
	dirContents, err := os.ReadDir(pathToSearch)
	if err != nil {
		log.Fatal(err)
	}
	if present, idx := containsIdx(dirContents, "Apple Media Services information"); present {
		if dirContents[idx].IsDir() {
			parentsPresent["Apple Media Services information"] = []int{1, idx} // An int '1' in [0] represents true, and a '0' represents false.
			fmt.Println("Found '" + boldHiRed("Apple Media Services information") + "'.")
		} else {
			fmt.Println("'" + boldHiRed("Apple Media Services information") + "' was present as a " + boldHiRed("file") + ", not a " + boldHiRed("folder") + ". Ensure you typed the correct path.")
			os.Exit(1)
		}
	} else {
		fmt.Println("'" + boldHiRed("Apple Media Services information") + "' was not present. Make sure you typed the right path " + boldHiRed("and") + " have unzipped the file if you received it as a .zip.")
		os.Exit(1)
	}
	dirContents, err = os.ReadDir(pathToSearch + "/Apple Media Services information")
	if err != nil {
		log.Fatal(err)
	}
	if present, idx := containsIdx(dirContents, "Apple_Media_Services.zip"); present {
		if !dirContents[idx].IsDir() {
			parentsPresent["Apple_Media_Services.zip"] = []int{1, idx} // An int '1' in [0] represents true, and a '0' represents false.
		}
	}
	if present, idx := containsIdx(dirContents, "Apple_Media_Services"); present {
		if dirContents[idx].IsDir() {
			parentsPresent["Apple_Media_Services"] = []int{1, idx} // An int '1' in [0] represents true, and a '0' represents false.
			fmt.Println("Found '" + boldHiRed("Apple_Media_Services") + "'.")
		} else {
			fmt.Println("'" + boldHiRed("Apple_Media_Services") + "' was present as a " + boldHiRed("file") + ", not a " + boldHiRed("folder") + ". You may have accidentally renamed or deleted required files/folders.")
			os.Exit(1)
		}
	} else {
		if parentsPresent["Apple_Media_Services.zip"][0] == 1 {
			fmt.Println("You have not unzipped '" + boldHiRed("Apple_Media_Services.zip") + "', as '" + boldHiRed("Apple_Media_Services") + "' was not present.")
			os.Exit(1)
		} else {
			fmt.Println("'" + boldHiRed("Apple_Media_Services") + "' was not present. You may have accidentally renamed or deleted it.")
			os.Exit(1)
		}
	}
	dirContents, err = os.ReadDir(pathToSearch + "/Apple Media Services information/Apple_Media_Services")
	if err != nil {
		log.Fatal(err)
	}
	if present, idx := containsIdx(dirContents, "Apple Music Activity"); present {
		if dirContents[idx].IsDir() {
			parentsPresent["Apple Music Activity"] = []int{1, idx} // An int '1' in [0] represents true, and a '0' represents false.
			fmt.Println("Found '" + boldHiRed("Apple Music Activity") + "'.")
		} else {
			fmt.Println("'" + boldHiRed("Apple Music Activity") + "' was present as a " + boldHiRed("file") + ", not a " + boldHiRed("folder") + ". You may have accidentally renamed or deleted required files/folders.")
			os.Exit(1)
		}
	} else {
		fmt.Println("'" + boldHiRed("Apple Music Activity") + "' was not present. You may have accidentally renamed or deleted it.")
		os.Exit(1)
	}

	dirContents, err = os.ReadDir(pathToSearch + "/Apple Media Services information/Apple_Media_Services/Apple Music Activity")
	if err != nil {
		log.Fatal(err)
	}

	// Marks which CSV files must be present in the 'Apple Music Activity' directory, and their indices within said folder. -1 means the CSV is not present or has not yet been searched for.
	// Note to self: If I want to add more required CSVs to this list, make sure to add them to 'contentsOfCSVs' as well.
	var presentCSVs = map[string]int{
		"Apple Music - Top Content.csv": -1,
		"Apple Music Play Activity.csv": -1,
	}

	for idx, entry := range dirContents {
		if _, ok := presentCSVs[entry.Name()]; ok && !entry.IsDir() {
			presentCSVs[entry.Name()] = idx
		}
	}
	for key, csv := range presentCSVs {
		if csv == -1 {
			fmt.Println("Missing required CSV file '" + boldHiRed(key) + "'. You may have accidentally renamed or deleted said CSV file.")
			os.Exit(1)
		} else {
			fmt.Println("Found '" + boldHiRed(key) + "'.")
		}
	}

	// Contains the entire contents of the CSV files within.
	var contentsOfCSVs = map[string][][]string{
		"Apple Music - Top Content.csv": {},
		"Apple Music Play Activity.csv": {},
	}

	for key := range contentsOfCSVs {
		var contents, err = readAllCSV(key, pathToSearch+"/Apple Media Services information/Apple_Media_Services/Apple Music Activity")
		if err != nil {
			log.Fatalln(err)
		}
		contentsOfCSVs[key] = contents
		fmt.Println("Loaded '" + boldHiRed(key) + "'.")
	}

	var yearForReview = time.Now().Year()
	fmt.Print("\nEnter which year for " + boldHiRed("Year in Review") + " to review. (Leave blank to default to " + boldHiRed(yearForReview) + ") ")
	var reader = bufio.NewReader(os.Stdin)
	tmpYear, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	tmpYear = inputCleanse(tmpYear)
	if tmpYear == "" {
		fmt.Println("Year to Review: " + boldHiRed(yearForReview))
	} else {
		var regex = regexp.MustCompile(`^\d{4}$`)
		if regex.MatchString(tmpYear) {
			tmpYearInt, err := strconv.Atoi(tmpYear)
			if err != nil {
				log.Fatalln(err)
			}
			if tmpYearInt < 2015 { // Year when Apple Music was released.
				fmt.Println("Cannot review years before " + boldHiRed("2015") + ", as that was before Apple Music existed!")
				os.Exit(1)
			} else if tmpYearInt > yearForReview { // 'yearForReview' is still set to current year at this point.
				fmt.Println("Cannot predict your future listening habits (unfortunately).")
				os.Exit(1)
			} else {
				yearForReview = tmpYearInt
				fmt.Println("Year to Review: " + boldHiRed(yearForReview))
			}
		} else {
			fmt.Println("Not a valid year (within the last few thousand years, at least 😆).")
			os.Exit(1)
		}
	}
}
