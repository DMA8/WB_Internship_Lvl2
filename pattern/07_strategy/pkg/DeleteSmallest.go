package pkg

type DecreaserSalarySmallest struct{}

func (d DecreaserSalarySmallest)Delete(s *SalaryTable){
	var NameToDelete string
	min := 1<<62
	for name, salary := range s.Salary {
		if salary < min{
			min = salary
			NameToDelete = name
		}
	}
	delete(s.Salary, NameToDelete)
}