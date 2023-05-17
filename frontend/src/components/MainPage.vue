<template>
  <p style="font-size: 30px;font-weight: 800;text-align: left;margin: 30px 40px;">最新数字藏品</p>
    <div style="margin: 0 40px;column-count: 6;column-gap: 30px;">
      
      <el-card v-for="(item, index) in simpleList.data" :key="index" class="box-card" 
          style="width: 200px;height: 220px; border-radius: 20px;margin-bottom: 20px;" 
          :body-style="{ padding: '0px'}"
          @click="buy(item)"
          >
        <img
          :src="item.url"
          class="image"
          style="height: 120px;width: auto;margin: auto;"
        />
        <div style="padding: 14px;">
          <span style="display: block;font-weight: bold;">{{item.workname}}</span>
          <span class="author" style="display: block;">@{{item.authorname}}</span>
          <div class="bottom" style="font-weight: bold;">
            <span>IPC:{{item.price}}</span>
            <span>{{item.createat.substring(0, 10)}}</span>
          </div>
        </div>
    </el-card>
    </div>
    
    <el-dialog v-model="outerVisible" width="60%" :title="detailInfo.title">
        <template #default>
          <div style="display:flex;flex-direction:row;">
            <div>
              <img
              :src="detailInfo.upImg"
              class="image"
              style="width: auto;height: 250px; margin: 0;"
              />
            </div>
            <div style="margin: auto;align-items: center;">
              <div>
                <el-avatar
                :size="60"
                src="https://www.nintendo.com.hk/switch/arzga/assets/img/top/chara_copy_1.png"/>
              </div>
              <span class="detailSpan" style="display: block;">拥有者：@{{detailInfo.owner}}</span>
              <span class="detailSpan" style="display: block;">作者：@{{detailInfo.author}}</span>
              <span class="detailSpan" style="display: block;">Token：{{detailInfo.token}}</span>
              <span class="detailSpan" style="display: block;">作品描述：{{detailInfo.description}}</span>
              <span class="detailSpan" style="display: block;">售价（IPC）：{{detailInfo.price}}</span>
            </div>
          </div>
          <el-dialog
            v-model="innerVisible"
            width="30%"
            title="确认购买"
            append-to-body
          >
            <el-input v-model="BuyPassword" placeholder="请输入你的账户密码"/>
            <template #footer>
          <div>
            <n-button round  @click="innerVisible = false" style="margin-right: 20px;">
              <template #icon>
                <n-icon>
                  <CloseCircleOutline />
                </n-icon>
              </template>
              取消
            </n-button>
            <n-button color="#ff4400" round  @click="buyAction" style="margin-right: 10px;" :loading="buyLoading">
              <template #icon>
                <n-icon>
                  <CheckmarkCircleOutline />
                </n-icon>
              </template>
              确认
            </n-button>
          </div>
        </template>
          </el-dialog>
        </template>
        <template #footer>
          
          <div class="dialog-footer">
            <n-button round style="margin-right: 20px;">
              <template #icon>
                <n-icon>
                  <DownloadOutline />
                </n-icon>
              </template>
              <a style="text-decoration: none;color: black;" :href="downloadurl" download>下载</a>
            </n-button>
            <n-button round @click="outerVisible = false" style="margin-right: 20px;">
              <template #icon>
                <n-icon>
                  <CloseCircleOutline />
                </n-icon>
              </template>
              取消
            </n-button>
            <n-button color="#386ac3" round  @click="innerVisible = true" style="margin-right: 10px;">
              <template #icon>
                <n-icon>
                  <BagCheck />
                </n-icon>
              </template>
              立即购买
            </n-button>
          </div>
        </template>
      </el-dialog>

</template>
<script lang="ts" setup>
import { ref,reactive,onMounted,provide,readonly, inject } from 'vue'
import { BagCheck, CloseCircleOutline, CheckmarkCircleOutline,DownloadOutline} from '@vicons/ionicons5'
import {buyCollection} from '../axios/index'
import {login} from '../axios/index'
import { ElNotification } from 'element-plus' 
const outerVisible = ref(false)
const innerVisible = ref(false)
const buyLoading = ref(false)


/* const simpleList = reactive({
    data: []
  })
   */

const simpleList = inject('simpleList')
const detailInfo = reactive({
  title : String,
  description : String,
  upImg:String,
  author:String,
  fileName:String,
  price:String,
  hash:String,
  owner:String,
  token:String,
  id:String
})
const downloadurl = ref('')
/* onMounted(() => {
    getAllCollection().then((resp) => {
        console.log(resp.data)
        simpleList.data = resp.data
    }).catch((err)=>{
        console.log(err)
    })
}) */
const buy = (item) => {
  detailInfo.title = item.workname;
  detailInfo.description = item.workdescription;
  detailInfo.upImg = item.url;
  detailInfo.author = item.authorname;
  detailInfo.fileName = item.filename;
  detailInfo.price = item.price;
  detailInfo.hash = item.hash;
  detailInfo.owner = item.owners; 
  detailInfo.token = item.hash;
  detailInfo.id = item.ID;
  downloadurl.value = "http://127.0.0.1:4007/ipfs/" + detailInfo.hash + "?download=true&filename=" + detailInfo.fileName;
  outerVisible.value=true;
}

const BuyPassword = ref('')

const buyAction = () => {
  buyLoading.value = true;
  let formData = new FormData();
    formData.append("cid", String(detailInfo.id))
    formData.append("toUser", String(detailInfo.owner))
    formData.append("pass", BuyPassword.value)
    buyCollection(formData).then((resp) => {
              openSuccess("购买成功")
              console.log(resp.data)
              buyLoading.value = false;
          }).catch((err)=>{
              openError("购买失败")
              console.log(err)
              buyLoading.value = false;
          })
  }

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

.author {
  font-size: 12px;
  color: #999;
}

.bottom {
  margin-top: 13px;
  margin-left: 2px;
  margin-right: 2px;
  line-height: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.button {
  padding: 0;
  min-height: auto;
}


.image {
  width: 100%;
  display: block;
}

.detailSpan {
  margin-top:10px;
}

</style>