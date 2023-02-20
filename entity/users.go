package entity

type Users struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	UserCourses []UserCourses `json:"user_courses" gorm:"foreignKey:UsersID"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserResponse struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	UserCourses []UserCourses `json:"user_courses" gorm:"foreignKey:UsersID"`
}
