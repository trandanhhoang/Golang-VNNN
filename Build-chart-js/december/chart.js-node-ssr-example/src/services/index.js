import {
  SvgCanvas,
  Rect2D,
  SvgCanvas2DGradient,
} from "red-agate-svg-canvas/modules";

import * as ChartJs from "chart.js";

const fs = require("fs");
import sharp from "sharp";

// Get the global scope.
// If running on a node, "g" points to a "global" object.
// When running on the browser, "g" points to the "window" object.
const g = Function("return this")();

let counter = 1;

function getRandomInt(max) {
  return Math.floor(Math.random() * max);
}

export const buildChart = async (chart) => {
  // console.log(JSON.stringify(chart, null, 4));

  const datasets = chart.data.datasets.map((dataset, item) => {
    if (dataset.backgroundColor !== undefined) {
      return dataset;
    }

    const data = dataset.data;

    const mainColor =
      "rgba(" +
      getRandomInt(255) +
      "," +
      getRandomInt(255) +
      "," +
      getRandomInt(255);
    const colorBackground = mainColor + "," + "0.2)";
    const colorBorder = mainColor + "," + "1)";

    return {
      label: dataset.label,
      data: dataset.data,
      backgroundColor: data.map(() => colorBackground),
      borderColor: data.map(() => colorBorder),
      borderWidth: 1,
    };
  });

  // SvgCanvas has a "CanvasRenderingContext2D"-compatible interface.
  const ctx = new SvgCanvas();

  // SvgCanvas lacks the canvas property.
  ctx.canvas = {
    width: 800,
    height: 400,
    style: {
      width: "800px",
      height: "400px",
    },
  };

  // SvgCanvas does not have font glyph information,
  // so manually set the ratio of (font height / font width).
  ctx.fontHeightRatio = 2;

  // Chart.js needs a "HTMLCanvasElement"-like interface that has "getContext()" method.
  // "getContext()" should returns a "CanvasRenderingContext2D"-compatible interface.
  const el = { getContext: () => ctx };

  //Adding options

  let options = {
    scales: {
      yAxes: [
        {
          ticks: {
            beginAtZero: true,
          },
        },
      ],
    },
    // If "devicePixelRatio" is not set, Chart.js get the devicePixelRatio from "window" object.
    // node.js environment has no window object.
    devicePixelRatio: 1,
    // Disable animations.
    animation: false,
    // events: [],
    responsive: false,
  };

  if (chart.options !== undefined) {
    options = {
      ...chart.options,
      ...options,
    };
  }

  const chartDetail = {
    type: chart.type,
    data: {
      labels: chart.data.labels,
      datasets,
    },
    options,
  };
  console.log(JSON.stringify(chartDetail, null, 4));

  // Chart.js needs the "CanvasGradient" in the global scope.
  const savedGradient = g.CanvasGradient;
  g.CanvasGradient = SvgCanvas2DGradient;
  try {
    const chart = new ChartJs.Chart(el, chartDetail);
  } finally {
    if (savedGradient) {
      g.CanvasGradient = savedGradient;
    }
  }

  // Render as SVG.
  const svgString = ctx.render(new Rect2D(0, 0, 800, 400), "px");

  //write to file in path "output/output.svg"
  fs.writeFileSync("output/output.svg", svgString);
  // convert svg to jpeg
  const path = await convertSvgToJpeg();

  console.log("path ", path);

  return Promise.resolve(path);
};

const convertSvgToJpeg = async () => {
  const imgInput = "output/output.svg";
  const imgOutput = "output/output" + counter + ".jpeg";

  try {
    await sharp(imgInput)
      .jpeg()
      .toFile(imgOutput)
      .then((info) => {
        console.log(info);
        counter++;
      })
      .catch((err) => {
        console.log(err);
        return err;
      });
  } catch (err) {
    return err;
  }

  return Promise.resolve(imgOutput);
};
