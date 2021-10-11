package metrics

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

type UserId int
type UserMap map[UserId]*User

type UserData struct {
	Users    UserMap
	Payments []Payment
}

type Address struct {
	fullAddress string
	zip         int
}

type DollarAmount struct {
	dollars, cents uint64
}

type Payment struct {
	amount DollarAmount
	time   time.Time
}

type User struct {
	id       UserId
	name     string
	age      int
	address  Address
	payments []Payment
}

func AverageAge(users UserMap) float64 {
	average, count := 0.0, 0.0
	for _, u := range users {
		count += 1
		average += (float64(u.age) - average) / count
	}
	return average
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

	users := make(UserMap, len(userLines))
	for _, line := range userLines {
		id, _ := strconv.Atoi(line[0])
		name := line[1]
		age, _ := strconv.Atoi(line[2])
		address := line[3]
		zip, _ := strconv.Atoi(line[3])
		users[UserId(id)] = &User{UserId(id), name, age, Address{address, zip}, []Payment{}}
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

	return UserData{Users: users, Payments: payments}
}
