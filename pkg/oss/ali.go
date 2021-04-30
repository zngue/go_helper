package oss

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"github.com/xinliangnote/go-util/md5"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type UploadResonse struct {
	OldUrl string
	NewUrl string
	Error  error
}

type fileUpload struct {
	FileExt string
	bucket  *oss.Bucket
}

type FileUpload interface {
	FilePathAndName(dirs ...string) string
	RangeFileName() string
	RandString(len int) string
	RandomInt(start int, end int) int
	OssClient() (*oss.Bucket, error)
	OssUpload(path string, byt []byte) (string, error)
	UploadFileByUrl(imageArr []string) ([]*UploadResonse, error)
	UploadFileByByte(byt [][]byte) ([]*UploadResonse, error)
	SetExt(byt string)
}

func (f *fileUpload) SetExt(dirs string) {
	f.FileExt = dirs
}
func (f *fileUpload) FilePathAndName(dirs ...string) string {
	path := "hot"
	if len(dirs) >= 1 {
		path = dirs[0]
	}
	if f.FileExt == "" {
		f.FileExt = "png"
	}
	format := time.Now().Format("2006-01-02")
	return fmt.Sprintf("d/file/%s/%s/%s.%s", path, format, f.RangeFileName(), f.FileExt)
}

/*
*@Author Administrator
*@Date 26/4/2021 15:21
*@desc 生成随机文件名
 */
func (f *fileUpload) RangeFileName() string {
	one := f.RandString(100) + cast.ToString(f.RandomInt(10, 999999999)) + cast.ToString(time.Now().Unix())
	s := md5.MD5(one)
	return s
}

/*
*@Author Administrator
*@Date 26/4/2021 15:21
*@params len 字符串长度
*@desc 生成随机字符串
 */
func (f *fileUpload) RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

/*
*@Author Administrator
*@Date 26/4/2021 15:21
*@params start 开始数字
*@params end 结束数字
*@desc 生成随机数
 */
func (f *fileUpload) RandomInt(start int, end int) int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(end - start)
	random = start + random
	return random
}

/*
*@Author Administrator
*@Date 26/4/2021 15:21
*@params path 文件上传路径
*@params byt 上传文件的二进制文件
*@desc oss 文件上传
 */
func (f *fileUpload) OssUpload(path string, byt []byte) (string, error) {
	client, err := f.OssClient()
	if err != nil {
		return "", err
	}
	errs := client.PutObject(path, strings.NewReader(string(byt)))
	if errs != nil {
		return "", errs
	}
	return viper.GetString("oss.domain") + path, nil
}

func (f *fileUpload) GetUrlFileType(url string) error {
	head, err := http.Head(url)
	if err != nil {
		return err
	}
	FileType, ok := head.Header["Content-Type"]
	if !ok {
		return errors.New("url header content type is not exit")
	}
	if len(FileType) < 2 {
		return errors.New("url header content type is not ull")
	}
	f.FileExt = FileType[1]
	return nil
}
func (f *fileUpload) GetUrlFileByte(url string) ([]byte, error) {

	res, uerr := http.Get(url)
	if uerr != nil {
		return nil, uerr
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (f *fileUpload) UrlToByte(url string) (string, error) {
	if len(f.FileExt) == 0 {
		err := f.GetUrlFileType(url)
		if err != nil {
			return "", err
		}
	}
	nano := time.Now().Unix()
	fileByte, ferry := f.GetUrlFileByte(url)
	unixNano := time.Now().Unix()
	if ferry != nil {
		return "", ferry
	}
	uploadUrl, fer := f.OssUpload(f.FilePathAndName(), fileByte)

	upTime := time.Now().Unix()

	fmt.Println(url, "上传时间", upTime-unixNano, "下载时间", unixNano-nano)
	if fer != nil {
		return "", fer
	}
	return uploadUrl, nil
}
func (f *fileUpload) UploadFileByByte(byt [][]byte) ([]*UploadResonse, error) {
	var up []*UploadResonse
	var wg errgroup.Group
	for _, bytes := range byt {
		wg.Go(func() error {
			pathName := f.FilePathAndName()
			toByte, err := f.OssUpload(pathName, bytes)
			response := &UploadResonse{
				NewUrl: toByte,
				Error:  err,
			}
			up = append(up, response)
			return nil
		})
	}
	wg.Wait()
	return up, nil
}

/*
*@Author Administrator
*@Date 26/4/2021 15:21
*@params imageArr 图片地址数组
*@desc 通过url上传
 */
func (f *fileUpload) UploadFileByUrl(imageArr []string) ([]*UploadResonse, error) {
	var up []*UploadResonse
	var wg errgroup.Group
	for _, url := range imageArr {
		if url == "" {
			continue
		}
		response := &UploadResonse{
			OldUrl: url,
		}
		wg.Go(func() error {
			toByte, err := f.UrlToByte(url)
			response.NewUrl = toByte
			response.Error = err
			up = append(up, response)
			return nil
		})
	}
	wg.Wait()
	return up, nil
}

/*
*@Author Administrator
*@Date 26/4/2021 15:21
*@desc oss 上传实例
 */
func (f *fileUpload) OssClient() (*oss.Bucket, error) {
	if f.bucket != nil {
		return f.bucket, nil
	}
	endpoint := viper.GetString("oss.endpoint")
	accessKeySecret := viper.GetString("oss.accessKeySecret")
	accessKeyID := viper.GetString("oss.accessKeyID")
	bucketName := viper.GetString("oss.bucketName")
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return nil, err
	}
	bucket, errs := client.Bucket(bucketName)
	if errs != nil {
		return nil, errs
	}
	f.bucket = bucket
	return bucket, nil
}

/*
*@Author Administrator
*@Date 26/4/2021 16:29
*@desc
 */
func NewUploadFile() FileUpload {
	return new(fileUpload)
}
