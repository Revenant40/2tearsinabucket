package main

import (

  "net/http"
  "fmt"
  "flag"
  "os"
  "bufio"
  "log"
  "strconv"
  "time"
  "net"
  "os/exec"

)

var timeout = time.Duration(55 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
    return net.DialTimeout(network, addr, timeout)
}


func getStatusCode(url string) string {

  transport := http.Transport{
        Dial: dialTimeout,
    }

    client := http.Client{
        Transport: &transport,
    }

	resp, err := client.Get(url)
	if err != nil {
	  return err.Error()
	}
  defer resp.Body.Close()
	return strconv.Itoa(resp.StatusCode)
}

func main() {

  target := flag.String("t"," ", "target")
	wordlist := flag.String("w","bucket-names.txt", "wordlist")

	var words = fmt.Sprintf("%s", *wordlist)

  flag.Parse()

  fmt.Println(`
            /   /(
           /(  ((_)    2 TEARS IN A BUCKET
          (#_)`)
  fmt.Println("\n ++ Created By: Revanent")
  fmt.Println(" ++ ASCII Art By: VK")
  fmt.Println(" ++ Target Set To:", *target)
  fmt.Println(" ++ Wordlist Set To:", *wordlist)
  fmt.Println("\n+++ Only checking for buckets with a 200 or 403 response code +++\n")
	
  if *target == " " {
		fmt.Println("Use the -t flag to set your target (e.g ./2tearsinabucket -t 'targets name')\n")
		os.Exit(1)
	}

  readFile, err := os.Open(words)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

  defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

    var t = fmt.Sprintf("http://%s%s.s3.amazonaws.com",*target,fileScanner.Text())

    var lastCheck = getStatusCode(t)
		var url = fmt.Sprintf("s3://%s%s",*target,fileScanner.Text())

    if lastCheck == "200" || lastCheck == "403" {
        fmt.Println(t)
        fmt.Println("Got status code", lastCheck, "\n" )
    }

    if lastCheck == "200" {
      app := "aws"
      arg0 := "s3"
      arg1 := "ls"
      arg2 := url
      cmd := exec.Command(app, arg0, arg1, arg2)
      stdout, err := cmd.Output()

        if err != nil {
          fmt.Println(err.Error())
          return
        }
          fmt.Println(string(stdout) + "\n")

      }
	}

}
