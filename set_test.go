package structkit

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type PostSet struct {
	Title    string
	TitlePtr *string
	Creator  UserSet
	Editor   *UserSet
	Comments []CommentSet
}

type UserSet struct {
	Name            string
	unexportedField bool
}

type CommentSet struct {
	Content string
}

type SetTestSuite struct {
	suite.Suite
}

func (test *SetTestSuite) TestSetErr() {
	// nil
	err := Set(nil, "Title", "Post Title")
	test.Error(err)

	// nil pointer
	nilPayload := new(PostSet)
	nilPayload = nil
	err = Set(nilPayload, "Title", "Post Title")
	test.Error(err)

	// not a pointer
	err = Set(PostSet{}, "Title", "Post Title")
	test.Error(err)

	// not a slice or struct
	err = Set(new(int), "Title", "Post Title")
	test.Error(err)

	buildCache := PostSet{Creator: UserSet{Name: "John Doe", unexportedField: false}}

	err = Set(&buildCache, "unknown", "Post Title")
	test.Error(err)

	err = Set(&buildCache, "TitlePtr", "Post Title")
	test.Error(err)
}

func (test *SetTestSuite) TestSet() {
	buildCache := PostSet{}

	expected := "Post Title"

	err := Set(&buildCache, "Title", "Post Title")
	test.NoError(err)
	test.Equal(expected, buildCache.Title)

	titlePtr := new(string)
	*titlePtr = "Post Title"

	err = Set(&buildCache, "TitlePtr", titlePtr)
	test.NoError(err)
	test.Equal(titlePtr, buildCache.TitlePtr)

	err = Set(&buildCache, "Creator.Name", "John Doe")
	test.NoError(err)
	test.Equal("John Doe", buildCache.Creator.Name)
}

func TestSetTestSuite(t *testing.T) {
	suite.Run(t, new(SetTestSuite))
}
