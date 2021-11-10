module ex5_sitemap

go 1.17

require github.com/adrianprayoga/gophercises/ex4_link v0.0.0-20211003063809-fbc6f05fc886

require golang.org/x/net v0.0.0-20210929193557-e81a3d93ecf6 // indirect

replace (
	github.com/adrianprayoga/gophercises/ex4_link v0.0.0-20211003063809-fbc6f05fc886 => ../ex4_link
)