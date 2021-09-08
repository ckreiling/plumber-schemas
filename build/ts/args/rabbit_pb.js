// source: args/rabbit.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.protos.args.RabbitConn', null, global);
goog.exportSymbol('proto.protos.args.RabbitReadArgs', null, global);
goog.exportSymbol('proto.protos.args.RabbitWriteArgs', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protos.args.RabbitConn = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.args.RabbitConn, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.args.RabbitConn.displayName = 'proto.protos.args.RabbitConn';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protos.args.RabbitReadArgs = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.args.RabbitReadArgs, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.args.RabbitReadArgs.displayName = 'proto.protos.args.RabbitReadArgs';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protos.args.RabbitWriteArgs = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.args.RabbitWriteArgs, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.args.RabbitWriteArgs.displayName = 'proto.protos.args.RabbitWriteArgs';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protos.args.RabbitConn.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.args.RabbitConn.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.args.RabbitConn} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.args.RabbitConn.toObject = function(includeInstance, msg) {
  var f, obj = {
    address: jspb.Message.getFieldWithDefault(msg, 1, ""),
    useTls: jspb.Message.getBooleanFieldWithDefault(msg, 2, false),
    insecureTls: jspb.Message.getBooleanFieldWithDefault(msg, 3, false)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protos.args.RabbitConn}
 */
proto.protos.args.RabbitConn.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.args.RabbitConn;
  return proto.protos.args.RabbitConn.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.args.RabbitConn} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.args.RabbitConn}
 */
proto.protos.args.RabbitConn.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setAddress(value);
      break;
    case 2:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setUseTls(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setInsecureTls(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protos.args.RabbitConn.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.args.RabbitConn.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.args.RabbitConn} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.args.RabbitConn.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAddress();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getUseTls();
  if (f) {
    writer.writeBool(
      2,
      f
    );
  }
  f = message.getInsecureTls();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
};


/**
 * optional string address = 1;
 * @return {string}
 */
proto.protos.args.RabbitConn.prototype.getAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.args.RabbitConn} returns this
 */
proto.protos.args.RabbitConn.prototype.setAddress = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional bool use_tls = 2;
 * @return {boolean}
 */
proto.protos.args.RabbitConn.prototype.getUseTls = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 2, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.args.RabbitConn} returns this
 */
proto.protos.args.RabbitConn.prototype.setUseTls = function(value) {
  return jspb.Message.setProto3BooleanField(this, 2, value);
};


/**
 * optional bool insecure_tls = 3;
 * @return {boolean}
 */
proto.protos.args.RabbitConn.prototype.getInsecureTls = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 3, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.args.RabbitConn} returns this
 */
proto.protos.args.RabbitConn.prototype.setInsecureTls = function(value) {
  return jspb.Message.setProto3BooleanField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protos.args.RabbitReadArgs.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.args.RabbitReadArgs.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.args.RabbitReadArgs} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.args.RabbitReadArgs.toObject = function(includeInstance, msg) {
  var f, obj = {
    exchangeName: jspb.Message.getFieldWithDefault(msg, 1, ""),
    queueName: jspb.Message.getFieldWithDefault(msg, 2, ""),
    bindingKey: jspb.Message.getFieldWithDefault(msg, 3, ""),
    queueExclusive: jspb.Message.getBooleanFieldWithDefault(msg, 4, false),
    queueDeclare: jspb.Message.getBooleanFieldWithDefault(msg, 5, false),
    queueDurable: jspb.Message.getBooleanFieldWithDefault(msg, 6, false),
    autoAck: jspb.Message.getBooleanFieldWithDefault(msg, 7, false),
    consumerTag: jspb.Message.getFieldWithDefault(msg, 8, ""),
    queueDelete: jspb.Message.getBooleanFieldWithDefault(msg, 9, false)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protos.args.RabbitReadArgs}
 */
proto.protos.args.RabbitReadArgs.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.args.RabbitReadArgs;
  return proto.protos.args.RabbitReadArgs.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.args.RabbitReadArgs} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.args.RabbitReadArgs}
 */
proto.protos.args.RabbitReadArgs.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setExchangeName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setQueueName(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setBindingKey(value);
      break;
    case 4:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setQueueExclusive(value);
      break;
    case 5:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setQueueDeclare(value);
      break;
    case 6:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setQueueDurable(value);
      break;
    case 7:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setAutoAck(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setConsumerTag(value);
      break;
    case 9:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setQueueDelete(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protos.args.RabbitReadArgs.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.args.RabbitReadArgs.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.args.RabbitReadArgs} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.args.RabbitReadArgs.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getExchangeName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getQueueName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getBindingKey();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getQueueExclusive();
  if (f) {
    writer.writeBool(
      4,
      f
    );
  }
  f = message.getQueueDeclare();
  if (f) {
    writer.writeBool(
      5,
      f
    );
  }
  f = message.getQueueDurable();
  if (f) {
    writer.writeBool(
      6,
      f
    );
  }
  f = message.getAutoAck();
  if (f) {
    writer.writeBool(
      7,
      f
    );
  }
  f = message.getConsumerTag();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getQueueDelete();
  if (f) {
    writer.writeBool(
      9,
      f
    );
  }
};


/**
 * optional string exchange_name = 1;
 * @return {string}
 */
proto.protos.args.RabbitReadArgs.prototype.getExchangeName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.args.RabbitReadArgs} returns this
 */
proto.protos.args.RabbitReadArgs.prototype.setExchangeName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string queue_name = 2;
 * @return {string}
 */
proto.protos.args.RabbitReadArgs.prototype.getQueueName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.args.RabbitReadArgs} returns this
 */
proto.protos.args.RabbitReadArgs.prototype.setQueueName = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string binding_key = 3;
 * @return {string}
 */
proto.protos.args.RabbitReadArgs.prototype.getBindingKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.args.RabbitReadArgs} returns this
 */
proto.protos.args.RabbitReadArgs.prototype.setBindingKey = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional bool queue_exclusive = 4;
 * @return {boolean}
 */
proto.protos.args.RabbitReadArgs.prototype.getQueueExclusive = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 4, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.args.RabbitReadArgs} returns this
 */
proto.protos.args.RabbitReadArgs.prototype.setQueueExclusive = function(value) {
  return jspb.Message.setProto3BooleanField(this, 4, value);
};


/**
 * optional bool queue_declare = 5;
 * @return {boolean}
 */
proto.protos.args.RabbitReadArgs.prototype.getQueueDeclare = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 5, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.args.RabbitReadArgs} returns this
 */
proto.protos.args.RabbitReadArgs.prototype.setQueueDeclare = function(value) {
  return jspb.Message.setProto3BooleanField(this, 5, value);
};


/**
 * optional bool queue_durable = 6;
 * @return {boolean}
 */
proto.protos.args.RabbitReadArgs.prototype.getQueueDurable = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 6, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.args.RabbitReadArgs} returns this
 */
proto.protos.args.RabbitReadArgs.prototype.setQueueDurable = function(value) {
  return jspb.Message.setProto3BooleanField(this, 6, value);
};


/**
 * optional bool auto_ack = 7;
 * @return {boolean}
 */
proto.protos.args.RabbitReadArgs.prototype.getAutoAck = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 7, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.args.RabbitReadArgs} returns this
 */
proto.protos.args.RabbitReadArgs.prototype.setAutoAck = function(value) {
  return jspb.Message.setProto3BooleanField(this, 7, value);
};


/**
 * optional string consumer_tag = 8;
 * @return {string}
 */
proto.protos.args.RabbitReadArgs.prototype.getConsumerTag = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.args.RabbitReadArgs} returns this
 */
proto.protos.args.RabbitReadArgs.prototype.setConsumerTag = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional bool queue_delete = 9;
 * @return {boolean}
 */
proto.protos.args.RabbitReadArgs.prototype.getQueueDelete = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 9, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.args.RabbitReadArgs} returns this
 */
proto.protos.args.RabbitReadArgs.prototype.setQueueDelete = function(value) {
  return jspb.Message.setProto3BooleanField(this, 9, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protos.args.RabbitWriteArgs.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.args.RabbitWriteArgs.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.args.RabbitWriteArgs} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.args.RabbitWriteArgs.toObject = function(includeInstance, msg) {
  var f, obj = {
    exchangeName: jspb.Message.getFieldWithDefault(msg, 1, ""),
    routingKey: jspb.Message.getFieldWithDefault(msg, 2, ""),
    appId: jspb.Message.getFieldWithDefault(msg, 3, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protos.args.RabbitWriteArgs}
 */
proto.protos.args.RabbitWriteArgs.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.args.RabbitWriteArgs;
  return proto.protos.args.RabbitWriteArgs.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.args.RabbitWriteArgs} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.args.RabbitWriteArgs}
 */
proto.protos.args.RabbitWriteArgs.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setExchangeName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setRoutingKey(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setAppId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protos.args.RabbitWriteArgs.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.args.RabbitWriteArgs.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.args.RabbitWriteArgs} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.args.RabbitWriteArgs.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getExchangeName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getRoutingKey();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getAppId();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional string exchange_name = 1;
 * @return {string}
 */
proto.protos.args.RabbitWriteArgs.prototype.getExchangeName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.args.RabbitWriteArgs} returns this
 */
proto.protos.args.RabbitWriteArgs.prototype.setExchangeName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string routing_key = 2;
 * @return {string}
 */
proto.protos.args.RabbitWriteArgs.prototype.getRoutingKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.args.RabbitWriteArgs} returns this
 */
proto.protos.args.RabbitWriteArgs.prototype.setRoutingKey = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string app_id = 3;
 * @return {string}
 */
proto.protos.args.RabbitWriteArgs.prototype.getAppId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.args.RabbitWriteArgs} returns this
 */
proto.protos.args.RabbitWriteArgs.prototype.setAppId = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


goog.object.extend(exports, proto.protos.args);
