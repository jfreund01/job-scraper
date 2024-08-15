package main

import (
  //"flag";
  "strings";
  //"os";
  "github.com/jfreund01/job-scraper/src/job_scraper";
  //"github.com/jfreund01/job-scraper/src/setup";
  "github.com/rivo/tview";
  "github.com/gdamore/tcell/v2";
)

var (
  actionFlag string
)

func removeEmptyAndTrimEdges(input []string) []string {
    var result []string
    for _, str := range input {
        trimmed := strings.TrimSpace(str)
        if trimmed != "" {
            result = append(result, trimmed)
        }
    }
    return result
}

func DrawCLI() {
  app := tview.NewApplication()

  textArea := tview.NewTextArea().
    SetLabel("Keywords")

  outputArea := tview.NewTextView().
    SetLabel("Output")

  form := tview.NewForm().
    AddButton("Quit", func() { app.Stop() }).
    AddFormItem(textArea).
    SetFieldBackgroundColor(tcell.NewRGBColor(56,70,90)).
    AddButton("Search", func(){
      // get Keywords
      keywords := textArea.GetText()
      keywords_slice := strings.Split(keywords, "\n")
      cleaned_keywords := removeEmptyAndTrimEdges(keywords_slice) 
      // print keywords
      output := "Scraping for: " + strings.Join(cleaned_keywords, ", ")
      outputArea.SetText(output)
      job_scraper.ScrapeJobs(cleaned_keywords, outputArea)
    }).
    AddFormItem(outputArea).
    SetFieldTextColor(tcell.NewRGBColor(255,255,255)).
    SetButtonBackgroundColor(tcell.NewRGBColor(56,70,90)).
    SetLabelColor(tcell.NewRGBColor(255,255,255))

  form.SetBackgroundColor(tcell.NewRGBColor(51,56,82)).
    SetTitle("Scrape Jobs").
    SetBorder(true).
    SetTitleAlign(tview.AlignLeft)
    //SetFieldBackgroundColor(tcell.NewRGBColor(201,20,251))
    //SetButtonBackgroundColor(tcell.NewRGBColor(201,20,251))
  if err := app.SetRoot(form, true).Run(); err != nil {
    panic(err)
  }
}

func main() {

  DrawCLI()
//  flag.StringVar(&actionFlag, "action", "", "Action to perform")
//  flag.Parse()
//
//  if actionFlag == "scrape" {
//    // check if setup.cfg exists
//    _, err := os.Stat("setup.cfg")
//    if os.IsNotExist(err) {
//      fmt.Println("setup.cfg does not exist. Run 'go run main.go -action=setup' to create it")
//      return
//    }
//    job_scraper.ScrapeJobs()
//  } else if actionFlag == "setup" {
//    setup.Setup()
//  } else {
//    fmt.Println("Invalid action")
//  }
}
