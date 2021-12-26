# how to test

- cd testAPI
- test all API
  - go test main_test.go -v
- test each API
  - go test main_test.go -run name_API -v
  - exam: go test main_test.go -run PostOrder -v
    - Will add a line have data of "labels" to cancel order later.
  - exam: go test main_test.go -run Cancel -v
    - Will read the first line, delete it and use this data to cancel order
