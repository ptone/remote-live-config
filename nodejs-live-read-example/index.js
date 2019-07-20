var SelfReloadJSON = require('self-reload-json');

var config = new SelfReloadJSON('/tmp/demo.json');

config.on('updated' , (j) => console.log('config updated'))

function checkConfig() {
    console.log(config.msg);
  };
  
  function main() {
    setInterval(checkConfig, 1000);
  };
  
  main();
