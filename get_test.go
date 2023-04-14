package structkit

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Post struct {
	Title    string
	Creator  User
	Editor   *User
	Comments []Comment
}

type User struct {
	Name string

	unexportedField bool
}

type Comment struct {
	Content string
}

type GetTestSuite struct {
	suite.Suite
}

func (test *GetTestSuite) TestGetNil() {
	buildCache := Post{Title: "Post Title", Creator: User{Name: "John Doe", unexportedField: false}, Editor: &User{Name: "Jane Doe"}}
	// Found value cp embedded struct
	result := Get(buildCache, "Editor.Name")
	test.Equal(result, buildCache.Editor.Name)

	payload := Post{Title: "Post Title", Creator: User{Name: "John Doe"}, Editor: nil}
	result = Get(payload, "Editor.Name")
	test.Equal(result, nil)

	nilPayload := new(Post)
	nilPayload = nil
	result = Get(nilPayload, "Editor.Name")
	test.Equal(result, nil)

	result = Get("test", "Editor.Name")
	test.Equal(result, nil)
}

func (test *GetTestSuite) TestGetSlice() {
	payload := Post{Title: "Post Title", Comments: []Comment{{Content: "Hello"}}}
	result := Get(payload, "Comments.[0].Content")
	test.Equal(result, "Hello")

	result = Get(payload, "Comments.[0].Content")
	test.Equal("Hello", result)

	result = Get(payload, "Comments.[a].Content")
	test.Equal(result, nil)

	result = Get(payload, "Comments.[a.Content")
	test.Equal(result, nil)

	result = Get(payload, "Comments.a.Content")
	test.Equal(result, nil)

	result = Get(payload, "Comments.[10].Content")
	test.Equal(result, nil)
}

func (test *GetTestSuite) TestGet() {
	payload := Post{Title: "Post Title", Creator: User{Name: "John Doe"}, Editor: &User{Name: "Jane Doe"}}

	// Invalide payload
	result := Get(nil, "Title")
	test.Equal(result, nil)

	// Found value
	result = Get(payload, "Title")
	test.Equal(result, payload.Title)

	result = Get(payload, "Title")
	test.Equal(result, payload.Title)

	// Found value cp embedded struct
	result = Get(payload, "Creator.Name")
	test.Equal(result, payload.Creator.Name)

	result = Get(payload, "Creator.Name")
	test.Equal(result, payload.Creator.Name)

	// Not found value
	result = Get(payload, "unexportedField")
	test.Equal(result, nil)

	// Found value with pointer of payload
	result = Get(&payload, "Title")
	test.Equal(result, payload.Title)

	// Found value cp embedded struct
	result = Get(payload, "Creator.Name")
	test.Equal(result, payload.Creator.Name)

	// Found value cp embedded struct
	result = Get(payload, "Editor.Name")
	test.Equal(result, payload.Editor.Name)
}

func TestGetTestSuite(t *testing.T) {
	suite.Run(t, new(GetTestSuite))
}

func BenchmarkGet(b *testing.B) {
	payload := Post{Title: "Post Title", Creator: User{Name: "John Doe"}}

	for i := 0; i < b.N; i++ {
		Get(payload, "Creator")
	}
}

func BenchmarkGetEmbeddedValue(b *testing.B) {
	payload := Post{Title: "Post Title", Creator: User{Name: "John Doe"}}

	for i := 0; i < b.N; i++ {
		Get(payload, "Creator.Name")
	}
}

func BenchmarkGetEmbeddedSliceValue(b *testing.B) {
	payload := Post{Title: "Post Title", Comments: []Comment{{Content: "Hello"}}}

	for i := 0; i < b.N; i++ {
		Get(payload, "Comments.[0].Content")
	}
}
