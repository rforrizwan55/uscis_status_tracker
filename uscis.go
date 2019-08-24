package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

var username string
var password string

func main() {
	statusMap := make(map[string]string)
	start(statusMap)
	ticker := time.NewTicker(60 * time.Minute)
	for range ticker.C {
		//fmt.Println("Tick at", t)
		start(statusMap)
	}
}

func start(statusMap map[string]string) {
	var center = []string{"wac", "eac"}
	var uniqueStatuses []string
	var statuses []string

	var cc int
	var start int64
	var limit int
	flag.IntVar(&cc, "center", 0, "wac = 0 & eac = 1; default:0")
	flag.Int64Var(&start, "start", 1919054100, "start of the series; default:1919054100")
	flag.IntVar(&limit, "limit", 100, "no. of cases; default:100")
	flag.StringVar(&username, "email", "", "gmail address for sending and receiving status")
	flag.StringVar(&password, "pass", "", "gmail password for sending and receiving status")

	flag.Parse()
	for i := 0; i < limit; i++ {

		var caseId = center[cc] + strconv.FormatInt(start+int64(i), 10)
		var statusData = callApi(caseId)
		var parsedData = parseHtml(statusData)

		if len(parsedData) > 0 {
			statuses = append(statuses, parsedData)

			if !contains(uniqueStatuses, parsedData) {
				uniqueStatuses = append(uniqueStatuses, parsedData)
			}

			var oldStatus = statusMap[caseId]
			if oldStatus != "" && oldStatus != parsedData {
				var statusString = fmt.Sprintf("Change in status for %s - %s -----> %s\n", caseId, statusMap[caseId], parsedData)
				fmt.Printf(statusString)
				sendEmail(statusString, "Change status")
			}

			statusMap[caseId] = parsedData
		}

		//fmt.Printf("%s - %s\n", caseId, parsedData)

		if i == limit-1 {
			var summary = ""
			var s = time.Now().String()
			summary = summary + "-----------SUMMARY at " + s + "---------\n"
			for _, v := range uniqueStatuses {
				summary = summary + fmt.Sprintf("%s - %d\n", v, occurrence(statuses, v))
			}
			summary = summary + "--------map size : " + strconv.Itoa(len(statusMap)) + "---------\n"
			summary = summary + "-----------SUMMARY---------\n"

			fmt.Println(summary)
			sendEmail(summary, "Summary - "+s)
			statuses = []string{}
		}
	}
}

func callApi(x string) string {
	r, err := http.Post("https://egov.uscis.gov/casestatus/mycasestatus.do?appReceiptNum="+x, "", nil)
	if err != nil {
		panic(err)
	}
	data, _ := ioutil.ReadAll(r.Body)

	return string(data)
}

func parseHtml(data string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	content := string("")
	if err != nil {
		panic(err)
	}

	doc.Find(".rows.text-center").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		content = s.Find("h1").Text()
	})

	return content
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func occurrence(data []string, item string) int {
	var count = 0
	for _, v := range data {
		if v == item {
			count++
		}
	}

	return count
}

func sendEmail(body string, subject string) {
	auth := smtp.PlainAuth(
		"",
		username,
		password,
		"smtp.gmail.com",
	)
	msg := []byte("To: " + username + "\r\nSubject:" +
		subject + "\r\n" +
		"\r\n" + body +
		"\r\n")
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		username,
		[]string{username},
		[]byte(msg),
	)
	if err != nil {
		log.Fatal(err)
	}
}
