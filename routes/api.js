var express = require('express');
var router = express.Router();
var http = require('http');
var querystring = require('querystring');
/*
//Used for routes that must be authenticated.
function isAuthenticated (req, res, next) {

	if (req.isAuthenticated()){
		return next();
	}

    // if the user is not authenticated then redirect him to the login page
    return res.redirect('/#login');
};

//Register the authentication middleware
router.use('/photos', isAuthenticated);
*/
router.route('/photos')

/* GET photos listing. */
.get(function(req, res) {

	var pageNumber = req.query.current_page
		//get the user details using apis 
	var options = {
		host: "localhost",
		port: "1334",
		path: "/photos?current_page=" + pageNumber,
		method: "GET",
	}

	var request = http.request(options, function(ress) {

		console.log('getting the images1');
		var buffer = "",
			output;
		ress.on("data", function(chunk) {
			buffer += chunk;
		});

		ress.on("end", function(err) {
			output = JSON.parse(buffer);

			if (output.Status == "success") {
				// console.log(output);
				res.send(output);
			} else {
				res.send("some problem while getting the images");
			}
		});
	});

	request.on('error', function(e) {
		console.log('problem with request: ' + e.message);
	});
	// write data to request body
	// request.write(options);
	request.end();

})

.post(function(req, res) {
	res.send('to do this method is not defined')
})

router.route('/photos/:id')

.get(function(req, res) {

	var photo = {
		'id': '1',
		'url': 'https://c2.staticflickr.com/8/7149/13853152865_ae866a8ea3_z.jpg'
	};
	res.send({
		state: 'success',
		'photo': photo
	});
})

//define Content-Type:application/json while giving the request 
router.route('/votes')

.post(function(req, res) {
	var postData = querystring.stringify({
		'user_id': req.body.user_id,
		'photo_id': req.body.photo_id,
		'value': req.body.vote
	});

	var options = {
		hostname: 'localhost',
		port: 1334,
		path: '/votes',
		method: 'POST',
		headers: {
			'Content-Type': 'application/x-www-form-urlencoded',
			'Content-Length': postData.length
		}
	};

     console.log(options);

	var req = http.request(options, function(ress) {

		var output, buffer = '';
		ress.on('data', function(chunk) {
			buffer += chunk;
		});

		ress.on("end", function(err) {
			output = JSON.parse(buffer);

			if (output.Status == "success") {
				console.log(output);
				res.send(output.Vote);
			} else {
				res.send("some problem while voting a image");
			}
		});
	});

	req.on('error', function(e) {
		console.log('problem with request: ' + e.message);
	});

	// write data to request body
	req.write(postData);
	req.end();
	// res.send("some problem while updating votes count");
})


router.route('/islogin')

.get(function(req, res) {

	if (req.user) {
		res.send({
			state: 'success',
			user: req.user
		});
	} else {
		res.send({
			state: 'unsuccess',
			message: 'user is not logged in'
		});
	}
})

module.exports = router;