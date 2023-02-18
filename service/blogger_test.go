package service

import (
	"blog/server/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateBlogger(t *testing.T) *model.Blogger {
	blogger1 := &model.Blogger{
		Username: "testuser",
		Password: "testpassword",
		Nickname: "testuser",
		Sign: "testsign",
		Profile: "testprofile", 
		Img: "testimg",
	}
	err := bloggerService.Create(blogger1)
	require.NoError(t, err)
	return blogger1
}

func TestCreateBlogger(t *testing.T) {
	CreateBlogger(t)
}

func TestFindIdFirst(t *testing.T) {
	result, err := Service.FindIdFirst() 
	require.NoError(t, err)
	require.Equal(t, result.Id, 1)
}

func TestFindByName(t *testing.T) {
	blogger := CreateBlogger(t)
    result, err := Service.FindByName(blogger) 
    require.NoError(t, err)
    require.Equal(t, result.Username, blogger.Username)
    require.Equal(t, result.Password, blogger.Password)
    require.Equal(t, result.Nickname, blogger.Nickname)
    require.Equal(t, result.Sign, blogger.Sign)
    require.Equal(t, result.Profile, blogger.Profile)
    require.Equal(t, result.Img, blogger.Img)
}