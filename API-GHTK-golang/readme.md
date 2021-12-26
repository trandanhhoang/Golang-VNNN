# last version

- Make Models better by use omittype and \*struct ?
- Make more test
- What is mapstructure ? Learn it.
- Do I need xbApiRespond ?

# version 3

- Move data of order from a json file

# version 2

- Add test function, use can read how to use them below.
- Learning how to write,read file(main_test.go). Delete file.(ultils/Popline.go)
- Models have no change. I will make them better soon

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

# version 1

- Because version 1 so I'm quick build without file .env to hide "Token".
- Mapstructure have not been use
- I think I will use "github.com/stretchr/testify/assert" soon. It help me make more test-case to test the API.
- After you check if I do correct, I will add them soon.
