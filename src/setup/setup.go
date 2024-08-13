package setup

import (
  "fmt";
  "os";
  "log";
  "bufio";
)

func Setup() {
  // Create an setup file with the following contents:
  // keyword1
  // keyword2

  file, err := os.Create("setup.cfg")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close() 
  
  initialized_string := "keyword1\nkeyword2\n"
  writer := bufio.NewWriter(file)
  _, err = writer.WriteString(initialized_string)
  if err != nil {
    log.Fatal(err)
  }
  writer.Flush()

  

  fmt.Println("Created setup.txt")
  fmt.Println("Please enter the keywords you would like to search for in the setup.txt file")
}

