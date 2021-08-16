# GoResearch
My research of Golang. Includes study projects and maybe some useless stuff, too (to be honest, both at the same time)
# Projects list
## dirTree
Print the directory tree of the given directory recursively. If the flag `-f` is given, prints files and their sizes in bytes. To launch, use:

`go run main.go <directory> [-f]`

Example output (`testdata/` directory in the rep):

```
├───project
│       ├───file.txt (19b)
│       └───gopher.png (70372b)
├───static
│       ├───a_lorem
│       │       ├───dolor.txt (empty)
│       │       ├───gopher.png (70372b)
│       │       └───ipsum
│       │               └───gopher.png (70372b)
│       ├───css
│       │       └───body.css (28b)
│       ├───empty.txt (empty)
│       ├───html
│       │       └───index.html (57b)
│       ├───js
│       │       └───site.js (10b)
│       └───z_lorem
│               ├───dolor.txt (empty)
│               ├───gopher.png (70372b)
│               └───ipsum
│                       └───gopher.png (70372b)
├───zline
│       ├───empty.txt (empty)
│       └───lorem
│               ├───dolor.txt (empty)
│               ├───gopher.png (70372b)
│               └───ipsum
│                       └───gopher.png (70372b)
└───zzfile.txt (empty)
```
