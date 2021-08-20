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
## pipelineHash
Project consists of two main solutions: create `ExecutePipeline()` function that executes `func job(in, out chan interface{})` functions like in unix-pipeline - each function's output serves as next one's input and calculate some sort of hash from the given data. The hash calculates using pipeline as follows (crc32, md5 - corresponding algorithms, data - some input variable):

1) `SingleHash()`: crc32(data) + "~" + crc32(md5(data))
2) `MultiHash()`: crc32(data + "0") + ... + crc32(data + "5")
3) `CombineResults()`: sort `MultiHash()` output and concatenate them using "_"

Features:

1) `crc32` takes 1s to execute, `md5` takes 10ms, but if it is being run more than one at a time, overheats lock for 1s
2) We have only 3 seconds to calculate everything

So, pipeline and hash functions have to be run in parallel to calculate everything in time.

Debug output for input values `[0, 1, 1, 2, 3, 5, 8]`:

```
func SingleHash(): data = 1
func SingleHash(): crc32(data) = 2212294583
func SingleHash(): md5(data) = c4ca4238a0b923820dcc509a6f75849b
func SingleHash(): crc32(md5(data)) = 709660146
func SingleHash(): result = 2212294583~709660146
func MultiHash(): data = 2212294583~709660146

...
func MultiHash(): result = 1696913515191343735512658979631549563179965036907783101867
func MultiHash(): th = 5 crc32(th+data)= 241521304
func MultiHash(): th = 4 crc32(th+data)= 424490584
func MultiHash(): th = 0 crc32(th+data)= 1173136728
func MultiHash(): th = 1 crc32(th+data)= 1388626328
func MultiHash(): th = 3 crc32(th+data)= 2090076184
func MultiHash(): th = 2 crc32(th+data)= 1807510744
func MultiHash(): result = 1173136728138862632818075107442090076184424490584241521304
func CombineResult(): result = 1173136728138862632818075107442090076184424490584241521304_1696913515191343735512658979631549563179965036907783101867_27225454331033649287118297354036464389062965355426795162684_29568666068035183841425683795340791879727309630931025356555_3994492081516972096677631278379039212655368881548151736_4958044192186797981418233587017209679042592862002427381542_4958044192186797981418233587017209679042592862002427381542
```

## integralCacl
Calculate the integral of the hard-coded function (it is being passed to `calculateIntegral()` function as argument). Takes the section to integrate on as input, shows the calculated value on the screen. Uses Simpson's method for definite integrals along with goroutines, which amount corresponds to CPU cores. Usage example:
```
$ go run main.go
Integral calculation using goroutines, enter the length: 1 10000
Integral of f(x) from 1 to 10000 equals 9520.928265989416
```
