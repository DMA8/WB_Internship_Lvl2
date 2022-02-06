package pkg

type Optimizer interface{
	Delete(*SalaryTable)
}

type SalaryTable struct{
	Salary				map[string]int
	OptimizationAlgo	Optimizer
}

func (s *SalaryTable)Optimize(){
	if s.OptimizationAlgo != nil {
		s.OptimizationAlgo.Delete(s)
	}
}