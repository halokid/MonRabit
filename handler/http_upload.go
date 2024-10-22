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
var HttpDtUploadPath string

func init()  {
  var err error
  HttpUploadPath, err = utils.Cfg.GetValue("comm", "HttpUploadPath")
  utils.CheckErr(err)
  HttpDtUploadPath, err = utils.Cfg.GetValue("comm", "HttpDtUploadPath")
  utils.CheckErr(err)
}

func HttpUploadHandle(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "acting upload ....")
  formFile, header, err := r.FormFile("uploadfile")
  if err != nil {
    fmt.Fprintf(w, "from file name ERROR..")
  }
  defer formFile.Close()

  //create save file
  //httpUploadPath, err := utils.Cfg.GetValue("comm", "HttpUploadPath")
  destFile, err := os.Create(HttpUploadPath + "/" + header.Filename)
  if err != nil {
    fmt.Fprintf(w, "destFile ERROR..")
  }
  defer destFile.Close()

  // save file
  _, err = io.Copy(destFile, formFile)
  if err != nil {
    fmt.Fprintf(w, "copy file ERROR..")
  }

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
  utils.CheckErr(err, "HttpDtUploadPath error...")
  destFile, err := os.Create(HttpDtUploadPath + "/" + folderPath + "/" + header.Filename)
  utils.CheckErr(err)
  defer destFile.Close()

  // save file
  _, err = io.Copy(destFile, formFile)
  utils.CheckErr(err)

  //fmt.Fprintf(w, "upload success")
  fmt.Fprintf(w, string(HttpDtUploadPath + "/" + folderPath + "/" + header.Filename))

}



func CreateDataFolder(folderPath string) error  {
  dateFolder := strings.Split(folderPath, "/")[0]
  timeFolder := strings.Split(folderPath, "/")[1]

  //httpUploadPath, _ := utils.Cfg.GetValue("comm", "HttpUploadPath")
  if !utils.PathExists(HttpDtUploadPath + "/" + dateFolder) {
    err := os.Mkdir(HttpDtUploadPath + "/" + dateFolder, os.ModePerm)
    utils.CheckErr(err)
  }

  err := os.Mkdir(HttpDtUploadPath + "/" + dateFolder + "/" + timeFolder, os.ModePerm)
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
  if utils.PathExists( HttpDtUploadPath + "/" + dateFolder + "/" + timeFolder) {
    fmt.Println("folder exists")
    return true, folderPath
  } else {
    fmt.Println("folder not exists")
    return false, folderPath
  }
}



