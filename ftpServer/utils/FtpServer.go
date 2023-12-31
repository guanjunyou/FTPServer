package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

//var ftpClient *ftp.ServerConn
//
//func InitClient() {
//	ftpClient, _ = newClient()
//}
//func newClient() (*ftp.ServerConn, error) {
//	server := config.Config.VideoServer.Ip
//	host := config.Config.VideoServer.Host
//	ftpClient, err := ftp.Dial(fmt.Sprintf("%s:%s", server, host))
//	if err != nil {
//		fmt.Println("无法连接到服务器", err)
//		return nil, err
//	}
//	// 登录到 FTP 服务器
//	err = ftpClient.Login(config.Config.VideoServer.User, config.Config.VideoServer.Password)
//	if err != nil {
//		fmt.Println("登录失败:", err)
//		return nil, err
//	}
//	return ftpClient, nil
//}

/*
*
localPath:本地路劲，dir:Home/douyin/dir
*/
//func Upload(localPath string, dir string) (string, error) {
//	file, err := os.Open(localPath)
//	if err != nil {
//		fmt.Println("无法打开本地文件:", err)
//		return "", err
//	}
//
//	//获取当前工作目录
//	currentDir, err1 := ftpClient.CurrentDir()
//	if err1 != nil {
//		return "", err1
//	}
//	//拼接路径
//	savePath := path.Join(currentDir, dir)
//	err = ftpClient.MakeDir(savePath)
//	if err != nil {
//		fmt.Println("创建文件夹失败")
//		//return "", err
//	}
//	dstName := filepath.Base(file.Name())
//	dstPath := path.Join(savePath, dstName)
//	err = ftpClient.Stor(dstPath, file)
//	if err != nil {
//		fmt.Println("存储文件失败: ", err)
//		return "", err
//	}
//	return dstPath, nil
//}

func SaveFileToFileSystem(c *gin.Context, file *multipart.FileHeader, targetDir string) (string, error) {
	// 获取文件的名称
	fileName := file.Filename

	// 创建目标目录（如果不存在）
	if err1 := os.MkdirAll(targetDir, 0755); err1 != nil {
		log.Println("创建文件夹失败")
		//return err
	}

	dstPath := filepath.Join(targetDir, fileName)
	err2 := c.SaveUploadedFile(file, dstPath)
	if err2 != nil {
		return "", err2
	}
	return fileName, nil
}
