package service

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"mime/multipart"
	"monitor/extend/conf"
	"monitor/extend/utils"
	"os"
	"path"
	"strings"
)

type UploadService struct{}

// GetImgPath 获取图片相对目录
func (us *UploadService) GetImgPath() string {
	return conf.ServerConf.StaticRootPath
}

// GetImgFullPath 获取图片完整目录
func (us *UploadService) GetImgFullPath() string {
	return conf.ServerConf.StaticRootPath + conf.ServerConf.UploadImagePath
}

func (us *UploadService) GetImgName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.MakeSha1(fileName)
	return fileName + ext
}

func (us *UploadService) GetImgFullURL(name string) string {
	return conf.ServerConf.PrefixURL + conf.ServerConf.UploadImagePath + name
}

func (us *UploadService) CheckImgExt(fileName string) bool {
	ext := path.Ext(fileName)
	for _, allowExt := range conf.ServerConf.ImageFormats {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func (us *UploadService) CheckImgSize(f multipart.File) bool {
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error().Msg(err.Error())
		return false
	}
	//单位转换
	const convertRatio float64 = 1024 * 1024
	fileSize := float64(len(content)) / convertRatio
	return fileSize <= conf.ServerConf.UploadLimit

}

func (us *UploadService) CheckImgPath(path string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd: %v", err)
	}
	isExist, err := utils.IsExist(dir + "/" + path)
	if err != nil {
		return fmt.Errorf("dir is exists: %v", err)
	}
	if isExist == false {
		err := os.MkdirAll(dir+"/"+path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("mkdir err : %v", err)
		}
	}
	isPerm := utils.IsPerm(path)
	if isPerm {
		return fmt.Errorf("permission denied src:%s", path)
	}
	return nil
}
