db.createCollection("products");
db.products.insertMany([
  {
    _id: new ObjectId("56955ca46063c5600627f393"),
    name: "shirt",
    size: "x",
    quantity: 10,
    length: "50cm",
  },
  {
    _id: new ObjectId("56955ca46063c5600627f394"),
    name: "shirt",
    size: "x",
    quantity: 10,
    length: "30cm",
  },
  {
    _id: new ObjectId("56955ca46063c5600627f395"),
    name: "pant",
    size: "xl",
    quantity: 15,
    length: "40cm",
  },
  {
    _id: new ObjectId("56955ca46063c5600627f396"),

    name: "pant",
    size: "xxl",
    quantity: 15,
    length: "50cm",
  },
  {
    _id: new ObjectId("56955ca46063c5600627f397"),
    name: "pant",
    size: "x",
    quantity: 15,
    length: "90cm",
  },
]);
