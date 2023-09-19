// package: grid
// file: grid.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class GridNode extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getName(): string;
  setName(value: string): void;

  getType(): GridNodeTypeEnumMap[keyof GridNodeTypeEnumMap];
  setType(value: GridNodeTypeEnumMap[keyof GridNodeTypeEnumMap]): void;

  getPower(): number;
  setPower(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GridNode.AsObject;
  static toObject(includeInstance: boolean, msg: GridNode): GridNode.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GridNode, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GridNode;
  static deserializeBinaryFromReader(message: GridNode, reader: jspb.BinaryReader): GridNode;
}

export namespace GridNode {
  export type AsObject = {
    id: number,
    name: string,
    type: GridNodeTypeEnumMap[keyof GridNodeTypeEnumMap],
    power: number,
  }
}

export class GetNodeRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetNodeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetNodeRequest): GetNodeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetNodeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetNodeRequest;
  static deserializeBinaryFromReader(message: GetNodeRequest, reader: jspb.BinaryReader): GetNodeRequest;
}

export namespace GetNodeRequest {
  export type AsObject = {
    id: number,
  }
}

export class UpdateNodeRequest extends jspb.Message {
  hasNode(): boolean;
  clearNode(): void;
  getNode(): GridNode | undefined;
  setNode(value?: GridNode): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateNodeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateNodeRequest): UpdateNodeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateNodeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateNodeRequest;
  static deserializeBinaryFromReader(message: UpdateNodeRequest, reader: jspb.BinaryReader): UpdateNodeRequest;
}

export namespace UpdateNodeRequest {
  export type AsObject = {
    node?: GridNode.AsObject,
  }
}

export class GetNodeResponse extends jspb.Message {
  hasNode(): boolean;
  clearNode(): void;
  getNode(): GridNode | undefined;
  setNode(value?: GridNode): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetNodeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetNodeResponse): GetNodeResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetNodeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetNodeResponse;
  static deserializeBinaryFromReader(message: GetNodeResponse, reader: jspb.BinaryReader): GetNodeResponse;
}

export namespace GetNodeResponse {
  export type AsObject = {
    node?: GridNode.AsObject,
  }
}

export class UpdateNodeResponse extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateNodeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateNodeResponse): UpdateNodeResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateNodeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateNodeResponse;
  static deserializeBinaryFromReader(message: UpdateNodeResponse, reader: jspb.BinaryReader): UpdateNodeResponse;
}

export namespace UpdateNodeResponse {
  export type AsObject = {
    id: number,
  }
}

export class NodeLink extends jspb.Message {
  getSource(): number;
  setSource(value: number): void;

  getTarget(): number;
  setTarget(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NodeLink.AsObject;
  static toObject(includeInstance: boolean, msg: NodeLink): NodeLink.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NodeLink, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NodeLink;
  static deserializeBinaryFromReader(message: NodeLink, reader: jspb.BinaryReader): NodeLink;
}

export namespace NodeLink {
  export type AsObject = {
    source: number,
    target: number,
  }
}

export class GetGridResponse extends jspb.Message {
  clearNodesList(): void;
  getNodesList(): Array<GridNode>;
  setNodesList(value: Array<GridNode>): void;
  addNodes(value?: GridNode, index?: number): GridNode;

  clearLinksList(): void;
  getLinksList(): Array<NodeLink>;
  setLinksList(value: Array<NodeLink>): void;
  addLinks(value?: NodeLink, index?: number): NodeLink;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetGridResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetGridResponse): GetGridResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetGridResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetGridResponse;
  static deserializeBinaryFromReader(message: GetGridResponse, reader: jspb.BinaryReader): GetGridResponse;
}

export namespace GetGridResponse {
  export type AsObject = {
    nodesList: Array<GridNode.AsObject>,
    linksList: Array<NodeLink.AsObject>,
  }
}

export interface GridNodeTypeEnumMap {
  POWER_STATION: 0;
  SOLAR_PLANT: 1;
  SUBSTATION: 2;
  TRANSMISSION: 3;
  CONSUMER: 4;
}

export const GridNodeTypeEnum: GridNodeTypeEnumMap;

