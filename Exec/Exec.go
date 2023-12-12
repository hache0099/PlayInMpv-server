package Exec

import (
  //"fmt"
  "os"
  "os/exec"
  "syscall"
)

func OpenLink(link string) err{
  playerBin, err = exec.LookPath("mpv")
  if err != nil{
    return err
  }

  args := []string{"mpv","--profile=yt-720", link}

  env := os.Environ()

  execErr = syscall.Exec(playerBin, args, env)

  if execErr != nil{
    return execErr
  }

  return nil
  
}
