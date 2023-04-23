<template>

  <n-button v-if="isLogin" color="#fb7299" round style="margin: auto 10px;" @click="centerDialogVisible = true">
        <template #icon>
          <n-icon>
            <cl />
          </n-icon>
        </template>
        上传作品
      </n-button>
  <el-dialog 
    v-model="centerDialogVisible" 
    title="作品上传" 
    width="35%" 
    center
    style="border-radius: 45px">
    <div style="display: flex;align-items: center;flex-direction: column;">
      <el-form 
        ref="ruleFormRef"
        :inline="true"  
        :rules="rules"
        :model="ruleForm"
        style="display: flex;align-items: center;flex-direction: column;padding: 0 30px;"
      >
        <el-form-item label="请输入作品名" prop="workName">
          <el-input v-model="ruleForm.workName" placeholder="起个好听的名字啦~" />
        </el-form-item>
        <el-form-item label="请输入作品描述" prop="workDescription" >
          <el-input v-model="ruleForm.workDescription" placeholder="介绍一下你的作品~" />
        </el-form-item>
        <el-form-item label="请输入作品价格" prop="price" >
          <el-input v-model="ruleForm.price" placeholder="你的作品价值~" />
        </el-form-item>
      </el-form>
      <el-image style="width: 400px; height: 240px; margin-top: 20px;" :src="ruleForm.url" fit="cover">
        <template #error>
          <div class="image-slot">
            <el-icon><icon-picture /></el-icon>
          </div>
        </template>
      </el-image>
    </div>
    <template #footer>
      <span class="dialog-footer">
        <el-upload
          ref="uploadRef"
          class="upload-demo"
          :multiple="false"
          action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
          :auto-upload="false"
          :http-request="uploadFile"
          :on-change="handleChange"
          :on-remove="handleChange"
        >
          <template #trigger >
            <el-button type="primary">选择文件</el-button>
          </template>
          <el-button class="ml-3" type="success" style="margin-right: 8px;" 
            :loading="isGenerated"
            @click="generateThumbnail"
            :disabled="ruleForm.workDescription === ''">
            生成封面
          </el-button>
          <el-button class="ml-3" type="success" style="margin-right: 8px;" @click="submitUpload"
           :disabled="canUpload" :loading="isUploading">
            上传
          </el-button>
          <el-button class="ml-3" color="#626aef" type="primary" :icon="Delete" @click="doClear">
            清空
          </el-button>
        </el-upload>
      </span>
    </template>
  </el-dialog>

  </template>
  <script lang="ts" setup>
  import { ref, reactive } from 'vue'
  import {uploadCollection, getAIThumbnail, getAllCollection} from '../axios/index'
  import {CloudUpload as cl} from '@vicons/ionicons5'
  import type { UploadInstance, FormInstance, FormRules } from 'element-plus'
  import { Delete} from '@element-plus/icons-vue'
  import { Picture as IconPicture } from '@element-plus/icons-vue'
  import { ElNotification } from 'element-plus' 
    // 3.服务器每次返回信息都会执行一次onmessage方法
    const openSuccess = (ms) => {
      ElNotification({
          title: '通知',
          dangerouslyUseHTMLString: true,
          message: ms,
          duration: 7000,
          type: 'success',
          position: 'bottom-right',
      })
    }
    const openError = (ms) => {
      ElNotification({
          title: '通知',
          dangerouslyUseHTMLString: true,
          message: ms,
          duration: 7000,
          type: 'error',
          position: 'bottom-right',
      })
    }
  const ruleFormRef = ref<FormInstance>()
  const isGenerated = ref(false)
  const isUploading = ref(false)
  
  const checkFunc = (rule: any, value: any, callback: any) => {
    if (value === '') {
      return callback(new Error('你为什么不输入啊啊啊啊啊啊啊'))
    }
    setTimeout(() => {}, 1000)
  }
  const rules = reactive<FormRules>({
    workName: [{ validator: checkFunc, trigger: 'blur' }],
    workDescription: [{ validator: checkFunc, trigger: 'blur' }],
    price: [{ validator: checkFunc, trigger: 'blur' }],
  })
  const ruleForm = reactive({
    workName: '',
    workDescription: '',
    url : '',
    price : 0,
  })
  const canUpload = ref(true);
  const uploadRef = ref<UploadInstance>()
  const centerDialogVisible = ref(false)
  const submitUpload = () => {
    uploadRef.value!.submit()
  }
  const doClear= ()=>{
    uploadRef.value.clearFiles()
    canUpload.value=true;
  }

  const handleChange= (file, fileList)=>{
    if (fileList.length > 0) {
        console.log('已选择文件');
        canUpload.value=false;
      } else {
        console.log('未选择文件');
        canUpload.value=true;
      }
  }
  const generateThumbnail = () => {
    isGenerated.value = true
    openSuccess("开始生成")
    getAIThumbnail(ruleForm.workDescription).then((resp) => {
            console.log(resp.data)
            ruleForm.url=resp.data.Data.data.output_imgs[0]
            isGenerated.value = false
            openSuccess("生成完成")
        }).catch((err)=>{
            console.log(err)
            isGenerated.value = false
            openError("生成失败")
        })
  }
  const uploadFile =(fileObj)=>{
    if(ruleForm.url === '') {
      openError("请先生成封面")
      return
    }
    isUploading.value = true
    openSuccess("开始上传")
    let formData = new FormData();
    formData.set("workName", ruleForm.workName)
    formData.append("workDescription", ruleForm.workDescription)
    formData.append("file", fileObj.file);
    formData.append("url", ruleForm.url);
    formData.append("price", String(ruleForm.price));
    uploadCollection(formData).then((resp) => {
            console.log(resp.data)
            isUploading.value = false
            openSuccess("上传完成")
            centerDialogVisible.value=false
        }).catch((err)=>{
            console.log(err)
            isUploading.value = false
            openError("上传失败")
        })
  }
  defineProps({
      isLogin: Boolean
  })
</script>
<style>
.dialog-footer button:first-child {
  margin-right: 10px;
}

.image-slot {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  background: var(--el-fill-color-light);
  color: var(--el-text-color-secondary);
  font-size: 30px;
}
.image-slot .el-icon {
  font-size: 30px;
}
</style>