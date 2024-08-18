# gossg

## Goals

1. Write your templates in go templ
2. Write your content in markdown
3. Generate static HTML
4. Serve HTML with go

## Dependencies

- go 1.22.5
- air
- templ
- make

## Commands

### Develoment mode

- `make templ`: start templ file generation in watch mode
- `make dev`: start dev server with hot reload
  [http://localhost:7331](http://localhost:7331)

### SSR (Server Side Rendering) mode

- `make build`: build go binary
- `make ssr`: run go server

### SSG (Static Site Generation) mode

- `make gen`: generate static HTML files
- `make ssg`: start prod server [http://localhost:8080](http://localhost:8080)
  (SSG mode)

## Themes

You can add a new theme under `themes/mytheme/` (replace `mytheme` with the name
of the theme), whith the following structure:

```
themes
└── mytheme
    ├── static
    └── templates
```

Put static files like `css`, `js`, images etc. under `themes/mytheme/static`.

Put `templ` components under `themes/mytheme/templates`.

Every theme should have the following `templ` components exported from the root
directory of the theme, as you can see in the `pages/page.go`:

```go
type (
	NotFoundPage func() templ.Component
	IndexPage    func(files []*pkg.FileData) templ.Component
	PostPage     func(post *pkg.FileData) templ.Component
	TagPage      func(tag string, files []*pkg.FileData) templ.Component
	CategoryPage func(category string, files []*pkg.FileData) templ.Component
)
```

### Switching themes

When running the following command, the files from the theme folder gets copied
over the appropriate directories:

`make theme name=mytheme`
