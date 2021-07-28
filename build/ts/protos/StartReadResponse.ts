// Original file: protos/read.proto

import type { Status as _protos_common_Status, Status__Output as _protos_common_Status__Output } from '../protos/common/Status';

export interface StartReadResponse {
  'read_id'?: (string);
  'status'?: (_protos_common_Status | null);
}

export interface StartReadResponse__Output {
  'read_id': (string);
  'status': (_protos_common_Status__Output | null);
}
