package models

type OccurrenceAudio struct {
	ID           string `json:"id" gorm:"primaryKey"`
	AudioData    []byte `json:"audioData"`
	OccurrenceID string `json:"occurrenceId"`
}
