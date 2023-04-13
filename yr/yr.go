package yr

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
	"github.com/Trym123/funtemps/conv"
	"bufio"
	"os"
	"io"
)

func CelsiusToFahrenheitString(celsius string)(string, error){
	var fahrFloat float64
	var err error
	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFahrenheit(celsiusFloat)
	}
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString, err
}

// Forutsetter at vi kjenner strukturen i filen og denne implementasjon 
// er kun for filer som inneholder linjer hvor det fjerde element
// på linjen er verdien for temperaturaaling i grader celsius
func CelsiusToFahrenheitLine(line string)(string, error){

    dividedString := strings.Split(line, ";")
	var err error
	
	if (len(dividedString) == 4) {
		dividedString[3], err = CelsiusToFahrenheitString(dividedString[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(dividedString, ";"), nil
	
	
	//return "Kjevik;SN39040;18.03.2022 01:50;42.8", err
}
func StudentString(string) (string, error) {
	var err error

	tekst := ("Endringen er gjort av Trym")
	return tekst, err
}

func StudentLine(line string) (string, error) {

	dividedString := strings.Split(line, ";")
	var err error

	if (len(dividedString) == 4) {
		dividedString[3], err = StudentString(dividedString[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(dividedString, ";"), nil
}


func CountLines(filename string) (int, error) {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()
	//LAger en ny scanner som leser filen linje etter linje
	scanner := bufio.NewScanner(file)
	//looper gjennom filen og teller linjene
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning file: %v", err)
	}

	return lineCount, nil
}

func EditLastLine(filename string) error {
	//Åpner filen for reading og writing
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()
	//Finner slutten på filen
	_, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		return fmt.Errorf("error seeking to end of file: %v", err)
	}
	//Leser siste linje
	reader := bufio.NewReader(file)
	lastLine, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return fmt.Errorf("error reading last line: %v", err)
	}
	// Endrer den siste linjen ved å slette de siste 3 tegnene
	newLastLine := lastLine[:len(lastLine)-3]
	// Skriver den modifiserte siste linjen
	newLastLine += "Made by Trym\n"
	_, err = file.WriteAt([]byte(newLastLine), int64(len(newLastLine)-len(lastLine)))
	if err != nil {
		return fmt.Errorf("error writing modified last line: %v", err)
	}

	return nil
}

func CalculateAverageFourthElement(filePath string) (float64, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize variables to keep track of the sum and count of fourth elements
	sum := 0.0
	count := 0

	// Loop through each line in the file
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		if lineNumber == 1 || lineNumber == 27 {
			continue
		}

		// Split the line into fields
		line := scanner.Text()
		fields := strings.Split(line, ";")
		if len(fields) < 4 {
			return 0, fmt.Errorf("line %d has less than 4 fields", lineNumber)
		}

		// Convert the fourth field to a float and add it to the sum
		num, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return 0, fmt.Errorf("error converting field %d in line %d to float: %v", 3, lineNumber, err)
		}
		sum += num
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	// Calculate the average of the fourth elements
	if count == 0 {
		return 0, fmt.Errorf("no valid lines found")
	}
	average := sum / float64(count)

	return average, nil
}

