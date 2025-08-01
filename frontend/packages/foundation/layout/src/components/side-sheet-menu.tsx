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
 
import { IconButton } from '@coze-arch/coze-design';
import { IconSideFoldOutlined } from '@coze-arch/bot-icons';

import { useOpenGlobalLayoutSideSheet } from './global-layout/hooks';

// 用于在移动端模式开启侧边栏
export const SideSheetMenu = () => {
  const open = useOpenGlobalLayoutSideSheet();

  return (
    <IconButton
      color="secondary"
      icon={<IconSideFoldOutlined className="coz-fg-primary text-base" />}
      onClick={open}
    />
  );
};

export default SideSheetMenu;
