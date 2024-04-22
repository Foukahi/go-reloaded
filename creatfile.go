package stock

import "os"

func CreatFile(txt string, file string) {
	os.WriteFile(file, []byte(txt), 0644)
}
