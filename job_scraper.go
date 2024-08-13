
package main

import (
  "os";
  "log";
  "bufio";
  "strings";
  "math/rand";
  "fmt";
  "time";
  "crypto/tls";
  "net/http";
  // "os/exec";
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

func RandomString(string_list []string) string {
  randInt := rand.Intn(len(string_list))
  return string_list[randInt]
}

func GenerateProxies() ([]string) {
  // cmd := exec.Command("python3", "proxies.py")
  
  // err := cmd.Run()

  // if err != nil {
  //  log.Fatal(err)
  // }

  file, err := os.Open("proxies_list.txt")
  
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  
  var proxy_list []string

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    proxy_list = append(proxy_list, scanner.Text())
  }
 
  fmt.Println("Loaded", len(proxy_list), "proxies")

  return proxy_list
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

  for _, url := range JOB_BOARD_URLS {
    fmt.Println(url)
    full_url := "https://www.linkedin.com/jobs/search?keywords=Software%20Engineer&location=United%20States&geoId=103644278&f_JT=F&f_E=2&f_PP=102448103&f_TPR=&position=1&pageNum=0"  
    c.Visit(full_url)
  }

//  for _, url := range job_links {
//    // fmt.Println("Visiting job: ", url)
//    time.Sleep(5 * time.Second)
//    c.Visit(url)
//  }

  parsed := ParseJobLinks(job_links)

  fmt.Println("Total errors:", error_count, "out of", len(job_links), "links")

}


