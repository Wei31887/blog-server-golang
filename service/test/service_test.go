package test

import (
	"blog/server/service"
	"fmt"
	"testing"
)

var blog service.Blog
// var blogType service.BlogType
// var blogger service.Blogger

func TestBlog(t *testing.T) {
	fmt.Println("Testing query of blog ...")
	t.Run("Test find next blog", TestFindNextBlog)
}

func TestFindNextBlog(t *testing.T) {
	blog.Id = 2
	resBlog, err := blog.FindNextBlog()
	if err != nil {
		t.Error("Can't find the next blog", err)
	}
	fmt.Println(resBlog)

}