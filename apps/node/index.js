/**
 * Created by Mr.Dung on 13/06/2017.
 */

var Excel = require('exceljs');
const EXCEL_FILE = "test2.xlsx";
const LIMIT_ROW = 1000000;
const LABEL = {
  time: "Time:",
};

function read() {
  console.time("Time:");
  var workbook = new Excel.Workbook();
  workbook.xlsx.readFile(EXCEL_FILE).then(function () {
    var worksheet = workbook.getWorksheet(1);
    var total = 0;
    worksheet.eachRow({
      inclideEmpty: true
    }, function(row, index) {
      var value =  row.values[2];
      console.log(value);
      if (!isNaN(value)) {
        total += value;
      }
    });
    console.log("Total: ", total);
    console.timeEnd("Time:")
  }).catch(function (err) {
    console.log(err);
  });
}

function write() {

}

// exec function
switch (process.argv[2]) {
  case "read":
    read();
    break;
  case "write":
    write();
    break;
  default:
    read();
}
