var express = require('express');
var router = express.Router();


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
	var photos=[
	{
		'id':'1',
		'url':'https://c2.staticflickr.com/8/7149/13853152865_ae866a8ea3_z.jpg'
	},
	{
		'id':'2',
		'url':'https://c2.staticflickr.com/8/7149/13853152865_ae866a8ea3_z.jpg'
	}];
	res.send(photos);
})

.post(function(req,res){
	res.send('to do this method is not defined')
})

router.route('/photos/:id')

.get(function(req,res){

	var photo={
		'id':'1',
		'url':'https://c2.staticflickr.com/8/7149/13853152865_ae866a8ea3_z.jpg'
	};
	res.send({state: 'success', 'photo':photo});
})

module.exports = router;
