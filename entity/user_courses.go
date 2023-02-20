package entity

type UserCourses struct {
	Id        int `json:"id" gorm:"primaryKey"`
	UsersID   int `json:"users_id" gorm:"foreignKey:UsersID;"`
	Users     Users
	CoursesID int `json:"courses_id" gorm:"foreignKey:CoursesID;"`
	Courses   Courses
}

type CreateUserCourseRequest struct {
	UsersID   int `json:"users_id"`
	CoursesID int `json:"courses_id"`
}

type UpdateUserCourseRequest struct {
	UsersID   int `json:"users_id"`
	CoursesID int `json:"courses_id"`
}

type GetUserCourseResponse struct {
	Id        int `json:"id"`
	UsersID   int `json:"users_id"`
	CoursesID int `json:"courses_id"`
}
