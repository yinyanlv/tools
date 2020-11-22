package createproj

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

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

func RenderAndCopyFiles(tplVar TemplateVar) error {
	fInfoList, err := ioutil.ReadDir(TempDir)
	if err != nil {
		return err
	}
	len := len(fInfoList)

	if len != 1 {
		return errors.New("临时文件夹中不存在或存在多个项目")
	}
	f := fInfoList[0]
	name := f.Name()

	var destDirName string

	// 当不存在projectName时默认采用模版项目名称
	if tplVar.ProjectName == "" {
		destDirName = name
	} else {
		destDirName = tplVar.ProjectName
	}

	pwd, err := os.Getwd()
	if err != nil {
		return nil
	}
	destPath := path.Join(pwd, destDirName)

	// 如果目标文件夹已经存在，删除已存在的目标文件夹
	if Exist(destPath) {
		err = os.RemoveAll(destPath)
		if err != nil {
			return err
		}
	}

	err = os.Mkdir(destDirName, os.ModePerm)
	if err != nil {
		return err
	}

	srcPath := path.Join(pwd, TempDir, name)

	err = WalkAndHandle(srcPath, destPath, tplVar)
	if err != nil {
		return err
	}

	return nil
}

func WalkAndHandle(srcPath, destPath string, tplVar TemplateVar) error {
	fInfoList, err := ioutil.ReadDir(srcPath)
	if err != nil {
		return nil
	}

	if !Exist(destPath) {
		err = os.Mkdir(destPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for _, item := range fInfoList {
		name := item.Name()
		newSrcPath := path.Join(srcPath, name)
		newDestPath := path.Join(destPath, name)

		if item.IsDir() {
			// 跳过.git目录
			if name == ".git" {
				continue
			}
			err = WalkAndHandle(newSrcPath, newDestPath, tplVar)
			if err != nil {
				return err
			}
		} else {
			err = Copy(newSrcPath, newDestPath, tplVar)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Copy(srcPath, destPath string, tplVar TemplateVar) error {
	fmt.Println("---")
	fmt.Printf("正在处理文件：%s\n", srcPath)

	content, err := ioutil.ReadFile(srcPath)
	str := string(content)
	if tplVar.ProjectName != "" {
		str = strings.Replace(str, "{{__project_name__}}", tplVar.ProjectName, -1)
	}
	if err != nil {
		return err
	}
	fmt.Printf("正在复制文件：%s\n", srcPath)

	err = ioutil.WriteFile(destPath, []byte(str), 0777)
	if err != nil {
		return err
	}
	fmt.Printf("创建文件成功：%s\n", destPath)

	return nil
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)

	return err == nil || os.IsExist(err)
}
