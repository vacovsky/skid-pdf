(function () {
    var app = angular.module('skidpdf', []);
    app.Root = '/pdf';

    app.controller('skidpdfcontrol', function ($scope, $http) {
        $scope.message = "";
        $scope.formSelected = "simpleGET";
        $scope.pdfRequest = {
            "url": "",
            "data": "",
            "grayscale": false,
            "landscape": false,
            "headers": {},
            "postParams": {}
        };

        $scope.basePdfRequest = {
            "url": "",
            "data": "",
            "grayscale": false,
            "landscape": false,
            "headers": {},
            "postParams": {}
        };

        $scope.addHeader = function () {};

        $scope.addPostData = function () {

        };

        $scope.makeRequest = function () {
            if ($scope.pdfRequest.url == "") {
                $scope.message = "Please provide a target URL.";
            } else {
                if ($scope.formSelected == "simpleGET") {
                    $http.get("/pdf?uri=" + $scope.pdfRequest.url + "&grayscale=" + $scope.pdfRequest.grayscale + "&landscape=" + $scope.pdfRequest.landscape, {
                        responseType: "arraybuffer"
                    }).then(function (response) {
                        var file = new Blob([response.data], {
                            type: 'application/pdf'
                        });
                        var fileURL = URL.createObjectURL(file);
                        var a = document.createElement('a');
                        a.href = fileURL;
                        a.target = '_blank';
                        a.download = $scope.formSelected + ".pdf";
                        document.body.appendChild(a);
                        a.click();
                    }).then(function () {
                        $scope.message = "";
                    });
                } else if ($scope.formSelected == "complexGET" || $scope.formSelected == "complexPOST") {
                    $http.post("/pdf", $scope.pdfRequest, {
                        responseType: "arraybuffer"
                    }).then(function (response) {
                        var file = new Blob([response.data], {
                            type: 'application/pdf'
                        });
                        var fileURL = URL.createObjectURL(file);
                        var a = document.createElement('a');
                        a.href = fileURL;
                        a.target = '_blank';
                        a.download = $scope.formSelected + ".pdf";
                        document.body.appendChild(a);
                        a.click();
                    });
                } else {
                    $scope.message = "Please select a job type."
                }
            }
        };



        $scope.resetForm = function () {
            $scope.pdfRequest = $scope.basePdfRequest;
        };


    });

})();