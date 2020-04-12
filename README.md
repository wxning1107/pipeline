# Pipeline

Pipeline is a simple pipeline written in Go. It builds a parallel processing pipeline in Go language to realize an external sort pipeline.

## Introduction

Pipeline reads the data source from the file(stdin/stdout). The reading is in blocks. Using internal sorting to sort each block. And then merge them two by two.

Each merged node is sorted for big data. Then merging them two by two. And finally it will be written to the file.

Finally, the final version provides external services for external sort.

## Document introduction
[pipeline introduction](http://wxning.com/2020/04/02/pipeline/)

