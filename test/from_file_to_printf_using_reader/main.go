package main

import (
  "io"
  "log"
  "os"
  "sync"
)

type ThreadSafeFile struct {
  mu sync.Mutex
  handle *os.File
}

func (f *ThreadSafeFile) write(out *os.File) {
  // lock to prevent the case where the file pointer is rewound while another thread is reading
  f.mu.Lock()
  // rewind
  f.handle.Seek(0, 0)
  if _, err := io.Copy(out, f.handle); err != nil {
    log.Fatal(err)
  }
  f.mu.Unlock()

  // flush output buffer
  out.Sync()
}

func (f *ThreadSafeFile) close() {
  f.handle.Close()
}

var gIndexHtmlFile = &ThreadSafeFile {}

func init() {
  indexFileName := "index.html"
  var err error

  gIndexHtmlFile.handle, err = os.Open(indexFileName)
  if err != nil { log.Fatal(err) }
}

func printContent(wg *sync.WaitGroup) {
  defer wg.Done()
  gIndexHtmlFile.write(os.Stdout)
}

func main() {
  defer gIndexHtmlFile.close()

  var wg sync.WaitGroup
  // consider as HTTP handler calls this for each request
  for i := 0; i < 10; i++ {
    wg.Add(1)
    go printContent(&wg)
  }
  wg.Wait()
}