const Mustache = require("mustache");

const template = process.argv[2];
const data = JSON.parse(process.argv[3]);

process.stdout.write(Mustache.render(template, data));