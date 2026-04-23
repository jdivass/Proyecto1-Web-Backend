package utils
import (
	"os"
)
func EnsureUploadsDir() error {
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		return os.Mkdir("./uploads", os.ModePerm)
	}
	return nil
}