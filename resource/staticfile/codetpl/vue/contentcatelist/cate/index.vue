<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :isPadding="false" :footer="false" :loading="loading" width="1000px" @height-change="onHeightChange" :minHeight="modelHeight" :title="getTitle" >
    <div class="modalbox" :style="{'min-height':`${windHeight}px`}">
      <div class="table-content">
        <a-row style="margin-bottom: 10px">
        <a-col :span="16" >
          <a-space>
            <a-input :style="{width:'160px'}"  v-model="formModel.name" placeholder="名称" allow-clear />
            <a-range-picker v-model="formModel.createdTime" :style="{width:'200px'}" />
            <a-button type="primary" @click="search">
              <template #icon>
                <icon-search />
              </template>
              查询
            </a-button>
            <a-button @click="reset">
              {{ $t('searchTable.form.reset') }}
            </a-button>
          </a-space>
        </a-col>
        <a-col
          :span="8"
           style="text-align: right;"
        >
        <a-space>
          <a-button type="primary" @click="createData">
            <template #icon>
              <icon-plus />
            </template>
            {{ $t('searchTable.operation.create') }}
          </a-button>
          </a-space>
        </a-col>
      </a-row>
        <a-table
          row-key="id"
          :loading="loading"
          :pagination="pagination"
          :columns="(columns as TableColumnData[])"
          :data="renderData"
          :bordered="{wrapper:true,cell:true}"
          size="medium"
          :default-expand-all-rows="true"
          ref="artable"
          @page-change="handlePaageChange" 
          @page-size-change="handlePaageSizeChange" 
        >
          <template #createtime="{ record }">
            {{dayjs(record.createtime*1000).format("YYYY-MM-DD")}}
          </template>
          <template #status="{ record }">
            
            <a-switch type="round" v-model="record.status" :checked-value="0" :unchecked-value="1" @change="handleStatus(record)">
                <template #checked>
                  开
                </template>
                <template #unchecked>
                  关
                </template>
              </a-switch>
          </template>
          <template #operations="{ record }">
            <Icon icon="svgfont-bianji1" class="iconbtn" @click="handleEdit(record)" :size="18" color="#0960bd"></Icon>
            <a-divider direction="vertical" />
            <a-popconfirm content="您确定要删除吗?" @ok="handleDel(record)">
              <Icon icon="svgfont-icon7" class="iconbtn" :size="18" color="#ed6f6f"></Icon>
            </a-popconfirm>
          </template>
        </a-table>
      </div>
    </div>
       <!--表单-->
   <AddForm ref="addFormRef" @success="search"/>
  </BasicModal>
</template>
<script lang="ts">
  import { defineComponent, ref, computed, unref,reactive} from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import useLoading from '@/hooks/loading';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import { Pagination } from '@/types/global';
  //数据
  import { columns} from './data';
  import dayjs from 'dayjs';
  //api
  import { getList,upStatus, del} from './api';
  import { IconPicker ,Icon} from '@/components/Icon';
  import { Message } from '@arco-design/web-vue';
  import AddForm from './AddForm.vue';
  export default defineComponent({
    name: 'cateindex',
    components: { BasicModal,IconPicker,Icon,AddForm },
    emits: ['success'],
    setup(_, { emit }) {
      const isUpdate = ref(false);
      const modelHeight= ref(620);
      const windHeight= ref(620);
      //表格
      const renderData = ref([]);
      const { loading, setLoading } = useLoading(true);
      const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
          setModalProps({ confirmLoading: false });
          isUpdate.value = !!data?.isUpdate;
          fetchData()
      });
      //查询字段
      const generateFormModel = () => {
        return {
          title: '',
          name: '',
          createdTime: [],
          status: '',
        };
      };
      const formModel = ref(generateFormModel());
      //加载数据
      const fetchData = async () => {
          setLoading(true);
          try {
            const data= await getList({page:pagination.current,pageSize:pagination.pageSize,...formModel.value});
            renderData.value = data.items;
            pagination.current = data.page;
            pagination.total = data.total;
          } catch (err) {
            // you can report use errorHandler or other
          } finally {
            setLoading(false);
          }
        };
      const getTitle = computed(() => (!unref(isUpdate) ? '分类管理' : '编辑数据'));
     //点击确认
       //监听高度
       const onHeightChange=(val:any)=>{
        windHeight.value=val
      }
      //表格
        //分页
      const basePagination: Pagination = {
        current: 1,
        pageSize: 10,
      };
      const pagination = reactive({
        ...basePagination,
        showTotal:true,
        showPageSize:true,
      });
      //分页
      const handlePaageChange = (page:any) => {
        pagination.current=page
        fetchData();
      }
      //分页总数
      const handlePaageSizeChange = (pageSize:any) => {
        pagination.pageSize=pageSize
        fetchData();
      }
       //更新状态
      const handleStatus=async(record:any)=>{
        try {
            Message.loading({content:"更新状态中",id:"upStatus"})
          const res= await upStatus({id:record.id,status:record.status});
          if(res){
            Message.success({content:"更新状态成功",id:"upStatus"})
          }
        }catch (error) {
          Message.clear("top")
        } 
      }
      //删除数据
      const handleDel=async(record:any)=>{
          try {
              Message.loading({content:"删除中",id:"upStatus"})
            const res= await del({ids:[record.id]});
            if(res){
              fetchData();
              Message.success({content:"删除成功",id:"upStatus"})
            }
          }catch (error) {
            Message.clear("top")
          } 
      }
      //新增数据
      const addFormRef=ref()
      const createData=async()=>{
        addFormRef.value.ShowModal({
          isUpdate: false,
          record:null
        })
      }
      //编辑数据
      const handleEdit=async(record:any)=>{
        addFormRef.value.ShowModal({
          isUpdate: true,
          record:record
        })
      }
       //查找
      const search = () => {
        fetchData();
      };
      const reset = () => {
        formModel.value = generateFormModel();
        fetchData();
      };
      return { 
        registerModal, 
        getTitle, 
        columns,
        loading,
        dayjs,
        modelHeight,
        onHeightChange,windHeight,
        pagination,handlePaageChange,handlePaageSizeChange,
        handleStatus,handleDel,handleEdit,renderData,formModel,
        search,reset,createData,
        addFormRef,
      };
    },
  });
</script>
<style lang="less" scoped>
  .modalbox{
    padding: 10px;
    .table-content{

    }
  }
</style>

