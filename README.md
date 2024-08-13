# Job Scraper

This is a simple job scraper that scrapes job positings from job board websites, and saves the job postings in a CSV file. The job scraper is built using Go and the Colly web scraping framework.

### Installation
1. Install Go
2. Install the Colly web scraping framework by running `go get -u github.com/gocolly/colly`
3. Clone the repository

### Usage
1. Run the job scraper setup by running `go run main.go --action setup`
2. Edit the `setup.cfg` file to specify the keywords you want to search for
3. Run the job scraper by running `go run main.go --action scrape`

### Future improvements
- [ ] Add support for more job boards
- [ ] Add support for more job fields through the use of editable config files
- [ ] Add support for more output formats (e.g. JSON, XML)
- [ ] Add support for more output destinations (e.g. database, cloud storage)

### Supported Job Boards
- [LinkedIn](https://www.linkedin.com/)

### Future Supported Job Boards
- [ ] [Indeed](https://www.indeed.com/)
- [ ] [Dice](https://www.dice.com/)
- [ ] [Glassdoor](https://www.glassdoor.com/)
