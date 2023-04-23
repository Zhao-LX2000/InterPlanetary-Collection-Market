<template>
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
      @select="handleSelect"
      
      style="margin:0 40px;background-color:transparent"
    >
    
     <el-image style="width: 350px; max-height:40px; margin: auto;" :src="logoimg" fit="fill" />
      <div class="flex-grow" />
      <!-- *******************************************. -->
      <Upload :isLogin="isLogin"/>
<!--       <n-button color="#8a2be2" round style="margin: auto 20px;" @click="searchBalance">
        <template #icon>
          <n-icon>
            <se />
          </n-icon>
        </template>
        查询余额
      </n-button> -->
      <el-popover
      :width="300"
      popper-style="box-shadow: rgb(14 18 22 / 35%) 0px 10px 38px -10px, rgb(14 18 22 / 20%) 0px 10px 20px -15px; padding: 20px;"
      >
      
      <template #reference>
        <el-avatar
        src="https://www.nintendo.com.hk/switch/arzga/assets/img/top/chara_multiplay_1.png"
        style="margin: auto 20px;"
      />
      </template>
      <template #default>
        <div v-if="isLogin"
          class="demo-rich-conent"
          style="display: flex; gap: 16px; flex-direction: column;align-items: center;"
        >
          <el-avatar
            :size="60"
            src="https://www.nintendo.com.hk/switch/arzga/assets/img/top/chara_copy_1.png"
            style="margin-bottom: 8px"
          />
            <p
              class="demo-rich-content__name"
              style="margin: 0; font-weight: 500"
            >
              @{{ loginUserName }}
            </p>
          <div style="font-weight: 600;">
            <span style="margin-right: 20px;">ETH:{{ETH}}</span>
            <span>IPC:{{IPC}}</span>
          </div>
          <div style="font-size: 16px;">
            <n-button color="#8a2be2" round  @click="loginOut" style="margin-right: 10px;">
              <template #icon>
                <n-icon>
                  <ExitOutline />
                </n-icon>
              </template>
              退出登录
            </n-button>
            <n-button color="#1ece9b" round  @click="searchBalance">
              <template #icon>
                <n-icon>
                  <Reload />
                </n-icon>
              </template>
              更新余额
            </n-button>
          </div>
        </div>
        <div v-else
          class="demo-rich-conent"
          style="display: flex; width: auto;align-items: center;"
        >
          <div>
            <el-avatar
            :size="60"
            src="https://www.nintendo.com.hk/switch/arzga/assets/img/top/chara_copy_1.png"/>
          </div>
          <div style="text-align: center; margin-left: 30px;">
            <el-button size="large" type="primary" @click="dialogLoginFormVisible = true">登录</el-button>
            <el-button size="large" type="primary" @click="dialogRegisterFormVisible = true">注册</el-button>
          </div>
        </div>
      </template>
    </el-popover>
      <el-menu-item index="1"> <p style="margin: 0;font-weight: 700;font-size: large;">Center</p></el-menu-item>
    </el-menu>
    <el-dialog v-model="dialogLoginFormVisible" title="登录" width="20%" :center="true">
        <el-form :model="loginForm" >
          <el-form-item label="用户名" label-width="70px">
            <el-input v-model="loginForm.username" autocomplete="off" />
          </el-form-item>
          <el-form-item label="密码"  label-width="70px">
            <el-input v-model="loginForm.password" autocomplete="off" />
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer" style="align-items: center;">
            <el-button size="large" @click="dialogLoginFormVisible = false">取消</el-button>
            <el-button size="large"  type="primary" @click="loginConfirm" :loading=isLoading>确认</el-button>
          </span>
        </template>
      </el-dialog>

    <el-dialog v-model="dialogRegisterFormVisible" title="注册" width="20%" :center="true">
      <el-form :model="registerForm" >
        <el-form-item label="用户名" label-width="70px">
          <el-input v-model="registerForm.username" autocomplete="off" />
        </el-form-item>
        <el-form-item label="密码"  label-width="70px">
          <el-input v-model="registerForm.password" autocomplete="off" />
        </el-form-item>
        <el-form-item label="确认密码"  label-width="70px">
          <el-input v-model="registerForm.password2" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer" style="align-items: center;">
          <el-button size="large" @click="dialogRegisterFormVisible = false">取消</el-button>
          <el-button size="large"  type="primary" @click="registerConfirm" :loading-icon="Eleme" :loading=isLoading>确认</el-button>
        </span>
      </template>
    </el-dialog>
  </template>
  
  <script lang="ts" setup>
  import {ref, reactive, onMounted} from 'vue'
  import {login, register, queryForIPCBalance, queryForETHBalance} from '../axios/index'
  import logoimg from '../assets/CIPCM星际藏品集市.png'
  import { Eleme } from '@element-plus/icons-vue'
  import Upload from './Upload.vue'
  import { ElNotification } from 'element-plus' 
  import { Reload, ExitOutline} from '@vicons/ionicons5'
  import { log } from 'console'
  const activeIndex = ref('1')
  const handleSelect = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
  }
  const dialogLoginFormVisible = ref(false)
  const dialogRegisterFormVisible = ref(false)
  const isLoading = ref(false)
  const isLogin = ref(false)
  const loginUserName = ref('')
  const ETH = ref(0)
  const IPC = ref(0)
  onMounted(() => {
    const token = sessionStorage.getItem('token')
    // 如果token存在，则在请求头中添加Authorization字段
    if (!token) {
      console.log("false")
      isLogin.value=false
      return false
    } else {
      console.log("true")
      loginUserName.value=sessionStorage.getItem('user')
      searchBalance()
      isLogin.value=true
      return true
    }
  })
  
  const loginOut = () => {
    sessionStorage.removeItem('token')
    sessionStorage.removeItem('user')
    loginUserName.value='';
    isLogin.value=false
    openSuccess("成功退出")
  }
  const loginConfirm = () => {
    login(loginForm).then((resp) => {
            console.log(resp.data)
            if(resp.data.code != 0){
              openError("登录失败")
              return
            }
            dialogLoginFormVisible.value = false
            const token = resp.data.token;
            sessionStorage.setItem('token', token);
            sessionStorage.setItem('user',loginForm.username);
            searchBalance();
            loginUserName.value=loginForm.username;
            loginForm.username='';
            loginForm.password='';
            isLogin.value=true
            openSuccess("登录成功")
        }).catch((err)=>{
            console.log(err)
            dialogLoginFormVisible.value = false
            openError("登录失败")
        })
  }
  const registerConfirm = () => {
    isLoading.value = true
    register(registerForm).then((resp) => {
            console.log(resp.data)
            if(resp.data.message != 'err_code=0, err_msg=Success'){
              openError("注册失败")
              isLoading.value = false
              return
            }
            dialogRegisterFormVisible.value = false
            isLoading.value = false
            registerForm.password=''
            registerForm.password2=''
            registerForm.username=''
            openSuccess("注册成功")
        }).catch((err)=>{
            console.log(err)
            dialogRegisterFormVisible.value = false
            isLoading.value = false
            openError("注册失败")
        })
  }
  const searchBalance = () => {
    queryForIPCBalance().then((resp) => {
            console.log("IPC:")
            console.log(resp.data)
            IPC.value=resp.data.Balance
            queryForETHBalance().then((resp) => {
              console.log("ETH:")
              console.log(resp.data)
              ETH.value=parseFloat(resp.data.Balance) / 1000000000000000000;
              openSuccess("更新完成")
            }).catch((err)=>{
                console.log(err)
                openError("余额查询失败")
            })
        }).catch((err)=>{
            console.log(err)
            openError("余额查询失败")
        })
  }
  
  const loginForm = reactive({
    username: '',
    password: '',
  })
  const registerForm = reactive({
    username: '',
    password: '',
    password2: '',
  })
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
  </script>
  
  <style>
  .flex-grow {
    flex-grow: 1;
    
  }
  .block {
    align-items: center;
    display: flex;
  }
  </style>
  