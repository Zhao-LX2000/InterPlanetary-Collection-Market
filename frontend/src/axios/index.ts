import axios from 'axios'

axios.defaults.timeout = 50000

axios.interceptors.request.use(config => {

  if(config.url.startsWith('http://127.0.0.1:4007/ipfs/')){
    console.log("here")
    config.headers.content
    return config
    
  }
  // 从localStorage中获取token
  const token = sessionStorage.getItem('token');

  // 如果token存在，则在请求头中添加Authorization字段
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config
}, error => {
  return Promise.error(error)
})

function login(loginData:any) {
    return axios.post(`http://localhost:8080/v1/user/login`, loginData)
}

function register(registerData:any) {
  return axios.post(`http://localhost:8080/v1/user/register`, registerData)
}

function queryForIPCBalance() {
  return axios.get(`http://localhost:8080/v1/artwork/getTokenBalance`)
}

function queryForETHBalance() {
  return axios.get(`http://localhost:8080/v1/artwork/getBalance`)
}

function uploadCollection(CollectionData:any) {
  return axios.post(`http://localhost:8080/v1/file/upload`, CollectionData, { headers: {
    "Content-type": "multipart/form-data"
  }})
}

function downloadFromIPFS(hash, filename) {
  return axios.get(`http://127.0.0.1:4007/ipfs/` + hash + "?download=true&filename=" + filename
  ,{   headers: { 
    'User-Agent': 'Apifox/1.0.0 (https://www.apifox.cn)', 
    'Accept': '*/*', 
    'Host': '127.0.0.1:4007', 
    'Connection': 'keep-alive'
 }})
}

function buyCollection(CollectionData) {
  return axios.post(`http://localhost:8080/v1/artwork/BuyCollection`,CollectionData)
}


function getAIThumbnail(input:string) {

  var data = JSON.stringify({
    "task": "text-to-image-synthesis",
    "inputs": [
      input
    ],
    "parameters": {
       "method": "text2image",
       "area": 983040,
       "aspect_ratio": 1.6666666666666667,
       "batch_size": 1
    },
    "urlPaths": {
       "inUrls": [
          {
             "value": input,
             "type": "text",
             "displayProps": {
                "size": "small"
             },
             "displayType": "OnlyTextArea",
             "name": "text",
             "title": "输入prompt",
             "validator": {
                "max_words": 75
             }
          },
          null
       ],
       "outUrls": [
          {
             "outputKey": "output_imgs"
          }
       ]
    }
  });
  
  return axios.post(`https://decoder.modelscope.cn/api/v1/models/damo/cv_diffusion_text-to-image-synthesis/widgets/predictions`, data, { headers: {
    "Content-type": "application/json"
  }})
}

function getAllCollection() {
  return axios.get(`http://localhost:8080/v1/file/list`)
}


export {
    login,
    register,
    uploadCollection,
    queryForIPCBalance,
    queryForETHBalance,
    getAIThumbnail,
    getAllCollection,
    downloadFromIPFS,
    buyCollection
}