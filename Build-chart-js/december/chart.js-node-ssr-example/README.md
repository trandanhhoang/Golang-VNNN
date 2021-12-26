# File image in folder output/\*.{svg,jpeg}

# Directory structure

```
.
├── api
│   ├── chart.js -> Call service buildChart
│   ├── index.js
│   └── schema.json -> Rule to handle request
├── index.js
└── services
    ├── color-factory.js -> color-data that added when user send without color field 
    └── index.js -> Contain main logic of buildChart function
```