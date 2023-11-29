import { computed } from 'vue';
import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
export const columns = computed<TableColumnData[]>(() => [
{},
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
      fixed: 'right',
      width: 100,
      align:"center"
    },
  ]);