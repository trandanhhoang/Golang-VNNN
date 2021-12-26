export const register = (registed, req) => {
  try {
    if (registed === "/chart/demo") {
      validate(req.data.labels, ["string", "number"]);
      for (const dataset of req.data.datasets) {
        validate(dataset.data, ["string", "number"]);
      }
    }
  } catch (e) {
    console.log(e);
    throw "Client send invalid value";
  }
};

export const validate = (req, types) => {
  for (const label of req) {
    if (types.filter((typ) => typeof label === typ).length > 0) {
      continue;
    } else {
      throw "Validate fail";
    }
  }
};
