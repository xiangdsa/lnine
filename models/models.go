package models

type BzSeoEcmsDownload struct {
	ID      int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	Classid int64 `gorm:"column:classid" db:"classid" json:"classid" form:"classid"`
	//Ttid       int64  `gorm:"column:ttid" db:"ttid" json:"ttid" form:"ttid"`
	//Onclick    int64  `gorm:"column:onclick" db:"onclick" json:"onclick" form:"onclick"`
	//Plnum      int64  `gorm:"column:plnum" db:"plnum" json:"plnum" form:"plnum"`
	//Totaldown  int64  `gorm:"column:totaldown" db:"totaldown" json:"totaldown" form:"totaldown"`
	//Newspath   string `gorm:"column:newspath" db:"newspath" json:"newspath" form:"newspath"`
	//Filename   string `gorm:"column:filename" db:"filename" json:"filename" form:"filename"`
	//Userid     int64  `gorm:"column:userid" db:"userid" json:"userid" form:"userid"`
	//Username   string `gorm:"column:username" db:"username" json:"username" form:"username"`
	//Firsttitle int64  `gorm:"column:firsttitle" db:"firsttitle" json:"firsttitle" form:"firsttitle"`
	//Isgood     int64  `gorm:"column:isgood" db:"isgood" json:"isgood" form:"isgood"`
	//Ispic      int64  `gorm:"column:ispic" db:"ispic" json:"ispic" form:"ispic"`
	//Istop      int64  `gorm:"column:istop" db:"istop" json:"istop" form:"istop"`
	//Isqf       int64  `gorm:"column:isqf" db:"isqf" json:"isqf" form:"isqf"`
	//Ismember   int64  `gorm:"column:ismember" db:"ismember" json:"ismember" form:"ismember"`
	//Isurl      int64  `gorm:"column:isurl" db:"isurl" json:"isurl" form:"isurl"`
	//Truetime   int64  `gorm:"column:truetime" db:"truetime" json:"truetime" form:"truetime"`
	//Lastdotime int64  `gorm:"column:lastdotime" db:"lastdotime" json:"lastdotime" form:"lastdotime"`
	//Havehtml   int64  `gorm:"column:havehtml" db:"havehtml" json:"havehtml" form:"havehtml"`
	//Groupid    int64  `gorm:"column:groupid" db:"groupid" json:"groupid" form:"groupid"`
	//Userfen    int64  `gorm:"column:userfen" db:"userfen" json:"userfen" form:"userfen"`
	//Titlefont  string `gorm:"column:titlefont" db:"titlefont" json:"titlefont" form:"titlefont"`
	//Titleurl   string `gorm:"column:titleurl" db:"titleurl" json:"titleurl" form:"titleurl"`
	//Stb        int64  `gorm:"column:stb" db:"stb" json:"stb" form:"stb"`
	//Fstb       int64  `gorm:"column:fstb" db:"fstb" json:"fstb" form:"fstb"`
	//Restb      int64  `gorm:"column:restb" db:"restb" json:"restb" form:"restb"`
	//Keyboard   string `gorm:"column:keyboard" db:"keyboard" json:"keyboard" form:"keyboard"`
	Title string `gorm:"column:title" db:"title" json:"title" form:"title"`
	//Newstime   int64  `gorm:"column:newstime" db:"newstime" json:"newstime" form:"newstime"`
	//Titlepic   string `gorm:"column:titlepic" db:"titlepic" json:"titlepic" form:"titlepic"`
	//Eckuid     int64  `gorm:"column:eckuid" db:"eckuid" json:"eckuid" form:"eckuid"`
	//Softfj     string `gorm:"column:softfj" db:"softfj" json:"softfj" form:"softfj"`
	//Language   string `gorm:"column:language" db:"language" json:"language" form:"language"`
	//Softtype   string `gorm:"column:softtype" db:"softtype" json:"softtype" form:"softtype"`
	//Softsq     string `gorm:"column:softsq" db:"softsq" json:"softsq" form:"softsq"`
	//Star       int64  `gorm:"column:star" db:"star" json:"star" form:"star"`
	//Filesize   string `gorm:"column:filesize" db:"filesize" json:"filesize" form:"filesize"`
	Softsay string `gorm:"column:softsay" db:"softsay" json:"softsay" form:"softsay"`
	//Ftitle     string `gorm:"column:ftitle" db:"ftitle" json:"ftitle" form:"ftitle"`
	//Dname      string `gorm:"column:dname" db:"dname" json:"dname" form:"dname"`
	//Diggtop    int64  `gorm:"column:diggtop" db:"diggtop" json:"diggtop" form:"diggtop"`
	//Diggdown   int64  `gorm:"column:diggdown" db:"diggdown" json:"diggdown" form:"diggdown"`
	//Morepic    string `gorm:"column:morepic" db:"morepic" json:"morepic" form:"morepic"`
	//Tishi      string `gorm:"column:tishi" db:"tishi" json:"tishi" form:"tishi"`
	//Smalltext  string `gorm:"column:smalltext" db:"smalltext" json:"smalltext" form:"smalltext"`
	//Infozm     string `gorm:"column:infozm" db:"infozm" json:"infozm" form:"infozm"`
	//Fmimg      string `gorm:"column:fmimg" db:"fmimg" json:"fmimg" form:"fmimg"`
}

//Bz_Seo_ecms_az_data_1
type BzSeoEcmsAzData1 struct {
	ID         int64  `gorm:"column:id" db:"id" json:"id" form:"id"`
	Classid    int64  `gorm:"column:classid" db:"classid" json:"classid" form:"classid"`
	//Keyid      string `gorm:"column:keyid" db:"keyid" json:"keyid" form:"keyid"`
	//Dokey      int64  `gorm:"column:dokey" db:"dokey" json:"dokey" form:"dokey"`
	//Newstempid int64  `gorm:"column:newstempid" db:"newstempid" json:"newstempid" form:"newstempid"`
	//Closepl    int64  `gorm:"column:closepl" db:"closepl" json:"closepl" form:"closepl"`
	//Haveaddfen int64  `gorm:"column:haveaddfen" db:"haveaddfen" json:"haveaddfen" form:"haveaddfen"`
	//Infotags   string `gorm:"column:infotags" db:"infotags" json:"infotags" form:"infotags"`
	//Writer     string `gorm:"column:writer" db:"writer" json:"writer" form:"writer"`
	//Befrom     string `gorm:"column:befrom" db:"befrom" json:"befrom" form:"befrom"`
	Newstext   string `gorm:"column:newstext" db:"newstext" json:"newstext" form:"newstext"`
}

type BzSeoEcmsIos struct {
	ID int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	//Classid    int64  `gorm:"column:classid" db:"classid" json:"classid" form:"classid"`
	//Ttid       int64  `gorm:"column:ttid" db:"ttid" json:"ttid" form:"ttid"`
	//Onclick    int64  `gorm:"column:onclick" db:"onclick" json:"onclick" form:"onclick"`
	//Plnum      int64  `gorm:"column:plnum" db:"plnum" json:"plnum" form:"plnum"`
	//Totaldown  int64  `gorm:"column:totaldown" db:"totaldown" json:"totaldown" form:"totaldown"`
	//Newspath   string `gorm:"column:newspath" db:"newspath" json:"newspath" form:"newspath"`
	//Filename   string `gorm:"column:filename" db:"filename" json:"filename" form:"filename"`
	//Userid     int64  `gorm:"column:userid" db:"userid" json:"userid" form:"userid"`
	//Username   string `gorm:"column:username" db:"username" json:"username" form:"username"`
	//Firsttitle int64  `gorm:"column:firsttitle" db:"firsttitle" json:"firsttitle" form:"firsttitle"`
	//Isgood     int64  `gorm:"column:isgood" db:"isgood" json:"isgood" form:"isgood"`
	//Ispic      int64  `gorm:"column:ispic" db:"ispic" json:"ispic" form:"ispic"`
	//Istop      int64  `gorm:"column:istop" db:"istop" json:"istop" form:"istop"`
	//Isqf       int64  `gorm:"column:isqf" db:"isqf" json:"isqf" form:"isqf"`
	//Ismember   int64  `gorm:"column:ismember" db:"ismember" json:"ismember" form:"ismember"`
	//Isurl      int64  `gorm:"column:isurl" db:"isurl" json:"isurl" form:"isurl"`
	//Truetime   int64  `gorm:"column:truetime" db:"truetime" json:"truetime" form:"truetime"`
	//Lastdotime int64  `gorm:"column:lastdotime" db:"lastdotime" json:"lastdotime" form:"lastdotime"`
	//Havehtml   int64  `gorm:"column:havehtml" db:"havehtml" json:"havehtml" form:"havehtml"`
	//Groupid    int64  `gorm:"column:groupid" db:"groupid" json:"groupid" form:"groupid"`
	//Userfen    int64  `gorm:"column:userfen" db:"userfen" json:"userfen" form:"userfen"`
	//Titlefont  string `gorm:"column:titlefont" db:"titlefont" json:"titlefont" form:"titlefont"`
	//Titleurl   string `gorm:"column:titleurl" db:"titleurl" json:"titleurl" form:"titleurl"`
	//Stb        int64  `gorm:"column:stb" db:"stb" json:"stb" form:"stb"`
	//Fstb       int64  `gorm:"column:fstb" db:"fstb" json:"fstb" form:"fstb"`
	//Restb      int64  `gorm:"column:restb" db:"restb" json:"restb" form:"restb"`
	//Keyboard   string `gorm:"column:keyboard" db:"keyboard" json:"keyboard" form:"keyboard"`
	Title      string `gorm:"column:title" db:"title" json:"title" form:"title"`
	//Newstime   int64  `gorm:"column:newstime" db:"newstime" json:"newstime" form:"newstime"`
	//Titlepic   string `gorm:"column:titlepic" db:"titlepic" json:"titlepic" form:"titlepic"`
	//Eckuid     int64  `gorm:"column:eckuid" db:"eckuid" json:"eckuid" form:"eckuid"`
	//Ftitle     string `gorm:"column:ftitle" db:"ftitle" json:"ftitle" form:"ftitle"`
	//Smalltext  string `gorm:"column:smalltext" db:"smalltext" json:"smalltext" form:"smalltext"`
	//Diggtop    int64  `gorm:"column:diggtop" db:"diggtop" json:"diggtop" form:"diggtop"`
	//Dname      string `gorm:"column:dname" db:"dname" json:"dname" form:"dname"`
	//Daxiao     string `gorm:"column:daxiao" db:"daxiao" json:"daxiao" form:"daxiao"`
	//Banbe      string `gorm:"column:banbe" db:"banbe" json:"banbe" form:"banbe"`
	//Xitong     string `gorm:"column:xitong" db:"xitong" json:"xitong" form:"xitong"`
	//Star       int64  `gorm:"column:star" db:"star" json:"star" form:"star"`
	//Yuyan      string `gorm:"column:yuyan" db:"yuyan" json:"yuyan" form:"yuyan"`
	//Shouquan   string `gorm:"column:shouquan" db:"shouquan" json:"shouquan" form:"shouquan"`
	//Yijuhua    string `gorm:"column:yijuhua" db:"yijuhua" json:"yijuhua" form:"yijuhua"`
	//Down       string `gorm:"column:down" db:"down" json:"down" form:"down"`
	//Morepic    string `gorm:"column:morepic" db:"morepic" json:"morepic" form:"morepic"`
	//Yugao      int64  `gorm:"column:yugao" db:"yugao" json:"yugao" form:"yugao"`
	//Tbsm       string `gorm:"column:tbsm" db:"tbsm" json:"tbsm" form:"tbsm"`
	//Fmimg      string `gorm:"column:fmimg" db:"fmimg" json:"fmimg" form:"fmimg"`
	//Fmtitle    string `gorm:"column:fmtitle" db:"fmtitle" json:"fmtitle" form:"fmtitle"`
}

type BzSeoEcmsNews struct {
	ID int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	//Classid    int64  `gorm:"column:classid" db:"classid" json:"classid" form:"classid"`
	//Ttid       int64  `gorm:"column:ttid" db:"ttid" json:"ttid" form:"ttid"`
	//Onclick    int64  `gorm:"column:onclick" db:"onclick" json:"onclick" form:"onclick"`
	//Plnum      int64  `gorm:"column:plnum" db:"plnum" json:"plnum" form:"plnum"`
	//Totaldown  int64  `gorm:"column:totaldown" db:"totaldown" json:"totaldown" form:"totaldown"`
	//Newspath   string `gorm:"column:newspath" db:"newspath" json:"newspath" form:"newspath"`
	//Filename   string `gorm:"column:filename" db:"filename" json:"filename" form:"filename"`
	//Userid     int64  `gorm:"column:userid" db:"userid" json:"userid" form:"userid"`
	//Username   string `gorm:"column:username" db:"username" json:"username" form:"username"`
	//Firsttitle int64  `gorm:"column:firsttitle" db:"firsttitle" json:"firsttitle" form:"firsttitle"`
	//Isgood     int64  `gorm:"column:isgood" db:"isgood" json:"isgood" form:"isgood"`
	//Ispic      int64  `gorm:"column:ispic" db:"ispic" json:"ispic" form:"ispic"`
	//Istop      int64  `gorm:"column:istop" db:"istop" json:"istop" form:"istop"`
	//Isqf       int64  `gorm:"column:isqf" db:"isqf" json:"isqf" form:"isqf"`
	//Ismember   int64  `gorm:"column:ismember" db:"ismember" json:"ismember" form:"ismember"`
	//Isurl      int64  `gorm:"column:isurl" db:"isurl" json:"isurl" form:"isurl"`
	//Truetime   int64  `gorm:"column:truetime" db:"truetime" json:"truetime" form:"truetime"`
	//Lastdotime int64  `gorm:"column:lastdotime" db:"lastdotime" json:"lastdotime" form:"lastdotime"`
	//Havehtml   int64  `gorm:"column:havehtml" db:"havehtml" json:"havehtml" form:"havehtml"`
	//Groupid    int64  `gorm:"column:groupid" db:"groupid" json:"groupid" form:"groupid"`
	//Userfen    int64  `gorm:"column:userfen" db:"userfen" json:"userfen" form:"userfen"`
	//Titlefont  string `gorm:"column:titlefont" db:"titlefont" json:"titlefont" form:"titlefont"`
	Titleurl string `gorm:"column:titleurl" db:"titleurl" json:"titleurl" form:"titleurl"`
	//Stb        int64  `gorm:"column:stb" db:"stb" json:"stb" form:"stb"`
	//Fstb       int64  `gorm:"column:fstb" db:"fstb" json:"fstb" form:"fstb"`
	//Restb      int64  `gorm:"column:restb" db:"restb" json:"restb" form:"restb"`
	//Keyboard   string `gorm:"column:keyboard" db:"keyboard" json:"keyboard" form:"keyboard"`
	//Title      string `gorm:"column:title" db:"title" json:"title" form:"title"`
	//Newstime   int64  `gorm:"column:newstime" db:"newstime" json:"newstime" form:"newstime"`
	//Titlepic   string `gorm:"column:titlepic" db:"titlepic" json:"titlepic" form:"titlepic"`
	//Eckuid     int64  `gorm:"column:eckuid" db:"eckuid" json:"eckuid" form:"eckuid"`
	//Ftitle     string `gorm:"column:ftitle" db:"ftitle" json:"ftitle" form:"ftitle"`
	//Smalltext  string `gorm:"column:smalltext" db:"smalltext" json:"smalltext" form:"smalltext"`
	//Diggtop    int64  `gorm:"column:diggtop" db:"diggtop" json:"diggtop" form:"diggtop"`
}
type BzSeoEcmsVideo struct {
	ID int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	//Classid    int64  `gorm:"column:classid" db:"classid" json:"classid" form:"classid"`
	//Ttid       int64  `gorm:"column:ttid" db:"ttid" json:"ttid" form:"ttid"`
	//Onclick    int64  `gorm:"column:onclick" db:"onclick" json:"onclick" form:"onclick"`
	//Plnum      int64  `gorm:"column:plnum" db:"plnum" json:"plnum" form:"plnum"`
	//Totaldown  int64  `gorm:"column:totaldown" db:"totaldown" json:"totaldown" form:"totaldown"`
	//Newspath   string `gorm:"column:newspath" db:"newspath" json:"newspath" form:"newspath"`
	//Filename   string `gorm:"column:filename" db:"filename" json:"filename" form:"filename"`
	//Userid     int64  `gorm:"column:userid" db:"userid" json:"userid" form:"userid"`
	//Username   string `gorm:"column:username" db:"username" json:"username" form:"username"`
	//Firsttitle int64  `gorm:"column:firsttitle" db:"firsttitle" json:"firsttitle" form:"firsttitle"`
	//Isgood     int64  `gorm:"column:isgood" db:"isgood" json:"isgood" form:"isgood"`
	//Ispic      int64  `gorm:"column:ispic" db:"ispic" json:"ispic" form:"ispic"`
	//Istop      int64  `gorm:"column:istop" db:"istop" json:"istop" form:"istop"`
	//Isqf       int64  `gorm:"column:isqf" db:"isqf" json:"isqf" form:"isqf"`
	//Ismember   int64  `gorm:"column:ismember" db:"ismember" json:"ismember" form:"ismember"`
	//Isurl      int64  `gorm:"column:isurl" db:"isurl" json:"isurl" form:"isurl"`
	//Truetime   int64  `gorm:"column:truetime" db:"truetime" json:"truetime" form:"truetime"`
	//Lastdotime int64  `gorm:"column:lastdotime" db:"lastdotime" json:"lastdotime" form:"lastdotime"`
	//Havehtml   int64  `gorm:"column:havehtml" db:"havehtml" json:"havehtml" form:"havehtml"`
	//Groupid    int64  `gorm:"column:groupid" db:"groupid" json:"groupid" form:"groupid"`
	//Userfen    int64  `gorm:"column:userfen" db:"userfen" json:"userfen" form:"userfen"`
	//Titlefont  string `gorm:"column:titlefont" db:"titlefont" json:"titlefont" form:"titlefont"`
	Titleurl string `gorm:"column:titleurl" db:"titleurl" json:"titleurl" form:"titleurl"`
	//Stb        int64  `gorm:"column:stb" db:"stb" json:"stb" form:"stb"`
	//Fstb       int64  `gorm:"column:fstb" db:"fstb" json:"fstb" form:"fstb"`
	//Restb      int64  `gorm:"column:restb" db:"restb" json:"restb" form:"restb"`
	//Keyboard   string `gorm:"column:keyboard" db:"keyboard" json:"keyboard" form:"keyboard"`
	//Title      string `gorm:"column:title" db:"title" json:"title" form:"title"`
	//Newstime   int64  `gorm:"column:newstime" db:"newstime" json:"newstime" form:"newstime"`
	//Titlepic   string `gorm:"column:titlepic" db:"titlepic" json:"titlepic" form:"titlepic"`
	//Eckuid     int64  `gorm:"column:eckuid" db:"eckuid" json:"eckuid" form:"eckuid"`
	//Ftitle     string `gorm:"column:ftitle" db:"ftitle" json:"ftitle" form:"ftitle"`
	//Smalltext  string `gorm:"column:smalltext" db:"smalltext" json:"smalltext" form:"smalltext"`
	//Diggtop    int64  `gorm:"column:diggtop" db:"diggtop" json:"diggtop" form:"diggtop"`
	//Newstext   string `gorm:"column:newstext" db:"newstext" json:"newstext" form:"newstext"`
	//Diggdown   int64  `gorm:"column:diggdown" db:"diggdown" json:"diggdown" form:"diggdown"`
}

type BzSeoEnewszt struct {
	Ztid int64 `gorm:"column:ztid" db:"ztid" json:"ztid" form:"ztid"`
	//Ztname      string `gorm:"column:ztname" db:"ztname" json:"ztname" form:"ztname"`
	//Onclick     int64  `gorm:"column:onclick" db:"onclick" json:"onclick" form:"onclick"`
	//Ztnum       int64  `gorm:"column:ztnum" db:"ztnum" json:"ztnum" form:"ztnum"`
	//Listtempid  int64  `gorm:"column:listtempid" db:"listtempid" json:"listtempid" form:"listtempid"`
	Ztpath string `gorm:"column:ztpath" db:"ztpath" json:"ztpath" form:"ztpath"`
	//Zttype      string `gorm:"column:zttype" db:"zttype" json:"zttype" form:"zttype"`
	//Zturl       string `gorm:"column:zturl" db:"zturl" json:"zturl" form:"zturl"`
	//Classid     int64  `gorm:"column:classid" db:"classid" json:"classid" form:"classid"`
	//Islist      int64  `gorm:"column:islist" db:"islist" json:"islist" form:"islist"`
	//Maxnum      int64  `gorm:"column:maxnum" db:"maxnum" json:"maxnum" form:"maxnum"`
	//Reorder     string `gorm:"column:reorder" db:"reorder" json:"reorder" form:"reorder"`
	//Intro       string `gorm:"column:intro" db:"intro" json:"intro" form:"intro"`
	//Ztimg       string `gorm:"column:ztimg" db:"ztimg" json:"ztimg" form:"ztimg"`
	//Zcid        int64  `gorm:"column:zcid" db:"zcid" json:"zcid" form:"zcid"`
	//Showzt      int64  `gorm:"column:showzt" db:"showzt" json:"showzt" form:"showzt"`
	//Ztpagekey   string `gorm:"column:ztpagekey" db:"ztpagekey" json:"ztpagekey" form:"ztpagekey"`
	//Classtempid int64  `gorm:"column:classtempid" db:"classtempid" json:"classtempid" form:"classtempid"`
	//Myorder     int64  `gorm:"column:myorder" db:"myorder" json:"myorder" form:"myorder"`
	//Usezt       int64  `gorm:"column:usezt" db:"usezt" json:"usezt" form:"usezt"`
	//Yhid        int64  `gorm:"column:yhid" db:"yhid" json:"yhid" form:"yhid"`
	//Endtime     int64  `gorm:"column:endtime" db:"endtime" json:"endtime" form:"endtime"`
	//Closepl     int64  `gorm:"column:closepl" db:"closepl" json:"closepl" form:"closepl"`
	//Checkpl     int64  `gorm:"column:checkpl" db:"checkpl" json:"checkpl" form:"checkpl"`
	//Restb       int64  `gorm:"column:restb" db:"restb" json:"restb" form:"restb"`
	//Usernames   string `gorm:"column:usernames" db:"usernames" json:"usernames" form:"usernames"`
	//Addtime     int64  `gorm:"column:addtime" db:"addtime" json:"addtime" form:"addtime"`
	//Pltempid    int64  `gorm:"column:pltempid" db:"pltempid" json:"pltempid" form:"pltempid"`
}

type BzSeoEnewspage struct {
	ID int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	//Title           string `gorm:"column:title" db:"title" json:"title" form:"title"`
	//Pagetext        string `gorm:"column:pagetext" db:"pagetext" json:"pagetext" form:"pagetext"`
	Path string `gorm:"column:path" db:"path" json:"path" form:"path"`
	//Classid         int64  `gorm:"column:classid" db:"classid" json:"classid" form:"classid"`
	//Pagetitle       string `gorm:"column:pagetitle" db:"pagetitle" json:"pagetitle" form:"pagetitle"`
	//Pagekeywords    string `gorm:"column:pagekeywords" db:"pagekeywords" json:"pagekeywords" form:"pagekeywords"`
	//Pagedescription string `gorm:"column:pagedescription" db:"pagedescription" json:"pagedescription" form:"pagedescription"`
	//Tempid          int64  `gorm:"column:tempid" db:"tempid" json:"tempid" form:"tempid"`
}
type BzSeoEnewsclass struct {
	Classid int64 `gorm:"column:classid" db:"classid" json:"classid" form:"classid"`
	//Bclassid      int64  `gorm:"column:bclassid" db:"bclassid" json:"bclassid" form:"bclassid"`
	//Classname     string `gorm:"column:classname" db:"classname" json:"classname" form:"classname"`
	Classpath string `gorm:"column:classpath" db:"classpath" json:"classpath" form:"classpath"`
}

//Bz_Seo_ecms_az
type BzSeoEcmsAz struct {
	ID      int64  `gorm:"column:id" db:"id" json:"id" form:"id"`
	Classid int64  `gorm:"column:classid" db:"classid" json:"classid" form:"classid"`
	//Ttid    int64  `gorm:"column:ttid" db:"ttid" json:"ttid" form:"ttid"`
	//Onclick int64  `gorm:"column:onclick" db:"onclick" json:"onclick" form:"onclick"`
	//Plnum   int64  `gorm:"column:plnum" db:"plnum" json:"plnum" form:"plnum"`
	Title   string `gorm:"column:title" db:"title" json:"title" form:"title"`
}
