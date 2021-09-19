package main

import (
	"fmt"
	"math/rand"
	"time"
)

func remove(arr []string, i int) []string {
	return append(arr[:i], arr[i+1:]...)
}

func shuffle(arr []string) []string {
	result := make([]string, 0)

	for len(arr) > 0 {
		k := rand.Intn(len(arr))
		result = append(result, arr[k])
		arr = remove(arr, k)
	}

	return result
}

func main() {
	rand.Seed(time.Now().Unix())

	classes := []string{
		"Barbarian",
		"Crusader",
		"Demon Hunter",
		"Monk",
		"Necromancer",
		"Witch Doctor",
		"Wizard",
	}

	gender := []string{
		"Male",
		"Female",
	}

	acts := []string{
		"Act I",
		"Act II",
		"Act III",
		"Act IV",
		"Act V",
	}

	event := []string{
		"Nephalem Rift",
		"Greater Nephalem Rift",
		"Bounties",
	}

	shuffled := shuffle(classes)
	fmt.Println(len(shuffled))
	for _, name := range shuffled {
		fmt.Println(name)
	}
}
