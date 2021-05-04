package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestGetFromToRange(t *testing.T) {
	toRange, to := GetFromToRange(time.Now(), time.Hour*24*30)
	fmt.Printf("%v \n", toRange)
	fmt.Printf("%v \n", to)
}

func TestGetDayFromNowRange(t *testing.T) {
	f, to := GetDayFromNowRange(2)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", to)
}

func TestGetFormToRecently(t *testing.T) {
	recently, to := GetFormToRecently(time.Minute * 60)
	fmt.Printf("%v        %v\n", recently, to)
}
