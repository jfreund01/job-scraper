
package main

import (
  "time";
  "encoding/csv";
  "os";
  "log";
  "bufio";
  "strings";
  "math/rand";
  "fmt";
  "crypto/tls";
  "net/http";
  "github.com/gocolly/colly";
)

const LINKEDIN_JOB_PREFIX string = "https://www.linkedin.com/jobs/view/"

var job_search_terms = []string{
  "Software Engineer",
  "Simulation Engineer",
  "Game Developer",
  "Python",
  "C++",
  "C",
  "C developer",
  "Software Developer",
}

var JOB_BOARD_URLS = []string{ 
  "https://www.linkedin.com/jobs/search?keywords=%s&location=United States&geoId=103644278&f_JT=F&f_E=2&f_PP=102448103&f_TPR=&position=1&pageNum=0",
} 

type Job struct {
  JobTitle  string  `json:"job_title"`
  JobID     string  `json:"job_id"`
  JobLink   string  `json:"job_link"`
}

func WriteToCSV(jobs []Job) {
  file, err := os.Create("jobs.csv")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  writer := csv.NewWriter(file)
  defer writer.Flush()

  for _, job := range jobs {
    err := writer.Write([]string{job.JobTitle, job.JobID, job.JobLink})
    if err != nil {
      log.Fatal(err)
    }
  }
  fmt.Println("Wrote", len(jobs), "jobs to jobs.csv")
}

func RandomString(string_list []string) string {
  randInt := rand.Intn(len(string_list))
  return string_list[randInt]
}

func GetUserAgents() ([]string) {
  file, err := os.Open("user_agent_list.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  var userAgentList []string

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    userAgentList = append(userAgentList, scanner.Text())
  }
  
  fmt.Println("Loaded", len(userAgentList), "user agents")

  return userAgentList
}

func ParseJobLinks(link string) Job {
  split_strings := strings.Split(link, "/")[5]
  split_strings = strings.Split(split_strings, "?")[0]
  split_jobs_title := strings.Split(split_strings, "-")
  job := Job{} 
  if len (split_strings) > 0 {
    job.JobID = split_jobs_title[len(split_jobs_title)-1]
    split_jobs_title = split_jobs_title[:len(split_jobs_title)-1]
    job.JobTitle = strings.Join(split_jobs_title, " ")
  }
  job.JobLink = link
  //fmt.Println("Job title: ", job_title)
  return job
}

func main() {
  c := colly.NewCollector()
  c.WithTransport(&http.Transport{ 
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  })
  
  user_agent_list := GetUserAgents()
  // proxy_list := GenerateProxies()

  job_links := make([]string, 0, 50)
  error_count := 0
  c.OnResponse(func(r *colly.Response) {
    fmt.Println("Visited", r.Request.URL)
  })

  c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    link := e.Attr("href")
    if strings.HasPrefix(link, LINKEDIN_JOB_PREFIX) {
      // fmt.Println("Job link: ", link)
      job_links = append(job_links, link)
      // c.Visit(link)
    }
  })
  
  c.OnError(func(r *colly.Response, err error) {
    // body := string(r.Body)
    fmt.Println("Error on request: ")
    fmt.Println(err)
    error_count += 1
    time.Sleep(30 * time.Second)
  })

  c.OnRequest(func(r *colly.Request) {
    user := RandomString(user_agent_list)
    // proxy := RandomString(proxy_list)
    // fmt.Println("Proxy: ", proxy)
    fmt.Println("User agent: ", user)
    r.Headers.Set("User-Agent", user)
    // if err != nil {
    //  fmt.Println("Error setting proxy")
    //  fmt.Println(err)
    // }

    fmt.Println("Visiting : ", r.URL.String())
  })

  for _, keyword := range job_search_terms {
    fmt.Println("Searching for: ", keyword)
    url := fmt.Sprintf(JOB_BOARD_URLS[0], keyword)
    url = strings.Replace(url, " ", "%20", -1)
    c.Visit(url)
  }

//  for _, url := range job_links {
//    // fmt.Println("Visiting job: ", url)
//    time.Sleep(5 * time.Second)
//    c.Visit(url)
//  }
  
  jobs := make([]Job, 0, 50)
  
  for _, url := range job_links {
    jobs = append(jobs, ParseJobLinks(url))
  }
  
  fmt.Println("Total jobs found:", len(jobs))

  fmt.Println("Total errors:", error_count, "out of", len(job_links), "links")

  WriteToCSV(jobs)
}


