package metrics

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
)

type UserData struct {
	Ages     []uint8
	Payments []uint32
}

func AverageAge(ages []uint8) float64 {
	average0, average1 := 0.0, 0.0
	weight := 1.0 / float64(len(ages))
	i := 0

	for i = 0; i < len(ages)-1; i += 2 {
		average0 += float64(ages[i]) * weight
		average1 += float64(ages[i+1]) * weight
	}

	for _ = i; i < len(ages); i++ {
		average0 += float64(ages[i]) * weight
	}

	return average0 + average1
}

func AveragePaymentAmount(payments []uint32) float64 {
	average0, average1 := 0.0, 0.0
	weight := 1.0 / float64(len(payments))
	i := 0

	for i = 0; i < len(payments)-1; i += 2 {
		average0 += float64(payments[i]) * weight
		average1 += float64(payments[i+1]) * weight
	}

	for _ = i; i < len(payments); i++ {
		average0 += float64(payments[i]) * weight
	}

	return (average0 + average1) * 0.01
}

// Compute the standard deviation of payment amounts
func StdDevPaymentAmount(payments []uint32) float64 {
	mean := AveragePaymentAmount(payments) * 100
	squaredDiffs, count := 0.0, 0.0
	for _, p := range payments {
		count += 1
		amount := float64(p)
		diff := amount - mean
		squaredDiffs += diff * diff
	}
	return math.Sqrt(squaredDiffs/count) * 0.01
}

func LoadData() UserData {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Unable to read users.csv", err)
	}
	reader := csv.NewReader(f)
	userLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse users.csv as csv", err)
	}

	ages := make([]uint8, len(userLines))
	for i, line := range userLines {
		age, _ := strconv.Atoi(line[2])
		ages[i] = uint8(age)
	}

	f, err = os.Open("payments.csv")
	if err != nil {
		log.Fatalln("Unable to read payments.csv", err)
	}
	reader = csv.NewReader(f)
	paymentLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse payments.csv as csv", err)
	}

	payments := make([]uint32, len(paymentLines))
	for i, line := range paymentLines {
		paymentCents, _ := strconv.Atoi(line[0])
		payments[i] = uint32(paymentCents)
	}

	return UserData{Ages: ages, Payments: payments}
}
