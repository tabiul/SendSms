'use strict'
smsApp.controller('smsController', function ($scope, $http, $location) {
    $scope.sendSms = function () {
        if($scope.phoneNumber === undefined) {
            $scope.errorMessage = "Phone Number is required";
            return;
        }
        if($scope.message === undefined) {
            $scope.errorMessage = "message is required";
            return;
        }
        $scope.maxSize = 160 * 3;
        if($scope.message.length > $scope.maxSize) {
            $scope.errorMessage = "maximum message size is " + maxSize;
            return;
        }
        var data = {
            phoneNumber: $scope.phoneNumber,
            message: $scope.message
        };
        var res = $http.post('sms/send', angular.toJson(data));
        res.success(function (data, status, headers, config) {
            if (status == 200) {
                $scope.responses = data.responses;
            }
        });
        res.error(function (response, status, headers, config) {
            $scope.errorMessage = response
        });
    }

});
