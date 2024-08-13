package main

import (
  "flag";
  "fmt";
  "os";
  "github.com/jfreund01/job-scraper/src/job_scraper";
  "github.com/jfreund01/job-scraper/src/setup";
)

var (
  actionFlag string
)

func main() {
  flag.StringVar(&actionFlag, "action", "", "Action to perform")
  flag.Parse()

  if actionFlag == "scrape" {
    // check if setup.cfg exists
    _, err := os.Stat("setup.cfg")
    if os.IsNotExist(err) {
      fmt.Println("setup.cfg does not exist. Run 'go run main.go -action=setup' to create it")
      return
    }
    job_scraper.ScrapeJobs()
  } else if actionFlag == "setup" {
    setup.Setup()
  } else {
    fmt.Println("Invalid action")
  }
}
