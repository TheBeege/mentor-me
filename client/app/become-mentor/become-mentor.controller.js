'use strict';

angular.module('mentorMeApp')
  .controller('BecomeMentorCtrl', function ($scope, $http) {
    var init = function() {
      kik.getUser(function (user) {
        if ( !user ) {
          // user denied access to their information
          $scope.error = "We cannot register you without your kik user information.";
        } else {
          $scope.username = user.username; // 'string'
          $scope.fullName = user.fullName;
          $scope.pic = user.pic;
          $scope.thumbnail = user.thumbnail;
        }
      })

      // check if they're already a mentor
      $http({
        url: '/api/mentors/byusername/' + user.username,
        method: 'GET'
      }).success(function(mentors) {
        if(mentors.length == 1) {
          // they're already a mentor, we should render the edit view
          $scope.alreadyMentor = true;
          
        } else {
          $scope.alreadyMentor = false;
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
        $http.post('/api/mentors', $scope.mentor);
        $scope.mentor = {};
      } else {
        // They shouldn't have been able to do this
      }
    };

    init();
  });
