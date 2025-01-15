package data

import "time"

type Comment struct {
	Commentid   int32     `json:"comment_id"`
	Pageid      int32     `json:"page_id"`
	Userid      int32     `json:"user_id"`
	Createdat   time.Time `json:"created_at"`
	Editedbool  bool      `json:"edited_bool"`
	Upvotes     int32     `json:"up_votes"`
	Downvotes   int32     `json:"down_votes"`
	Commentdata string    `json:"comment_data"`
	Parentid    int32     `json:"parent_id"`
}

type Page struct {
	Pageurl       string    `json:"page_url"`
	Pageid        int32     `json:"page_id"`
	Commentscount int32     `json:"comments_count"`
	Createdat     time.Time `json:"created_at"`
}

type User struct {
	Userid    int32  `json:"user_id"`
	Username  string `json:"user_name"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Emailid   string `json:"email_id"`
}

type Userrelation struct {
	Vieweruserid    int32  `json:"viewer_user_id"`
	Commentoruserid int32  `json:"commentor_user_id"`
	Tag             string `json:"tag"`
	Positivity      bool   `json:"positivity"`
}
