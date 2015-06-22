var express = require('express');
var router = express.Router();


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

router.route('/photos')

/* GET photos listing. */
.get(function(req, res) {
	res.send('to do return all photos');
})

.post(function(req,res){
	res.send('to do this method is not defined')
})

router.route('/photos/:id')

.get(function(req,res){
	res.send('send individual photo details'+req.params.id)
})

module.exports = router;
