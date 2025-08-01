/*
 * Copyright 2025 coze-dev Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
 
// THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.
/* eslint-disable */
/* tslint:disable */
// @ts-nocheck

import * as audit from './namespaces/audit';
import * as counter from './namespaces/counter';
import * as flow_platform_audit_common from './namespaces/flow_platform_audit_common';
import * as marketplace_common from './namespaces/marketplace_common';
import * as report_admin_api from './namespaces/report_admin_api';
import * as report_common from './namespaces/report_common';
import * as report_public_api from './namespaces/report_public_api';
import * as user_common from './namespaces/user_common';
import * as user_public_api from './namespaces/user_public_api';
import * as user_rpc from './namespaces/user_rpc';

export {
  audit,
  counter,
  flow_platform_audit_common,
  marketplace_common,
  report_admin_api,
  report_common,
  report_public_api,
  user_common,
  user_public_api,
  user_rpc,
};
export * from './namespaces/audit';
export * from './namespaces/counter';
export * from './namespaces/flow_platform_audit_common';
export * from './namespaces/marketplace_common';
export * from './namespaces/report_admin_api';
export * from './namespaces/report_common';
export * from './namespaces/report_public_api';
export * from './namespaces/user_common';
export * from './namespaces/user_public_api';
export * from './namespaces/user_rpc';

export type Int64 = string | number;

export default class BasicApiService<T> {
  private request: any = () => {
    throw new Error('BasicApiService.request is undefined');
  };
  private baseURL: string | ((path: string) => string) = '';

  constructor(options?: {
    baseURL?: string | ((path: string) => string);
    request?<R>(
      params: {
        url: string;
        method: 'GET' | 'DELETE' | 'POST' | 'PUT' | 'PATCH';
        data?: any;
        params?: any;
        headers?: any;
      },
      options?: T,
    ): Promise<R>;
  }) {
    this.request = options?.request || this.request;
    this.baseURL = options?.baseURL || '';
  }

  private genBaseURL(path: string) {
    return typeof this.baseURL === 'string'
      ? this.baseURL + path
      : this.baseURL(path);
  }

  /**
   * POST /api/user/v2/update_profile_check
   *
   * == 用户信息 ==
   *
   * -- http --
   */
  PublicUpdateUserProfileCheck(
    req?: user_public_api.UpdateUserProfileCheckRequest,
    options?: T,
  ): Promise<user_public_api.UpdateUserProfileCheckResponse> {
    const _req = req || {};
    const url = this.genBaseURL('/api/user/v2/update_profile_check');
    const method = 'POST';
    const data = { user_unique_name: _req['user_unique_name'] };
    return this.request({ url, method, data }, options);
  }

  /** GET /api/user/v2/get_profile */
  PublicGetUserProfile(
    req?: user_public_api.GetUserProfileRequest,
    options?: T,
  ): Promise<user_public_api.GetUserProfileResponse> {
    const _req = req || {};
    const url = this.genBaseURL('/api/user/v2/get_profile');
    const method = 'GET';
    const params = { user_id: _req['user_id'], bid: _req['bid'] };
    const headers = { Cookie: _req['Cookie'] };
    return this.request({ url, method, params, headers }, options);
  }

  /** POST /api/user/v2/update_profile */
  PublicUpdateUserProfile(
    req?: user_public_api.UpdateUserProfileRequest,
    options?: T,
  ): Promise<user_public_api.UpdateUserProfileResponse> {
    const _req = req || {};
    const url = this.genBaseURL('/api/user/v2/update_profile');
    const method = 'POST';
    const data = {
      user_unique_name: _req['user_unique_name'],
      name: _req['name'],
      avatar: _req['avatar'],
      signature: _req['signature'],
    };
    const headers = { Cookie: _req['Cookie'] };
    return this.request({ url, method, data, headers }, options);
  }

  /**
   * GET /api/report/get_meta
   *
   * 查询举报元数据
   */
  GetReportMeta(
    req?: report_public_api.GetReportMetaRequest,
    options?: T,
  ): Promise<report_public_api.GetReportMetaResponse> {
    const url = this.genBaseURL('/api/report/get_meta');
    const method = 'GET';
    return this.request({ url, method }, options);
  }

  /**
   * GET /api/report/query
   *
   * 查询举报信息
   */
  ReportQuery(
    req?: report_admin_api.ReportQueryRequest,
    options?: T,
  ): Promise<report_admin_api.ReportQueryResponse> {
    const _req = req || {};
    const url = this.genBaseURL('/api/report/query');
    const method = 'GET';
    const params = {
      object_id_list: _req['object_id_list'],
      object_type: _req['object_type'],
      report_time_begin: _req['report_time_begin'],
      report_time_end: _req['report_time_end'],
      report_uid: _req['report_uid'],
      page_num: _req['page_num'],
      page_size: _req['page_size'],
    };
    return this.request({ url, method, params }, options);
  }

  /**
   * POST /api/report/submit
   *
   * == 举报 ==
   *
   * -- http --
   *
   * 举报
   */
  ReportSubmit(
    req?: report_public_api.ReportSubmitRequest,
    options?: T,
  ): Promise<report_public_api.ReportSubmitResponse> {
    const _req = req || {};
    const url = this.genBaseURL('/api/report/submit');
    const method = 'POST';
    const data = {
      object_type: _req['object_type'],
      object_id: _req['object_id'],
      detail: _req['detail'],
    };
    const headers = { Cookie: _req['Cookie'] };
    return this.request({ url, method, data, headers }, options);
  }

  /** GET /api/report/get_report_times */
  GetReportTimes(
    req?: report_admin_api.GetReportTimesRequest,
    options?: T,
  ): Promise<report_admin_api.GetReportTimesResponse> {
    const _req = req || {};
    const url = this.genBaseURL('/api/report/get_report_times');
    const method = 'GET';
    const params = {
      object_id_list: _req['object_id_list'],
      object_type: _req['object_type'],
    };
    return this.request({ url, method, params }, options);
  }
}
/* eslint-enable */
