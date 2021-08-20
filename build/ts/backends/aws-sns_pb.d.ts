// package: protos.backends
// file: backends/aws-sns.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class AWSSNS extends jspb.Message { 
    getTopicArn(): string;
    setTopicArn(value: string): AWSSNS;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AWSSNS.AsObject;
    static toObject(includeInstance: boolean, msg: AWSSNS): AWSSNS.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AWSSNS, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AWSSNS;
    static deserializeBinaryFromReader(message: AWSSNS, reader: jspb.BinaryReader): AWSSNS;
}

export namespace AWSSNS {
    export type AsObject = {
        topicArn: string,
    }
}
