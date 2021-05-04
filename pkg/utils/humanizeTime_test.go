package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestRelTimeEn(t *testing.T) {
	fmt.Printf("%v \n", TimeFromNow(time.Now().Add(-time.Second*50)))
	fmt.Printf("%v \n", TimeFromNow(time.Now().Add(-time.Minute*50)))
	fmt.Printf("%v \n", TimeFromNow(time.Now().Add(-time.Hour*50)))
	fmt.Printf("%v \n", TimeFromNow(time.Now().Add(-time.Hour*500)))
	fmt.Printf("%v \n", TimeFromNow(time.Now().Add(-time.Hour*5000)))
	fmt.Printf("%v \n", TimeFromNow(time.Now().Add(-time.Hour*50000)))
	fmt.Printf("%v \n", TimeFromNow(time.Now().Add(-time.Hour*500000)))
}
