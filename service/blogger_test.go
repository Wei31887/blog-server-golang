package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateBlogger(t *testing.T) *Blogger {
	blogger1 := &Blogger{
		Username: "testuser",
		Password: "testpassword",
		Nickname: "testuser",
		Sign: "testsign",
		Profile: "testprofile", 
		Img: "testimg",
	}
	err := blogger1.Create()
	require.NoError(t, err)
	return blogger1
}

func TestCreateBlogger(t *testing.T) {
	CreateBlogger(t)
}

func TestFindIdFirst(t *testing.T) {
	var blogger *Blogger
	result, err := blogger.FindIdFirst() 
	require.NoError(t, err)
	require.Equal(t, result.Id, 1)
}

func TestFindByName(t *testing.T) {
	blogger := CreateBlogger(t)
    result, err := blogger.FindByName() 
    require.NoError(t, err)
    require.Equal(t, result.Username, blogger.Username)
    require.Equal(t, result.Password, blogger.Password)
    require.Equal(t, result.Nickname, blogger.Nickname)
    require.Equal(t, result.Sign, blogger.Sign)
    require.Equal(t, result.Profile, blogger.Profile)
    require.Equal(t, result.Img, blogger.Img)
}