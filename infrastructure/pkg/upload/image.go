package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/thoas/go-funk"
	"github.com/viafcccy/garden-be/global"
	"github.com/viafcccy/garden-be/infrastructure/pkg/file"
	"github.com/viafcccy/garden-be/infrastructure/pkg/util"
)

func GetImageFullUrl(name string) string {
	return global.Gconfig.Image.ImagePrefixUrl + "/" + GetImagePath() + name
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

func GetImagePath() string {
	return global.Gconfig.Image.ImageSavePath
}

func GetImageFullPath() string {
	return global.Gconfig.Image.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	return funk.ContainsString(global.Gconfig.Image.ImageAllowExts, ext)
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		global.GLog.Warnln(err)
		return false
	}

	return size <= global.Gconfig.Image.ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
