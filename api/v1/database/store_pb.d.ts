// package: database
// file: api/v1/database/store.proto

import * as jspb from "google-protobuf";
import * as api_v1_database_session_pb from "../../../api/v1/database/session_pb";
import * as api_v1_database_kv_pb from "../../../api/v1/database/kv_pb";
import * as api_v1_database_raft_control_pb from "../../../api/v1/database/raft_control_pb";

export class DeleteResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteResponse): DeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteResponse;
  static deserializeBinaryFromReader(message: DeleteResponse, reader: jspb.BinaryReader): DeleteResponse;
}

export namespace DeleteResponse {
  export type AsObject = {
  }
}

export class DeleteRequest extends jspb.Message {
  hasPayload(): boolean;
  clearPayload(): void;
  getPayload(): api_v1_database_kv_pb.KeyValue | undefined;
  setPayload(value?: api_v1_database_kv_pb.KeyValue): void;

  hasSession(): boolean;
  clearSession(): void;
  getSession(): api_v1_database_session_pb.Session | undefined;
  setSession(value?: api_v1_database_session_pb.Session): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRequest): DeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRequest;
  static deserializeBinaryFromReader(message: DeleteRequest, reader: jspb.BinaryReader): DeleteRequest;
}

export namespace DeleteRequest {
  export type AsObject = {
    payload?: api_v1_database_kv_pb.KeyValue.AsObject,
    session?: api_v1_database_session_pb.Session.AsObject,
  }
}

export class PutRequest extends jspb.Message {
  hasPayload(): boolean;
  clearPayload(): void;
  getPayload(): api_v1_database_kv_pb.KeyValue | undefined;
  setPayload(value?: api_v1_database_kv_pb.KeyValue): void;

  hasSession(): boolean;
  clearSession(): void;
  getSession(): api_v1_database_session_pb.Session | undefined;
  setSession(value?: api_v1_database_session_pb.Session): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PutRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PutRequest): PutRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PutRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PutRequest;
  static deserializeBinaryFromReader(message: PutRequest, reader: jspb.BinaryReader): PutRequest;
}

export namespace PutRequest {
  export type AsObject = {
    payload?: api_v1_database_kv_pb.KeyValue.AsObject,
    session?: api_v1_database_session_pb.Session.AsObject,
  }
}

export class GetRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  hasSession(): boolean;
  clearSession(): void;
  getSession(): api_v1_database_session_pb.Session | undefined;
  setSession(value?: api_v1_database_session_pb.Session): void;

  getClusterid(): number;
  setClusterid(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetRequest): GetRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRequest;
  static deserializeBinaryFromReader(message: GetRequest, reader: jspb.BinaryReader): GetRequest;
}

export namespace GetRequest {
  export type AsObject = {
    key: string,
    session?: api_v1_database_session_pb.Session.AsObject,
    clusterid: number,
  }
}

export class GetResponse extends jspb.Message {
  clearResultsList(): void;
  getResultsList(): Array<api_v1_database_kv_pb.KeyValue>;
  setResultsList(value: Array<api_v1_database_kv_pb.KeyValue>): void;
  addResults(value?: api_v1_database_kv_pb.KeyValue, index?: number): api_v1_database_kv_pb.KeyValue;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetResponse): GetResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetResponse;
  static deserializeBinaryFromReader(message: GetResponse, reader: jspb.BinaryReader): GetResponse;
}

export namespace GetResponse {
  export type AsObject = {
    resultsList: Array<api_v1_database_kv_pb.KeyValue.AsObject>,
  }
}
