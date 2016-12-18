package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Votes struct {
	Total      float64
	Republican float64
	Democrat   float64
	Other      float64
}

type County struct {
	ID      int
	Name    string
	Votes08 Votes
	Votes12 Votes
	Votes16 Votes
}

func GetFlippedCounties0816(counties []County) ([]County, []County) {
	var rToD []County
	var dToR []County
	for _, c := range counties {
		if c.Votes08.Democrat < c.Votes08.Republican && c.Votes16.Democrat > c.Votes16.Republican {
			rToD = append(rToD, c)
		}
		if c.Votes08.Democrat > c.Votes08.Republican && c.Votes16.Democrat < c.Votes16.Republican {
			dToR = append(dToR, c)
		}
	}
	return rToD, dToR
}

func GetFlippedCounties1216(counties []County) ([]County, []County) {
	var rToD []County
	var dToR []County
	for _, c := range counties {
		if c.Votes12.Democrat < c.Votes12.Republican && c.Votes16.Democrat > c.Votes16.Republican {
			rToD = append(rToD, c)
		}
		if c.Votes12.Democrat > c.Votes12.Republican && c.Votes16.Democrat < c.Votes16.Republican {
			dToR = append(dToR, c)
		}
	}
	return rToD, dToR
}

func getNum(s string) float64 {
	d, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("BAD: %v", err)
	}
	return d
}

func ParseRow(row string) County {
	var county County
	pieces := strings.Split(row, ",")
	id, err := strconv.Atoi(pieces[0])
	if err != nil {
		log.Fatalf("DERP: %v", err)
	}
	county.ID = id
	county.Name = pieces[1]
	county.Votes08.Total = getNum(pieces[2])
	county.Votes08.Democrat = getNum(pieces[3])
	county.Votes08.Republican = getNum(pieces[4])
	county.Votes08.Other = getNum(pieces[5])
	county.Votes12.Total = getNum(pieces[6])
	county.Votes12.Democrat = getNum(pieces[7])
	county.Votes12.Republican = getNum(pieces[8])
	county.Votes12.Other = getNum(pieces[9])
	county.Votes16.Total = getNum(pieces[10])
	county.Votes16.Democrat = getNum(pieces[11])
	county.Votes16.Republican = getNum(pieces[12])
	county.Votes16.Other = getNum(pieces[13])
	return county
}

func main() {
	file, err := os.Open("US_County_Level_Presidential_Results_08-16.csv")
	if err != nil {
		log.Fatal("ono! %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var counties []County
	for scanner.Scan() {
		counties = append(counties, ParseRow(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("poo %v", err)
	}

	rToD, dToR := GetFlippedCounties0816(counties)
	log.Printf("08-16: R->D: %d, D->R: %d", len(rToD), len(dToR))

	rToD, dToR = GetFlippedCounties1216(counties)
	log.Printf("12-16: R->D: %d, D->R: %d", len(rToD), len(dToR))
}
