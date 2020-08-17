package ebook

import (
	"context"
	"crypto/md5"
	"ebook/ebook/api/ebook"
	"ebook/ebook/conf"
	"ebook/ebook/internel/utils/response"
	"ebook/ebook/internel/utils/rpc"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func Add(ctx *gin.Context)  {
	_, err := os.Stat(conf.PREVIEW_UPLOAD_DIR)
	fmt.Println("===== err: ",err)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(conf.PREVIEW_UPLOAD_DIR, os.ModePerm)
			fmt.Println("===== err: ",err)
		}
	}
	
	_, err = os.Stat(conf.EBOOK_UPLOAD_DIR)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(conf.EBOOK_UPLOAD_DIR, os.ModePerm)
		}
	}
	
	ebookName := ctx.PostForm("ebookName")
	englishName := ctx.PostForm("englishName")
	aliasName := ctx.PostForm("aliasName")
	category := ctx.PostForm("category")
	//categoryUint, _ := strconv.ParseUint(category, 10, 64)
	publishDate := ctx.PostForm("publishDate")
	keyWords := ctx.PostForm("keyWords")
	
	pic, err := ctx.FormFile("pic")
	picType := pic.Header.Get("Content-Type")
	picSize := pic.Size
	picName := pic.Filename
	picSrc, err := pic.Open()
	if err != nil {
		response.Error(ctx, "PREVIEW_OPEN_ERROR")
		return
	}
	defer picSrc.Close()
	picHash := md5.New()
	io.Copy(picHash, picSrc)
	picHashValue := hex.EncodeToString(picHash.Sum([]byte("")))
	
	previewDir := picHashValue + "." + strings.Split(picType, "/")[1]
	previewPath := path.Join(conf.PREVIEW_UPLOAD_DIR, previewDir)
	dist, err := os.Create(previewPath)
	if err != nil {
		response.Error(ctx, "PREVIEW_PATH_CREATE_ERROR")
		return
	}
	picSrc, err = pic.Open()
	defer picSrc.Close()
	if err != nil {
		response.Error(ctx, "PREVIEW_OPEN_ERROR")
		return
	}
	io.Copy(dist, picSrc)

	
	file, err := ctx.FormFile("file")
	fileType := file.Header.Get("Content-Type")
	fileSize := file.Size
	fileName := file.Filename
	
	fileSrc, err := file.Open()
	if err != nil {
		response.Error(ctx, "EBOOK_OPEN_ERROR")
		return
	}
	defer fileSrc.Close()
	hash := md5.New()
	io.Copy(hash, fileSrc)
	fileHashValue := hex.EncodeToString(hash.Sum([]byte("")))
	
	ebookDir := fileHashValue + "." + strings.Split(fileType, "/")[1]
	filePath := path.Join(conf.EBOOK_UPLOAD_DIR, ebookDir)
	dist, err = os.Create(filePath)
	if err != nil {
		response.Error(ctx, "EBOOK_PATH_CREATE_ERROR")
		return
	}
	fileSrc, err = file.Open()
	if err != nil {
		response.Error(ctx, "EBOOK_OPEN_ERROR")
		return
	}
	defer fileSrc.Close()
	io.Copy(dist, fileSrc)
	
	
	addParams := &ebook.AddRequest{
		UserId:               "52fdfc072182654f163f5f0f9a621d72",
		EbookName:            ebookName,
		EnglishName:          englishName,
		AliasName:            aliasName,
		Category:             category,
		PublishDate:          publishDate,
		KeyWords:             keyWords,
		PreviewType:          picType,
		PreviewSize:          picSize,
		PreviewDir:           previewDir,
		PreviewUploadName:    picName,
		PreviewHashValue:     picHashValue,
		EbookType:            fileType,
		EbookSize:            fileSize,
		EbookDir:             ebookDir,
		EbookUploadName:      fileName,
		EbookHashValue:       fileHashValue,
	}
	addRes, err := rpc.EbookRpc().Add(context.Background(), addParams)
	if err != nil {
		log.Println("rpc request error: ", err.Error())
		response.Error(ctx, "RPC_REQUEST_ERROR")
		return
	}
	if addRes.Errno != 0 {
		log.Println("request params error: ", addRes.Errmsg)
		response.Error(ctx, addRes.Errmsg)
		return
	}
	response.OK(ctx, addRes.Data)
	return
}

func Get(ctx *gin.Context)  {
	fmt.Println("SignOut")
}

func GetList(ctx *gin.Context)  {
	fmt.Println("SignOut")
}

func Download(ctx *gin.Context)  {
	var requestBody ViewRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		log.Println("request params error: ", err.Error())
		response.Error(ctx, "PARAMS_ERROR")
		return
	}
	getEbookParams := &ebook.GetRequest{
		Id:                   requestBody.Id,
		UserId:               "52fdfc072182654f163f5f0f9a621d72",
	}
	fmt.Println(getEbookParams)
	ebook, err := rpc.EbookRpc().Get(context.Background(), getEbookParams)
	if err != nil {
		log.Println("request params error: ", err.Error())
		response.Error(ctx, "RPC_REQUEST_ERROR")
		return
	}
	if ebook.Errno != 0 {
		log.Println("get ebook error: ", ebook.Errmsg)
		response.Error(ctx, ebook.Errmsg)
		return
	}
	
	dir := ""
	if requestBody.Type == "image" {
		dir = path.Join("/image", ebook.Data.PreviewDir)
		ctx.Header("Content-Disposition", "attachment; filename=" + ebook.Data.PreviewDir)
	}
	if requestBody.Type == "ebook" {
		dir = path.Join("/file", ebook.Data.EbookDir)
		ctx.Header("Content-Disposition", "attachment; filename=" + ebook.Data.EbookDir)
	}
	fmt.Println(dir)
	fmt.Println(ctx.GetHeader("Content-Disposition"))
	ctx.Redirect(301, dir)
	ctx.Abort()
	return
}

func View(ctx *gin.Context)  {
	var requestBody ViewRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		log.Println("request params error: ", err.Error())
		response.Error(ctx, "PARAMS_ERROR")
		return
	}
	getEbookParams := &ebook.GetRequest{
		Id:                   requestBody.Id,
		UserId:               "52fdfc072182654f163f5f0f9a621d72",
	}
	fmt.Println(getEbookParams)
	ebook, err := rpc.EbookRpc().Get(context.Background(), getEbookParams)
	if err != nil {
		log.Println("request params error: ", err.Error())
		response.Error(ctx, "RPC_REQUEST_ERROR")
		return
	}
	if ebook.Errno != 0 {
		log.Println("get ebook error: ", ebook.Errmsg)
		response.Error(ctx, ebook.Errmsg)
		return
	}
	
	dir := ""
	if requestBody.Type == "image" {
		dir = path.Join("/image", ebook.Data.PreviewDir)
	}
	if requestBody.Type == "ebook" {
		dir = path.Join("/file", ebook.Data.EbookDir)
	}
	fmt.Println(dir)
	ctx.Redirect(301, dir)
	ctx.Abort()
	return
}
