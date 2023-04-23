<template>
    <div class="common-layout">
      <el-container>
        <el-header class="bgt">
          <TopBar/>
          <MainTop/>
        </el-header>
        <el-main>
          <MainPage/>
        </el-main>
      </el-container>
    </div>
  <Notification/>
</template>

<script setup lang="ts">
import { ref,reactive,onMounted,provide,readonly } from 'vue'
import {getAllCollection} from './axios/index'
import MainPage from './components/MainPage.vue'
import TopBar from './components/TopBar.vue'
import Notification from './components/Notification.vue'
import MainTop from './components/MainTop.vue'

const simpleList = reactive({
    data: []
  })
onMounted(() => {
    getAllCollection().then((resp) => {
        console.log(resp.data)
        simpleList.data = resp.data
    }).catch((err)=>{
        console.log(err)
    })
})

provide('simpleList', simpleList)


/* 
export default {
  components:{
    TopBar,
    Notification,
    MainTop,
    MainPage
  }
} */

</script>

<style scoped>
.el-header{
  height: 400px;
}
.el-main {
    padding-top: 10px;
}

.bgt{
  background-image: url('./assets/bg.png');
  background-size: cover;
  background-repeat: no-repeat;
  background-attachment:scroll;
}
</style>