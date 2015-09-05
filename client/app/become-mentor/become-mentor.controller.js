'use strict';

angular.module('mentorMeApp')
  .controller('BecomeMentorCtrl', function ($scope, $http) {
    var init = function() {
      kik.getUser(function (user) {
        if ( !user ) {
          // user denied access to their information
        } else {
          $scope.username = user.username; // 'string'
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
        $http.post('/api/mentors', $scope.mentor);
        $scope.mentor = {};
      } else {
        // They shouldn't have been able to do this
      }
    };

    init();
  });
