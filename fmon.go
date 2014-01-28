package main

import (
  "github.com/howeyc/fsnotify"
  "log"
  "flag"
  "os/exec"
  "fmt"
  "time"
)

func run(command []string) {
  cmd := exec.Command(command[0], command[1:]...)
  buf, err := cmd.CombinedOutput()
  fmt.Printf("%s\n", buf)
  if err != nil {
    log.Fatal(err)
  }
  time.Sleep(time.Second)
}

func main() {
  pathPtr := flag.String("path", ".", "path of the target dir")
  flag.Parse()
  command := flag.Args()

  if len(command) == 0 {
    log.Fatal("usage: fmon.exe [-path targetDir] {command}")
  }

  done := make(chan bool)

  watcher, err := fsnotify.NewWatcher()
  if err != nil {
    log.Fatal(err)
  }
  defer watcher.Close()
  log.Println("start watching dir: ", *pathPtr)

  // Process events
  go func() {
    for {
      select {
      case ev := <-watcher.Event:
        log.Println("event:", ev)
        run(command)
      case err := <-watcher.Error:
        log.Println("error:", err)
      }
    }
  }()

  err = watcher.Watch(*pathPtr)
  if err != nil {
    log.Fatal(err)
  }

	<-done
}
