package createproj

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Create(url string, tplVar TemplateVar) error {
	err := removeTempDir()
	if err != nil {
		return err
	}

	err = makeTempDir()
	if err != nil {
		return err
	}

	err = Fetch(url)
	if err != nil {
		return err
	}

	err = RenderAndCopyFiles(tplVar)
	if err != nil {
		return err
	}

	err = removeTempDir()
	if err != nil {
		return err
	}

	return nil
}

func Fetch(url string) error {
	var stdout, stderr bytes.Buffer
	fmt.Printf("开始从 %s 获取代码...\n", url)
	cmd := exec.Command("git", "clone", url)
	cmd.Dir = fmt.Sprintf("%s", TempDir)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	out := stdout.String() + stderr.String()
	fmt.Printf(out)
	fmt.Printf("从 %s 获取代码成功。\n", url)
	return nil
}
