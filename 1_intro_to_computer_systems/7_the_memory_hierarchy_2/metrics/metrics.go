package metrics

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

type UserData struct {
	Ages     []uint8
	Payments []Payment
}

type DollarAmount struct {
	dollars, cents uint64
}

type Payment struct {
	amount DollarAmount
	time   time.Time
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

func AveragePaymentAmount(payments []Payment) float64 {
	average, count := 0.0, 0.0
	for _, p := range payments {
		count += 1
		amount := float64(p.amount.dollars) + float64(p.amount.cents)/100
		average += (amount - average) / count
	}
	return average
}

// Compute the standard deviation of payment amounts
func StdDevPaymentAmount(payments []Payment) float64 {
	mean := AveragePaymentAmount(payments)
	squaredDiffs, count := 0.0, 0.0
	for _, p := range payments {
		count += 1
		amount := float64(p.amount.dollars) + float64(p.amount.cents)/100
		diff := amount - mean
		squaredDiffs += diff * diff
	}
	return math.Sqrt(squaredDiffs / count)
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

	payments := make([]Payment, len(paymentLines))
	for i, line := range paymentLines {
		paymentCents, _ := strconv.Atoi(line[0])
		datetime, _ := time.Parse(time.RFC3339, line[1])
		payments[i] = Payment{
			DollarAmount{uint64(paymentCents / 100), uint64(paymentCents % 100)},
			datetime,
		}
	}

	return UserData{Ages: ages, Payments: payments}
}
