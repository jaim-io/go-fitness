package dto

type MuscleGroupDTO struct {
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageData   []byte `json:"image_data"`
}
