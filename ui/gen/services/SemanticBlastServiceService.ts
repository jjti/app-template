/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { rpcStatus } from '../models/rpcStatus';
import type { sblastBlastRequest } from '../models/sblastBlastRequest';
import type { sblastBlastResponse } from '../models/sblastBlastResponse';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class SemanticBlastServiceService {

    /**
     * @param body
     * @returns sblastBlastResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public static semanticBlastServiceBlast(
        body: sblastBlastRequest,
    ): CancelablePromise<sblastBlastResponse | rpcStatus> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/v1/blast',
            body: body,
        });
    }

}
