package mongo

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestMongoDB(test *testing.T) {
	//mongopool := Initialize("172.16.000.000:27017", 100).Database("uu").Collection("cu")
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	// key : 364e_ve005050001012_SD6_2021
	forever := make(chan struct{})
	//today := time.Now().Format("20060102")

	//initUser(50000) // userNumber만큼의 유저 랜덤으로 권한 주고 생성
	println(randomNumber(1, 24))
	println(randomNumber(1, 24))
	println(randomNumber(1, 24))
	println(randomNumber(1, 24))

	<-forever

}

func randomNumbers(num int) []int {
	var result []int
	for i := 0; i < num; i++ {
		result = append(result, randomNumber(21000, 36000))

	}
	return result
}

func randomNumber(min, max int) int { // cno
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func randomString(length int) string { // id
	rand.Seed(time.Now().UTC().UnixNano())
	charSet := []rune("abcdedfghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVXYZ1234567890")
	var output strings.Builder
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteRune(randomChar)
	}
	return output.String()
}
