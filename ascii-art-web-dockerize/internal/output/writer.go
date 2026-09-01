package output

import "os"

// WriteOutput writes the given content to the specified file.
// It creates the file if it does not exist, or truncates it if it does.
func WriteOutput(filename, content string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}
