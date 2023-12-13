package Exec

import (
  //"fmt"
  "os"
  "os/exec"
  "syscall"
)

func OpenLink(link string) error{
  playerBin, err := exec.LookPath("mpv")
  if err != nil{
    return err
  }

  args := []string{"mpv","--profile=720p", link}

  env := os.Environ()

  execErr := syscall.Exec(playerBin, args, env)

  if execErr != nil{
    return execErr
  }

  return nil
  
}
