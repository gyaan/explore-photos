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
          console.log('user name not found with username'+username);
          return done(null, false);    
      }
            //is valid password
            if(!isValidPassword(users[username],password)){
               console.log('invalid username and password');
               return done(null,false) 
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
            console.log("user name already exist")
            return done(null,false);
        }
           //add user to db

           users[username]={
            username:username,
            password:createHash(password)
        };      

          console.log(users[username].username + 'registerd successfully');
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