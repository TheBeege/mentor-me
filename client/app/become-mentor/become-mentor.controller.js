'use strict';

angular.module('mentorMeApp')
.controller('BecomeMentorCtrl', function ($scope, $http, $state) {
  $scope.alerts = new array();

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
            $scope.alerts.push({
              type: 'danger',
              msg: 'Failed to grab your info. Try reloading?'
            })
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
          $scope.alerts.push({
            type: 'success',
            msg: 'Successfully set you up as a mentor!'
          })
        }, function errorCallback(response) {
          $scope.alerts.push({
            type: 'danger',
            msg: 'Something went wrong setting you up. Try again?'
          })
        }
      );
    } else {
      $scope.alerts.push({
        type: 'danger',
        msg: "You shouldn't be here. Trying to mess with things?"
      })
    }
  };

  $scope.updateMentor = function() {
    $http.put('/api/mentors/' + $scope.mentor._id, $scope.mentor).then(
      function successCallback(response) {
        $scope.alerts.push({
          type: 'success',
          msg: 'Successfully updated your info'
        })
      }, function errorCallback(response) {
        $scope.alerts.push({
          type: 'danger',
          msg: 'Something went wrong updating your info. Try again?'
        })
      }
    );
  }

  $scope.closeAlert = function(index) {
    $scope.alerts.splice(index, 1);
  }

  init();
});
