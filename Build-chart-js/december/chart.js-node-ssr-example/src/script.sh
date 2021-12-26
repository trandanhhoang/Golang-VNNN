curl -H "Content-Type: application/json" -X POST -d'{
  "type": "bar",
  "data": {
    "labels": [2012, 2013, 2014, 2015, 2016],
    "datasets": [
      {
        "label": "Users",
        "data": [120, 60, 50, 180, 120]
      }
    ]
  }
}' http://localhost:3000/api/chart/demo
