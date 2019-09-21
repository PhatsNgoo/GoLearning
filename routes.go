// routes.go

package main

func initializeRoutes() {
	//handler the index route
	router.GET("/",showIndexPage)

	//handle GET request at /article/view/some_article_id
	router.GET("/article/view/:articleID",getArticle)
}