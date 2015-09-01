'use strict';

angular.module('mentorMeApp')
  .controller('BecomeMentorCtrl', function ($scope, $http) {
    $scope.message = 'Hello';

    $scope.addMentor = function() {
      if($scope.mentor === {}) {
        return;
      }
      $scope.mentor.tags = $scope.mentor.tags.split(",");
      $http.post('/api/mentors', $scope.mentor);
      $scope.mentor = {};
    };
  });
