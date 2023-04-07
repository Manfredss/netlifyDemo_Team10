import { createApp } from 'vue'
import App from './App.vue'
import axios from 'axios'
import bodyParser from 'body-parser'
import naive from 'naive-ui'
// import Vue from 'vue'
// import router from './router'
// import "lib-flexible"
// import autoprefixer from 'autoprefixer'
// import postcssConfig from '../postcss.config'
import './assets/main.css'


axios.interceptors.request.use((config)=>{
    config.headers.Authorization = localStorage.getItem('Token')
    return config
}, error => {
    return Promise.reject(error)
})
// App.use(bodyParser.json())

const app = createApp(App)
app.use(naive)
app.mount('#app');
// createApp(App).mount('#app');

// const bodyParser = require('body-parser')
// App.use(bodyParser.urlencoded({ extended: false }))
// App.use(bodyParser.json())
