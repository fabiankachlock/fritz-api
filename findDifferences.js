const fs = require("node:fs");

const orig = fs.readFileSync("response.json", "utf8").toString();
const other = fs.readFileSync("my.json", "utf8").toString();

const compareArray = (path, orig, other) => {
  for (let i = 0; i < orig.length; i++) {
    if (typeof orig[i] === "object" && "UID" in orig[i]) {
      const origItemID = orig[i].UID;
      const otherItem = other.find((item) => item.UID === origItemID);
      compare([...path, origItemID], orig[i], otherItem);
    } else if (typeof orig[i] === "object" && other[i]) {
      compare([...path, i.toString()], orig[i], other[i]);
    } else if (orig[i] !== other[i]) {
      console.log(`${[...path, i.toString()].join(".")} is different`);
    }
  }
};

const compare = (path, orig, other) => {
  for (const key in orig) {
    if (!other.hasOwnProperty(key)) {
      console.log(`${[...path, key].join(".")} is missing`);
    } else if (Array.isArray(orig[key])) {
      compareArray([...path, key], orig[key], other[key]);
    } else if (typeof orig[key] === "object") {
      compare([...path, key], orig[key], other[key]);
    } else {
      if (orig[key] !== other[key]) {
        console.log(`${[...path, key].join(".")} is different`);
      }
    }
  }
};

compare([], JSON.parse(orig), JSON.parse(other));
