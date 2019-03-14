package utils

import (
  "github.com/Unknwon/goconfig"
  "log"
  "os"
)

var Cfg *goconfig.ConfigFile
func init()  {
  //var err error
  cfg, err := goconfig.LoadConfigFile("config.ini")
  CheckErr(err, "no config.ini file...")
  Cfg = cfg
}

// panic error, quit the program
func CheckErr(err error, msg ...string) {
  if err != nil {
    log.Println(msg)
    panic(err.Error())
  }
}

/**
only debug info, output debug info & save log file
 */
func DebugLog(content string) {
  //DebugFlag, err := Cfg.GetValue("comm", "DebugFlag")
  //CheckErr(err)
  DebugFlag := Cfg.MustBool("comm", "DebugFlag")
  if DebugFlag {
    log.Println(content)
  }

  LogFlag := Cfg.MustBool("comm", "LogFlag")
  if LogFlag {
    logFileHandle, err := os.OpenFile(LogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
    CheckErr(err)
    _, _ = logFileHandle.WriteString(content + "\n\n")
  }
}


/**
check os path exist or not
 */
func PathExists(path string) bool {
  _, err := os.Stat(path)
  if err == nil {
    return true
  }
  return false
}


/**
create folder
 */
func CreateFolder()  {

}






