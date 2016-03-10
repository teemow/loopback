package loop

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Attach(name, path string) (string, error) {
	imagePath := fmt.Sprintf("%s/%s.img", path, name)

	_, err := os.Stat(imagePath)
	if err != nil {
		return "", fmt.Errorf("image doesn't exists: %s", imagePath)
	}

	var output string
	output, err = losetup([]string{"-f", "--show", imagePath})

	return trim(output), err
}

func Find(name, path string) (string, error) {
	imagePath := fmt.Sprintf("%s/%s.img", path, name)

	device, err := losetup([]string{"-O", "NAME", "-n", "-j", imagePath})
	if err != nil {
		return "", fmt.Errorf("Can't find the device: %s", err)
	}

	return trim(device), nil
}

func Detach(device string) error {
	_, err := losetup([]string{"-d", device})
	if err != nil {
		return fmt.Errorf("Can't detach device: %s", err)
	}

	return nil
}

func Unmount(device string) error {
	umount, err := exec.LookPath("umount")
	if err != nil {
		return fmt.Errorf("umount not found.")
	}

	var output []byte
	output, err = exec.Command(umount, device).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed: %s %v: %s (%s)", umount, device, output, err)
	}

	return nil
}

func List() (string, error) {
	return list(true)
}

func ListWithoutHeadings() (string, error) {
	return list(false)
}

func Format(device, fsType string) error {
	mkfs, err := exec.LookPath(fmt.Sprintf("mkfs.%s", fsType))
	if err != nil {
		return fmt.Errorf("mkfs for %s not found.", fsType)
	}

	var output []byte
	output, err = exec.Command(mkfs, device).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed: %s %v: %s (%s)", mkfs, device, output, err)
	}

	return nil
}

func Mount(device, mountPath string) error {
	mount, err := exec.LookPath("mount")
	if err != nil {
		return fmt.Errorf("mount not found.")
	}

	args := []string{device, mountPath}

	output, err := exec.Command(mount, args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed: %s %v: %s (%s)", mount, strings.Join(args, " "), output, err)
	}

	return nil
}

func list(showHeadings bool) (string, error) {
	args := []string{"--list"}

	if showHeadings == false {
		args = append(args, "--noheadings")
	}

	return losetup(args)
}

func trim(in string) string {
	return strings.Trim(in, "\n")
}

func losetup(args []string) (string, error) {
	losetup, err := exec.LookPath("losetup")
	if err != nil {
		return "", fmt.Errorf("losetup not found.")
	}

	output, err := exec.Command(losetup, args...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed: %s %v: %s (%s)", losetup, strings.Join(args, " "), output, err)
	}

	return string(output[:]), nil
}
