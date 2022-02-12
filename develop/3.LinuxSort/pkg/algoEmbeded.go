package pkg

import "sort"

// Sorter for embeding sorting methods
type Sorter interface{
	Sort(interface{})
}

type stringSlice	[]string
type float64Slice	[]float64
type intsSlice		[]int


func (p stringSlice)Len()int {return len(p)}
func (p float64Slice)Len()int {return len(p)}
func (p intsSlice)Len()int {return len(p)}


func (p stringSlice)Less(i, j int) bool{return p[i] < p[j]}
func (p float64Slice)Less(i, j int) bool{return p[i] < p[j]}
func (p intsSlice)Less(i, j int) bool{return p[i] < p[j]}


func (p stringSlice)Swap(i, j int){ p[i], p[j] = p[j], p[i]}
func (p float64Slice)Swap(i, j int){ p[i], p[j] = p[j], p[i]}
func (p intsSlice)Swap(i, j int){ p[i], p[j] = p[j], p[i]}

//QuickSort sorts any datatype with predeclared methods Less, Swap, and Len
type QuickSort struct{}

//Sort methods sorts given data. In works with []string, []float64, []int
func (q QuickSort)Sort(a interface{}){
	switch a.(type) {
	case []string:
		var s stringSlice
		s = a.([]string)
		sort.Sort(s)
	case []float64:
		var s float64Slice
		s = a.([]float64)
		sort.Sort(s)
	case []int:
		var s intsSlice
		s = a.([]int)
		sort.Sort(s)
	}
}
