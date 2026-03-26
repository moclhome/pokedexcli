package internal

import (
	"bootdev/go/pokedexcli/internal/pokeapi"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	fileNameStart string = "userdata_"
	fileNameEnd   string = ".txt"
)

var printer Printer

/*func getFileForReading(filename string) (*os.File, error) {
	var f *os.File
	var err error
	if _, err = os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		printer.Println("no file yet")
		f, err = os.Create(filename)
	} else {
		printer.Println("File found!")
		f, err = os.Open(filename)
	}
	if err != nil {
		return nil, fmt.Errorf("Error while reading file: '%v\n", err)
	}
	return f, nil
}*/

func FetchUserDataFromFile(c *Config, user string) (bool, error) {
	printer = c.Printer
	var f *os.File
	var savedData []byte
	var found bool
	filename := fileNameStart + user + fileNameEnd

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		f, err = os.Create(filename)
	} else {
		savedData, err = os.ReadFile(filename)
		if err != nil {
			return false, fmt.Errorf("Error while reading file: '%v\n", err)
		}
	}
	/*f, err := getFileForReading(userDataFileName)
	if err != nil {
		return false, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	n, err := f.Read(savedData)*/
	//c.Printer.Printf("Read bytes: %d\n", n)
	c.File = f // TODO don't safe, fetch two times

	if len(savedData) > 0 {
		if err := fillDataIntoStructure(savedData, c.CaughtPokemons); err != nil {
			return false, err
		}
		found = (len(c.CaughtPokemons) > 0)
	} else {
		printer.Printf("filename: %s; nothing read\n", f.Name())
	}
	return found, nil
}

func fillDataIntoStructure(dataToStore []byte, pokeMap map[string]pokeapi.Pokemon) error {
	var readUserData []pokeapi.Pokemon
	if err := json.Unmarshal(dataToStore, &readUserData); err != nil {
		return fmt.Errorf("Error while unmarshaling user data: '%v\n", err)
	}
	for _, pokemon := range readUserData {
		pokeMap[pokemon.Name] = pokemon
		CurrentCompletionData["pokemon"] = append(CurrentCompletionData["pokemon"], pokemon.Name)
	}
	return nil
}

func WriteUserDataToFile(data []byte, user string) error {
	filename := fileNameStart + user + fileNameEnd
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v\n", err)
	}
	/*written, err := c.File.WriteString(string(data)) TODO: This writes to file. But we have to handle empty pokedex + only writing for this user.
	if err != nil {
		c.Printer.Printf("Error writing to file: %v\n", err)
		return err
	}
	//c.File.Sync()
	*/
	return nil
}

func testFileHandling() {
	f, err := os.Create("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("Hello, How are you?")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
