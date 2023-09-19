// package: grid
// file: grid.proto

var grid_pb = require("./grid_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var GridService = (function () {
  function GridService() {}
  GridService.serviceName = "grid.GridService";
  return GridService;
}());

GridService.GetGrid = {
  methodName: "GetGrid",
  service: GridService,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: grid_pb.GetGridResponse
};

GridService.GetNode = {
  methodName: "GetNode",
  service: GridService,
  requestStream: false,
  responseStream: false,
  requestType: grid_pb.GetNodeRequest,
  responseType: grid_pb.GetNodeResponse
};

GridService.UpdateNode = {
  methodName: "UpdateNode",
  service: GridService,
  requestStream: false,
  responseStream: false,
  requestType: grid_pb.UpdateNodeRequest,
  responseType: grid_pb.UpdateNodeResponse
};

exports.GridService = GridService;

function GridServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

GridServiceClient.prototype.getGrid = function getGrid(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(GridService.GetGrid, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

GridServiceClient.prototype.getNode = function getNode(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(GridService.GetNode, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

GridServiceClient.prototype.updateNode = function updateNode(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(GridService.UpdateNode, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.GridServiceClient = GridServiceClient;

