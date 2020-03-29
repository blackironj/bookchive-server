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
	Pages      *int    `db:"pages" json:"pages"`
}

//Libraries Model
type Libraries struct {
	UK       int    `db:"uk" json:"uk"`
	UserUUID string `db:"user_uuid" json:"user_uuid"`
	BookID   string `db:"book_id"  json:"book_id"`
	AddedDT  *int64 `db:"added_dt" json:"added_dt"`
}

//BookInLibrary Model
type BookInLibrary struct {
	UK         int     `db:"uk" json:"uk"`
	BookID     string  `db:"book_id" json:"book_id"`
	AddedDT    *int64  `db:"added_dt" json:"added_dt"`
	Title      string  `db:"title" json:"title"`
	Authors    string  `db:"authors" json:"authors"`
	Categories *string `db:"categories" json:"categories"`
	Thumbnail  *string `db:"thumbnail" json:"thumbnail"`
}
