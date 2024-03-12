module github.com/blogrenderer

go 1.22.0

replace github.com/blogposts => ../blogposts

require (
	github.com/approvals/go-approval-tests v0.0.0-20220530063708-32d5677069bd
	github.com/blogposts v0.0.0-00010101000000-000000000000
)

require github.com/gomarkdown/markdown v0.0.0-20231222211730-1d6d20845b47 // indirect
