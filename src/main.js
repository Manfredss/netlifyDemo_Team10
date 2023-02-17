import { createApp } from 'vue'
import App from './App.vue'
import axios from 'axios'
import bodyParser from 'body-parser'
// import Vue from 'vue'
// import router from './router'

import './assets/main.css'


axios.interceptors.request.use((config)=>{
    config.headers.Authorization = localStorage.getItem('Token')
    return config
}, error => {
    return Promise.reject(error)
})
// App.use(bodyParser.json())

createApp(App).mount('#app');
// const bodyParser = require('body-parser')
// App.use(bodyParser.urlencoded({ extended: false }))
// App.use(bodyParser.json())
