/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as status_pb from './status_pb';

import {
  Message,
  MessageRequest,
  PostRequest} from './message_pb';

export class MessageServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoSubscribe = new grpcWeb.AbstractClientBase.MethodInfo(
    Message,
    (request: MessageRequest) => {
      return request.serializeBinary();
    },
    Message.deserializeBinary
  );

  subscribe(
    request: MessageRequest,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/MessageService/Subscribe',
      request,
      metadata || {},
      this.methodInfoSubscribe);
  }

  methodInfoPost = new grpcWeb.AbstractClientBase.MethodInfo(
    status_pbStatus,
    (request: PostRequest) => {
      return request.serializeBinary();
    },
    status_pbStatus.deserializeBinary
  );

  post(
    request: PostRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: status_pbStatus) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/MessageService/Post',
      request,
      metadata || {},
      this.methodInfoPost,
      callback);
  }

}

