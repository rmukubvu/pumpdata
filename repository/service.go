package repository

import (
	"fmt"
	"strings"
	"time"
)

func expirationCallback(key string, value interface{}) {
	fmt.Printf("This key(%s) with value %s has expired\n", key, value)
	//do the sending of the notifications here...
	str := fmt.Sprintf("%v", value)
	split := strings.Split(str, "|")
	sendEmail(split[0], split[1], split[2])
}

func serviceExpiryWithNotify(key, value string, duration time.Duration) {
	//serviceCache.SetWithTTL(key,value,duration)
	serviceCache.SetWithTTL(key, value, 30*time.Second)
}
