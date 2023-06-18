/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { protobufAny } from './protobufAny';

export type rpcStatus = {
    code?: number;
    message?: string;
    details?: Array<protobufAny>;
};

