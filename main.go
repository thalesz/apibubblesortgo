package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Função BubbleSort com acompanhamento dos passos
func bubbleSort(arr []int) ([]int, [][]int) {
	n := len(arr)
	var steps [][]int // Para armazenar os passos intermediários

	// Adiciona o estado inicial do array
	steps = append(steps, append([]int(nil), arr...))

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Troca os elementos
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		// Adiciona o array após cada iteração externa
		steps = append(steps, append([]int(nil), arr...))
	}
	return arr, steps
}

func main() {
	r := gin.Default()

	// Endpoint que recebe os números, aplica o Bubble Sort e retorna o resultado
	r.GET("/bubblesort", func(c *gin.Context) {
		// Pega os parâmetros da query string
		numbers := c.DefaultQuery("numbers", "")

		if numbers == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Você precisa passar os números na query string"})
			return
		}

		// Converte a string para um slice de inteiros
		numStrings := strings.Split(numbers, ",")
		var nums []int
		for _, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Todos os valores devem ser números inteiros válidos"})
				return
			}
			nums = append(nums, num)
		}

		// Aplica o Bubble Sort
		sorted, steps := bubbleSort(nums)

		// Retorna o array ordenado e os passos intermediários
		c.JSON(http.StatusOK, gin.H{
			"array_ordenado": sorted,
			"passos":         steps,
		})
	})

	// Inicia o servidor na porta 8080
	r.Run(":8080")
}
