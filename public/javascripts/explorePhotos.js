var app = angular.module('explorePhotos', ['ngRoute', 'ngResource','infinite-scroll']).run(function($http, $rootScope) {


  //first check if user is already loggedin or not 


  $rootScope.authenticated = false;

  $rootScope.current_user = 'Guest';


  $http.get('/api/islogin').success(function(data) {

    if (data.state == 'success') {
      $rootScope.authenticated = true;
      $rootScope.current_user = data.user.Username;
      $location.path('/');
    } else {
      $scope.error_message = data.message;
    }

  });

  $rootScope.signout = function() {
    $http.get('auth/signout');
    $rootScope.authenticated = false;
    $rootScope.current_user = 'Guest';
  };
});

app.config(function($routeProvider) {
  $routeProvider
  //the timeline display
    .when('/', {
      templateUrl: 'main.html',
      controller: 'mainController'
    })
    //the login display
    .when('/login', {
      templateUrl: 'login.html',
      controller: 'authController'
    })
    //the signup display
    .when('/signup', {
      templateUrl: 'register.html',
      controller: 'authController'
    });
});

//similarly we have to do for upvote and downvote

app.factory('photosService', function($resource) {
  return $resource('/api/photos/:id');
});

app.factory('votesService', function($resource) {
  return $resource('/api/votes/:id');
});


app.controller('mainController', function($scope, $http, $rootScope, $location, photosService) {

  $scope.photos = photosService.query();
  console.log($scope.photos);
});

app.controller('authController', function($scope, $http, $rootScope, $location) {
  $scope.user = {
    username: '',
    password: ''
  };
  $scope.error_message = '';

  $scope.login = function() {
    $http.post('/auth/login', $scope.user).success(function(data) {

      if (data.state == 'success') {
        $rootScope.authenticated = true;
        $rootScope.current_user = data.user.Username;
        $location.path('/');
      } else {
        $scope.error_message = data.message;
      }

    });
  };

  $scope.register = function() {
    $scope.error_message = 'registeration request for ' + $scope.user.username;
    $http.post('auth/signup', $scope.user).success(function(data) {
      if (data.state == 'success') {
        $rootScope.authenticated = true;
        console.log(data);
        console.log($rootScope.authenticated);
        $rootScope.current_user = data.user.Username;
        $location.path('/');
      } else {
        $scope.error_message = data.message;
      }

    });

  };
});

/* apart form angulr */

//We are using $(window).load here because we want to wait until the images are loaded  
$(window).load(function() {
  //for each description div...  
  $('div.image_description').each(function() {
    //...set the opacity to 0...  
    $(this).css('opacity', 0);
    //..set width same as the image...  
    $(this).css('width', $(this).siblings('img').width());
    //...get the parent (the wrapper) and set it's width same as the image width... '  
    $(this).parent().css('width', $(this).siblings('img').width());
    //...set the display to block  
    $(this).css('display', 'block');
  });

  $('div.image_wrapper').hover(function() {
    //when mouse hover over the wrapper div  
    //get it's children elements with class description '  
    //and show it using fadeTo  
    $(this).children('.image_description').stop().fadeTo(500, 0.8);
  }, function() {
    //when mouse out of the wrapper div  
    //use fadeTo to hide the div  
    $(this).children('.image_description').stop().fadeTo(500, 0);
  });

});