package mediaFileHandle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-commweb/common"
	"go-commweb/config"
	"go-commweb/global"
	log "go-commweb/log"
	"os"
	"path"
	"path/filepath"
	"strings"
)


// 上传文件接口
func MediaFileAdd(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		log.LOGGER.Error("err %v", err)
		return
	}

	for filePath, files := range form.File {
		log.LOGGER.Info("filePath [%s] upload file count [%d]", filePath, len(files))
		if len(form.File) <= 0 {
			log.LOGGER.Error("the upload file is empty")
			continue
		}
		for _, file := range files {
			log.LOGGER.Info("filepath [%s] filename [%s]", filePath, file.Filename)
			filename := filepath.Base(file.Filename)
			absPath := path.Join( config.ConfigParam.UpLoadConfig.Path,  filePath , filename)
			absPath = strings.Replace(absPath, "\\", "/", -1)
			if common.PathExists(absPath) {
				log.LOGGER.Info("file already exist %s",absPath)
				continue
			}
			parentPath := fmt.Sprintf("%s%s", config.ConfigParam.UpLoadConfig.Path,  filePath)
			if !common.PathExists(parentPath) {
				log.LOGGER.Info("file parent path is not exist, make it")
				os.MkdirAll(parentPath, 777) // 创建文件夹 附带权限
				os.Chmod(parentPath, 0777)
			}
			log.LOGGER.Info("filename [%s] absPath [%s]", filename, absPath)
			if err := c.SaveUploadedFile(file, absPath); err != nil {
				common.ResponseJson(c, global.RequestSuccess, global.SaveFileError, "保存文件失败", "")
				continue
			}
		}
	}
	common.ResponseJson(c, global.RequestSuccess, global.RequestSuccess, "ok", "")
}


func MediaFileDelete(c *gin.Context) {
	// delete
	directory := c.PostForm("directory")
	filename := c.PostForm("filename")

	log.LOGGER.Info("delete directory [%s] filename [%s]", directory, filename)

	if directory == "" || filename == "" {
		reason := fmt.Sprintf("miss argument directory %s filename %s", directory, filename)
		log.LOGGER.Error("%s", reason)
		common.ResponseJson(c, global.RequestSuccess, global.ParamMissingError, "参数不全", "")
		return
	}
	absPath := path.Join( config.ConfigParam.UpLoadConfig.Path,  directory , filename)
	absPath = strings.Replace(absPath, "\\", "/", -1)
	if common.PathExists(absPath) {
		log.LOGGER.Info("remove file %s",absPath)
		os.Remove(absPath)
	}else{
		common.ResponseJson(c, global.RequestSuccess, global.FileNotExists, "文件不存在", "")
		return
	}
	common.ResponseJson(c, global.RequestSuccess, global.RequestSuccess, "ok", "")
}