<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :isPadding="false" :loading="loading" width="800px" @height-change="onHeightChange" :minHeight="modelHeight" :title="getTitle" @ok="handleSubmit">
    <div class="addFormbox" :style="{'min-height':`${windHeight}px`}">
      <div class="tabs-content" >
        <a-form ref="formRef" :model="formData" auto-label-width>
          <div class="content_box">
              <!--基础信息-->
              <a-scrollbar   style="overflow: auto;" :style="{height:`${windHeight}px`}">
                <div class="besecontent" >
                  <a-row :gutter="16">
<!--replaceTpl-->
                  </a-row>
                </div>
              </a-scrollbar>
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
  import { BasicModal, useModalInner,useModal } from '/@/components/Modal';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import useLoading from '@/hooks/loading';
  import { cloneDeep } from 'lodash-es';
  //api
  import { save } from './api';
  import { Message } from '@arco-design/web-vue';
  import type { RequestOption} from '@arco-design/web-vue/es/upload/interfaces';
  import { userUploadApi } from '@/api/common';
  import FileManage from '@/views/datacenter/attachment/components/FileManage.vue';
  export default defineComponent({
    name: 'AddBook',
    components: { BasicModal,FileManage },
    emits: ['success'],
    setup(_, { emit }) {
      const [registerFileModal, { openModal:openFileModal }] = useModal();
      const visibleimage=ref(false);
      const isUpdate = ref(false);
      const modelHeight= ref(420);
      const windHeight= ref(420);
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
          setModalProps({ confirmLoading: false });
          isUpdate.value = !!data?.isUpdate;
          if (unref(isUpdate)) {
            formData.value=cloneDeep(data.record)
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
            Message.loading({content:"提交中",id:"upStatus",duration:0})
            let savedata=cloneDeep(unref(formData))
            await save(savedata);
            Message.success({content:"提交成功",id:"upStatus",duration:1500})
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
                Message.success({content:"上传成功",id:"upStatus",duration:2000})
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
        OYoptions:[
          { label: '否', value: 0 },
          { label: '是', value: 1 },
        ],
        SHoptions:[
            { label: '正常', value: 0 },
            { label: '禁用', value: 1 },
        ],
        customupFile,
        modelHeight,
        editorRef,
        onHeightChange,windHeight,
        registerFileModal,selectImg,UpImage,visibleimage,
      };
    },
  });
</script>
<style lang="less" scoped>
  @import '@/assets/style/formlayer.less';
  
  .addpadding{
    padding: 0px 20px;
  }
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
            opacity: 0;
            transition: opacity 0.1s cubic-bezier(0, 0, 1, 1);
            .opbtn{
              cursor: pointer;
            }
        }
      }
      .upload-picture-card{
        width: 100%;
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        user-select: none;
        cursor: pointer;
        .upload-picture-card-text{
           text-align: center;
           color:  var(--color-neutral-5);
        }
      }
    }
  }
</style>

