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
func SingleHash(): data = 0
func SingleHash(): crc32(data) = 4108050209
func SingleHash(): md5(data) = cfcd208495d565ef66e7dff9f98764da
func SingleHash(): crc32(md5(data)) = 502633748
func SingleHash(): result = 4108050209~502633748
func MultiHash(): data = 4108050209~502633748
func SingleHash(): data = 2
func SingleHash(): crc32(data) = 450215437
func SingleHash(): md5(data) = c81e728d9d4c2f636f067f89cc14862c
func SingleHash(): crc32(md5(data)) = 1933333237
func SingleHash(): result = 450215437~1933333237
func MultiHash(): data = 450215437~1933333237
func SingleHash(): data = 1
func SingleHash(): crc32(data) = 2212294583
func SingleHash(): md5(data) = c4ca4238a0b923820dcc509a6f75849b
func SingleHash(): crc32(md5(data)) = 709660146
func SingleHash(): result = 2212294583~709660146
func MultiHash(): data = 2212294583~709660146
func SingleHash(): data = 5
func SingleHash(): crc32(data) = 2226203566
func SingleHash(): md5(data) = e4da3b7fbbce2345d7772b0674a318d5
func SingleHash(): crc32(md5(data)) = 3690458478
func SingleHash(): result = 2226203566~3690458478
func MultiHash(): data = 2226203566~3690458478
func SingleHash(): data = 3
func SingleHash(): crc32(data) = 1842515611
func SingleHash(): md5(data) = eccbc87e4b5ce2fe28308fd9f2a7baf3
func SingleHash(): crc32(md5(data)) = 1684880638
func SingleHash(): result = 1842515611~1684880638
func MultiHash(): data = 1842515611~1684880638
func SingleHash(): data = 8
func SingleHash(): crc32(data) = 4194326291
func SingleHash(): md5(data) = c9f0f895fb98ab9159f51fd0297e236d
func SingleHash(): crc32(md5(data)) = 2004971030
func SingleHash(): result = 4194326291~2004971030
func MultiHash(): data = 4194326291~2004971030
func MultiHash(): th = 0 crc32(th+data)= 495804419
func MultiHash(): th = 5 crc32(th+data)= 2427381542
func MultiHash(): th = 2 crc32(th+data)= 4182335870
func MultiHash(): th = 4 crc32(th+data)= 259286200
func MultiHash(): th = 3 crc32(th+data)= 1720967904
func MultiHash(): th = 1 crc32(th+data)= 2186797981
func MultiHash(): result = 4958044192186797981418233587017209679042592862002427381542
func MultiHash(): th = 5 crc32(th+data)= 1025356555
func MultiHash(): th = 2 crc32(th+data)= 1425683795
func MultiHash(): th = 0 crc32(th+data)= 2956866606
func MultiHash(): th = 1 crc32(th+data)= 803518384
func MultiHash(): th = 4 crc32(th+data)= 2730963093
func MultiHash(): th = 3 crc32(th+data)= 3407918797
func MultiHash(): result = 29568666068035183841425683795340791879727309630931025356555
func MultiHash(): th = 5 crc32(th+data)= 795162684
func MultiHash(): th = 2 crc32(th+data)= 1182973540
func MultiHash(): th = 0 crc32(th+data)= 2722545433
func MultiHash(): th = 4 crc32(th+data)= 2965355426
func MultiHash(): th = 3 crc32(th+data)= 3646438906
func MultiHash(): th = 1 crc32(th+data)= 1033649287
func MultiHash(): result = 27225454331033649287118297354036464389062965355426795162684
func MultiHash(): th = 0 crc32(th+data)= 495804419
func MultiHash(): th = 3 crc32(th+data)= 1720967904
func MultiHash(): th = 5 crc32(th+data)= 2427381542
func MultiHash(): th = 4 crc32(th+data)= 259286200
func MultiHash(): th = 2 crc32(th+data)= 4182335870
func MultiHash(): th = 1 crc32(th+data)= 2186797981
func MultiHash(): result = 4958044192186797981418233587017209679042592862002427381542
func MultiHash(): th = 4 crc32(th+data)= 1265536888
func MultiHash(): th = 0 crc32(th+data)= 399449208
func MultiHash(): th = 3 crc32(th+data)= 783790392
func MultiHash(): th = 1 crc32(th+data)= 15169720
func MultiHash(): th = 5 crc32(th+data)= 1548151736
func MultiHash(): th = 2 crc32(th+data)= 966776312
func MultiHash(): result = 3994492081516972096677631278379039212655368881548151736
func MultiHash(): th = 5 crc32(th+data)= 783101867
func MultiHash(): th = 0 crc32(th+data)= 1696913515
func MultiHash(): th = 2 crc32(th+data)= 1265897963
func MultiHash(): th = 4 crc32(th+data)= 965036907
func MultiHash(): th = 3 crc32(th+data)= 1549563179
func MultiHash(): th = 1 crc32(th+data)= 1913437355
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
