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
 
import { type FC } from 'react';

import classNames from 'classnames';
import { I18n } from '@coze-arch/i18n';
import { Input } from '@coze-arch/coze-design';

import { type LiteralValueInputProps } from './type';

import styles from './styles.module.less';
export const InputString: FC<LiteralValueInputProps> = ({
  className,
  value,
  defaultValue,
  disabled,
  testId,
  onChange,
  onBlur,
  onFocus,
  placeholder,
  validateStatus,
  style,
}) => (
  <Input
    className={classNames(className, styles['input-string'])}
    data-testid={testId}
    disabled={disabled}
    defaultValue={defaultValue as string}
    value={value as string}
    onChange={onChange}
    onBlur={e => onBlur?.(e.target.value)}
    onFocus={onFocus}
    placeholder={
      placeholder || I18n.t('workflow_detail_node_input_selectvalue')
    }
    validateStatus={validateStatus}
    style={{ ...style, padding: '0 6px' }}
    size="small"
  />
);
