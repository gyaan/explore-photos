var LocalStrategy   = require('passport-local').Strategy;
var bCrypt = require('bcrypt-nodejs');
//temporary data store
var users = {};
module.exports = function(passport){

    // Passport needs to be able to serialize and deserialize users to support persistent login sessions
    passport.serializeUser(function(user, done) {
        console.log('serializing user:',user.username);
        return done(null, user.username);
    });

    passport.deserializeUser(function(username, done) {

        return done(null,users[username]);

    });

    passport.use('login', new LocalStrategy({
        passReqToCallback : true
    },
    function(req, username, password, done) { 

        if(!users[username]){
          return done('user name not found', false);    
      }
            //is valid password
            if(!isValidPassword(users[username],password)){
             return done('user password is not valid',false) 
         }

             //successfull logged in 

             console.log("successfully signed in");
             return done(null,users[username])

         }
         ));

    passport.use('signup', new LocalStrategy({
            passReqToCallback : true // allows us to pass back the entire request to the callback
        },
        function(req, username, password, done) {

          //check if user is there or not 

          if(users[username]){
            return done(null,"username is already exit");
        }
           //add user to db

           users[username]={
            username:username,
            password:createHash(password)
        };

        return done(null,users[username]);

    })
    );

    var isValidPassword = function(user, password){
        return bCrypt.compareSync(password, user.password);
    };
    // Generates hash using bCrypt
    var createHash = function(password){
        return bCrypt.hashSync(password, bCrypt.genSaltSync(10), null);
    };

};