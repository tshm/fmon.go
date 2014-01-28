package main

import (
  "github.com/howeyc/fsnotify"
  "log"
  "flag"
  "os"
  "os/exec"
  "os/signal"
  "syscall"
  "fmt"
  "time"
)

func run(command []string) {
  cmd := exec.Command(command[0], command[1:]...)
  buf, err := cmd.CombinedOutput()
  fmt.Printf("%s\n", buf)
  if err != nil {
    log.Println(err)
  }
}

func main() {
  pathPtr := flag.String("path", ".", "path of the target dir")
  deadtimePtr := flag.Uint("deadtime", 1000, "deadtime of the trigger")
  flag.Parse()
  command := flag.Args()

  if len(command) == 0 {
    log.Fatal("usage: fmon.exe [-path targetDir] [-deadtime 1000] {command}")
  }

  done := make(chan os.Signal, 1)
  signal.Notify(done,
    os.Kill,
    os.Interrupt,
    syscall.SIGHUP,
    syscall.SIGINT,
    syscall.SIGTERM,
    syscall.SIGQUIT)

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
        time.Sleep(time.Duration(*deadtimePtr) * time.Millisecond)
        // flush channel
        chanlen := len(watcher.Event)
        for i := 0; i <= chanlen; i++ {
          <- watcher.Event
        }
      case err := <-watcher.Error:
        log.Println("error:", err)
      }
    }
  }()

  err = watcher.Watch(*pathPtr)
  if err != nil {
    log.Fatal(err)
  }

  <- done
  log.Println("Exiting.")
}
