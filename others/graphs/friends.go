package main

import (
	"fmt"
	"strings"
)

func contains(s []string, present string) bool {
	for _, el := range s {
		if el == present {
			return true
		}
	}
	return false
}

func networkOfFriends(network map[string][]string, person string, listOfFriends []string) []string {
	firstDegreeFriends := network[person]
	if len(firstDegreeFriends) == 0 {
		return listOfFriends
	}

	for _, friend := range firstDegreeFriends {
		if contains(listOfFriends, friend) {
			continue
		}

		listOfFriends = networkOfFriends(network, friend, append(listOfFriends, friend))
	}

	return listOfFriends
}

func numberOfFriends(relationships []string, people []string) map[string]int {
	result := map[string]int{}

	network := map[string][]string{}

	for _, edge := range relationships {
		friends := strings.Split(edge, ":")

		network[friends[0]] = append(network[friends[0]], friends[1])
		network[friends[1]] = append(network[friends[1]], friends[0])

		// for idx, friend := range friends {
		// 	otherIdx := (idx + 1) % 2
		// 	otherFriend := friends[otherIdx]

		// 	listOfFriends := network[friend]

		// 	if !contains(listOfFriends, otherFriend) {
		// 		network[friend] = append(listOfFriends, otherFriend)
		// 	}
		// }
	}

	// fmt.Println(network)

	for _, person := range people {
		listOfFriends := networkOfFriends(network, person, []string{person})

		result[person] = len(listOfFriends) - 1
	}

	return result
}

func main() {
	fmt.Println(numberOfFriends([]string{"Anne:Barbara", "Barbara:Ivan", "Vinny:Vlad"}, []string{"Anne", "Vinny"})) // Anne:2, Vinny: 1

	fmt.Println(numberOfFriends([]string{"Octavia:Zoltan", "Zoltan:Marko", "Marko:Diego", "Diego:Andres"}, []string{"Octavia", "Diego"})) // Octavia: 4, Diego: 4

	fmt.Println(numberOfFriends([]string{"Vanja:Sergio", "Sergio:Pedro", "Pedro:Martin", "Martin:Pablo", "Pablo:Sergio", "Jelena:Ivan", "Jelena:Alan", "Alan:Tomislav"}, []string{"Tomislav", "Martin"})) // Tomislav: 3, Martin: 4

	fmt.Println(numberOfFriends([]string{"Alvaro:Alex", "Roman:Nikola", "Octavia:Barbara", "Joao:Marko", "Luis:Vanja", "Gabriel:Gustavo", "Alan:Pablo", "Ivan:Andres", "Artem:Anne", "Martin:Alessandro", "Sebastian:Vinny", "Eduardo:Francis", "Zoltan:Vlad"}, []string{"Zoltan", "Ivan"})) // Zoltan:1, Ivan: 1

	fmt.Println(numberOfFriends([]string{"David:Francis", "Francis:Alessandro", "Alessandro:Ivan", "Tomislav:Ivan", "Anne:Tomislav", "Anne:David", "Francis:Tomislav"}, []string{"Francis", "David"})) // Francis: 5, David: 5

	fmt.Println(numberOfFriends([]string{"Alessandro:Anna", "Anna:Anne", "Anne:Barbara", "Barbara:David", "David:Francis", "Francis:Eduardo", "Eduardo:Anna", "Eduardo:Alessandro", "Luis:Marko", "Joao:Vlad", "Vlad:Luka", "Luka:Nikola", "Nikola:Roman", "Vlad:Roman", "Vlad:Vinny", "Vinny:Roman", "Vlad:Andres", "Vinny:Ivan"}, []string{"Barbara", "Joao"})) // Barbara: 6, Joao: 7
}
