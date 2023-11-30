package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var danes = time.Now()
var danesf = time.Time.Format(danes, "02.01.2006")
var seznam = make(map[string]string)
var dni int

func main() {
	dateList := []string{"13.11.1968", "22.02.1968", "19.11.2019", "02.12.1990", "01.11.1996"}

	fmt.Printf("%-15s %-15s %-15s %-10s\n", "RD", "Datum", "danes", "dni")
	for _, dt := range dateList {
		tryDate(dt)
	}
}

func tryDate(mydate string) {

	var rd string = mydate
	var newY string

	thisDateList := strings.Split(danesf, ".")
	thisd := thisDateList[0]
	thism := thisDateList[1]
	thisY := thisDateList[2]

	myDateList := strings.Split(rd, ".")
	myd := myDateList[0]
	mym := myDateList[1]
	// myY := myDateList[2]

	var YY int
	if myd >= thisd && mym >= thism {
		newY = thisY
	} else if myd < thisd && mym > thism {
		newY = thisY
	} else {
		YY, _ = strconv.Atoi(thisY)
		YY += 1
		newY = strconv.Itoa(YY)
	}

	myRDf := fmt.Sprintf("%s.%s.%s", myd, mym, newY)
	myRDfinal, _ := time.Parse("02.01.2006", myRDf)
	diff := myRDfinal.Sub(danes)

	if myRDf == danesf {
		dni = 0
	} else {
		dni = int(diff.Hours()/24) + 1
	}

	fmt.Printf("%-15s %-15s %-15s %-10v\n", mydate, myRDf, danesf, dni)

	/*
		myRDs := item.rd
		myRD := strings.Split(myRDs, ".")
		myRDf := fmt.Sprintf("%s.%s.%s", myRD[0], myRD[1], thisY)
		danes2 := danes.AddDate(1, 0, 0)

		var YY string
		if myRD[0] < thisDateList[0] || myRD[1] < thisDateList[1] {
			YY = time.Time.Format(danes2, "2006")
		} else {
			YY = time.Time.Format(danes, "2006")
		}

		myRDf = fmt.Sprintf("%s.%s.%s", myRD[0], myRD[1], YY)

		prsd = myRDf
		myRDfinal, _ := time.Parse("02.01.2006", myRDf)

		diff := myRDfinal.Sub(danes)

		//v1.0.06
		if myRDf == time.Time.Format(danes, "02.01.2006") {
			dni = 0
		} else {
			// v1.0.5 ( +1)
			dni = int(diff.Hours()/24) + 1
		}

		now, _ := strconv.Atoi(YY)
		then, _ := strconv.Atoi(myRD[2])

		prsyrs = now - then
		prsdni = dni

		item.dt = prsd
		item.yrs = prsyrs
		item.dys = prsdni

		seznam2 = append(seznam2, *item)

		}
	*/
}
