package entity

type CoursesCategories struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Courses []Courses `json:"courses" gorm:"foreignKey:CoursesCategoriesID"`
}

type CreateCoursesCategoriesRequest struct {
	Name string `json:"name"`
}

type UpdateCoursesCategoriesRequest struct {
	Name string `json:"name"`
}

type GetCoursesCategoriesResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
