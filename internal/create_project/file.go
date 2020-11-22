package createproj

import "os"

func makeTempDir() error {
	err := os.Mkdir(TempDir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func removeTempDir() error {
	err := os.RemoveAll(TempDir)
	if err != nil {
		return err
	}
	return nil
}
