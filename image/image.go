package image

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func Create(name, path string, gigaBytes int) error {
	_, err := os.Stat(path)
	if err != nil {
		err = os.MkdirAll(path, 0600)
		if err != nil {
			return fmt.Errorf("couldn't create image folder at %s", path)
		}
	}

	if Exists(name, path) {
		return fmt.Errorf("image already exists: %s", imagePath(name, path))
	}

	var dd string
	dd, err = exec.LookPath("dd")
	if err != nil {
		return fmt.Errorf("dd not found.")
	}

	args := []string{
		"if=/dev/zero",
		fmt.Sprintf("of=%s", imagePath(name, path)),
		"bs=1M",
		fmt.Sprintf("count=%d", gigaBytes*1024),
	}

	output, err := exec.Command(dd, args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed: %s %v: %s (%s)", dd, strings.Join(args, " "), output, err)
	}

	return nil
}

func List(path string) ([]string, error) {
	_, err := os.Stat(path)
	if err != nil {
		return []string{}, fmt.Errorf("couldn't find image folder at %s", path)
	}

	files, _ := ioutil.ReadDir(path)
	list := []string{}
	for _, f := range files {
		list = append(list, f.Name())
	}

	return list, nil
}

func Destroy(name, path string) error {
	if Exists(name, path) {
		return fmt.Errorf("image doesn't exist: %s", imagePath(name, path))
	}

	return os.Remove(imagePath(path, name))
}

func Exists(name, path string) bool {
	if _, err := os.Stat(imagePath(name, path)); err != nil {
		return false
	}

	return true
}

func imagePath(name, path string) string {
	return fmt.Sprintf("%s/%s.img", path, name)
}
