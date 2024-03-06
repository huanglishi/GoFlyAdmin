<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :isPadding="false" :loading="loading" width="1000px" @height-change="onHeightChange" :minHeight="modelHeight" :title="getTitle" @ok="handleSubmit">
    <div class="addFormbox" :style="{'min-height':`${windHeight}px`}">
      <div class="tabs-header" v-if="isEditor">
        <div class="tabs-nav-wrap">
            <div class="tap_item" v-for="iten in tapList" :class="{item_active:activeKey==iten.id}" @click="()=>{activeKey=iten.id}">
                <div class="label">{{iten.name}}</div>
            </div>
        </div>
        <div class="tabs-bar" :style="{top: `${(activeKey-1)*64}px`,height: `64px`}"></div>
      </div>
      <div class="tabs-content" :class="{addpadding:!isEditor}">
        <a-form ref="formRef" :model="formData" auto-label-width>
          <div class="content_box">
              <!--基础信息-->
              <a-scrollbar v-show="activeKey==1" style="overflow: auto;" :style="{height:`${windHeight}px`}">
                <div class="besecontent" >
                  <a-row :gutter="16">
                    <a-col :span="12">
                      <a-form-item field="cid" label="选择分类" validate-trigger="input" >
                        <a-select v-model="formData.cid" :options="cateList" placeholder="请选择分类" />
                      </a-form-item>
                    </a-col>
<!--replaceTpl-->  
                  </a-row>
                </div>
              </a-scrollbar>
              <!--高级信息-->
              <div class="hcontent" v-show="activeKey==2" :style="{height:`${windHeight}px`}">
                <Editor :minHeight="windHeight" ref="editorRef" @updata="handleEditUpdta"/>
              </div>
          </div>
        </a-form>
      </div>
    </div>
    <!--附件-->
    <FileManage @register="registerFileModal" @success="selectImg"/>
  </BasicModal>
</template>
<script lang="ts">
  import { defineComponent, ref, computed, unref} from 'vue';
  import { BasicModal, useModalInner,useModal} from '/@/components/Modal';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import useLoading from '@/hooks/loading';
  import { cloneDeep } from 'lodash-es';
  //api
  import { save,getContent } from './api';
  import { getCate } from './cate/api';
  import { Message ,TreeNodeData} from '@arco-design/web-vue';
  import type { RequestOption} from '@arco-design/web-vue/es/upload/interfaces';
  import { userUploadApi } from '@/api/common';
  import Editor from "@/components/Editor/Main.vue"; // @ is an alias to /src
  import FileManage from '@/views/datacenter/attachment/components/FileManage.vue';
  export default defineComponent({
    name: 'AddBook',
    components: { BasicModal,Editor,FileManage },
    emits: ['success'],
    setup(_, { emit }) {
      const [registerFileModal, { openModal:openFileModal }] = useModal();
      const visibleimage=ref(false);
      //判断是否存在编辑器
      const isEditor=ref(true);
      const isUpdate = ref(false);
      const cateList = ref<TreeNodeData[]>([]);
      const activeKey= ref(1);
      const modelHeight= ref(620);
      const windHeight= ref(620);
      //表单
      const formRef = ref<FormInstance>();
      //表单字段
      const basedata={
            id:0,
replaceField:null
        }
      const formData = ref(basedata)
      //编辑器
      const editorRef = ref();
      const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
          formRef.value?.resetFields()
          activeKey.value=1
          setModalProps({ confirmLoading: false });
          const mdata = await getCate({});
          const parntList_df : any=[{value: 0,label: "未选分类"}];
          if(mdata){
            cateList.value=parntList_df.concat(mdata)
          }else{
            cateList.value=parntList_df
          }
          isUpdate.value = !!data?.isUpdate;
          if (unref(isUpdate)) {
            formData.value=cloneDeep(data.record)
            const mewdata = await getContent({id:data.record.id});
            formData.value=Object.assign({},formData.value,mewdata)
            if(editorRef.value)
            editorRef.value.setVal(mewdata.content)
          }else{
            formData.value=basedata
          }
      });
      const getTitle = computed(() => (!unref(isUpdate) ? '新增数据' : '编辑数据'));
     //点击确认
     const { loading, setLoading } = useLoading();
     const handleSubmit = async () => {
      try {
          const res = await formRef.value?.validate();
          if (!res) {
            setLoading(true);
            Message.loading({content:"提交中",id:"upStatus",duration:2000})
            let savedata=cloneDeep(unref(formData))
            await save(savedata);
            Message.success({content:"提交成功",id:"upStatus",duration:2000})
            closeModal()
            emit('success');
            setLoading(false);
          }
        } catch (error) {
          setLoading(false);
          Message.clear("top")
        }
      };
      //上传附件
      const customupFile = (options: RequestOption) => {
            Message.loading({content:"上传中",id:"upStatus",duration:0})
          const controller = new AbortController();
            (async function requestWrap() {
              const {
                onProgress,
                onError,
                onSuccess,
                fileItem,
              } = options;
              onProgress(20);
              const onUploadProgress = (event: ProgressEvent) => {
                let percent;
                if (event.total > 0) {
                  percent = (event.loaded / event.total) * 100;
                }
                onProgress(parseInt(String(percent), 10), event);
              };
              try {
                //开始手动上传
                const filename=fileItem?.name||""
                const resdata = await userUploadApi({ name: 'file', file: fileItem.file as Blob, filename,data:{cid:0}},onUploadProgress);
                //更新附件
                if(resdata){
                  formData.value['replaceFile']=resdata.url
                }
                Message.success({content:"上传成功",id:"upStatus",duration:1500})
              } catch (error) {
                onError(error);
                Message.error({content:"上传失败",id:"upStatus",duration:2000})
              }
            })();
            return {
              abort() {
                controller.abort();
              },
            };
      };
       //上传图片
       const UpImage=()=>{
        openFileModal(true, {
          filetype:"image",
          getnumber: "one",//one 单张
          openfrom: "use",//manage管理 use选择使用
        });
      }
      //选择附件返回
      const selectImg=(item:any)=>{
        if(item.type=="more"){
             item.list.forEach((img:any) => {
               console.log("多张附件返回:",img)
             });
          }else if(item.type=="one"){
            formData.value['replaceimage']=item.url
          }
      }
      //编辑器返回数据
      const handleEditUpdta=(val:string)=>{
        formData.value["content"]=val
      }
       //监听高度
       const onHeightChange=(val:any)=>{
        windHeight.value=val
      }
      return { 
        registerModal, 
        getTitle, 
        handleSubmit,
        formRef,
        loading,
        formData,
        cateList,
        OYoptions:[
          { label: '否', value: 0 },
          { label: '是', value: 1 },
        ],
        tapList:[
          {id:1,name:"基础内容"},
          {id:2,name:"详细内容"},
        ],
        activeKey,
        customupFile,
        handleEditUpdta,
        modelHeight,
        editorRef,
        onHeightChange,windHeight,
        isEditor,
        registerFileModal,selectImg,UpImage,visibleimage,
      };
    },
  });
</script>
<style lang="less" scoped>
  @import '@/assets/style/formlayer.less';
  .upfilezip{
    display: flex;
    .upbtn{
      padding-right: 10px;
    }
    .showfile{
      flex: 1;
      height: 32px;
      line-height: 32px;
      a{
        text-decoration: none;
      }
    }
  }
  //上传图片
  .upimagebox{
    display: flex;
    .imagebtn{
      position: relative;
      width: 160px;
      height: 90px;
      background-color: var(--color-neutral-1);
      border-radius: 4px;
      overflow: hidden;
      -ms-flex-negative: 0;
      flex-shrink: 0;
      //预览
      .upload-show-picture{
        position: relative;
        box-sizing: border-box;
        width: 100%;
        height: 100%;
        overflow: hidden;
        display: flex;
        align-items: center;
        justify-content: center;
        img{
          height: 100%;
        }
        &:hover{
          .upload-show-picture-mask{
             opacity: 1;
          }
        }
        .upload-show-picture-mask{
            position: absolute;
            top: 0;
            right: 0;
            bottom: 0;
            left: 0;
            color: var(--color-white);
            font-size: 16px;
            line-height: 90px;
            text-align: center;
            background: rgba(0, 0, 0, 0.5);
            cursor: pointer;
            opacity: 0;
            transition: opacity 0.1s cubic-bezier(0, 0, 1, 1);
        }
      }
      .upload-picture-card{
        width: 100%;
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        .upload-picture-card-text{
           text-align: center;
           color:  var(--color-neutral-5);
        }
      }
    }
  }
</style>

