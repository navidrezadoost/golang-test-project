package common

import (
	"car/handlers/config"
	"golang.org/x/exp/rand"
	"math"
	"strconv"
	"time"
)

func GenerateOtp() string {
	cfg := config.GetConfig()
	rand.Seed(uint64(time.Now().UnixNano()))
	m := int(math.Pow(10, float64(cfg.Otp.Digits-1)))
	i := int(math.Pow(10, float64(cfg.Otp.Digits)) - 1)
	var number = rand.Intn(i-m) + m
	return strconv.Itoa(number)
}
