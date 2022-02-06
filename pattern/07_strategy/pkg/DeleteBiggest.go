package pkg

type DecreaserSalaryBiggest struct{}

func (d DecreaserSalaryBiggest)Delete(s *SalaryTable){
	var NameToDelete string
	max := 0
	for name, salary := range s.Salary{
		if salary > max{
			max = salary
			NameToDelete = name
		}
	}
	delete(s.Salary, NameToDelete)
}