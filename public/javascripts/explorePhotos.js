var app = angular.module('explorePhotos', ['ngRoute', 'ngResource']).run(function($http,$rootScope) {
  $rootScope.authenticated = false;
  $rootScope.current_user = 'Guest';
  
   $rootScope.signout = function(){
    $http.get('auth/signout');
    $rootScope.authenticated = false;
    $rootScope.current_user = 'Guest';
  };
});

app.config(function($routeProvider){
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

app.factory('photosService', function($resource){
	return $resource('/api/photos/:id');
});

app.controller('mainController', function($scope, $http, $rootScope, $location,photosService){
	
    $scope.photos=photosService.query();

    console.log($scope.photos);
});

app.controller('authController', function($scope, $http, $rootScope, $location){
	$scope.user = {username: '', password: ''};
	$scope.error_message = '';

	$scope.login = function(){
		$http.post('/auth/login',$scope.user).success(function(data){

        	if(data.state=='success'){
        		$rootScope.authenticated = true;
        		$rootScope.current_user=data.user.username;
        		$location.path('/');
        	}
        	else{
        		$scope.error_message=data.message;
        	}
        	
        });
	};

	$scope.register = function(){
		$scope.error_message = 'registeration request for ' + $scope.user.username;
         $http.post('auth/signup',$scope.user).success(function(data){
            if(data.state=='success'){
                $rootScope.authenticated = true;
        		$rootScope.current_user=data.user.username;
        		$location.path('/');	
               }
               else{
             $scope.error_message=data.message;  	
               }

         });

	};
});