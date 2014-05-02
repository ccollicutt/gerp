#gerp

Half the time I try to run grep I end up typing gerp. I thought a good learning experience with golang would be to write a `grep`-like command line application that does the kind of grepping I usually do, which is looking for something in all the files in a directory, or also all the directories underneath as well (ie. recursive).

[![Gobuild Download](http://gobuild.io/badge/github.com/curtisgithub/gerp/download.png)](http://gobuild.io/github.com/curtisgithub/gerp)


##Example

This will check all the files in the test directory to see if they contain the pattern "hi".

```bash
curtis$ ./gerp hi test/
test/test_grep.txt: hi
```

##Issues

* The pattern "h" won't work b/c the cli package thinks it's asking for help
  * Eg. `gerp h somefile.txt` doesn't work
* Can't search a single file yet (doesn't know it's being pointed to a single file instead of a directory)
* Doesn't know a file is a binary file (feature I'd like to add)