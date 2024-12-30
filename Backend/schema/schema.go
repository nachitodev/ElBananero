package schema

import (
	"gorm.io/gorm"
)

// Video represents a video entity in the system
type Video struct {
	gorm.Model               // This already includes ID, CreatedAt, UpdatedAt, and DeletedAt
	Titulo      string       `gorm:"type:varchar(255)" json:"titulo"`
	Categoria   string       `gorm:"type:varchar(255)" json:"categoria"`
	Hashtags    string       `gorm:"type:varchar(255)" json:"hashtags"`
	URLVideo    string       `gorm:"type:varchar(255)" json:"url_video"`
	Comentarios []Comentario `gorm:"foreignKey:VideoID" json:"comentarios,omitempty"`
}

// Comentario represents a comment associated with a video
type Comentario struct {
	gorm.Model
	VideoID uint   `gorm:"index" json:"video_id"`
	Nombre  string `gorm:"type:varchar(255)" json:"nombre"`
	Texto   string `gorm:"type:varchar(255)" json:"texto"`
}
