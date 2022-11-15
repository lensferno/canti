package commons

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 这段代码大部分直接搬运自nps，经过一定的修改

func GetDefaultInstallPath() (path string) {
	if IsWindows() {
		return `C:\Programs Files\canti`
	} else {
		return `/etc`
	}
}

func CopyProgram(srcPath, binFile string) string {
	binPath, _ := filepath.Abs(os.Args[0])
	if IsWindows() {
		copyForWindows(srcPath, binFile)
	} else {
		binPath = copyForUnix(srcPath, binFile)
	}

	chMod(binPath, 0755)
	return binPath
}

func copyForUnix(srcPath, binFile string) (targetPath string) {
	_, err := copyFile(filepath.Join(srcPath, binFile), "/usr/bin/"+binFile)
	if err != nil {
		_, err := copyFile(filepath.Join(srcPath, binFile), "/usr/local/bin/"+binFile)
		if err != nil {
			log.Fatal("复制文件时发生错误：", err)
		} else {
			return "/usr/local/bin/" + binFile
		}
	} else {
		return "/usr/bin/" + binFile
	}

	return ""
}

func copyForWindows(srcPath, binFile string) {
	copyFile(filepath.Join(srcPath, binFile+".exe"), filepath.Join(GetAppPath(), binFile+".exe"))
}

func MkidrDirAll(path string, v ...string) {
	for _, item := range v {
		if err := os.MkdirAll(filepath.Join(path, item), 0755); err != nil {
			log.Fatalf("Failed to create directory %s error:%s", path, err.Error())
		}
	}
}

func CopyDir(srcPath string, destPath string) error {
	//检测目录正确性
	if srcInfo, err := os.Stat(srcPath); err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		if !srcInfo.IsDir() {
			e := errors.New("SrcPath is not the right directory!")
			return e
		}
	}
	if destInfo, err := os.Stat(destPath); err != nil {
		return err
	} else {
		if !destInfo.IsDir() {
			e := errors.New("DestInfo is not the right directory!")
			return e
		}
	}
	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			destNewPath := strings.Replace(path, srcPath, destPath, -1)
			log.Println("copy file ::" + path + " to " + destNewPath)
			copyFile(path, destNewPath)
			if !IsWindows() {
				chMod(destNewPath, 0766)
			}
		}
		return nil
	})
	return err
}

//生成目录并拷贝文件
func copyFile(src, dest string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()
	//分割path目录
	destSplitPathDirs := strings.Split(dest, string(filepath.Separator))

	//检测时候存在目录
	destSplitPath := ""
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs)-1 {
			destSplitPath = destSplitPath + dir + string(filepath.Separator)
			b, _ := pathExists(destSplitPath)
			if b == false {
				log.Println("mkdir:" + destSplitPath)
				//创建目录
				err := os.Mkdir(destSplitPath, os.ModePerm)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}
	dstFile, err := os.Create(dest)
	if err != nil {
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

//检测文件夹路径时候存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func chMod(name string, mode os.FileMode) {
	if !IsWindows() {
		os.Chmod(name, mode)
	}
}

func GetAppPath() string {
	if path, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
		return path
	}
	return os.Args[0]
}
