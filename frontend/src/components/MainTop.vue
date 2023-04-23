<template>
  <n-tabs 
      default-value="all" 
      style="margin: 10px 40px;"
      size="large"
      >
    <n-tab name="all" tab="全部">
      <n-gradient-text     :gradient="{
                              from: 'rgb(255, 255, 255)',
                              to: 'rgb(255, 255, 255)'
                            }" 
        type="success" :size="16" >
        全部
      </n-gradient-text>
    </n-tab>
    <n-tab name="images" tab="图片">
      <n-gradient-text     :gradient="{
                              from: 'rgb(255, 255, 255)',
                              to: 'rgb(255, 255, 255)'
                            }" 
        type="success" :size="16" >
        图片
      </n-gradient-text>
    </n-tab>
    <n-tab name="text" tab="文本">
      <n-gradient-text     :gradient="{
                              from: 'rgb(255, 255, 255)',
                              to: 'rgb(255, 255, 255)'
                            }"  
        type="success" :size="16" >
        文本
      </n-gradient-text>
    </n-tab>
    <n-tab name="video" tab="视频">
      <n-gradient-text     :gradient="{
                              from: 'rgb(255, 255, 255)',
                              to: 'rgb(255, 255, 255)'
                            }" 
        type="success" :size="16" >
        视频
      </n-gradient-text>
    </n-tab>
  </n-tabs>
  <el-carousel :interval="4000" type="card" height="250px" style="margin: 0 40px;">
    <el-carousel-item v-for="(item, index) in simpleList.data" :key="index">
      <img :src="item.url" alt="" style="width: 100%;
                                      object-fit: cover;
                                      top: 50%;
                                      left: 50%;
                                      min-width: 100%;
                                      position: absolute;
                                      min-height: 100%;
                                      transform: translate(-50%, -50%);"
                                      >
      <!-- <h3 text="2xl" justify="center">{{ item }}</h3> -->
    </el-carousel-item>
  </el-carousel>
</template>

<script lang="ts" setup>
import { reactive,onMounted,inject } from 'vue'
const simpleList = inject('simpleList')
</script>

<style scoped>
.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}

.n-gradient-text {
  font-weight: bold;
}

.light-green {
  height: 200px;
  background-color: rgba(0, 128, 0, 0.12);
}
.green {
  height: 200px;
  background-color: rgba(0, 128, 0, 0.24);
}

</style>