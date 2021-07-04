package sorting

func SortIns(Arr []int) []int {
	for i := 1; i < len(Arr); i++ {
		for j := i; j > 0 && Arr[j-1] > Arr[j]; j-- {
			Arr[j-1], Arr[j] = Arr[j], Arr[j-1]
		}
	}
	return Arr
}

func BubbleSort(Arr []int) []int {
	swapped := true

	for swapped {

		swapped = false

		for i := 0; i < len(Arr)-1; i++ {
			if Arr[i+1] < Arr[i] {

				Arr[i+1], Arr[i] = Arr[i], Arr[i+1]

				swapped = true

			}

		}

	}

	return Arr

}
