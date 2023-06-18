/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { rpcStatus } from '../models/rpcStatus';
import type { serverBlastRequest } from '../models/serverBlastRequest';
import type { serverBlastResponse } from '../models/serverBlastResponse';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class SemanticBlastServiceService {

    /**
     * @param body
     * @returns serverBlastResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public static semanticBlastServiceBlast(
        body: serverBlastRequest,
    ): CancelablePromise<serverBlastResponse | rpcStatus> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/v1/blast',
            body: body,
        });
    }

}
