'use strict';

describe('Controller: NoPermissionsCtrl', function () {

  // load the controller's module
  beforeEach(module('mentorMeApp'));

  var NoPermissionsCtrl, scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    NoPermissionsCtrl = $controller('NoPermissionsCtrl', {
      $scope: scope
    });
  }));

  it('should ...', function () {
    expect(1).toEqual(1);
  });
});
