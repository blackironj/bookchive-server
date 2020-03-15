package model

//Users Model
type Users struct {
	UUID     string  `db:"uuid"`
	Email    string  `db:"email"`
	Name     *string `db:"name"`
	SigninDT *int64  `db:"signin_dt"`
}

//Books Model
type Books struct {
	ID         string  `db:"id" json:"id" binding:"required"`
	Title      string  `db:"title" json:"title" binding:"required"`
	Subtitle   *string `db:"subtitle" json:"subtitle"`
	Authors    string  `db:"authors" json:"authors" binding:"required"`
	Publisher  string  `db:"publisher" json:"publisher"`
	Categories *string `db:"categories" json:"categories"`
	Thumbnail  *string `db:"thumbnail" json:"thumbnail"`
	Pages      *string `db:"pages" json:"pages"`
}
