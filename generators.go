package main

import (
	"math/rand"
	"strconv"

	"github.com/brianvoe/gofakeit"
	"github.com/bxcodec/faker"
	"github.com/zazab/zhash"
)

func GenerateName(hash *zhash.Hash, path ...string) {
	hash.Set(faker.Name(), path...)
}

func GenerateCompany(hash *zhash.Hash, path ...string) {
	hash.Set(gofakeit.Company(), path...)
}

func GenerateEmail(hash *zhash.Hash, path ...string) {
	hash.Set(gofakeit.Email(), path...)
}

func GenerateIntString(hash *zhash.Hash, path ...string) {
	min := 100000
	max := 999999

	hash.Set(strconv.Itoa(rand.Intn(max-min)+min), path...)
}

func GenerateSenLicense(hash *zhash.Hash, path ...string) {
	min := 1000000
	max := 9999999

	hash.Set("SEN-L"+strconv.Itoa(rand.Intn(max-min)+min), path...)
}

func GenerateCity(hash *zhash.Hash, path ...string) {
	hash.Set(gofakeit.City(), path...)
}

func GenerateZip(hash *zhash.Hash, path ...string) {
	hash.Set(gofakeit.Zip(), path...)
}

func GenerateState(hash *zhash.Hash, path ...string) {
	hash.Set(gofakeit.State(), path...)
}
