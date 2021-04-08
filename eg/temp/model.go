package temp

import "time"

type AmpArticle struct {
	ID        int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Title     string `gorm:"column:title;type:varchar(255)" json:"title"`
	Content   string `gorm:"column:content;type:longtext;not null" json:"content"`
	Time      string `gorm:"column:time;type:varchar(30)" json:"time"`
	Cover     string `gorm:"column:cover;type:varchar(1024)" json:"cover"`
	Name      string `gorm:"column:name;type:varchar(200)" json:"name"`
	Catid     int    `gorm:"column:catid;type:int(10)" json:"catid"`
	AccountID int    `gorm:"column:account_id;type:int(10)" json:"account_id"`
	Aid       int    `gorm:"unique;column:aid;type:int(10)" json:"aid"` // aid唯
}

// ApsT [...]
type ApsT struct {
	ID   int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Name string `gorm:"column:name;type:varchar(255)" json:"name"`
}

// Article [...]
type Article struct {
	ID       int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Title    string `gorm:"index;column:title;type:varchar(255)" json:"title"`
	Content  string `gorm:"column:content;type:text" json:"content"`
	CateID   int    `gorm:"column:cate_id;type:int(10) unsigned" json:"cate_id"`
	CateName string `gorm:"column:cate_name;type:varchar(255)" json:"cate_name"`
}

// ArticleList [...]
type ArticleList struct {
	ID       int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Title    string `gorm:"unique;column:title;type:varchar(255)" json:"title"`
	Content  string `gorm:"column:content;type:text" json:"content"`
	CateID   int    `gorm:"column:cate_id;type:int(10) unsigned" json:"cate_id"`
	CateName string `gorm:"column:cate_name;type:varchar(255)" json:"cate_name"`
}

func (a *ArticleList) TableName() string {

	return "article_list"
}

// Asp [...]
type Asp struct {
	UId  int  `gorm:"primary_key;column:uid;type:int(10) unsigned;not null" json:"uid"`
	Cid  int  `gorm:"index;column:cid;type:int(10) unsigned" json:"cid"`
	ApsT ApsT `gorm:"association_foreignkey:cid;foreignkey:id" json:"aps_t_list"`
}

// Cate [...]
type Cate struct {
	ID    int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Name  string `gorm:"unique;column:name;type:varchar(255);not null" json:"name"`
	Pid   int    `gorm:"column:pid;type:int(11);not null" json:"pid"`
	Intro string `gorm:"column:intro;type:varchar(255);not null" json:"intro"`
}

// IDeaContent [...]
type IDeaContent struct {
	ID          int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Title       string `gorm:"column:title;type:varchar(255)" json:"title"`
	Keywords    string `gorm:"column:keywords;type:varchar(255)" json:"keywords"`
	CateName    string `gorm:"column:cate_name;type:varchar(100)" json:"cate_name"`
	CateID      int    `gorm:"column:cate_id;type:int(10) unsigned" json:"cate_id"`
	CreateTime  int    `gorm:"column:create_time;type:int(11)" json:"create_time"`
	Description string `gorm:"column:description;type:varchar(255)" json:"description"`
	Content     string `gorm:"column:content;type:text" json:"content"`
}

// IDeaTag [...]
type IDeaTag struct {
	ID       int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Name     string `gorm:"column:name;type:varchar(50)" json:"name"`
	Desc     string `gorm:"column:desc;type:varchar(255)" json:"desc"`
	Keywords string `gorm:"column:keywords;type:varchar(255)" json:"keywords"`
}

// Mp [...]
type Mp struct {
	ID    int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	URL   string `gorm:"column:url;type:varchar(1024)" json:"url"`
	Cover string `gorm:"column:cover;type:varchar(1024)" json:"cover"`
}

// MpAccountList [...]
type MpAccountList struct {
	ID        int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	URL       string `gorm:"column:url;type:varchar(255)" json:"url"`
	Name      string `gorm:"index;column:name;type:varchar(255)" json:"name"`
	Img       string `gorm:"column:img;type:varchar(255)" json:"img"`
	BaseURL   string `gorm:"column:base_url;type:varchar(255)" json:"base_url"`
	Pic       string `gorm:"column:pic;type:varchar(255)" json:"pic"`
	AccountID int    `gorm:"unique;column:account_id;type:int(11)" json:"account_id"`
	CateID    int    `gorm:"column:cate_id;type:int(11)" json:"cate_id"`
	Desc      string `gorm:"column:desc;type:varchar(255)" json:"desc"`
	AddTime   int    `gorm:"column:add_time;type:int(11)" json:"add_time"`
}

// MpArticleList [...]
type MpArticleList struct {
	ID        int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Title     string `gorm:"column:title;type:varchar(255)" json:"title"`
	Content   string `gorm:"column:content;type:longtext;not null" json:"content"`
	Time      string `gorm:"column:time;type:varchar(30)" json:"time"`
	Cover     string `gorm:"column:cover;type:varchar(1024)" json:"cover"`
	Name      string `gorm:"column:name;type:varchar(200)" json:"name"`
	Catid     int    `gorm:"column:catid;type:int(10) unsigned" json:"catid"`
	AccountID int    `gorm:"column:account_id;type:int(10)" json:"account_id"`
	Aid       int    `gorm:"unique;column:aid;type:int(10) unsigned" json:"aid"`
}

// MpCateList [...]
type MpCateList struct {
	ID   int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Cid  int    `gorm:"unique;column:cid;type:int(10) unsigned" json:"cid"`
	URL  string `gorm:"column:url;type:varchar(255)" json:"url"`
	Name string `gorm:"index;column:name;type:varchar(255)" json:"name"`
	Img  string `gorm:"column:img;type:varchar(255)" json:"img"`
}

// Newip [...]
type Newip struct {
	ID int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`
	IP string `gorm:"unique;column:ip;type:varchar(255)" json:"ip"`
}

// Opts [...]
type Opts struct {
	Key   string `gorm:"index;column:key;type:varchar(64);not null" json:"key"`
	Value string `gorm:"column:value;type:varchar(2048);not null" json:"value"`
	Intro string `gorm:"column:intro;type:varchar(255);not null" json:"intro"`
}

// Post [...]
type Post struct {
	ID              int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	CateID          int       `gorm:"column:cate_id;type:int(11);not null" json:"cate_id"`
	UserID          int       `gorm:"column:user_id;type:int(11) unsigned;not null" json:"user_id"`
	Type            int8      `gorm:"column:type;type:tinyint(4);not null" json:"type"`     // 0 为文章，1 为页面
	Status          int8      `gorm:"column:status;type:tinyint(4);not null" json:"status"` // 0 为草稿，1 为待审核，2 为已拒绝，3 为已经发布
	Title           string    `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Path            string    `gorm:"column:path;type:varchar(255);not null" json:"path"`   // URL 的 pathname
	Summary         string    `gorm:"column:summary;type:longtext;not null" json:"summary"` // 摘要
	MarkdownContent string    `gorm:"column:markdown_content;type:longtext;not null" json:"markdown_content"`
	Content         string    `gorm:"column:content;type:longtext;not null" json:"content"`
	AllowComment    int8      `gorm:"column:allow_comment;type:tinyint(4);not null" json:"allow_comment"` // 1 为允许， 0 为不允许
	CreateTime      time.Time `gorm:"index;column:create_time;type:datetime" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time;type:datetime;not null" json:"update_time"`
	IsPublic        int8      `gorm:"column:is_public;type:tinyint(4);not null" json:"is_public"` // 1 为公开，0 为不公开
	CommentNum      int       `gorm:"column:comment_num;type:int(11);not null" json:"comment_num"`
	Options         string    `gorm:"column:options;type:varchar(4096);not null" json:"options"` // 一些选项，JSON 结构
}

// PostTag [...]
type PostTag struct {
	ID     int `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	PostID int `gorm:"unique_index:post_tag;column:post_id;type:int(11);not null" json:"post_id"`
	TagID  int `gorm:"unique_index:post_tag;column:tag_id;type:int(11);not null" json:"tag_id"`
}

// Saveip [...]
type Saveip struct {
	ID   int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`
	IP   string `gorm:"column:ip;type:varchar(255)" json:"ip"`
	Name string `gorm:"column:name;type:varchar(25)" json:"name"`
}

// SyUser [...]
type SyUser struct {
	ID       int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Phone    int64  `gorm:"unique;column:phone;type:bigint(11)" json:"phone"`
	Password string `gorm:"column:password;type:varchar(50)" json:"password"`
	Token    string `gorm:"column:token;type:text" json:"token"`
	Ext      string `gorm:"column:ext;type:varchar(255)" json:"ext"`
	Status   int8   `gorm:"column:status;type:tinyint(1) unsigned" json:"status"`
}

// SyUserCopy1 [...]
type SyUserCopy1 struct {
	ID       int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Phone    int64  `gorm:"unique;column:phone;type:bigint(11)" json:"phone"`
	Password string `gorm:"column:password;type:varchar(50)" json:"password"`
	Token    string `gorm:"column:token;type:text" json:"token"`
	Ext      string `gorm:"column:ext;type:varchar(255)" json:"ext"`
}

// Tag [...]
type Tag struct {
	ID    int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Name  string `gorm:"unique;column:name;type:varchar(255);not null" json:"name"`
	Intro string `gorm:"column:intro;type:varchar(255);not null" json:"intro"`
}

// TtAccount [...]
type TtAccount struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Name       string `gorm:"column:name;type:varchar(255)" json:"name"`
	Cateid     int    `gorm:"column:cateid;type:int(11)" json:"cateid"`
	Desc       string `gorm:"column:desc;type:varchar(255)" json:"desc"`
	AccountNum string `gorm:"column:account_num;type:varchar(255)" json:"account_num"`
	Pic        string `gorm:"column:pic;type:varchar(255)" json:"pic"`
	URLs       string `gorm:"column:urls;type:varchar(1024)" json:"urls"`
	Status     int8   `gorm:"column:status;type:tinyint(1)" json:"status"`
}

// TtArticle [...]
type TtArticle struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Title      string `gorm:"unique;column:title;type:varchar(150)" json:"title"`
	Name       string `gorm:"column:name;type:varchar(255)" json:"name"`
	Content    string `gorm:"column:content;type:text" json:"content"`
	Catid      int    `gorm:"column:catid;type:int(11)" json:"catid"`
	Pic        string `gorm:"column:pic;type:varchar(255)" json:"pic"`
	CateName   string `gorm:"column:cate_name;type:varchar(50)" json:"cate_name"`
	CreateTime int    `gorm:"column:create_time;type:int(10) unsigned" json:"create_time"`
	AccountID  int    `gorm:"column:account_id;type:int(11)" json:"account_id"`
	Desc       string `gorm:"column:desc;type:varchar(512)" json:"desc"`
	Cover      string `gorm:"column:cover;type:varchar(255)" json:"cover"`
	URLs       string `gorm:"column:urls;type:varchar(255)" json:"urls"`
}

// TtCate [...]
type TtCate struct {
	ID   int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Cid  int    `gorm:"column:cid;type:int(10) unsigned" json:"cid"`
	URL  string `gorm:"column:url;type:varchar(255)" json:"url"`
	Name string `gorm:"index;column:name;type:varchar(255)" json:"name"`
	Img  string `gorm:"column:img;type:varchar(255)" json:"img"`
}

// TxWx [...]
type TxWx struct {
	ID    int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Catid int    `gorm:"column:catid;type:int(10) unsigned" json:"catid"`
	URL   string `gorm:"unique;column:url;type:varchar(1024)" json:"url"`
	Pic   string `gorm:"column:pic;type:varchar(1024)" json:"pic"`
}

// User [...]
type User struct {
	ID     int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Name   string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Num    string    `gorm:"unique;column:num;type:varchar(255);not null" json:"num"`
	Pass   string    `gorm:"column:pass;type:varchar(255);not null" json:"pass"`
	Role   int       `gorm:"column:role;type:int(11);not null" json:"role"`
	Email  string    `gorm:"unique;column:email;type:varchar(255);not null" json:"email"`
	Phone  string    `gorm:"column:phone;type:varchar(255);not null" json:"phone"`
	IP     string    `gorm:"column:ip;type:varchar(32);not null" json:"ip"`
	Ecount int       `gorm:"column:ecount;type:int(11);not null" json:"ecount"`
	Ltime  time.Time `gorm:"column:ltime;type:datetime;not null" json:"ltime"`
	Ctime  time.Time `gorm:"column:ctime;type:datetime;not null" json:"ctime"`
}

// ZwArticle [...]
type ZwArticle struct {
	ID       int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Title    string `gorm:"index;column:title;type:varchar(255)" json:"title"`
	Content  string `gorm:"column:content;type:text" json:"content"`
	CateID   int    `gorm:"column:cate_id;type:int(10) unsigned" json:"cate_id"`
	CateName string `gorm:"column:cate_name;type:varchar(255)" json:"cate_name"`
}
