package config

import (
	"os"
	"path/filepath"
)

// 项目主目录
var rootDir string

func GetRootDir() string {
	return rootDir
}

func init() {
	//初始化配置
	inferRootDir()
}

// 推断 Root目录
func inferRootDir() {
	pwd, err := os.Getwd() //获取执行目录
	if err != nil {
		panic(err)
	}

	var infer func(string) string
	infer = func(dir string) string {
		if exists(dir + "./main.go") {
			return dir
		}

		//查询dir上级目录
		parent := filepath.Dir(dir)
		return infer(parent)
	}

	rootDir = infer(pwd)

}

func exists(dir string) bool {
	//检测目录是否存在
	_, err := os.Stat(dir)
	return err == nil || os.IsExist(err)
}
