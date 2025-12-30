package utils

import (
	"go-record/core"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else {
		if os.IsNotExist(err) {
		} else {
			core.LOGGER.Error(err.Error())
		}
	}
	return false
}


//get home path for application.
func GetHomePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	if EnvMacDevelopment() {
		exPath = GetDevHomePath() + "/tmp"
	}

	//if EnvWinDevelopment() {
	//	exPath = GetDevHomePath() + "/tmp"
	//}

	return UniformPath(exPath)
}


//whether mac develop environment
func EnvMacDevelopment() bool {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	return strings.HasPrefix(ex, "/private/var/folders")

}

//get development home path.
func GetDevHomePath() string {

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get dev home path.")
	}

	//$DevHomePath/code/tool/util/util_file.go
	dir := GetDirOfPath(file)
	dir = GetDirOfPath(dir)
	dir = GetDirOfPath(dir)
	dir = GetDirOfPath(dir)

	return dir
}

//eg /var/www/xx.log -> /var/www
func GetDirOfPath(fullPath string) string {

	index1 := strings.LastIndex(fullPath, "/")
	//maybe windows environment
	index2 := strings.LastIndex(fullPath, "\\")
	index := index1
	if index2 > index1 {
		index = index2
	}

	return fullPath[:index]
}

//1. replace \\ to /
//2. clean path.
//3. trim suffix /
func UniformPath(p string) string {
	p = strings.Replace(p, "\\", "/", -1)
	p = path.Clean(p)
	p = strings.TrimSuffix(p, "/")
	return p
}