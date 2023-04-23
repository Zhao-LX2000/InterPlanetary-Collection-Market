import { createApp } from "vue";
import App from "./App.vue";
import axios from 'axios'
import naive from "naive-ui";
import ElementPlus from "element-plus";
// import all element css, uncommented next line
import "element-plus/dist/index.css";

// or use cdn, uncomment cdn link in `index.html`

import "~/styles/index.scss";
import 'uno.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
// If you want to use ElMessage, import it.
import "element-plus/theme-chalk/src/message.scss"
// 通用字体
import 'vfonts/Lato.css'
import router from './router/router.js'
const app = createApp(App);
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }
app.config.globalProperties.axios = axios
app.use(ElementPlus);
app.use(naive);
app.use(router).mount("#app");
