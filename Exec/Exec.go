package Exec

import (
  "os/exec"
)

func OpenLink(link string){
  playerBin, err := exec.LookPath("mpv")
  if err != nil{
    return
  }


  cmd := exec.Command(playerBin, "--profile=720p",link)
  if err := cmd.Run(); err != nil{
    return
  }
}
