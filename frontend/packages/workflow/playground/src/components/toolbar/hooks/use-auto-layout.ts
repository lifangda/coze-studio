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
 
import { useCallback } from 'react';

import {
  usePlaygroundTools,
  useService,
} from '@flowgram-adapter/free-layout-editor';

import { WorkflowLayoutShortcutsContribution } from '@/shortcuts';

export const useAutoLayout = () => {
  const tools = usePlaygroundTools();
  const autoLayoutShortcut = useService(WorkflowLayoutShortcutsContribution);
  const autoLayout = useCallback(async () => {
    tools.fitView();
    await autoLayoutShortcut.autoLayout();
    tools.fitView();
  }, [tools, autoLayoutShortcut]);
  return autoLayout;
};
