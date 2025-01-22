package main

import (
	"fmt"
	"sort"
)

func main() {
	var workerSlice Company
	workerSlice.AddWorkerInfo("Михаил", "директор", 200, 5)

	workerSlice.AddWorkerInfo("Андрей", "рабочий", 180, 3)
	workerSlice.AddWorkerInfo("Вася", "зам. директора", 180, 3)
	workerSlice.AddWorkerInfo("Игорь", "начальник цеха", 120, 2)
	out, _ := workerSlice.SortWorkers()
	for _, v := range out {
		fmt.Println(v)
	}

	// fmt.Println(workerSlice)
}

type Company []Worker
type Worker struct {
	Name            string
	Position        string
	Salary          uint
	ExperienceYears uint
}
type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	*c = append(*c, Worker{Name: name, Position: position, Salary: salary, ExperienceYears: experience})
	return nil
}

func (c *Company) SortWorkers() ([]string, error) {
	WorkerMap := map[string]int{"директор": 4, "зам. директора": 3, "начальник цеха": 2, "мастер": 1, "рабочий": 0}
	sort.Slice(*c, func(i, j int) bool {
		if (*c)[i].Salary*(*c)[i].ExperienceYears != (*c)[j].Salary*(*c)[j].ExperienceYears {
			return (*c)[i].Salary*(*c)[i].ExperienceYears > (*c)[j].Salary*(*c)[j].ExperienceYears
		} else {
			return WorkerMap[(*c)[i].Position] > WorkerMap[(*c)[j].Position]
		}
	})
	out := make([]string, len(*c))
	for i, v := range *c {
		out[i] = fmt.Sprintf("%s - %d - %s", v.Name, v.ExperienceYears*v.Salary*12, v.Position)
	}
	return out, nil
}
