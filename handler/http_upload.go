package handler

import (
  "net/http"
  "github.com/r00tjimmy/MonRabit/utils"
  "os"
  "io"
  "fmt"
  "strings"
  "time"
)

var HttpUploadPath string

func init()  {
  var err error
  HttpUploadPath, err = utils.Cfg.GetValue("comm", "HttpUploadPath")
  utils.CheckErr(err)
}

func HttpUploadHandle(w http.ResponseWriter, r *http.Request) {
  formFile, header, err := r.FormFile("uploadfile")
  utils.CheckErr(err)
  defer formFile.Close()

  //create save file
  //httpUploadPath, err := utils.Cfg.GetValue("comm", "HttpUploadPath")
  destFile, err := os.Create(HttpUploadPath + "/" + header.Filename)
  utils.CheckErr(err)
  defer destFile.Close()

  // save file
  _, err = io.Copy(destFile, formFile)
  utils.CheckErr(err)

  fmt.Fprintf(w, "upload success")
}



/**
upload file to data-time folder
 */
func HttpUploadDateTimeHandle(w http.ResponseWriter, r *http.Request) {
  // check path
  exists, folderPath := DateTimeFolderExists()
  if !exists {
    // folder not exists, create
    err := CreateDataFolder(folderPath)
    utils.CheckErr(err)
  }

  // upload file
  formFile, header, err := r.FormFile("uploadfile")
  utils.CheckErr(err)
  defer formFile.Close()

  //create save file
  //httpUploadPath, err := utils.Cfg.GetValue("comm", "HttpUploadPath")
  utils.CheckErr(err, "httpUploadPath error...")
  destFile, err := os.Create(HttpUploadPath + "/" + folderPath + "/" + header.Filename)
  utils.CheckErr(err)
  defer destFile.Close()

  // save file
  _, err = io.Copy(destFile, formFile)
  utils.CheckErr(err)

  fmt.Fprintf(w, "upload success")

}



func CreateDataFolder(folderPath string) error  {
  dateFolder := strings.Split(folderPath, "/")[0]
  timeFolder := strings.Split(folderPath, "/")[1]

  //httpUploadPath, _ := utils.Cfg.GetValue("comm", "HttpUploadPath")
  if !utils.PathExists(HttpUploadPath + "/" + dateFolder) {
    err := os.Mkdir(HttpUploadPath + "/" + dateFolder, os.ModePerm)
    utils.CheckErr(err)
  }

  err := os.Mkdir(HttpUploadPath + "/" + dateFolder + "/" + timeFolder, os.ModePerm)
  utils.CheckErr(err)
  return  err
}


/**
check date-time folder name
folder display like:
`-- 2018-09-11
    |-- 11-23
    |   |-- xxxxxxx.xml
    |   `-- yyyyyyyyyyyy.xml
    `-- 11-25
        |-- aaaaaaaaaaaa.xml
        `-- bbbbbbbbbbb.xml
3 directories, 8 files
 */
func DateTimeFolderExists() (bool, string) {
  dateTime := time.Now().String()
  //fmt.Println(date_time)

  folderSli := strings.Split(dateTime, " ")
  //fmt.Println(folder_sli)

  dateFolder := folderSli[0]
  //fmt.Println(date_folder)

  timeFolder := strings.Split(folderSli[1], ":")[0] + "-" + strings.Split(folderSli[1], ":")[1]
  //fmt.Println(time_folder)

  folderPath := dateFolder + "/" + timeFolder
  //httpUploadPath, _ := utils.Cfg.GetValue("comm", "HttpUploadPath")
  if utils.PathExists( HttpUploadPath + "/" + dateFolder + "/" + timeFolder) {
    fmt.Println("folder exists")
    return true, folderPath
  } else {
    fmt.Println("folder not exists")
    return false, folderPath
  }
}



