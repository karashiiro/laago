package laago

import (
	"os"
)

// IsLargeAddressAware returns whether or not the executable at the specified
// location is large address aware.
func IsLargeAddressAware(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Read the low byte of the file characteristics
	buf := make([]byte, 1)
	_, err = file.ReadAt(buf, laaOffset)
	if err != nil {
		return false, err
	}

	// Check the LAA flag
	return (buf[0] & peImageFileLargeAddressAware) != 0, nil
}

// SetLargeAddressAware sets the large address aware bit on the specified
// executable. This may cause some executable images to function incorrectly.
func SetLargeAddressAware(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the low byte of the file characteristics
	buf := make([]byte, 1)
	_, err = file.ReadAt(buf, laaOffset)
	if err != nil {
		return err
	}

	// Set LAA
	buf[0] |= peImageFileLargeAddressAware

	// Write back the new file characteristics
	_, err = file.WriteAt(buf, laaOffset)
	if err != nil {
		return err
	}

	return nil
}

// UnsetLargeAddressAware unsets the large address aware bit on the specified
// executable. This may cause some executable images to function incorrectly.
func UnsetLargeAddressAware(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the low byte of the file characteristics
	buf := make([]byte, 1)
	_, err = file.ReadAt(buf, laaOffset)
	if err != nil {
		return err
	}

	// Unset LAA
	buf[0] &= ^peImageFileLargeAddressAware

	// Write back the new file characteristics
	_, err = file.WriteAt(buf, laaOffset)
	if err != nil {
		return err
	}

	return nil
}
