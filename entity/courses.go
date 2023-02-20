package entity

type Courses struct {
	ID                  int    `json:"id"`
	Title               string `json:"title"`
	CoursesCategoriesID int    `json:"courses_categories_id"`
	CoursesCategories   CoursesCategories
	UserCourses         []UserCourses `json:"user_courses" gorm:"foreignKey:CoursesID"`
}

type CreateCourseRequest struct {
	Title               string `json:"title"`
	CoursesCategoriesID int    `json:"courses_categories_id"`
}

type UpdateCourseRequest struct {
	Title               string `json:"title"`
	CoursesCategoriesID int    `json:"courses_categories_id"`
}

type GetCourseResponse struct {
	ID                  int    `json:"id"`
	Title               string `json:"title"`
	CoursesCategoriesID int    `json:"courses_categories_id"`
}
