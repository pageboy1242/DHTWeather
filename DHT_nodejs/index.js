const sensor = require("node-dht-sensor");
const axios = require('axios')
const dateTime = require('node-datetime');
const cron = require('node-cron');

function readSensor() {
	console.log("Reading sensor");
	sensor.read(22, 4, function(err, temperature, humidity) {
		if(!err) {
			console.log('temp: %f C, humidity: %f %',
			 Number.parseFloat(temperature).toFixed(2),
			  Number.parseFloat(humidity).toFixed(2));
			
			var dt = dateTime.create();
			var formatted = dt.format('Y-m-d H:M:S');
			console.log(formatted);

			axios.post('http://10.0.0.6:9090/', {
					ID: 0,
					NodeId: 1,
					ReadingDate: formatted,
					Temperature: temperature,
					Humidity: humidity
			})
			.then((res) => {
					console.log(`statusCode: ${res.statusCode}`);
					return true;
					//console.log(res)
			})
			.catch((error) => {
					console.error(error);
					return false;
			});
		}
		else {
			console.log(err);
			return false;
		}
	});
}
var success = true;

	//console.log("Ready to sleep");
	//delay();
	
readSensor();

	
console.log("Sensor read complete");

	//delay();
//};
//async function delay () {
//  console.log('started');
//  await sleep(5000);
//  console.log('finished');
//}