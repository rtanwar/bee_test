<!DOCTYPE html>
<html>
<head>
	<title>Beego</title>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<script type="text/javascript" src="/static/js/jquery.js"></script>
	<script type="text/javascript" src="/static/js/underscore-min.js"></script>
	<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
	<script type="text/javascript" src="/static/js/angular.min.js"></script>
	<script type="text/javascript" src="/static/js/angular-route.min.js"></script>
	<script type="text/javascript" src="/static/js/restangular.js"></script>
	<script type="text/javascript" src="/static/js/angular-ui.js"></script>	
	<script type="text/javascript" src="/static/js/ui-bootstrap-tpls-0.11.2.js"></script>	
	<script type="text/javascript" src="/static/js/app.js"></script>	



	<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.css"/>
</head>

<body>
	<header class="hero-unit" style="color:black;background-color:#A9F16C">
		<div ng-app="countryApp" class="container"> 
			<script type="text/ng-template" id="list.html">
				<alert ng-repeat="alert in alerts" type="{{alert.type}}" close="closeAlert($index)">{{alert.msg}}</alert>
				<div class="controls">          
						<button id="button4id" data-loading-text="Loading..." ng-click="add_new()" name="button4id" class="btn btn-success">Add New</button>
				</div>
				<table class="table table-striped">					
					<thead>
						<tr>
							<th ng-repeat="header in headers">
								<sort-by  onsort="onSort" sortdir="search.sortDir" sortedby="search.sortedBy" sortvalue="{{header.value}}">{{ header.title }}</sort-by>
							</th>
						</tr>
					</thead>
					<tbody>
						<tr ng-repeat="country in countries">
							<td>{{country.Name}}</td>
							<td>{{country.Continent}}</td>
							<td>{{country.Region}}</td>
							<td>{{country.Population}}</td>
							<td><a href="#/edit/{{country.Id}}" ><b class="glyphicon glyphicon-pencil"></b></a></td>	
						</tr></tbody>
					</table>			
				</script>

				<script type="text/ng-template" id="/detail.html">
					<alert ng-repeat="alert in alerts" type="{{alert.type}}" close="closeAlert($index)">{{alert.msg}}</alert>
					<legend>Country</legend>

					<form autocomplete="off" id="myForm" name="myForm" class="form-horizontal" role="form">
						<fieldset>
							<div class="form-group">              
								<label class="col-sm-2 control-label" for ="Id">Id</label>
								<div class="col-sm-3 ">
									<input required="required"  id="Id" name="Id" size=30 ng-model="country.Id"  class="form-control input-sm">              
								</div>
								<label class="col-sm-2 control-label" for ="name">Name</label>
								<div class="col-sm-3 ">
									<input autofocus="autofocus" required="required"  id="Name" name="Name" size=30 ng-model="country.Name"  class="form-control input-sm"/>
								</div>
							</div>

						</fieldset>          
						<div class="panel-footer" style="overflow:hidden;">
							<div class="control-group">          
								<div class="controls">          
									<button id="button1id" data-loading-text="Saving..." ng-click='savecurrent($event);commit_task();' name="button1id" class="btn btn-success">Save</button>
									<button id="button2id" name="button2id" ng-click="go('/')" class="btn btn-warning" formnovalidate>Cancel</button>
									<button ng-disabled="deletedisabled" id="button3id" name="button3id" ng-click="deletecurrent();" class="btn  btn-danger" formnovalidate>Delete</button>
								</div>
							</div>

						</div> <!-- panel-footer -->
					</form>

				</script>
				<div class="row">
					<div class="hero-text">
						<h1>Welcome <<<.user>>>! on JSON details page!</h1>
						<div ng-view></div>
					</div>
				</div>
			</div>
		</header>


	</body>
	</html>
