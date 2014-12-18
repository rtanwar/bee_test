//angular.module('airline', []).config(demoRouter);
//create module with ngRoute	
var cdoApp = angular.module('cdoApp',['ngRoute','restangular','ui.bootstrap']);

/**
*
*/

cdoApp.config(function($routeProvider) {
	$routeProvider
		.when('/', {templateUrl: 'partials/cdo/list.html',
		 controller: 'cdoController'
		 })
		.when('/add', {
		 templateUrl: 'partials/cdo/detail.html',
		 controller: 'cdoEditController',		 
		 resolve:{data:function(){return {};}}
		 })
		.when('/edit/:id', {
		 templateUrl: 'partials/cdo/detail.html',
		 controller: 'cdoEditController',		 
		 resolve:{data:function($route,Restangular){					 
		 return Restangular.one('ci/cdo/details', $route.current.params.id).get()
		 }}		 
		 })

		.otherwise({redirectTo:'/'});				
})

//Assign route function to module
//cdoApp.config(demoRouter);
//cdoApp.config(['$routeProvider', demoRouter]);

//create controllers to hold controllers 
//var Controllers = {};

/**
*create controller
*/

 
cdoController = cdoApp.controller('cdoController',function($scope,$http,Restangular,$location, $routeParams,$route)
{

	$scope.headers = [
	{ title: 'S.No.',       value: 'id' },
    { title: 'Date of Intimation',       value: 'Intimation_Date' },
    { title: 'Insured name', value: 'Name_of_Insured' },
    { title: 'Registration No',     value: 'Vehicle_No'},
    { title: 'Edit Current'},
  	];
	$scope.newdata = {};
	$scope.search = {};
	$scope.alert={};
	$scope.alert=null;
	//intialize status of save button on details page 
	

	$scope.sortvalue= '=';
  	$scope.sortedby= '=';

	$scope.cellInputEditableTemplate = '<input ng-class="\'colt\' + col.index" ng-input="COL_FIELD" ng-model="COL_FIELD" ng-blur="updateEntity(row)" />';
    $scope.cellSelectEditableTemplate = '<select ng-class="\'colt\' + col.index" ng-input="COL_FIELD" ng-model="COL_FIELD" ng-options="id as name for (id, name) in statuses" ng-blur="updateEntity(row)" />';
	//$scope.name="";
	
	//$scope.names=[{name:'user1',city:'city45'},{name:'user12',city:'city98'},{name:'user98',city:'city2'}];	
	//UserFactory.getnames().success(function(data) {$scope.names = data;});	
	//RestangularProvider.setBaseUrl('')			
	
	//$scope.insurers=data.data;	
	//if (data.recordcount) {$scope.recordcount=data.recordcount; $scope.numpages = $scope.recordcount/10};		
	
	$scope.fetchResult = function () {
	$scope.showloading = true;
	$scope.insurers= {};
	return Restangular.all('ci/cdo/').getList($scope.search).then(function (data) {		
      if (data.data) 
	  {
		  $scope.insurers = data.data;
		  $scope.totalPages = (data.recordcount/10)-1;
		  $scope.recordcount = data.recordcount;		  
		  $scope.showloading = false;  	  
	  }
    });	
  };
	
	$scope.selectPage = function(page) {
		$scope.search.pageNumber = page;    	
		$scope.fetchResult();
    } 
	//$scope.$watch("search.searchText", function() {_.debounce(function() {$scope.selectPage(); console.log('changed')},100)}, true);
	

	
	function getone(id) {
	Restangular.one('ci/cdo/details', id).get().then(function(insurer) {
		$scope.insurer=insurer;
		console.debug(insurer);	
		//alert($scope.contact.firstname);
		});			
	}	
	
	
	

	$scope.onSort = function (sortedBy, sortDir) {
	console.debug(sortedBy) ;
	console.debug(sortDir) ;
    $scope.search.sortDir = sortDir;
    $scope.search.sortedBy = sortedBy;      
    $scope.search.pageNumber = 1;    
  	$scope.sortedby= sortedBy;
  	$scope.sortvalue=sortedBy;
    $scope.fetchResult().then(function () {
      //The request fires correctly but sometimes the ui doesn't update, that's a fix
     $scope.search.pageNumber = 1;
    });

  };
 
    $scope.selectPage(1);

  
});


cdoEditController = cdoApp.controller('cdoEditController',function($scope,$http,Restangular,$location, $routeParams,$route,data)
{
	$scope.insurer = data;
	console.debug(data.id);
	$scope.savedisabled =  false;	
	$scope.deletedisabled=(typeof data.id === 'undefined'); // if the id is not set means we are addding hence no delete	
	
	$scope.go = function ( path ) {		
	  $location.path( path );
	};
	
	$scope.savecurrent = function(){
		if ($scope.insurer.id)
			$scope.insurer.post().then(function(){$scope.alert={ type: 'success', msg: 'Record Saved!!' }; $location.path('/')});
		else
			//baseNames.post($scope.insurer).then(function(){$location.path('/')});
			Restangular.all('ci/cdo/details').post($scope.insurer).then(function(){$scope.alert={ type: 'success', msg: 'Record Added!!' };$location.path('/')});
		$scope.savedisabled =  false;		
	}	
	
	
	$scope.deletecurrent = function(){				
		$scope.insurer.post('DELETE').then(function(){$scope.alert={ type: 'danger', msg: 'Record Deleted!!' };$location.path('/')});				
	}

	$scope.Insurer_address = function(Insurer_name) {
		return $http.get("/ci/cdo/getlist?field=address&q="+Insurer_name).then(function(response){			
		return response.data;	
		});
	}
	$scope.Vehicles_particulars= function(vehicle_name) {
		return $http.get("/ci/cdo/getlist?field=vehicle_name&q="+vehicle_name).then(function(response){			
		return response.data;	
		});
	}
	
	
	$scope.Insurer_names = function(Insurer_name) {
		return $http.get("/ci/cdo/getlist?field=insurer&q="+Insurer_name).then(function(response){			
		return response.data;	
		});
	}
	
		$scope.surveyor_descriptions = function(desc) {
		return $http.get("/ci/cdo/getlist?field=surveyor_descriptions&q="+desc).then(function(response){			
			return response.data;	
		});
	}

	
});


cdoFinalEditController = cdoApp.controller('cdoAppFinalEditController',function($scope,$http,Restangular,$location, $routeParams,$route,data)
{
	
	$scope.spareheaders = [
    { title: 'Sr No.'},
	{ title: 'Description'},
    { title: 'Estimated' },
    { title: 'Dealer Price'},	
  	];
	$scope.labourheaders = [
    { title: 'Sr No'},
    { title: 'Description'},
    { title: 'Estimated' },
    { title: 'Dealer Price'},
	{ title: 'Allowed'},
  	];
	$scope.spares = [];
	$scope.labours = [];
	
	$scope.cellInputEditableTemplate = '<input ng-class="\'colt\' + col.index" ng-input="COL_FIELD" ng-model="COL_FIELD" ng-blur="updateEntity(row)" />';
    $scope.cellSelectEditableTemplate = '<select ng-class="\'colt\' + col.index" ng-input="COL_FIELD" ng-model="COL_FIELD" ng-options="id as name for (id, name) in statuses" ng-blur="updateEntity(row)" />';
	
		
	
	$scope.fetchResult = function () {
		$scope.showloading = true;				 
		data.spares =  Restangular.all('ci/cdo/sparedetails/'+$route.current.params.id).getList().then(function (data) {		
		  if (data)
		  {
			  $scope.spares = data;
			  $scope.spares_org = angular.copy($scope.spares);			  
			  $scope.showloading = false;  	  
		  }
		});	
		
		Restangular.all('ci/cdo/labourdetails/'+$route.current.params.id).getList().then(function (data) {		
		  if (data) 
		  {
			  $scope.labours = data;
			  $scope.labours_org = angular.copy(data);			  
			  $scope.showloading = false;  	  
		  }
		});	
		 
		Restangular.one('ci/cdo/details', $route.current.params.id).get().then(function (data) {		
		  if (data) 
		  {			  
			  $scope.insurer = data;			  
			  $scope.showloading = false;  	  
		  }
		});	
	};
	
	
	//$scope.insurer = data.insurer;				
	$scope.sortableOptions = {
		cursor: 'move',	
		//disabled: true  
	}			
	//$scope.spares = data.spares;	
	
	$scope.spares_deleted = [];
	$scope.labours_deleted = [];
	//$scope.labours = data.labours;		
		

	$scope.LabourGridOptions = { 
      data: 'labours', 
      enableCellSelection: true,
      canSelectRows: false,      	  
	  enableCellSelection: true,
        enableRowSelection: false,
        enableCellEdit: true,
      columnDefs: [
        {field: 'ItemNo', displayName: 'ItemNo ', enableCellEdit: true}, 
		{field: 'Description', displayName: 'ItemNo ', enableCellEdit: true}, 
		{field: 'DealerPrice', displayName: 'ItemNo ', enableCellEdit: true}, 
		{field: 'Estimated', displayName: 'ItemNo ', enableCellEdit: true}, 
		{field: 'Allowed', displayName: 'ItemNo ', enableCellEdit: true}, 
		{field: 'Stax', displayName: 'ItemNo ', enableCellEdit: true}]		
    };
	
	$scope.delete_spare = function(item)
	{		
		var index=$scope.spares.indexOf(item);							
		var itemd= $scope.spares.splice(index,1);
		
		if (itemd[0].id)
			$scope.spares_deleted = $scope.spares_deleted.concat(itemd);
		console.debug($scope.spares_deleted);
	}
	
	$scope.add_spare = function ()
	{
		$scope.spares.push({ReportId:$scope.insurer.id,Description: '',Estimated:'',DealerPrice:''});				
	}

	$scope.commit_spare = function ()
	{			
		$scope.ReSerialize();
		angular.forEach($scope.spares, function(spare) {
			if (!spare.id)
			{
				console.debug('New item:'+spare.Description);
				Restangular.all('ci/cdo/sparedetails/').post(spare);
			}
			//else if ($scope.getById($scope.spares_org,spare.id) != spare)
			else if (!angular.equals($scope.getById($scope.spares_org,spare.id) , spare))
			{
					console.debug('Updating'+spare.id);																		
					spare.post();
			}
		})
		
		//console.debug($scope.spares_deleted);
		angular.forEach($scope.spares_deleted, function(spare) { 				
 				console.debug('Deleted Item:'+spare.id+','+spare.Description);
				spare.post('DELETE');
		})
	}
	
	$scope.getById = function(input, id) {		
		var i=0, len=input.length;
		for (; i<len; i++) {
		  //convert both ids to numbers to be sure
		  if (+input[i].id == +id) {
			return input[i];
		  }
		}
		return null;
	}
	
	
	
	$scope.delete_labour = function(item)
	{		

		var index=$scope.labours.indexOf(item);							
		var itemd= $scope.labours.splice(index,1);
		
		if (itemd[0].id)
			$scope.labours_deleted = $scope.labours_deleted.concat(itemd);
		console.debug($scope.labours_deleted);
	}
	
	$scope.add_labour = function ()
	{
		$scope.labours.push({ReportId:$scope.insurer.id,Description: '',Estimated:'',DealerPrice:''});				
	}

	$scope.commit_labour = function ()
	{	
		$scope.ReSerialize();
		//console.debug($scope.labours);
		angular.forEach($scope.labours, function(labour) {
			if (!labour.id)
			{
				console.debug('New item:'+labour.Description);
				Restangular.all('ci/cdo/labourdetails/').post(labour);
			}
			else if (!angular.equals($scope.getById($scope.labours_org,labour.id) , labour))
			{
					console.debug('Updating'+labour.id);																		
					labour.post();
			}
		})
		
		//console.debug($scope.spares_deleted);
		angular.forEach($scope.labours_deleted, function(labour) { 				
 				console.debug('Deleted Item:'+labour.id+','+labour.Description);
				labour.post('DELETE');
		})
		//$scope.alert={ type: 'success', msg: 'Record Saved!!' };
	}
	
	$scope.spare_descriptions = function(desc) {
		return $http.get("/ci/cdo/getlist?field=spare_descriptions&q="+desc).then(function(response){			
			return response.data;	
		});
	}	
	
	$scope.spare_estimated_total = function() 
	{	
		var t= {};
		var total_E = 0;
		var total_D = 0;
		angular.forEach($scope.spares, function(item) {
			total_E+= item.Estimated*1;
			total_D+= item.DealerPrice*1;
	})
		t.Estimated = total_E;
		t.Dealer = total_D;
		return t;
	}      	
	
	$scope.labour_estimated_total = function() 
	{	
		
		var t= {};
		var total_E = 0;
		var total_D = 0;
		var total_A		= 0;
		
		angular.forEach($scope.labours, function(item) {		
			
			total_E+= item.Estimated*1;
			total_D+= item.DealerPrice*1;
			total_A+=item.Allowed*1;
	})
		t.Estimated = total_E;
		t.Dealer = total_D;
		t.Allowed = total_A;		
		return t;
	}    

  	
	$scope.ReSerialize = function()
	{
		id = 1;
		angular.forEach($scope.labours, function(item) {
				item.ItemNo = id++;
		})
		id = 1;
		angular.forEach($scope.spares, function(item) {
			item.ItemNo = id++;
		})	
	}
	  
	//$scope.$watch(function () { return angular.toJson($scope.spares)}, function() {
    //   console.debug($scope.spares);
   // });

	
	$scope.$on('ngGridEventStartCellEdit', function () {	
			elm.focus();
			elm.select();
		});
		
	$scope.fetchResult();			
	
});




angular.module('cdoApp').directive('sortBy', function () {
  return {
  	
    template : 
      '<a ng-click="sort(sortvalue)">'+
        '<span ng-transclude></span>'+ 
        '<span ng-show="sortedby == sortvalue">'+
        '<i ng-class="{true: \'glyphicon glyphicon-arrow-up\', false: \'glyphicon glyphicon-arrow-down\'}[sortdir == \'asc\']"></i></span>'+
      '</a>',
    restrict: 'E',
    transclude: true,
    replace: true,
    scope: {
      sortdir: '=',
      sortedby: '=',
      sortvalue: '@',
      onsort: '='
    },
    link: function (scope, element, attrs) {
      scope.sort = function () {
        if (scope.sortedby == scope.sortvalue)
          scope.sortdir = scope.sortdir == 'asc' ? 'desc' : 'asc';
        else {
          scope.sortedby = scope.sortvalue;
          scope.sortdir = 'asc';
        }
        scope.onsort(scope.sortedby, scope.sortdir);
      }
    }
  };
});
