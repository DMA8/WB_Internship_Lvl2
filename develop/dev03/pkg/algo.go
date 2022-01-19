package pkg

func QsortStrings(a []string, left, rigth int) {
	if rigth - left > 0 {
		pivotInd := partitionStrings(a, left, rigth)
		QsortStrings(a, left, pivotInd - 1)
		QsortStrings(a, pivotInd + 1, rigth)
	}
}

func partitionStrings(a []string, left, rigth int) int {
	var i, pivot, wall int
	pivot = rigth 
	wall = left 
	for i = left; i < rigth; i++ {
		if a[i] < a[pivot] {
			a[i], a[wall] = a[wall], a[i]
			wall++
		}
	}
	a[pivot], a[wall] = a[wall], a[pivot]
	return wall
}

func QsortInt(a []int, left, rigth int) {
	if rigth - left > 0 {
		pivotInd := partitionInts(a, left, rigth)
		QsortInt(a, left, pivotInd - 1)
		QsortInt(a, pivotInd + 1, rigth)
	}
}

func partitionInts(a []int, left, rigth int) int {
	var i, pivot, wall int
	pivot = rigth 
	wall = left 
	for i = left; i < rigth; i++ {
		if a[i] < a[pivot] {
			a[i], a[wall] = a[wall], a[i]
			wall++
		}
	}
	a[pivot], a[wall] = a[wall], a[pivot]
	return wall
}