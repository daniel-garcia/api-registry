package identity

import (
	"fmt"
	"io/fs"
	"testing"
)

func TestEmbed(t *testing.T) {

	fs.WalkDir(content, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // Handle errors that occurred during the walk.
		}
		fmt.Println(path) // Print the path of the current file or directory.
		return nil
	})

}
