package io

import "os"

//CreateDir 创建文件夹
func CreateDir(dirPath string) bool {
	if !Exist(dirPath) {
		if os.Mkdir(dirPath, 0777) != nil {
			return false
		}
	}
	return true
}

//Exist 文件夹是否存在
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
