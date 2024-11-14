package models

type Job struct {
	Id             uint    `gorm:"primaryKey;autoIncrement"`
	UserId         uint    `gorm:"not null"`
	Name           string  `gorm:"type:varchar(120);not null"`
	Description    string  `gorm:"type:varchar(240)"`
	Duration       uint    `gorm:"type:integer;not null"`
	AfterSalesDays uint    `gorm:"type:integer"`
	Price          float32 `gorm:"type:numeric(10,2);not null"`

	User User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

type CreateJobRequest struct {
	Name           string  `json:"name"             validate:"required,min=3,max=120"`
	Description    string  `json:"description"      validate:"omitempty,max=240"`
	Duration       uint    `json:"duration"         validate:"required,gt=0"`
	AfterSalesDays uint    `json:"after_sales_days" validate:"omitempty,gt=0"`
	Price          float64 `json:"price"            validate:"required,gt=0"`
}

type GetJobResponse struct {
	Id             uint    `json:"id"`
	UserId         uint    `json:"user_id"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Duration       uint    `json:"duration"`
	AfterSalesDays uint    `json:"after_sales_days"`
	Price          float32 `json:"price"`
}

type GetUserJobsResponse struct {
	Jobs []GetJobResponse `json:"jobs"`
}

type UpdateJobRequest struct {
	Name           string  `json:"name"             validate:"required,min=3"`
	Description    string  `json:"description"      validate:"omitempty,max=240"`
	Duration       uint    `json:"duration"         validate:"required,gt=0"`
	AfterSalesDays uint    `json:"after_sales_days" validate:"omitempty,gt=0"`
	Price          float64 `json:"price"            validate:"required,gt=0"`
}
