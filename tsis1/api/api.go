package api

type Player struct {
    Id          int    `json:"id"`
    Number      int    `json:"number"`
    FirstName   string `json:"first_name"`
    LastName    string `json:"last_name"`
    MarketValue int    `json:"market_value"`
    DateOfBirth string `json:"date_of_birth"`
    Nation      string `json:"nation"`
	Position    string `json:"position"`
}

var Players = []Player{
    {1, 11, "Mohamed.", "Salah", 65000000, "1992-06-15", "Egypt","at"},
    {2, 1, "Alisson", "Becker", 32000000, "1992-10-02", "Brazil","gk"},
    {3, 62, "Caoimhín", "Kelleher", 15000000, "1998-11-23", "Ireland","gk"},
    {4, 13, "Adrián", "", 600000, "1987-01-03", "Spain","gk"},
	{5,5,"Ibrahima","Konate",38000000,"1999-05-25","France","df"},
	{6,4,"Virgil","van Dijk",32000000,"1991-06-08","Netherlands","df"},
	{7,2,"Joe","Gomez",22000000,"1997-05-23","England","df"},
	{8,32,"Joel","Matip",8000000,"1991-08-08","Cameroon","df"},
	{9,78,"Jarell","Quansah",6000000,"2003-01-29","England","df"},
	{10,47,"Nathaniel","Phillips",4000000,"1997-03-21","England","df"},
	{11,26,"Andrew","Robertson",35000000,"1994-03-11","Scotland","df"},
	{12,21,"Konstantinos","Tsimikas",22000000,"1996-05-12","Greece","df"},
	{13,66,"Trent","Alexander-Arnold",70000000,"1998-10-07","England","df"},
	{14,22,"Calvin","Ramsay",4000000,"2003-07-31","Scotland","df"},
	{15,3,"Wataru","Endo",13000000,"1993-02-09","Japan","md"},
	{16,43,"Stefan","Bajcetic",11000000,"2004-10-22","Spain","md"},
	{17,8,"Dominik","Szoboszlai",75000000,"2000-10-25","Hungary","md"},
	{18,10,"Alexis","Mac Allister",65000000,"1998-12-24","Argentina","md"},
	{19,38,"Ryan","Gravenberch",35000000,"2002-05-16","Netherlands","md"},
	{20,17,"Curtis","Jones",30000000,"2001-01-31","England","md"},
	{21,6,"Thiago","Alcántara",10000000,"1991-04-11","Spain","md"},
	{22,19,"Harvey","Elliot",30000000,"2003-04-04","England","md"},
	{23,7,"Luiz","Diaz",75000000,"1997-01-13","Colombia","at"},
	{24,20,"Diogo","Jota",50000000,"1996-04-12","Portugal","at"},
	{25,50,"Ben","Doak",10000000,"2005-11-11","Scotland","at"},
	{26,9,"Darwin","Nunez",65000000,"1999-06-24","Uruguay","at"},
	{27,18,"Cody","Gakpo",50000000,"1999-05-07","Netherlands","at"},
	{28, 7, "Kenny", "Dalglish", 0, "1951-03-04", "Scotland", "at"},
    {29, 8, "Steven", "Gerrard", 0, "1980-05-30", "England", "md"},
    {30, 9, "Ian", "Rush", 0, "1961-10-20", "Wales", "at"},
    {31, 11, "Robbie", "Fowler", 0, "1975-04-09", "England", "at"},
    {32, 23, "Jamie", "Carragher", 0, "1978-01-28", "England", "df"},
    {33, 27, "Divock", "Origi", 0, "1995-04-18", "Belgium", "at"},
    {34, 21, "Lucas", "Leiva", 0, "1987-01-09", "Brazil", "md"},
    {35, 18, "Dirk", "Kuyt", 0, "1980-07-22", "Netherlands", "at"},
}
