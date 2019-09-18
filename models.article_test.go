package main

import "testing"

//Test func that get all the article
func TestGetAllArticle(t *testing.T) {
	alist:=GetAllArticle()

	//Check that the length of the returned list is same as the original list
	if len(alist)!=len(articleList) {
		t.Fail()
	}
	
	//Check identity of each member
	for i,v:=range alist {
		if v.Content!=articleList[i].Content || v.ID!=articleList[i].ID || v.Title!=articleList[i].Title {
			t.Fail()
			break
		}
	}
}
