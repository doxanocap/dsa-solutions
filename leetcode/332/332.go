package _32

import (
	"slices"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	graph := map[string][]string{}

	// строим направленный граф
	for _, ticket := range tickets {
		graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
	}

	// сортируем по лексикографическому порядку
	for key := range graph {
		sort.Strings(graph[key])
	}

	// инициализируем пустой маршрут с cap=len(tickets)+1,
	// для избежания лишних аллокаций памяти, при стековом аппенде в слайс
	itinerary := make([]string, 0, len(tickets)+1)
	itinerary = dfs("JFK", graph, itinerary) // запускаем DFS с начальной точкой "JFK"

	// реверсим массив, так как мы добавляли в слайс с конца(stack)
	slices.Reverse(itinerary)
	return itinerary
}

func dfs(
	currCity string,
	graph map[string][]string,
	itinerary []string,
) []string {

	// продолжаем DFS, пока есть соседние города
	for len(graph[currCity]) > 0 {
		nextCity := graph[currCity][0]              // берем следующий город
		graph[currCity] = graph[currCity][1:]       // удаляем его из графа
		itinerary = dfs(nextCity, graph, itinerary) // рекурсивно вызываем DFS для следующего города
	}

	// добавляем в слайс если дошли до конечного города по принципу LIFO
	itinerary = append(itinerary, currCity)

	// обязательно возвращаем новый слайс
	// это может и не самый эффективный способ(например могли работать с *[]string),
	// но явно более идиоматичный подход в контексте Go
	return itinerary
}
