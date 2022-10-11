package dto

type ExerciseDTO struct {
	Id           uint32   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	MuscleGroups []string `json:"muscle_groups"`
	ImageData    []byte   `json:"image_data"`
	VideoLink    string   `json:"video_link"`
}
