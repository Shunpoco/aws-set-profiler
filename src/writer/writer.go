package writer

import "os"

func WriteProfile(path, profile string) error {
	return os.WriteFile(path, []byte(profile), 0644)
}
