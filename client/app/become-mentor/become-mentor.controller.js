'use strict';

angular.module('mentorMeApp')
.controller('BecomeMentorCtrl', function ($scope, $http, $state) {
  var init = function() {
    kik.getUser(function (user) {
      if ( !user ) {
        // user denied access to their information
        $state.go('no-permissions');
        return; // not sure if go above will pre-empt this?
      } else {
        // check if they already have an account
        console.log("username before request: " + user.username);
        console.log("url: /api/mentors/byusername/" + user.username);
        $http.get('/api/mentors/byusername/' + user.username).then(
          function successCallback(response) {
            $scope.mentor = response.data;
            console.log("mentor: ");
            console.log("mentor.username: " + $scope.mentor.username);
            console.log($scope.mentor);
            $scope.mentor.tags = $scope.mentor.tags.join();

            $scope.mentor.alreadyExists = true;
            if ($scope.mentor === {}) {
              $scope.mentor.alreadyExists = false;
              $scope.mentor.username = user.username; // 'string'
              $scope.mentor.fullName = user.fullName;
              $scope.mentor.pic = user.pic;
              $scope.mentor.thumbnail = user.thumbnail;
            }
          }, function errorCallback(response) {
            // TODO: do stuff
            console.log("failed");
          }
        );
      }
    });
  };

  $scope.addMentor = function() {
    if ($scope.username) {
      if($scope.mentor === {}) {
        return;
      }
      $scope.mentor.tags = $scope.mentor.tags.split(",");
      $scope.mentor.username = $scope.username;
      $scope.mentor.fullName = $scope.fullName;
      $scope.mentor.pic = $scope.pic;
      $scope.mentor.thumbnail = $scope.thumbnail;
      $scope.mentor.active = true;
      $http.post('/api/mentors', $scope.mentor).then(
        function successCallback(response) {
          // inform them of success
        }, function errorCallback(response) {
          // do something
        }
      );
    } else {
      // They shouldn't have been able to do this
    }
  };

  $scope.updateMentor = function() {
    $http.put('/api/mentors/' + $scope.mentor._id, $scope.mentor).then(
      function successCallback(response) {
        // inform them of success
      }, function errorCallback(response) {
        // do something
      }
    );
  }

  init();
});
