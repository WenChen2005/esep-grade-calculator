package esepunittests

type GradeCalculator struct {
	grades       []Grade
	passFailMode bool
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator(passFailMode ...bool) *GradeCalculator {
	mode := false

	if len(passFailMode) > 0 {
		mode = passFailMode[0]
	}

	return &GradeCalculator{
		grades:       make([]Grade, 0),
		passFailMode: mode,
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()
	if gc.passFailMode {
		if numericalGrade >= 70 {
			return "Pass"
		}
		return "Fail"
	}

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := computeAverageByType(gc.grades, Assignment)
	exam_average := computeAverageByType(gc.grades, Exam)
	essay_average := computeAverageByType(gc.grades, Essay)

	weighted_grade := float64(assignment_average)*.5 +
		float64(exam_average)*.35 +
		float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverageByType(grades []Grade, gradeType GradeType) int {
	sum := 0
	count := 0

	for _, grade := range grades {
		if grade.Type == gradeType {
			sum += grade.Grade
			count++
		}
	}

	return sum / count
}
