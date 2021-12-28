package bluebench

import (
	"math/rand"
	"time"
)

const (
	STR1  string = "いい"
	STR2  string = "い青"
	STR3  string = "ケスベ"
	STR4  string = "ケスベチン"
	STR5  string = "ケベス"
	STR6  string = "ケンサ"
	STR7  string = "サスい"
	STR8  string = "サスケ"
	STR9  string = "サスチン"
	STR10 string = "サベ"
	STR11 string = "サン"
	STR12 string = "サンチ"
	STR13 string = "スケスケ"
	STR14 string = "スケベ"
	STR15 string = "スケベン"
	STR16 string = "スケン"
	STR17 string = "スベスベ"
	STR18 string = "スベチン"
	STR19 string = "チチン"
	STR20 string = "チチンチン"
	STR21 string = "チベン"
	STR22 string = "チベン青"
	STR23 string = "チン"
	STR24 string = "チンス"
	STR25 string = "チンチ"
	STR26 string = "チンチン"
	STR27 string = "チンベ"
	STR28 string = "チンベス"
	STR29 string = "ベス"
	STR30 string = "ベチン"
	STR31 string = "ベチンチン"
	STR32 string = "ベベ"
	STR33 string = "ベベベ"
	STR34 string = "ベベン"
	STR35 string = "ベン"
	STR36 string = "ベンい"
	STR37 string = "ベンスケ"
	STR38 string = "ベンチ"
	STR39 string = "ベンベ"
	STR40 string = "青"
	STR41 string = "青い"
	STR42 string = "青い青"
	STR43 string = "青ケ"
	STR44 string = "青サ"
	STR45 string = "青ス"
	STR46 string = "青スケ"
	STR47 string = "青チ"
	STR48 string = "青チン"
	STR49 string = "青ベ"
	STR50 string = "青ベチ"
	STR51 string = "青ベンチ"
	STR52 string = "青ン"
	STR53 string = "青ンチ"
	STR54 string = "青青"
)

func Shindan() string {
	candidates := [54]string{STR1, STR2, STR3, STR4, STR5, STR6, STR7, STR8, STR9, STR10, STR11, STR12, STR13, STR14, STR15, STR16, STR17, STR18, STR19, STR20, STR21, STR22, STR23, STR24, STR25, STR26, STR27, STR28, STR29, STR30, STR31, STR32, STR33, STR34, STR35, STR36, STR37, STR38, STR39, STR40, STR41, STR42, STR43, STR44, STR45, STR46, STR47, STR48, STR49, STR50, STR51, STR52, STR53, STR54}
	rand.Seed(time.Now().UnixNano())
	aoi := candidates[rand.Intn(54)]
	rand.Seed(time.Now().UnixNano())
	bench := candidates[rand.Intn(54)]
	rand.Seed(time.Now().UnixNano())
	sasuke := candidates[rand.Intn(54)]
	return aoi + bench + "／" + sasuke
}
