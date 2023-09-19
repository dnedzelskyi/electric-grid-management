// package: grid
// file: grid.proto

import * as grid_pb from "./grid_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "@improbable-eng/grpc-web";

type GridServiceGetGrid = {
  readonly methodName: string;
  readonly service: typeof GridService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof grid_pb.GetGridResponse;
};

type GridServiceGetNode = {
  readonly methodName: string;
  readonly service: typeof GridService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof grid_pb.GetNodeRequest;
  readonly responseType: typeof grid_pb.GetNodeResponse;
};

type GridServiceUpdateNode = {
  readonly methodName: string;
  readonly service: typeof GridService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof grid_pb.UpdateNodeRequest;
  readonly responseType: typeof grid_pb.UpdateNodeResponse;
};

export class GridService {
  static readonly serviceName: string;
  static readonly GetGrid: GridServiceGetGrid;
  static readonly GetNode: GridServiceGetNode;
  static readonly UpdateNode: GridServiceUpdateNode;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class GridServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getGrid(
    requestMessage: google_protobuf_empty_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: grid_pb.GetGridResponse|null) => void
  ): UnaryResponse;
  getGrid(
    requestMessage: google_protobuf_empty_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: grid_pb.GetGridResponse|null) => void
  ): UnaryResponse;
  getNode(
    requestMessage: grid_pb.GetNodeRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: grid_pb.GetNodeResponse|null) => void
  ): UnaryResponse;
  getNode(
    requestMessage: grid_pb.GetNodeRequest,
    callback: (error: ServiceError|null, responseMessage: grid_pb.GetNodeResponse|null) => void
  ): UnaryResponse;
  updateNode(
    requestMessage: grid_pb.UpdateNodeRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: grid_pb.UpdateNodeResponse|null) => void
  ): UnaryResponse;
  updateNode(
    requestMessage: grid_pb.UpdateNodeRequest,
    callback: (error: ServiceError|null, responseMessage: grid_pb.UpdateNodeResponse|null) => void
  ): UnaryResponse;
}

