'use strict';

angular.module('mentorMeApp')
  .controller('NavbarCtrl', function ($scope, $location) {

    if ($location.path() == '/') {
      $scope.location = 'Home'
    } else {
      var path_array = $location.path().split('/');
      $scope.location = path_array[path_array.length-1];
      $scope.location = $scope.location.replace("-", " ");

      // http://stackoverflow.com/questions/196972/convert-string-to-title-case-with-javascript/196991#196991
      $scope.location = $scope.location.replace(/\w\S*/g, function(txt){
        return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
      });
    }


    $scope.isActive = function(route) {
      return route === $location.path();
    };
  });