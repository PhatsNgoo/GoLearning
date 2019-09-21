package main

import "errors"

type article struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

//Store database in memory
var articleList=[]article{
	{
		ID:      1,
		Title:   "First article",
		Content: "This is the first article on web site",
	},
	{
		ID:      2,
		Title:   "Second article",
		Content: "This is the second article on web site",
	},
}
func GetArticleByID(id int)(*article,error){
	for _,a:=range articleList{
		if a.ID == id {
			return  &a,nil
		}
	}
	return nil,errors.New("Article not found")
}
//Return a list of memory
func GetAllArticle()  []article{
	return articleList
}
