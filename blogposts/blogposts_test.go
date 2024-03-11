package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/batmanonwheels/blogposts"
)

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)
	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		got  blogposts.Post
		want blogposts.Post
	}{
		{
			got: posts[0],
			want: blogposts.Post{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}, Body: `Hello
			World`}},
		{
			got: posts[1],
			want: blogposts.Post{Title: "Post 2", Description: "Description 2", Tags: []string{"rust", "borrow-checker"}, Body: `B
			L
			M`}},
	}

	t.Run("parses the file", func(t *testing.T) {

		for _, c := range cases {
			assertPost(t, c.got, c.want)
		}
	})

}
